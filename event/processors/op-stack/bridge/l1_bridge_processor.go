package bridge

import (
	"fmt"
	"github.com/cornerstone-labs/acorus/config/op-stack"
	contracts2 "github.com/cornerstone-labs/acorus/event/processors/op-stack/contracts"
	"github.com/cornerstone-labs/acorus/event/processors/op-stack/mantle/op-bindings/bindings"
	"github.com/google/uuid"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event/l1-l2"
	"github.com/cornerstone-labs/acorus/database/worker"
)

// L1ProcessInitiatedBridgeEvents will query the database for bridge events that have been initiated between
// the specified block range. This covers every part of the multi-layered stack:
//  1. OptimismPortal
//  2. L1CrossDomainMessenger
//  3. L1StandardBridge
func L1ProcessInitiatedBridgeEvents(log log.Logger, db *database.DB, l1Contracts op_stack.L1Contracts, fromHeight, toHeight *big.Int) error {
	// (1) OptimismPortal
	optimismPortalTxDeposits, err := contracts2.OptimismPortalTransactionDepositEvents(l1Contracts.OptimismPortalProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(optimismPortalTxDeposits) > 0 {
		log.Info("detected transaction deposits", "size", len(optimismPortalTxDeposits))
	}

	portalDeposits := make(map[logKey]*contracts2.OptimismPortalTransactionDepositEvent, len(optimismPortalTxDeposits))
	transactionDeposits := make([]l1_l2.L1TransactionDeposit, len(optimismPortalTxDeposits))
	l1ToL2s := make([]worker.L1ToL2, len(optimismPortalTxDeposits))
	for i := range optimismPortalTxDeposits {
		depositTx := optimismPortalTxDeposits[i]
		portalDeposits[logKey{depositTx.Event.BlockHash, depositTx.Event.LogIndex}] = &depositTx
		deptx := &types.DepositTx{
			SourceHash:          depositTx.DepositTx.SourceHash,
			From:                depositTx.DepositTx.From,
			To:                  depositTx.DepositTx.To,
			Mint:                depositTx.DepositTx.Mint,
			Value:               depositTx.DepositTx.Value,
			Gas:                 depositTx.DepositTx.Gas,
			IsSystemTransaction: depositTx.DepositTx.IsSystemTransaction,
			Data:                depositTx.DepositTx.Data,
		}
		transactionDeposits[i] = l1_l2.L1TransactionDeposit{
			SourceHash:           depositTx.DepositTx.SourceHash,
			L2TransactionHash:    types.NewTx(deptx).Hash(),
			InitiatedL1EventGUID: depositTx.Event.GUID,
			GasLimit:             depositTx.GasLimit,
			Tx:                   depositTx.Tx,
		}
		l1ToL2s[i] = worker.L1ToL2{
			GUID:              uuid.New(),
			L1BlockHash:       depositTx.Event.BlockHash,
			QueueIndex:        nil,
			L1TransactionHash: depositTx.DepositTx.SourceHash,
			L2TransactionHash: transactionDeposits[i].L2TransactionHash,
			L1TxOrigin:        common.Hash{},
			Status:            0,
			FromAddress:       depositTx.DepositTx.From,
			ToAddress:         *depositTx.DepositTx.To,
			L1TokenAddress:    common.Address{},
			L2TokenAddress:    common.Address{},
			ETHAmount:         depositTx.DepositTx.EthValue,
			ERC20Amount:       depositTx.DepositTx.Value,
			GasLimit:          depositTx.GasLimit,
			Timestamp:         int64(depositTx.Event.Timestamp),
		}
	}
	if len(transactionDeposits) > 0 {
		if err := db.BridgeTransactions.StoreL1TransactionDeposits(transactionDeposits); err != nil {
			return err
		}
		if err := db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s); err != nil {
			return err
		}
	}

	// (2) L1CrossDomainMessenger
	crossDomainSentMessages, err := contracts2.CrossDomainMessengerSentMessageEvents("l1", l1Contracts.L1CrossDomainMessengerProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(crossDomainSentMessages) > 0 {
		log.Info("detected sent messages", "size", len(crossDomainSentMessages))
	}

	sentMessages := make(map[logKey]*contracts2.CrossDomainMessengerSentMessageEvent, len(crossDomainSentMessages))
	bridgeMessages := make([]l1_l2.L1BridgeMessage, len(crossDomainSentMessages))
	for i := range crossDomainSentMessages {
		sentMessage := crossDomainSentMessages[i]
		sentMessages[logKey{sentMessage.Event.BlockHash, sentMessage.Event.LogIndex}] = &sentMessage
		// extract the deposit hash from the previous TransactionDepositedEvent
		portalDeposit, ok := portalDeposits[logKey{sentMessage.Event.BlockHash, sentMessage.Event.LogIndex - 1}]
		if !ok {
			log.Error("expected TransactionDeposit preceding SentMessage event", "tx_hash", sentMessage.Event.TransactionHash.String())
			return fmt.Errorf("expected TransactionDeposit preceding SentMessage event. tx_hash = %s", sentMessage.Event.TransactionHash.String())
		}

		bridgeMessages[i] = l1_l2.L1BridgeMessage{TransactionSourceHash: portalDeposit.DepositTx.SourceHash, BridgeMessage: sentMessage.BridgeMessage}
	}
	if len(bridgeMessages) > 0 {
		if err := db.BridgeMessages.StoreL1BridgeMessages(bridgeMessages); err != nil {
			return err
		}
	}

	// (3) L1StandardBridge
	initiatedBridges, err := contracts2.StandardBridgeInitiatedEvents("l1", l1Contracts.L1StandardBridgeProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(initiatedBridges) > 0 {
		log.Info("detected bridge deposits", "size", len(initiatedBridges))
	}

	bridgedTokens := make(map[common.Address]int)
	bridgeDeposits := make([]l1_l2.L1BridgeDeposit, len(initiatedBridges))
	l1ToL2s2 := make([]worker.L1ToL2, len(initiatedBridges))
	for i := range initiatedBridges {
		initiatedBridge := initiatedBridges[i]
		// extract the cross domain message hash & deposit source hash from the following events
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
			initiatedBridge.BridgeTransfer.CrossDomainMessageHash = &sentMessage.BridgeMessage.MessageHash
			bridgedTokens[initiatedBridge.BridgeTransfer.TokenPair.LocalTokenAddress]++
			bridgeDeposits[i] = l1_l2.L1BridgeDeposit{
				TransactionSourceHash: portalDeposit.DepositTx.SourceHash,
				BridgeTransfer:        initiatedBridge.BridgeTransfer,
			}
			l1ToL2s2[i].L1TransactionHash = portalDeposit.DepositTx.SourceHash
			l1ToL2s2[i].L1TokenAddress = initiatedBridge.BridgeTransfer.TokenPair.LocalTokenAddress
			l1ToL2s2[i].L2TokenAddress = initiatedBridge.BridgeTransfer.TokenPair.RemoteTokenAddress
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
			initiatedBridge.BridgeTransfer.CrossDomainMessageHash = &sentMessage.BridgeMessage.MessageHash
			bridgedTokens[initiatedBridge.BridgeTransfer.TokenPair.LocalTokenAddress]++
			bridgeDeposits[i] = l1_l2.L1BridgeDeposit{
				TransactionSourceHash: portalDeposit.DepositTx.SourceHash,
				BridgeTransfer:        initiatedBridge.BridgeTransfer,
			}
			l1ToL2s2[i].L1TransactionHash = portalDeposit.DepositTx.SourceHash
			l1ToL2s2[i].L1TokenAddress = initiatedBridge.BridgeTransfer.TokenPair.LocalTokenAddress
			l1ToL2s2[i].L2TokenAddress = initiatedBridge.BridgeTransfer.TokenPair.RemoteTokenAddress
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
			initiatedBridge.BridgeTransfer.CrossDomainMessageHash = &sentMessage.BridgeMessage.MessageHash
			bridgedTokens[initiatedBridge.BridgeTransfer.TokenPair.LocalTokenAddress]++
			bridgeDeposits[i] = l1_l2.L1BridgeDeposit{
				TransactionSourceHash: portalDeposit.DepositTx.SourceHash,
				BridgeTransfer:        initiatedBridge.BridgeTransfer,
			}
			l1ToL2s2[i].L1TransactionHash = portalDeposit.DepositTx.SourceHash
			l1ToL2s2[i].L1TokenAddress = initiatedBridge.BridgeTransfer.TokenPair.LocalTokenAddress
			l1ToL2s2[i].L2TokenAddress = initiatedBridge.BridgeTransfer.TokenPair.RemoteTokenAddress
		}

	}
	if len(bridgeDeposits) > 0 {
		if err := db.BridgeTransfers.StoreL1BridgeDeposits(bridgeDeposits); err != nil {
			return err
		}
		if err := db.L1ToL2.UpdateTokenPairs(l1ToL2s2); err != nil {
			return err
		}
	}

	return nil
}

// L1ProcessFinalizedBridgeEvent will query the database for all the finalization markers for all initiated
// bridge events. This covers every part of the multi-layered stack:
//  1. OptimismPortal (Bedrock prove & finalize steps)
//  2. L1CrossDomainMessenger (relayMessage marker)
//  3. L1StandardBridge (no-op, since this is simply a wrapper over the L1CrossDomainMessenger)
func L1ProcessFinalizedBridgeEvents(log log.Logger, db *database.DB, l1Contracts op_stack.L1Contracts, fromHeight, toHeight *big.Int) error {
	// (1) OptimismPortal (proven withdrawals)
	provenWithdrawals, err := contracts2.OptimismPortalWithdrawalProvenEvents(l1Contracts.OptimismPortalProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(provenWithdrawals) > 0 {
		log.Info("detected proven withdrawals", "size", len(provenWithdrawals))
	}

	for i := range provenWithdrawals {
		proven := provenWithdrawals[i]
		withdrawal, err := db.BridgeTransactions.L2TransactionWithdrawal(proven.WithdrawalHash)
		if err != nil {
			return err
		} else if withdrawal == nil {
			log.Error("missing indexed withdrawal on proven event!", "tx_hash", proven.Event.TransactionHash.String())
			return fmt.Errorf("missing indexed withdrawal! tx_hash = %s", proven.Event.TransactionHash.String())
		}

		if err := db.BridgeTransactions.MarkL2TransactionWithdrawalProvenEvent(proven.WithdrawalHash, provenWithdrawals[i].Event.GUID); err != nil {
			log.Error("failed to mark withdrawal as proven", "err", err, "tx_hash", proven.Event.TransactionHash.String())
			return err
		}
		if err := db.L2ToL1.UpdateReadyForClaimStatus(proven.WithdrawalHash, provenWithdrawals[i].Event.GUID); err != nil {
			log.Error("failed to mark l2_to_l1 as proven", "err", err, "tx_hash", proven.Event.TransactionHash.String())
			return err
		}
		log.Info("update ready for claim status success")
	}

	// (2) OptimismPortal (finalized withdrawals)
	finalizedWithdrawals, err := contracts2.OptimismPortalWithdrawalFinalizedEvents(l1Contracts.OptimismPortalProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(finalizedWithdrawals) > 0 {
		log.Info("detected finalized withdrawals", "size", len(finalizedWithdrawals))
	}

	for i := range finalizedWithdrawals {
		finalized := finalizedWithdrawals[i]
		withdrawal, err := db.BridgeTransactions.L2TransactionWithdrawal(finalized.WithdrawalHash)
		if err != nil {
			return err
		} else if withdrawal == nil {
			log.Error("missing indexed withdrawal on finalization event!", "tx_hash", finalized.Event.TransactionHash.String())
			return fmt.Errorf("missing indexed withdrawal on finalization! tx_hash: %s", finalized.Event.TransactionHash.String())
		}

		if err = db.BridgeTransactions.MarkL2TransactionWithdrawalFinalizedEvent(finalized.WithdrawalHash, finalized.Event.GUID, finalized.Success); err != nil {
			log.Error("failed to mark withdrawal as finalized", "err", err, "tx_hash", finalized.Event.TransactionHash.String())
			return err
		}
		if err := db.L2ToL1.UpdateL1FinalizeStatus(finalized.WithdrawalHash, finalizedWithdrawals[i].Event.GUID); err != nil {
			log.Error("failed to mark l2_to_l1 as finalized", "err", err, "tx_hash", finalized.Event.TransactionHash.String())
			return err
		}
	}

	// (3) L1CrossDomainMessenger
	crossDomainRelayedMessages, err := contracts2.CrossDomainMessengerRelayedMessageEvents("l1", l1Contracts.L1CrossDomainMessengerProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(crossDomainRelayedMessages) > 0 {
		log.Info("detected relayed messages", "size", len(crossDomainRelayedMessages))
	}

	relayedMessages := make(map[logKey]*contracts2.CrossDomainMessengerRelayedMessageEvent, len(crossDomainRelayedMessages))
	for i := range crossDomainRelayedMessages {
		relayed := crossDomainRelayedMessages[i]
		relayedMessages[logKey{BlockHash: relayed.Event.BlockHash, LogIndex: relayed.Event.LogIndex}] = &relayed
		message, err := db.BridgeMessages.L2BridgeMessage(relayed.MessageHash)
		if err != nil {
			return err
		} else if message == nil {
			log.Error("missing indexed L2CrossDomainMessenger message", "tx_hash", relayed.Event.TransactionHash.String())
			return fmt.Errorf("missing indexed L2CrossDomainMessager message. tx_hash %s", relayed.Event.TransactionHash.String())
		}

		if err := db.BridgeMessages.MarkRelayedL2BridgeMessage(relayed.MessageHash, relayed.Event.GUID); err != nil {
			log.Error("failed to relay cross domain message", "err", err, "tx_hash", relayed.Event.TransactionHash.String())
			return err
		}
	}

	// (4) L1StandardBridge
	finalizedBridges, err := contracts2.StandardBridgeFinalizedEvents("l1", l1Contracts.L1StandardBridgeProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(finalizedBridges) > 0 {
		log.Info("detected finalized bridge withdrawals", "size", len(finalizedBridges))
	}

	finalizedTokens := make(map[common.Address]int)
	for i := range finalizedBridges {
		// Nothing actionable on the database. However, we can treat the relayed message
		// as an invariant by ensuring we can query for a deposit by the same hash
		finalizedBridge := finalizedBridges[i]
		relayedMessage, ok := relayedMessages[logKey{finalizedBridge.Event.BlockHash, finalizedBridge.Event.LogIndex + 1}]
		if !ok {
			log.Error("expected RelayedMessage following BridgeFinalized event", "tx_hash", finalizedBridge.Event.TransactionHash.String())
			return fmt.Errorf("expected RelayedMessage following BridgeFinalized event. tx_hash = %s", finalizedBridge.Event.TransactionHash.String())
		}

		// Since the message hash is computed from the relayed message, this ensures the deposit fields must match
		withdrawal, err := db.BridgeTransfers.L2BridgeWithdrawalWithFilter(l1_l2.BridgeTransfer{CrossDomainMessageHash: &relayedMessage.MessageHash})
		if err != nil {
			return err
		} else if withdrawal == nil {
			log.Error("missing L2StandardBridge withdrawal on L1 finalization", "tx_hash", finalizedBridge.Event.TransactionHash.String())
			return fmt.Errorf("missing L2StandardBridge withdrawal on L1 finalization. tx_hash: %s", finalizedBridge.Event.TransactionHash.String())
		}
		if err := db.L2ToL1.UpdateClaimedStatus(withdrawal.TransactionWithdrawalHash); err != nil {
			log.Error("failed to mark l2_to_l1 as claimed", "err", err)
			return err
		}

		finalizedTokens[finalizedBridge.BridgeTransfer.TokenPair.LocalTokenAddress]++
	}
	return nil
}
