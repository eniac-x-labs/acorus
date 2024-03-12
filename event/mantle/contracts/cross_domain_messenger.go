package contracts

import (
	"fmt"
	"github.com/cornerstone-labs/acorus/event/mantle/op-bindings/bindings"
	"github.com/google/uuid"
	"math/big"

	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// Standard ABI types copied from golang ABI tests
	uint256Type, _ = abi.NewType("uint256", "", nil)
	bytesType, _   = abi.NewType("bytes", "", nil)
	addressType, _ = abi.NewType("address", "", nil)

	CrossDomainMessengerLegacyRelayMessageEncoding = abi.NewMethod(
		"relayMessage",
		"relayMessage",
		abi.Function,
		"external", // mutability
		false,      // isConst
		true,       // payable
		abi.Arguments{ // inputs
			{Name: "target", Type: addressType},
			{Name: "sender", Type: addressType},
			{Name: "data", Type: bytesType},
			{Name: "nonce", Type: uint256Type},
		},
		abi.Arguments{}, // outputs
	)
)

type CrossDomainMessengerSentMessageEvent struct {
	Event                *event.ContractEvent
	Version              uint16
	MessageCalldata      []byte
	MessageHash          common.Hash
	Nonce                *big.Int
	SentMessageEventGUID uuid.UUID
	GasLimit             *big.Int
	FromAddress          common.Address
	ToAddress            common.Address
	ETHAmount            *big.Int
	ERC20Amount          *big.Int
	Data                 utils.Bytes
	HashVersion          uint64
	Timestamp            uint64
}

type CrossDomainMessengerRelayedMessageEvent struct {
	Event       *event.ContractEvent
	MessageHash common.Hash
}

func CrossDomainMessengerSentMessageEvents(contractAddress common.Address, chainId string,
	db *database.DB, fromHeight, toHeight *big.Int) ([]CrossDomainMessengerSentMessageEvent, error) {
	crossDomainMessengerAbi, err := bindings.CrossDomainMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	sentMessageEventAbi := crossDomainMessengerAbi.Events["SentMessage"]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: sentMessageEventAbi.ID}
	sentMessageEvents, err := db.ContractEvents.ContractEventsWithFilter(chainId, contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}
	if len(sentMessageEvents) == 0 {
		return nil, nil
	}

	sentMessageExtensionEventAbi := crossDomainMessengerAbi.Events["SentMessageExtension1"]
	contractEventFilter = event.ContractEvent{ContractAddress: contractAddress, EventSignature: sentMessageExtensionEventAbi.ID}
	sentMessageExtensionEvents, err := db.ContractEvents.ContractEventsWithFilter(chainId, contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	if len(sentMessageExtensionEvents) > len(sentMessageEvents) {
		return nil, fmt.Errorf("mismatch. %d sent messages & %d sent message extensions", len(sentMessageEvents), len(sentMessageExtensionEvents))
	}

	numVersionZeroMessages := len(sentMessageEvents) - len(sentMessageExtensionEvents)
	crossDomainSentMessages := make([]CrossDomainMessengerSentMessageEvent, len(sentMessageEvents))
	for i := range sentMessageEvents {
		sentMessage := bindings.CrossDomainMessengerSentMessage{Raw: *sentMessageEvents[i].RLPLog}
		err = UnpackLog(&sentMessage, sentMessageEvents[i].RLPLog, sentMessageEventAbi.Name, crossDomainMessengerAbi)
		if err != nil {
			return nil, err
		}
		version, _ := DecodeVersionedNonce(sentMessage.MessageNonce)
		if i < numVersionZeroMessages && version != 0 {
			return nil, fmt.Errorf("expected version zero nonce. nonce %d tx_hash %s", sentMessage.MessageNonce, sentMessage.Raw.TxHash)
		}
		mntValue := big.NewInt(0)
		ethValue := big.NewInt(0)
		hashVersion := 0
		if version > 0 {
			sentMessageExtension := bindings.CrossDomainMessengerSentMessageExtension1{Raw: *sentMessageExtensionEvents[i].RLPLog}
			err = UnpackLog(&sentMessageExtension, sentMessageExtensionEvents[i].RLPLog, sentMessageExtensionEventAbi.Name, crossDomainMessengerAbi)
			if err != nil {
				return nil, err
			}
			mntValue = sentMessageExtension.MntValue
			ethValue = sentMessageExtension.EthValue
			hashVersion = 1
		}
		messageCalldata, err := CrossDomainMessageCalldata(crossDomainMessengerAbi, &sentMessage, mntValue, ethValue)
		if err != nil {
			return nil, err
		}
		crossDomainSentMessages[i] = CrossDomainMessengerSentMessageEvent{
			Event:                &sentMessageEvents[i],
			MessageCalldata:      messageCalldata,
			MessageHash:          crypto.Keccak256Hash(messageCalldata),
			Nonce:                sentMessage.MessageNonce,
			SentMessageEventGUID: sentMessageEvents[i].GUID,
			GasLimit:             sentMessage.GasLimit,
			FromAddress:          sentMessage.Sender,
			ToAddress:            sentMessage.Target,
			ETHAmount:            ethValue,
			Data:                 sentMessage.Message,
			Timestamp:            sentMessageEvents[i].Timestamp,
			HashVersion:          uint64(hashVersion),
			ERC20Amount:          mntValue,
		}
	}
	return crossDomainSentMessages, nil
}

func TransferMessageHashFromV0ToV1(db *database.DB) error {
	crossDomainMessengerAbi, err := bindings.CrossDomainMessengerMetaData.GetAbi()
	if err != nil {
		return err
	}
	l2SendMsgList := db.L2SentMessageEvent.L2SentMessageList()
	for i := range l2SendMsgList {
		l2SendMsg := l2SendMsgList[i]
		l2l1Transaction, err := db.L2ToL1.L2ToL1TransactionTxHash(l2SendMsg.TxHash)
		if err != nil {
			log.Error("get l2 to l1 by transaction hash fail", "err", err)
			continue
		}
		mntValue := l2l1Transaction.ETHAmount
		ethValue := l2l1Transaction.ERC20Amount
		sentMessage := bindings.CrossDomainMessengerSentMessage{
			Target:       l2SendMsg.Target,
			Sender:       l2SendMsg.Sender,
			Message:      []byte(l2SendMsg.Message),
			MessageNonce: l2SendMsg.MessageNonce,
			GasLimit:     l2SendMsg.GasLimit,
		}
		messageCalldata, err := CrossDomainMessageCalldata(crossDomainMessengerAbi, &sentMessage, mntValue, ethValue)
		if err != nil {
			log.Error("create message call data fail", "err", err)
			continue
		}
		fmt.Println("message call data hash===", crypto.Keccak256Hash(messageCalldata))
	}
	return nil
}

func CrossDomainMessengerRelayedMessageEvents(chainId string, contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) ([]CrossDomainMessengerRelayedMessageEvent, error) {
	crossDomainMessengerAbi, err := bindings.CrossDomainMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	relayedMessageEventAbi := crossDomainMessengerAbi.Events["RelayedMessage"]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: relayedMessageEventAbi.ID}
	relayedMessageEvents, err := db.ContractEvents.ContractEventsWithFilter(chainId, contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	crossDomainRelayedMessages := make([]CrossDomainMessengerRelayedMessageEvent, len(relayedMessageEvents))
	for i := range relayedMessageEvents {
		relayedMessage := bindings.CrossDomainMessengerRelayedMessage{Raw: *relayedMessageEvents[i].RLPLog}
		err = UnpackLog(&relayedMessage, relayedMessageEvents[i].RLPLog, relayedMessageEventAbi.Name, crossDomainMessengerAbi)
		if err != nil {
			return nil, err
		}

		crossDomainRelayedMessages[i] = CrossDomainMessengerRelayedMessageEvent{
			Event:       &relayedMessageEvents[i],
			MessageHash: relayedMessage.MsgHash,
		}
	}

	return crossDomainRelayedMessages, nil
}

func CrossDomainMessageCalldataV0(sentMsg *bindings.CrossDomainMessengerSentMessage) ([]byte, error) {
	inputBytes, err := CrossDomainMessengerLegacyRelayMessageEncoding.Inputs.Pack(sentMsg.Target, sentMsg.Sender, sentMsg.Message, sentMsg.MessageNonce)
	if err != nil {
		return nil, err
	}
	return append(CrossDomainMessengerLegacyRelayMessageEncoding.ID, inputBytes...), nil
}

func CrossDomainMessageCalldata(abi *abi.ABI, sentMsg *bindings.CrossDomainMessengerSentMessage, mntValue *big.Int, ethValue *big.Int) ([]byte, error) {
	version, _ := DecodeVersionedNonce(sentMsg.MessageNonce)
	switch version {
	case 0:
		inputBytes, err := CrossDomainMessengerLegacyRelayMessageEncoding.Inputs.Pack(sentMsg.Target, sentMsg.Sender, sentMsg.Message, sentMsg.MessageNonce)
		if err != nil {
			return nil, err
		}
		return append(CrossDomainMessengerLegacyRelayMessageEncoding.ID, inputBytes...), nil
	case 1:
		msgBytes, err := abi.Pack("relayMessage", sentMsg.MessageNonce, sentMsg.Sender, sentMsg.Target, mntValue, ethValue, sentMsg.GasLimit, sentMsg.Message)
		if err != nil {
			return nil, err
		}
		return msgBytes, nil
	}

	return nil, fmt.Errorf("unsupported cross domain messenger version: %d", version)
}
