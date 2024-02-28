package bridge

import (
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"

	common2 "github.com/cornerstone-labs/acorus/common"
	"github.com/cornerstone-labs/acorus/common/bigint"
	common4 "github.com/cornerstone-labs/acorus/database/common"

	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	common3 "github.com/cornerstone-labs/acorus/event/op_stack/common"
	"github.com/cornerstone-labs/acorus/event/op_stack/contracts"
)

func L2ProcessInitiatedBridgeEvents(db *database.DB, fromHeight, toHeight *big.Int) error {
	// (1) L2ToL1MessagePasser
	l2ToL1MPMessagesPassed, err := contracts.L2ToL1MessagePasserMessagePassedEvents(common3.L2ToL1MessagePasser, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(l2ToL1MPMessagesPassed) > 0 {
		log.Println("detected transaction withdrawals", "size", len(l2ToL1MPMessagesPassed))
	}
	messagesPassed := make(map[logKey]*contracts.L2ToL1MessagePasserMessagePassed, len(l2ToL1MPMessagesPassed))
	l2ToL1s := make([]worker.L2ToL1, len(l2ToL1MPMessagesPassed))
	for i := range l2ToL1MPMessagesPassed {
		messagePassed := l2ToL1MPMessagesPassed[i]
		messagesPassed[logKey{messagePassed.Event.BlockHash, messagePassed.Event.LogIndex}] = &messagePassed
		blockNumber, err := db.L2ToL1.GetBlockNumberFromHash("10", messagePassed.Event.BlockHash)
		if err != nil {
			log.Println("can not get l2 blockNumber", "blockHash", messagePassed.Event.BlockHash)
			return err
		}

		l2ToL1s[i] = worker.L2ToL1{

			L1BlockNumber:           bigint.Zero,
			L2BlockNumber:           blockNumber,
			MsgNonce:                messagePassed.Nonce,
			L2TransactionHash:       messagePassed.Event.TransactionHash,
			WithdrawTransactionHash: messagePassed.WithdrawalHash,
			L1ProveTxHash:           common.Hash{},
			L1FinalizeTxHash:        common.Hash{},
			Status:                  common2.L2ToL1Pending,
			ETHAmount:               messagePassed.ETHAmount,
			TokenAmounts:            messagePassed.ERC20Amount.String(),
			GasLimit:                messagePassed.GasLimit,
			TimeLeft:                new(big.Int).SetUint64(0),
			L1TokenAddress:          common.Address{},
			L2TokenAddress:          common.Address{},
			Timestamp:               int64(messagePassed.Event.Timestamp),
		}
	}
	if len(messagesPassed) > 0 {
		if err := db.L2ToL1.StoreL2ToL1Transactions("10", l2ToL1s); err != nil {
			return err
		}
	}

	// (2) L2CrossDomainMessenger
	crossDomainSentMessages, err := contracts.CrossDomainMessengerSentMessageEvents(common3.L2CrossDomainMessenger, "10", db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(crossDomainSentMessages) > 0 {
		log.Println("detected sent messages", "size", len(crossDomainSentMessages))
	}

	sentMessages := make(map[logKey]*contracts.CrossDomainMessengerSentMessageEvent, len(crossDomainSentMessages))
	l2ToL1Cds := make([]worker.L2ToL1, len(crossDomainSentMessages))
	for i := range crossDomainSentMessages {
		sentMessage := crossDomainSentMessages[i]
		sentMessages[logKey{sentMessage.Event.BlockHash, sentMessage.Event.LogIndex}] = &sentMessage
		l2ToL1Cds[i].L2TransactionHash = sentMessage.Event.TransactionHash
		l2ToL1Cds[i].MessageHash = sentMessage.MessageHash
	}
	if len(crossDomainSentMessages) > 0 {
		if err := db.L2ToL1.UpdateMessageHash("10", l2ToL1Cds); err != nil {
			return err
		}
	}

	// (3) L2StandardBridge
	initiatedBridges, err := contracts.StandardBridgeInitiatedEvents("10", common3.L2StandardBridge, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(initiatedBridges) > 0 {
		log.Println("detected bridge withdrawals", "size", len(initiatedBridges))
	}

	bridgedTokens := make(map[common.Address]int)
	l2ToL1s2 := make([]worker.L2ToL1, len(initiatedBridges))
	for i := range initiatedBridges {
		initiatedBridge := initiatedBridges[i]
		// extract the cross domain message hash & deposit source hash from the following events
		if initiatedBridge.Event.EventSignature.String() == common3.ETHWithdrawEventSignature {
			messagePassed, ok := messagesPassed[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 1}]
			if !ok {
				log.Println("expected MessagePassed following BridgeInitiated event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected MessagePassed following BridgeInitiated event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}
			sentMessage, ok := sentMessages[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 2}]
			if !ok {
				log.Println("expected SentMessage following MessagePassed event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected SentMessage following MessagePassed event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}

			initiatedBridge.CrossDomainMessageHash = &sentMessage.MessageHash
			bridgedTokens[initiatedBridge.LocalTokenAddress]++
			l2ToL1s2[i].WithdrawTransactionHash = messagePassed.WithdrawalHash
			l2ToL1s2[i].L2TransactionHash = initiatedBridge.Event.TransactionHash
			l2ToL1s2[i].FromAddress = initiatedBridge.FromAddress
			l2ToL1s2[i].ToAddress = initiatedBridge.ToAddress
			l2ToL1s2[i].ETHAmount = initiatedBridge.ETHAmount
			l2ToL1s2[i].TokenAmounts = initiatedBridge.ERC20Amount.String()
			l2ToL1s2[i].L1TokenAddress = initiatedBridge.LocalTokenAddress
			l2ToL1s2[i].L2TokenAddress = initiatedBridge.RemoteTokenAddress
			l2ToL1s2[i].AssetType = common4.ETH
		} else if initiatedBridge.Event.EventSignature.String() == common3.ERC20WithdrawEventSignature {
			messagePassed, ok := messagesPassed[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 1}]
			if !ok {
				log.Println("expected MessagePassed following BridgeInitiated event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected MessagePassed following BridgeInitiated event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}
			sentMessage, ok := sentMessages[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 2}]
			if !ok {
				log.Println("expected SentMessage following MessagePassed event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected SentMessage following MessagePassed event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}

			initiatedBridge.CrossDomainMessageHash = &sentMessage.MessageHash
			bridgedTokens[initiatedBridge.LocalTokenAddress]++
			l2ToL1s2[i].WithdrawTransactionHash = messagePassed.WithdrawalHash
			l2ToL1s2[i].L2TransactionHash = initiatedBridge.Event.TransactionHash
			l2ToL1s2[i].FromAddress = initiatedBridge.FromAddress
			l2ToL1s2[i].ToAddress = initiatedBridge.ToAddress
			l2ToL1s2[i].ETHAmount = initiatedBridge.ETHAmount
			l2ToL1s2[i].TokenAmounts = initiatedBridge.ERC20Amount.String()
			l2ToL1s2[i].L1TokenAddress = initiatedBridge.LocalTokenAddress
			l2ToL1s2[i].L2TokenAddress = initiatedBridge.RemoteTokenAddress
			l2ToL1s2[i].AssetType = common4.ERC20

		}
	}

	if len(l2ToL1s2) > 0 {
		if err := db.L2ToL1.UpdateTokenPairsAndAddress("10", l2ToL1s2); err != nil {
			return err
		}
	}
	return nil
}

func L2ProcessFinalizedBridgeEvents(db *database.DB, fromHeight, toHeight *big.Int) error {
	log.Println("L2ProcessFinalizedBridgeEvents", "fromHeight", fromHeight, "toHeight", toHeight)
	// (1) L2CrossDomainMessenger
	crossDomainRelayedMessages, err := contracts.CrossDomainMessengerRelayedMessageEvents("10", common3.L2CrossDomainMessenger, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(crossDomainRelayedMessages) > 0 {
		log.Println("detected relayed messages", "size", len(crossDomainRelayedMessages))
	}
	relayedMessages := make(map[logKey]*contracts.CrossDomainMessengerRelayedMessageEvent, len(crossDomainRelayedMessages))
	for i := range crossDomainRelayedMessages {
		relayed := crossDomainRelayedMessages[i]
		relayedMessages[logKey{BlockHash: relayed.Event.BlockHash, LogIndex: relayed.Event.LogIndex}] = &relayed
		message, err := db.L1ToL2.L1ToL2Transaction(strconv.FormatUint(global_const.OpChinId, 10), relayed.MessageHash)
		if err != nil {
			return err
		} else if message == nil {
			log.Println("missing indexed L1CrossDomainMessenger message", "tx_hash", relayed.Event.TransactionHash.String())
			continue
		}
	}
	if len(relayedMessages) > 0 {
		log.Println("relayedMessages list length", "relayedMessages", len(relayedMessages))
	}

	// (2) L2StandardBridge
	finalizedBridges, err := contracts.StandardBridgeFinalizedEvents("10", common3.L2StandardBridge, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(finalizedBridges) > 0 {
		log.Println("detected finalized bridge deposits", "size", len(finalizedBridges))
	}
	relayMessageList := make([]event.RelayMessage, len(finalizedBridges))

	finalizedTokens := make(map[common.Address]int)
	for i := range finalizedBridges {
		finalizedBridge := finalizedBridges[i]
		relayedMessage, ok := relayedMessages[logKey{finalizedBridge.Event.BlockHash, finalizedBridge.Event.LogIndex + 1}]
		if !ok {
			log.Println("expected RelayedMessage following BridgeFinalized event", "tx_hash", finalizedBridge.Event.TransactionHash.String())
			return fmt.Errorf("expected RelayedMessage following BridgeFinalized event. tx_hash = %s", finalizedBridge.Event.TransactionHash.String())
		}
		blockNumber := bigint.One
		relayL2BlockNumber, err := db.L2ToL1.GetBlockNumberFromHash(strconv.FormatUint(global_const.OpChinId, 10), relayedMessage.Event.BlockHash)
		if err != nil {
			return err
		}
		if relayL2BlockNumber != nil {
			blockNumber = relayL2BlockNumber
		}
		relayMessageList[i] = event.RelayMessage{
			BlockNumber:          blockNumber,
			RelayTransactionHash: relayedMessage.Event.TransactionHash,
			MessageHash:          relayedMessage.MessageHash,
			L1TokenAddress:       finalizedBridge.RemoteTokenAddress,
			L2TokenAddress:       finalizedBridge.LocalTokenAddress,
			ETHAmount:            finalizedBridge.ETHAmount,
			ERC20Amount:          finalizedBridge.ERC20Amount,
			Related:              false,
			Timestamp:            relayedMessage.Event.Timestamp,
		}
		finalizedTokens[finalizedBridge.LocalTokenAddress]++
	}
	if len(relayMessageList) > 0 {
		err = db.RelayMessage.StoreRelayMessage(strconv.FormatUint(global_const.OpChinId, 10), relayMessageList)
		if err != nil {
			return err
		}
		for tokenAddr, size := range finalizedTokens {
			log.Println("tokenAddr and size", "tokenAddr", tokenAddr, "size", size)
		}
	}
	return nil
}
