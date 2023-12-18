package bridge

import (
	"fmt"
	common4 "github.com/cornerstone-labs/acorus/event/processor/op-stack/common"
	"github.com/google/uuid"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	common3 "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	common2 "github.com/cornerstone-labs/acorus/common"
	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/processor/op-stack/bridge/ovm1"
	"github.com/cornerstone-labs/acorus/event/processor/op-stack/contracts"
	"github.com/cornerstone-labs/acorus/event/processor/op-stack/op-bindings/bindings"
)

func L1ProcessInitiatedBridgeEvents(log log.Logger, db *database.DB, fromHeight, toHeight *big.Int) error {
	// (1) OptimismPortal
	optimismPortalTxDeposits, err := contracts.OptimismPortalTransactionDepositEvents(common3.HexToAddress(common4.OptimismPortalProxy), db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(optimismPortalTxDeposits) > 0 {
		log.Info("detected transaction deposits", "size", len(optimismPortalTxDeposits))
	}

	mintedGWEI := bigint.Zero
	portalDeposits := make(map[logKey]*contracts.OptimismPortalTransactionDepositEvent, len(optimismPortalTxDeposits))
	l1ToL2s := make([]worker.L1ToL2, len(optimismPortalTxDeposits))
	for i := range optimismPortalTxDeposits {
		depositTx := optimismPortalTxDeposits[i]
		portalDeposits[logKey{depositTx.Event.BlockHash, depositTx.Event.LogIndex}] = &depositTx
		if depositTx.DepositTx.EthValue != nil {
			mintedGWEI = new(big.Int).Add(mintedGWEI, depositTx.DepositTx.EthValue)
		}
		blockNumber, err := db.L1ToL2.GetBlockNumberFromHash(depositTx.Event.BlockHash)
		if err != nil {
			log.Error("can not get l1 blockNumber", "blockHash", depositTx.Event.BlockHash)
			return err
		}
		l1ToL2s[i] = worker.L1ToL2{
			GUID:                  uuid.New(),
			L1BlockNumber:         blockNumber,
			L2BlockNumber:         bigint.One,
			QueueIndex:            nil,
			L1TransactionHash:     depositTx.Event.TransactionHash,
			L2TransactionHash:     common.Hash{},
			TransactionSourceHash: depositTx.DepositTx.SourceHash,
			MessageHash:           common.Hash{},
			L1TxOrigin:            common.Hash{},
			Status:                common2.L1ToL2Pending,
			L1TokenAddress:        common.Address{},
			L2TokenAddress:        common.Address{},
			ETHAmount:             depositTx.DepositTx.EthValue,
			ERC20Amount:           depositTx.DepositTx.Value,
			GasLimit:              depositTx.GasLimit,
			Timestamp:             int64(depositTx.Event.Timestamp),
		}
	}
	if len(l1ToL2s) > 0 {
		if err := db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s); err != nil {
			return err
		}
		mintedETH, _ := bigint.WeiToETH(mintedGWEI).Float64()
		log.Info("Minte ETH", "mintedETH", mintedETH)
	}

	// (2) L1CrossDomainMessenger
	crossDomainSentMessages, err := contracts.CrossDomainMessengerSentMessageEvents("l1", common3.HexToAddress(common4.L1CrossDomainMessengerProxy), db, fromHeight, toHeight)
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
		if err := db.L1ToL2.UpdateMessageHash(l1ToL2c2); err != nil {
			return err
		}
	}

	// (3) L1StandardBridge
	initiatedBridges, err := contracts.StandardBridgeInitiatedEvents("l1", common3.HexToAddress(common4.L1StandardBridgeProxy), db, fromHeight, toHeight)
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
			initiatedBridge.CrossDomainMessageHash = &sentMessage.MessageHash
			bridgedTokens[initiatedBridge.LocalTokenAddress]++

			l1ToL2s2[i].L1TransactionHash = portalDeposit.Event.TransactionHash
			l1ToL2s2[i].FromAddress = initiatedBridge.FromAddress
			l1ToL2s2[i].ToAddress = initiatedBridge.ToAddress
			l1ToL2s2[i].L1TokenAddress = initiatedBridge.LocalTokenAddress
			l1ToL2s2[i].L2TokenAddress = initiatedBridge.RemoteTokenAddress
			l1ToL2s2[i].ERC20Amount = initiatedBridge.ERC20Amount
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
			l1ToL2s2[i].ERC20Amount = initiatedBridge.ERC20Amount
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
			l1ToL2s2[i].ERC20Amount = initiatedBridge.ERC20Amount
			l1ToL2s2[i].L1TokenAddress = initiatedBridge.LocalTokenAddress
			l1ToL2s2[i].L2TokenAddress = initiatedBridge.RemoteTokenAddress
		}
	}
	if len(l1ToL2s2) > 0 {
		if err := db.L1ToL2.UpdateTokenPairAndAddress(l1ToL2s2); err != nil {
			return err
		}
	}
	return nil
}

func L1ProcessFinalizedBridgeEvents(log log.Logger, db *database.DB, fromHeight, toHeight *big.Int) error {
	// (1) OptimismPortal (proven withdrawals)
	provenWithdrawals, err := contracts.OptimismPortalWithdrawalProvenEvents(common3.HexToAddress(common4.OptimismPortalProxy), db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(provenWithdrawals) > 0 {
		log.Info("detected proven withdrawals", "size", len(provenWithdrawals))
	}

	skippedOVM1ProvenWithdrawals := 0
	l2ToL1Prove := make([]worker.L2ToL1, len(provenWithdrawals))
	for i := range provenWithdrawals {
		proven := provenWithdrawals[i]
		withdrawal, err := db.L2ToL1.L2ToL1TransactionWithdrawal(proven.WithdrawalHash)
		if err != nil {
			return err
		} else if withdrawal == nil {
			if _, ok := ovm1.PortalWithdrawalTransactions[proven.WithdrawalHash]; ok {
				skippedOVM1ProvenWithdrawals++
				continue
			}
			log.Warn("missing indexed withdrawal", "tx_hash", proven.Event.TransactionHash.String())
			continue
		}
		blockNumber := bigint.One
		proveL1BlockNumber, err := db.L1ToL2.GetBlockNumberFromHash(proven.Event.BlockHash)
		if err != nil {
			return err
		}
		if proveL1BlockNumber != nil {
			blockNumber = proveL1BlockNumber
		}
		l2ToL1Prove[i].WithdrawTransactionHash = proven.WithdrawalHash
		l2ToL1Prove[i].Status = common2.L2ToL1InChallengePeriod
		l2ToL1Prove[i].L1ProveTxHash = proven.Event.TransactionHash
		l2ToL1Prove[i].L1BlockNumber = blockNumber
	}
	if len(provenWithdrawals) > 0 {
		if err := db.L2ToL1.MarkL2ToL1TransactionWithdrawalProven(l2ToL1Prove); err != nil {
			return err
		}
		if skippedOVM1ProvenWithdrawals > 0 {
			log.Info("skipped OVM 1.0 proven withdrawals", "size", skippedOVM1ProvenWithdrawals)
		}
	}

	// (2) OptimismPortal (finalized withdrawals)
	finalizedWithdrawals, err := contracts.OptimismPortalWithdrawalFinalizedEvents(common3.HexToAddress(common4.OptimismPortalProxy), db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(finalizedWithdrawals) > 0 {
		log.Info("detected finalized withdrawals", "size", len(finalizedWithdrawals))
	}

	skippedOVM1FinalizedWithdrawals := 0
	l2ToL1Finalized := make([]worker.L2ToL1, len(finalizedWithdrawals))
	for i := range finalizedWithdrawals {
		finalized := finalizedWithdrawals[i]
		blockNumber := bigint.One
		finalizedBlockNumber, err := db.L1ToL2.GetBlockNumberFromHash(finalized.Event.BlockHash)
		if err != nil {
			return err
		}
		if finalizedBlockNumber != nil {
			blockNumber = finalizedBlockNumber
		}

		l2ToL1Finalized[i].L1BlockNumber = blockNumber
		l2ToL1Finalized[i].WithdrawTransactionHash = finalized.WithdrawalHash
		l2ToL1Finalized[i].Status = common2.L2ToL1Claimed
		l2ToL1Finalized[i].L1FinalizeTxHash = finalized.Event.TransactionHash
	}
	if len(finalizedWithdrawals) > 0 {
		if err := db.L2ToL1.MarkL2ToL1TransactionWithdrawalFinalized(l2ToL1Finalized); err != nil {
			return err
		}
		if skippedOVM1ProvenWithdrawals > 0 {
			log.Info("skipped OVM 1.0 finalized withdrawals", "size", skippedOVM1FinalizedWithdrawals)
		}
	}

	// (3) L1CrossDomainMessenger
	crossDomainRelayedMessages, err := contracts.CrossDomainMessengerRelayedMessageEvents("l1", common3.HexToAddress(common4.L1CrossDomainMessengerProxy), db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(crossDomainRelayedMessages) > 0 {
		log.Info("detected relayed messages", "size", len(crossDomainRelayedMessages))
	}

	l2ToL1Cs := make([]worker.L2ToL1, len(crossDomainRelayedMessages))
	relayedMessages := make(map[logKey]*contracts.CrossDomainMessengerRelayedMessageEvent, len(crossDomainRelayedMessages))
	skippedOVM1Messages := 0
	for i := range crossDomainRelayedMessages {
		relayed := crossDomainRelayedMessages[i]
		relayedMessages[logKey{BlockHash: relayed.Event.BlockHash, LogIndex: relayed.Event.LogIndex}] = &relayed
		l2ToL1Cs[i].MessageHash = relayed.MessageHash
	}
	if len(crossDomainRelayedMessages) > 0 {
		if skippedOVM1Messages > 0 {
			log.Info("skipped OVM 1.0 relayed L2CrossDomainMessenger withdrawals", "size", skippedOVM1Messages)
		}
	}
	// (4) L1StandardBridge
	finalizedBridges, err := contracts.StandardBridgeFinalizedEvents("l1", common3.HexToAddress(common4.L1StandardBridgeProxy), db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(finalizedBridges) > 0 {
		log.Info("detected finalized bridge withdrawals", "size", len(finalizedBridges))
	}

	l2ToL1Fb := make([]worker.L2ToL1, len(crossDomainRelayedMessages))
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
		l2ToL1Fb[i].L1TokenAddress = finalizedBridge.LocalTokenAddress
		l2ToL1Fb[i].L2TokenAddress = finalizedBridge.RemoteTokenAddress
		l2ToL1Fb[i].ETHAmount = finalizedBridge.ETHAmount
		l2ToL1Fb[i].ERC20Amount = finalizedBridge.ERC20Amount
		l2ToL1Fb[i].MessageHash = relayedMessage.MessageHash
		finalizedTokens[finalizedBridge.LocalTokenAddress]++
	}
	if len(finalizedBridges) > 0 {
		for tokenAddr, size := range finalizedTokens {
			if err := db.L2ToL1.UpdateTokenPairsAndAddress(l2ToL1Fb); err != nil {
				return err
			}
			log.Info("tokenAddr and size", "tokenAddr", tokenAddr, "size", size)
		}
	}
	return nil
}
