package bridge

import (
	"fmt"
	common2 "github.com/cornerstone-labs/acorus/event/processor/op-stack/common"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/worker"
	contracts2 "github.com/cornerstone-labs/acorus/event/processor/op-stack/contracts"
	"github.com/ethereum-optimism/optimism/indexer/bigint"
)

func LegacyL1ProcessInitiatedBridgeEvents(log log.Logger, db *database.DB, fromHeight, toHeight *big.Int) error {
	// (1) CanonicalTransactionChain
	ctcTxDepositEvents, err := contracts2.LegacyCTCDepositEvents(common.HexToAddress(common2.LegacyCanonicalTransactionChain), db, fromHeight, toHeight)
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
		blockNumber, err := db.L1ToL2.GetBlockNumberFromHash(depositTx.Event.BlockHash)
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
			L1TxOrigin:        common.Hash{},
			Status:            0,
			L1TokenAddress:    common.Address{},
			L2TokenAddress:    common.Address{},
			ETHAmount:         depositTx.ETHAmount,
			ERC20Amount:       depositTx.ERC20Amount,
			GasLimit:          depositTx.GasLimit,
			Timestamp:         int64(depositTx.Event.Timestamp),
		}
	}
	if len(ctcTxDepositEvents) > 0 {
		if err := db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s); err != nil {
			return err
		}
		mintedETH, _ := bigint.WeiToETH(mintedWEI).Float64()
		log.Info("mint eth amount", "mintedETH", mintedETH)
	}

	// (2) L1CrossDomainMessenger
	crossDomainSentMessages, err := contracts2.CrossDomainMessengerSentMessageEvents("l1", common.HexToAddress(common2.L1CrossDomainMessengerProxy), db, fromHeight, toHeight)
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
		if err := db.L1ToL2.UpdateMessageHash(l1ToL2c2); err != nil {
			return err
		}
	}

	// (3) L1StandardBridge
	initiatedBridges, err := contracts2.L1StandardBridgeLegacyDepositInitiatedEvents(common.HexToAddress(common2.L1StandardBridgeProxy), db, fromHeight, toHeight)
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
		l1ToL2s2[i].ERC20Amount = initiatedBridge.ERC20Amount
	}
	if len(initiatedBridges) > 0 {
		if err := db.L1ToL2.UpdateTokenPairAndAddress(l1ToL2s2); err != nil {
			return err
		}
		for tokenAddr, size := range bridgedTokens {
			log.Info("tokenAddr and size", "tokenAddr", tokenAddr, "size", size)
		}
	}
	return nil
}

func LegacyL2ProcessInitiatedBridgeEvents(log log.Logger, db *database.DB, fromHeight, toHeight *big.Int) error {
	// (1) L2CrossDomainMessenger
	crossDomainSentMessages, err := contracts2.CrossDomainMessengerSentMessageEvents("l2", common.HexToAddress(common2.L2CrossDomainMessenger), db, fromHeight, toHeight)
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

		l2CrossDomainMessenger := common.HexToAddress(common2.L2CrossDomainMessenger)
		withdrawalHash := crypto.Keccak256Hash(append(sentMessage.MessageCalldata, l2CrossDomainMessenger[:]...))
		blockNumber, err := db.L1ToL2.GetBlockNumberFromHash(sentMessage.Event.BlockHash)
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
			ERC20Amount:             sentMessage.ERC20Amount,
			GasLimit:                sentMessage.GasLimit,
			TimeLeft:                new(big.Int).SetUint64(0),
			L1TokenAddress:          common.Address{},
			L2TokenAddress:          common.Address{},
			Timestamp:               int64(sentMessage.Event.Timestamp),
		}
	}
	if len(crossDomainSentMessages) > 0 {
		if err := db.L2ToL1.StoreL2ToL1Transactions(l2ToL1Cs); err != nil {
			return err
		}
		withdrawnETH, _ := bigint.WeiToETH(withdrawnWEI).Float64()
		log.Info("withdraw eth", "withdrawnETH", withdrawnETH)
	}

	// (2) L2StandardBridge
	initiatedBridges, err := contracts2.L2StandardBridgeLegacyWithdrawalInitiatedEvents(common.HexToAddress(common2.L2StandardBridge), db, fromHeight, toHeight)
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
		if err := db.L2ToL1.UpdateTokenPairsAndAddress(l2ToL1Bs); err != nil {
			return err
		}
		for tokenAddr, size := range bridgedTokens {
			log.Info("tokenAddr and size", "tokenAddr", tokenAddr, "size", size)
		}
	}
	return nil
}

func LegacyL1ProcessFinalizedBridgeEvents(log log.Logger, db *database.DB, fromHeight, toHeight *big.Int) error {
	crossDomainRelayedMessages, err := contracts2.CrossDomainMessengerRelayedMessageEvents("l1", common.HexToAddress(common2.L1CrossDomainMessengerProxy), db, fromHeight, toHeight)
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
		if err = db.L2ToL1.MarkL2ToL1TransactionWithdrawalFinalized(l2ToL1Finalized); err != nil {
			return err
		}
	}

	return nil
}

func LegacyL2ProcessFinalizedBridgeEvents(log log.Logger, db *database.DB, fromHeight, toHeight *big.Int) error {
	// (1) L2CrossDomainMessenger
	crossDomainRelayedMessages, err := contracts2.CrossDomainMessengerRelayedMessageEvents("l2", common.HexToAddress(common2.L2CrossDomainMessenger), db, fromHeight, toHeight)
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
		if err := db.L1ToL2.UpdateMessageHash(L1ToL2Fz); err != nil {
			log.Error("failed to relay cross domain message", "err", err)
			return err
		}
	}
	return nil
}
