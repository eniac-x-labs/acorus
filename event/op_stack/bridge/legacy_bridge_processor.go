package bridge

import (
	"fmt"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/worker"
	common3 "github.com/cornerstone-labs/acorus/event/op_stack/common"
	contracts2 "github.com/cornerstone-labs/acorus/event/op_stack/contracts"
)

func LegacyL1ProcessInitiatedBridgeEvents(db *database.DB, fromHeight, toHeight *big.Int) error {
	// (1) CanonicalTransactionChain
	ctcTxDepositEvents, err := contracts2.LegacyCTCDepositEvents(common3.LegacyCanonicalTransactionChain, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(ctcTxDepositEvents) > 0 {
		log.Info("detected legacy transaction deposits", "size", len(ctcTxDepositEvents))
	}

	mintedWEI := bigint.Zero
	ctcTxDeposits := make(map[logKey]*contracts2.LegacyCTCDepositEvent, len(ctcTxDepositEvents))
	l1ToL2s := make([]worker.L1ToL2, len(ctcTxDeposits))
	for i := range ctcTxDepositEvents {
		depositTx := ctcTxDepositEvents[i]
		mintedWEI = new(big.Int).Add(mintedWEI, depositTx.ETHAmount)
		blockNumber, err := db.L1ToL2.GetBlockNumberFromHash("10", depositTx.Event.BlockHash)
		if err != nil {
			log.Error("can not get l1 blockNumber", "blockHash", depositTx.Event.BlockHash)
			return err
		}
		l1ToL2s[i] = worker.L1ToL2{
			L1BlockNumber:     blockNumber,
			QueueIndex:        nil,
			L1TransactionHash: depositTx.Event.TransactionHash,
			L2TransactionHash: common.Hash{},
			MessageHash:       common.Hash{},
			L1TxOrigin:        depositTx.FromAddress,
			Status:            0,
			L1TokenAddress:    common.Address{},
			L2TokenAddress:    common.Address{},
			ETHAmount:         depositTx.ETHAmount,
			GasLimit:          depositTx.GasLimit,
			Timestamp:         int64(depositTx.Event.Timestamp),
		}
	}
	if len(ctcTxDepositEvents) > 0 {
		if err := db.L1ToL2.StoreL1ToL2Transactions("10", l1ToL2s); err != nil {
			return err
		}
	}

	// (2) L1CrossDomainMessenger
	crossDomainSentMessages, err := contracts2.CrossDomainMessengerSentMessageEvents(common3.L1CrossDomainMessengerProxy, "1", db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(crossDomainSentMessages) > 0 {
		log.Info("detected legacy sent messages", "size", len(crossDomainSentMessages))
	}

	sentMessages := make(map[logKey]*contracts2.CrossDomainMessengerSentMessageEvent, len(crossDomainSentMessages))
	l1ToL2c2 := make([]worker.L1ToL2, len(crossDomainSentMessages))
	for i := range crossDomainSentMessages {
		sentMessage := crossDomainSentMessages[i]
		sentMessages[logKey{sentMessage.Event.BlockHash, sentMessage.Event.LogIndex}] = &sentMessage
		l1ToL2c2[i].L1TransactionHash = sentMessage.Event.TransactionHash
		l1ToL2c2[i].MessageHash = sentMessage.MessageHash
	}
	if len(crossDomainSentMessages) > 0 {
		if err := db.L1ToL2.UpdateTokenPairAndAddress("10", l1ToL2c2); err != nil {
			return err
		}
	}

	// (3) L1StandardBridge
	initiatedBridges, err := contracts2.L1StandardBridgeLegacyDepositInitiatedEvents(common3.L1StandardBridgeProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(initiatedBridges) > 0 {
		log.Info("detected iegacy bridge deposits", "size", len(initiatedBridges))
	}

	bridgedTokens := make(map[common.Address]int)
	l1ToL2s2 := make([]worker.L1ToL2, len(initiatedBridges))
	for i := range initiatedBridges {
		initiatedBridge := initiatedBridges[i]
		sentMessage, ok := sentMessages[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex - 1}]
		if !ok {
			log.Error("missing cross domain message for bridge transfer", "tx_hash", initiatedBridge.Event.TransactionHash.String())
			return fmt.Errorf("expected SentMessage preceding DepositInitiated event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
		}
		ctcTxDeposit, ok := ctcTxDeposits[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex - 2}]
		if !ok {
			log.Error("missing transaction deposit for bridge transfer", "tx_hash", initiatedBridge.Event.TransactionHash.String())
			return fmt.Errorf("expected TransactionEnqueued preceding DepostInitiated event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
		}

		initiatedBridge.CrossDomainMessageHash = &sentMessage.MessageHash
		bridgedTokens[initiatedBridge.LocalTokenAddress]++

		l1ToL2s2[i].L1TransactionHash = ctcTxDeposit.Event.TransactionHash
		l1ToL2s2[i].FromAddress = initiatedBridge.FromAddress
		l1ToL2s2[i].ToAddress = initiatedBridge.ToAddress
		l1ToL2s2[i].L1TokenAddress = initiatedBridge.LocalTokenAddress
		l1ToL2s2[i].L2TokenAddress = initiatedBridge.RemoteTokenAddress
	}
	if len(initiatedBridges) > 0 {
		if err := db.L1ToL2.UpdateTokenPairAndAddress("10", l1ToL2s2); err != nil {
			return err
		}
	}
	return nil
}

func LegacyL2ProcessInitiatedBridgeEvents(db *database.DB, fromHeight, toHeight *big.Int) error {
	// (1) L2CrossDomainMessenger
	crossDomainSentMessages, err := contracts2.CrossDomainMessengerSentMessageEvents(common3.L2CrossDomainMessenger, "10", db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(crossDomainSentMessages) > 0 {
		log.Info("detected legacy transaction withdrawals (via L2CrossDomainMessenger)", "size", len(crossDomainSentMessages))
	}

	l2ToL1Cs := make([]worker.L2ToL1, len(crossDomainSentMessages))
	withdrawnWEI := bigint.Zero
	sentMessages := make(map[logKey]*contracts2.CrossDomainMessengerSentMessageEvent, len(crossDomainSentMessages))
	for i := range crossDomainSentMessages {
		sentMessage := crossDomainSentMessages[i]
		sentMessages[logKey{sentMessage.Event.BlockHash, sentMessage.Event.LogIndex}] = &sentMessage
		withdrawnWEI = new(big.Int).Add(withdrawnWEI, sentMessage.ETHAmount)

		// To ensure consistency in the schema, we duplicate this as the "root" transaction withdrawal. The storage key in the message
		// passer contract is sha3(calldata + sender). The sender always being the L2CrossDomainMessenger pre-bedrock.
		withdrawalHash := crypto.Keccak256Hash(append(sentMessage.MessageCalldata, common3.L2CrossDomainMessenger[:]...))
		blockNumber, err := db.L1ToL2.GetBlockNumberFromHash("10", sentMessage.Event.BlockHash)
		if err != nil {
			log.Error("can not get l1 blockNumber", "blockHash", sentMessage.Event.BlockHash)
			return err
		}
		l2ToL1Cs[i] = worker.L2ToL1{
			L2BlockNumber:           blockNumber,
			MsgNonce:                sentMessage.Nonce,
			L2TransactionHash:       sentMessage.Event.TransactionHash,
			WithdrawTransactionHash: withdrawalHash,
			L1ProveTxHash:           common.Hash{},
			L1FinalizeTxHash:        common.Hash{},
			Status:                  0,
			ETHAmount:               sentMessage.ETHAmount,
			GasLimit:                sentMessage.GasLimit,
			TimeLeft:                new(big.Int).SetUint64(0),
			L1TokenAddress:          common.Address{},
			L2TokenAddress:          common.Address{},
			Timestamp:               int64(sentMessage.Event.Timestamp),
		}
	}
	if len(crossDomainSentMessages) > 0 {
		if err := db.L2ToL1.StoreL2ToL1Transactions("10", l2ToL1Cs); err != nil {
			return err
		}
	}

	// (2) L2StandardBridge
	initiatedBridges, err := contracts2.L2StandardBridgeLegacyWithdrawalInitiatedEvents(common3.L2StandardBridge, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(initiatedBridges) > 0 {
		log.Info("detected legacy bridge withdrawals", "size", len(initiatedBridges))
	}

	bridgedTokens := make(map[common.Address]int)

	l2ToL1Bs := make([]worker.L2ToL1, len(initiatedBridges))

	for i := range initiatedBridges {
		initiatedBridge := initiatedBridges[i]

		// extract the cross domain message hash & deposit source hash from the following events
		// Unlike bedrock, the bridge events are emitted AFTER sending the cross domain message
		// 	- Event Flow: TransactionEnqueued -> SentMessage -> DepositInitiated
		sentMessage, ok := sentMessages[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex - 1}]
		if !ok {
			log.Error("expected SentMessage preceding DepositInitiated event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
			return fmt.Errorf("expected SentMessage preceding DepositInitiated event. tx_hash = %s", initiatedBridge.Event.TransactionHash)
		}
		initiatedBridge.CrossDomainMessageHash = &sentMessage.MessageHash
		bridgedTokens[initiatedBridge.LocalTokenAddress]++
		l2ToL1Bs[i].MessageHash = sentMessage.MessageHash
		l2ToL1Bs[i].FromAddress = initiatedBridge.FromAddress
		l2ToL1Bs[i].ToAddress = initiatedBridge.ToAddress
		l2ToL1Bs[i].L1TokenAddress = initiatedBridge.LocalTokenAddress
		l2ToL1Bs[i].L2TokenAddress = initiatedBridge.RemoteTokenAddress
	}
	if len(initiatedBridges) > 0 {
		if err := db.L2ToL1.UpdateTokenPairsAndAddress("10", l2ToL1Bs); err != nil {
			return err
		}
	}
	return nil
}

func LegacyL1ProcessFinalizedBridgeEvents(db *database.DB, fromHeight, toHeight *big.Int) error {
	crossDomainRelayedMessages, err := contracts2.CrossDomainMessengerRelayedMessageEvents("1", common3.L1CrossDomainMessengerProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(crossDomainRelayedMessages) > 0 {
		log.Info("detected relayed messages", "size", len(crossDomainRelayedMessages))
	}

	l2ToL1Finalized := make([]worker.L2ToL1, len(crossDomainRelayedMessages))
	for i := range crossDomainRelayedMessages {
		relayedMessage := crossDomainRelayedMessages[i]

		l2ToL1Finalized[i].MessageHash = relayedMessage.MessageHash
		l2ToL1Finalized[i].L1FinalizeTxHash = relayedMessage.Event.TransactionHash
	}
	if len(crossDomainRelayedMessages) > 0 {
		if err = db.L2ToL1.MarkL2ToL1TransactionWithdrawalFinalized("10", l2ToL1Finalized); err != nil {
			return err
		}
	}

	return nil
}

func LegacyL2ProcessFinalizedBridgeEvents(db *database.DB, fromHeight, toHeight *big.Int) error {
	// (1) L2CrossDomainMessenger
	crossDomainRelayedMessages, err := contracts2.CrossDomainMessengerRelayedMessageEvents("10", common3.L2CrossDomainMessenger, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(crossDomainRelayedMessages) > 0 {
		log.Info("detected relayed legacy messages", "size", len(crossDomainRelayedMessages))
	}

	L1ToL2Fz := make([]worker.L1ToL2, len(crossDomainRelayedMessages))
	for i := range crossDomainRelayedMessages {
		relayedMessage := crossDomainRelayedMessages[i]
		L1ToL2Fz[i].MessageHash = relayedMessage.MessageHash
		L1ToL2Fz[i].L2TransactionHash = relayedMessage.Event.TransactionHash
	}
	if len(crossDomainRelayedMessages) > 0 {
		if err := db.L1ToL2.UpdateMessageHash("10", L1ToL2Fz); err != nil {
			log.Error("failed to relay cross domain message", "err", err)
			return err
		}
	}
	return nil
}
