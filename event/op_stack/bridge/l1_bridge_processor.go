package bridge

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"

	common2 "github.com/cornerstone-labs/acorus/common"
	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	common3 "github.com/cornerstone-labs/acorus/event/op_stack/common"
	"github.com/cornerstone-labs/acorus/event/op_stack/contracts"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
)

func L1ProcessInitiatedBridgeEvents(db *database.DB, fromHeight, toHeight *big.Int) error {
	// (1) OptimismPortal
	optimismPortalTxDeposits, err := contracts.OptimismPortalTransactionDepositEvents(common3.OptimismPortalProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(optimismPortalTxDeposits) > 0 {
		log.Info("detected transaction deposits", "size", len(optimismPortalTxDeposits))
	}

	portalDeposits := make(map[logKey]*contracts.OptimismPortalTransactionDepositEvent, len(optimismPortalTxDeposits))
	l1ToL2s := make([]worker.L1ToL2, len(optimismPortalTxDeposits))
	for i := range optimismPortalTxDeposits {
		depositTx := optimismPortalTxDeposits[i]
		portalDeposits[logKey{depositTx.Event.BlockHash, depositTx.Event.LogIndex}] = &depositTx
		blockNumber, err := db.L1ToL2.GetBlockNumberFromHash("10", depositTx.Event.BlockHash)
		if err != nil {
			log.Error("can not get l1 blockNumber", "blockHash", depositTx.Event.BlockHash)
			return err
		}
		l1ToL2s[i] = worker.L1ToL2{
			GUID:                  uuid.New(),
			L1BlockNumber:         blockNumber,
			L2BlockNumber:         bigint.Zero,
			QueueIndex:            nil,
			L1TransactionHash:     depositTx.Event.TransactionHash,
			L2TransactionHash:     common.Hash{},
			TransactionSourceHash: depositTx.DepositTx.SourceHash,
			MessageHash:           common.Hash{},
			L1TxOrigin:            depositTx.DepositTx.From,
			Status:                common2.L1ToL2Pending,
			L1TokenAddress:        common.Address{},
			L2TokenAddress:        common.Address{},
			ETHAmount:             depositTx.DepositTx.Value,
			GasLimit:              depositTx.GasLimit,
			Timestamp:             int64(depositTx.Event.Timestamp),
		}
	}
	if len(l1ToL2s) > 0 {
		if err := db.L1ToL2.StoreL1ToL2Transactions("10", l1ToL2s); err != nil {
			return err
		}
	}

	// (2) L1CrossDomainMessenger
	crossDomainSentMessages, err := contracts.CrossDomainMessengerSentMessageEvents(common3.L1CrossDomainMessengerProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(crossDomainSentMessages) > 0 {
		log.Info("detected sent messages", "size", len(crossDomainSentMessages))
	}

	sentMessages := make(map[logKey]*contracts.CrossDomainMessengerSentMessageEvent, len(crossDomainSentMessages))
	l1ToL2c2 := make([]worker.L1ToL2, len(crossDomainSentMessages))
	for i := range crossDomainSentMessages {
		sentMessage := crossDomainSentMessages[i]
		sentMessages[logKey{sentMessage.Event.BlockHash, sentMessage.Event.LogIndex}] = &sentMessage
		l1ToL2c2[i].L1TransactionHash = sentMessage.Event.TransactionHash
		l1ToL2c2[i].MessageHash = sentMessage.MessageHash
	}
	if len(l1ToL2c2) > 0 {
		if err := db.L1ToL2.UpdateMessageHash("10", l1ToL2c2); err != nil {
			return err
		}
	}

	// (3) L1StandardBridge
	initiatedBridges, err := contracts.StandardBridgeInitiatedEvents("l1", common3.L1StandardBridgeProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(initiatedBridges) > 0 {
		log.Info("detected bridge deposits", "size", len(initiatedBridges))
	}

	bridgedTokens := make(map[common.Address]int)
	l1ToL2s2 := make([]worker.L1ToL2, len(initiatedBridges))
	for i := range initiatedBridges {
		initiatedBridge := initiatedBridges[i]
		standardBridgeAbi, _ := bindings.StandardBridgeMetaData.GetAbi()
		if initiatedBridge.Event.EventSignature == standardBridgeAbi.Events["ETHBridgeInitiated"].ID {
			portalDeposit, ok := portalDeposits[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 1}]
			if !ok {
				log.Error("expected TransactionDeposit following BridgeInitiated event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected TransactionDeposit following BridgeInitiated event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}
			sentMessage, ok := sentMessages[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 2}]
			if !ok {
				log.Error("expected SentMessage following TransactionDeposit event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected SentMessage following TransactionDeposit event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}
			initiatedBridge.CrossDomainMessageHash = &sentMessage.MessageHash
			bridgedTokens[initiatedBridge.LocalTokenAddress]++

			l1ToL2s2[i].L1TransactionHash = portalDeposit.Event.TransactionHash
			l1ToL2s2[i].FromAddress = initiatedBridge.FromAddress
			l1ToL2s2[i].ToAddress = initiatedBridge.ToAddress
			l1ToL2s2[i].L1TokenAddress = initiatedBridge.LocalTokenAddress
			l1ToL2s2[i].L2TokenAddress = initiatedBridge.RemoteTokenAddress
		} else if initiatedBridge.Event.EventSignature == standardBridgeAbi.Events["MNTBridgeInitiated"].ID {
			portalDeposit, ok := portalDeposits[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 6}]
			if !ok {
				log.Error("expected TransactionDeposit following BridgeInitiated event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected TransactionDeposit following BridgeInitiated event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}
			sentMessage, ok := sentMessages[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 7}]
			if !ok {
				log.Error("expected SentMessage following TransactionDeposit event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected SentMessage following TransactionDeposit event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}
			initiatedBridge.CrossDomainMessageHash = &sentMessage.MessageHash
			bridgedTokens[initiatedBridge.LocalTokenAddress]++
			l1ToL2s2[i].L1TransactionHash = portalDeposit.Event.TransactionHash
			l1ToL2s2[i].FromAddress = initiatedBridge.FromAddress
			l1ToL2s2[i].ToAddress = initiatedBridge.ToAddress
			l1ToL2s2[i].L1TokenAddress = initiatedBridge.LocalTokenAddress
			l1ToL2s2[i].L2TokenAddress = initiatedBridge.RemoteTokenAddress
		} else if initiatedBridge.Event.EventSignature == standardBridgeAbi.Events["ERC20BridgeInitiated"].ID {
			portalDeposit, ok := portalDeposits[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 1}]
			if !ok {
				log.Error("expected TransactionDeposit following BridgeInitiated event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected TransactionDeposit following BridgeInitiated event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}
			sentMessage, ok := sentMessages[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 2}]
			if !ok {
				log.Error("expected SentMessage following TransactionDeposit event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected SentMessage following TransactionDeposit event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}
			initiatedBridge.CrossDomainMessageHash = &sentMessage.MessageHash
			bridgedTokens[initiatedBridge.LocalTokenAddress]++
			l1ToL2s2[i].L1TransactionHash = portalDeposit.Event.TransactionHash
			l1ToL2s2[i].FromAddress = initiatedBridge.FromAddress
			l1ToL2s2[i].ToAddress = initiatedBridge.ToAddress
			l1ToL2s2[i].L1TokenAddress = initiatedBridge.LocalTokenAddress
			l1ToL2s2[i].L2TokenAddress = initiatedBridge.RemoteTokenAddress
		}
	}
	if len(l1ToL2s2) > 0 {
		if err := db.L1ToL2.UpdateTokenPairAndAddress("10", l1ToL2s2); err != nil {
			return err
		}
	}
	return nil
}

// L1ProcessProvenBridgeEvents Optimism portal proven withdrawals
func L1ProcessProvenBridgeEvents(db *database.DB, fromHeight, toHeight *big.Int) error {
	provenWithdrawals, err := contracts.OptimismPortalWithdrawalProvenEvents(common3.OptimismPortalProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	log.Info("detected proven withdrawals", "size", len(provenWithdrawals))
	withdrawProvenList := make([]event.WithdrawProven, len(provenWithdrawals))
	for i := range provenWithdrawals {
		blockNumber := bigint.One
		proven := provenWithdrawals[i]
		proveL1BlockNumber, err := db.L1ToL2.GetBlockNumberFromHash("10", proven.Event.BlockHash)
		if err != nil {
			return err
		}
		if proveL1BlockNumber != nil {
			blockNumber = proveL1BlockNumber
		}
		withdrawProvenList[i] = event.WithdrawProven{
			GUID:                  uuid.New(),
			WithdrawHash:          proven.WithdrawalHash,
			ProvenTransactionHash: proven.Event.TransactionHash,
			Related:               false,
			BlockNumber:           blockNumber,
			MessageHash:           common.Hash{},
			L1TokenAddress:        common.Address{},
			L2TokenAddress:        common.Address{},
			ETHAmount:             bigint.Zero,
			ERC20Amount:           bigint.Zero,
			Timestamp:             proven.Event.Timestamp,
		}
	}
	if len(withdrawProvenList) > 0 {
		if err := db.WithdrawProven.StoreWithdrawProven(withdrawProvenList); err != nil {
			return err
		}
	}
	return nil
}

// L1ProcessFinalizedBridgeEvents OptimismPortal (finalized withdrawals) and L1CrossDomainMessenger
func L1ProcessFinalizedBridgeEvents(db *database.DB, fromHeight, toHeight *big.Int) error {
	finalizedWithdrawals, err := contracts.OptimismPortalWithdrawalFinalizedEvents(common3.OptimismPortalProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	log.Info("detected finalized withdrawals", "size", len(finalizedWithdrawals))
	withdrawFinalizedList := make([]event.WithdrawFinalized, len(finalizedWithdrawals))
	for i := range finalizedWithdrawals {
		finalized := finalizedWithdrawals[i]
		blockNumber := bigint.One
		finalizedBlockNumber, err := db.L1ToL2.GetBlockNumberFromHash("10", finalized.Event.BlockHash)
		if err != nil {
			return err
		}
		if finalizedBlockNumber != nil {
			blockNumber = finalizedBlockNumber
		}
		withdrawFinalizedList[i] = event.WithdrawFinalized{
			GUID:                     uuid.New(),
			WithdrawHash:             finalized.WithdrawalHash,
			FinalizedTransactionHash: finalized.Event.TransactionHash,
			Related:                  false,
			BlockNumber:              blockNumber,
			MessageHash:              common.Hash{},
			L1TokenAddress:           common.Address{},
			L2TokenAddress:           common.Address{},
			ETHAmount:                bigint.Zero,
			ERC20Amount:              bigint.Zero,
			Timestamp:                finalized.Event.Timestamp,
		}
	}
	if len(withdrawFinalizedList) > 0 {
		if err := db.WithdrawFinalized.StoreWithdrawFinalized(withdrawFinalizedList); err != nil {
			return err
		}
	}

	crossDomainRelayedMessages, err := contracts.CrossDomainMessengerRelayedMessageEvents("l1", common3.L1CrossDomainMessengerProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	log.Info("detected relayed messages", "size", len(crossDomainRelayedMessages))
	relayedMessages := make(map[logKey]*contracts.CrossDomainMessengerRelayedMessageEvent, len(crossDomainRelayedMessages))

	for i := range crossDomainRelayedMessages {
		relayed := crossDomainRelayedMessages[i]
		relayedMessages[logKey{BlockHash: relayed.Event.BlockHash, LogIndex: relayed.Event.LogIndex}] = &relayed
	}

	//  L1StandardBridge
	finalizedBridges, err := contracts.StandardBridgeFinalizedEvents("l1", common3.L1StandardBridgeProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	log.Info("detected finalized bridge withdrawals", "size", len(finalizedBridges))
	withdrawFinalizedBridgeList := make([]event.WithdrawFinalized, len(crossDomainRelayedMessages))
	finalizedTokens := make(map[common.Address]int)
	for i := range finalizedBridges {
		finalizedBridge := finalizedBridges[i]
		relayedMessage, ok := relayedMessages[logKey{finalizedBridge.Event.BlockHash, finalizedBridge.Event.LogIndex + 1}]
		if !ok {
			log.Error("expected RelayedMessage following BridgeFinalized event", "tx_hash", finalizedBridge.Event.TransactionHash.String())
			return fmt.Errorf("expected RelayedMessage following BridgeFinalized event. tx_hash = %s", finalizedBridge.Event.TransactionHash.String())
		}
		withdrawFinalizedBridgeList[i].L1TokenAddress = finalizedBridge.LocalTokenAddress
		withdrawFinalizedBridgeList[i].L2TokenAddress = finalizedBridge.RemoteTokenAddress
		withdrawFinalizedBridgeList[i].ETHAmount = finalizedBridge.ETHAmount
		withdrawFinalizedBridgeList[i].ERC20Amount = finalizedBridge.ERC20Amount
		withdrawFinalizedBridgeList[i].MessageHash = relayedMessage.MessageHash
		finalizedTokens[finalizedBridge.LocalTokenAddress]++
	}
	if len(finalizedBridges) > 0 {
		if err := db.WithdrawFinalized.UpdateWithdrawFinalizedInfo(withdrawFinalizedBridgeList); err != nil {
			return err
		}
	}
	return nil
}
