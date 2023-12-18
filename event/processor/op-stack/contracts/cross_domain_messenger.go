package contracts

import (
	"fmt"
	"github.com/google/uuid"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/utils"
	"github.com/cornerstone-labs/acorus/event/processor/op-stack/op-bindings/bindings"
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
			// The actual transaction on the legacy L1CrossDomainMessenger has a trailing
			// proof argument but is ignored for the `XDomainCallData` encoding
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
	Timestamp            uint64
}

type CrossDomainMessengerRelayedMessageEvent struct {
	Event       *event.ContractEvent
	MessageHash common.Hash
}

func CrossDomainMessengerSentMessageEvents(chainSelector string, contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) ([]CrossDomainMessengerSentMessageEvent, error) {
	crossDomainMessengerAbi, err := bindings.CrossDomainMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	sentMessageEventAbi := crossDomainMessengerAbi.Events["SentMessage"]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: sentMessageEventAbi.ID}
	sentMessageEvents, err := db.ContractEvents.ContractEventsWithFilter(contractEventFilter, chainSelector, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}
	if len(sentMessageEvents) == 0 {
		return nil, nil
	}

	sentMessageExtensionEventAbi := crossDomainMessengerAbi.Events["SentMessageExtension1"]
	contractEventFilter = event.ContractEvent{ContractAddress: contractAddress, EventSignature: sentMessageExtensionEventAbi.ID}
	sentMessageExtensionEvents, err := db.ContractEvents.ContractEventsWithFilter(contractEventFilter, chainSelector, fromHeight, toHeight)
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
		if version > 0 {
			sentMessageExtension := bindings.CrossDomainMessengerSentMessageExtension1{Raw: *sentMessageExtensionEvents[i].RLPLog}
			err = UnpackLog(&sentMessageExtension, sentMessageExtensionEvents[i].RLPLog, sentMessageExtensionEventAbi.Name, crossDomainMessengerAbi)
			if err != nil {
				return nil, err
			}
			mntValue = sentMessageExtension.MntValue
			ethValue = sentMessageExtension.EthValue
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
			ERC20Amount:          mntValue,
		}
	}

	return crossDomainSentMessages, nil
}

func CrossDomainMessengerRelayedMessageEvents(chainSelector string, contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) ([]CrossDomainMessengerRelayedMessageEvent, error) {
	crossDomainMessengerAbi, err := bindings.CrossDomainMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	relayedMessageEventAbi := crossDomainMessengerAbi.Events["RelayedMessage"]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: relayedMessageEventAbi.ID}
	relayedMessageEvents, err := db.ContractEvents.ContractEventsWithFilter(contractEventFilter, chainSelector, fromHeight, toHeight)
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
