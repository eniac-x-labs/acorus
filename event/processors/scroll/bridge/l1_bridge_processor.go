package bridge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"

	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	l1_l2 "github.com/cornerstone-labs/acorus/database/event/l1-l2"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/processors/scroll/abi"
	"github.com/cornerstone-labs/acorus/event/processors/scroll/utils"
)

func L1DepositETH(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1DepositETHSig := abi.L1DepositETHSig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1DepositETHSig}
	depositEvents, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l1ToL2s := make([]worker.L1ToL2, len(depositEvents))
	transactionDeposits := make([]l1_l2.L1TransactionDeposit, len(depositEvents))
	if len(depositEvents) > 0 {
		log.Info("detected eth transaction deposits", "size", len(depositEvents))
	}
	for i := range depositEvents {
		ethDepositTx := depositEvents[i]
		rlpLog := *ethDepositTx.RLPLog
		depositEvent := abi.DepositETH{}

		unpackErr := utils.UnpackLog(abi.L1ETHGatewayABI, &depositEvent, "DepositETH", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}

		transactionDeposits[i] = l1_l2.L1TransactionDeposit{
			SourceHash:           ethDepositTx.TransactionHash,
			InitiatedL1EventGUID: ethDepositTx.GUID,
			Tx: l1_l2.Transaction{
				FromAddress: depositEvent.From,
				ToAddress:   depositEvent.To,
				ETHAmount:   depositEvent.Amount,
				ERC20Amount: big.NewInt(0),
				Data:        depositEvent.Data,
				Timestamp:   ethDepositTx.Timestamp,
			},
			GasLimit: big.NewInt(0),
		}
		l1ToL2s[i] = worker.L1ToL2{
			GUID:              uuid.New(),
			L1BlockHash:       rlpLog.BlockHash,
			QueueIndex:        nil,
			L1TransactionHash: rlpLog.TxHash,
			L2TransactionHash: common.Hash{},
			L1TxOrigin:        common.Hash{},
			Status:            0,
			FromAddress:       depositEvent.From,
			ToAddress:         depositEvent.To,
			L1TokenAddress:    common.Address{},
			L2TokenAddress:    common.Address{},
			ETHAmount:         depositEvent.Amount,
			ERC20Amount:       big.NewInt(0),
			GasLimit:          big.NewInt(0),
			Timestamp:         int64(ethDepositTx.Timestamp),
			AssetType:         int64(common2.ETH),
			MsgHash:           common.Hash{},
		}
	}
	if len(l1ToL2s) > 0 {
		if err := db.BridgeTransactions.StoreL1TransactionDeposits(transactionDeposits); err != nil {
			return err
		}
		//marshal, _ := json.Marshal(l1ToL2s)
		if err := db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s); err != nil {
			return err
		}
	}
	return nil
}

func L1DepositERC20(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1DepositERC20Sig := abi.L1DepositERC20Sig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1DepositERC20Sig}
	depositErc20Events, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l1ToL2s := make([]worker.L1ToL2, len(depositErc20Events))
	transactionDeposits := make([]l1_l2.L1TransactionDeposit, len(depositErc20Events))

	for i := range depositErc20Events {
		depositTx := depositErc20Events[i]
		rlpLog := *depositTx.RLPLog
		depositErc20Event := abi.ERC20MessageEvent{}
		transactionDeposits[i] = l1_l2.L1TransactionDeposit{
			SourceHash:           depositTx.TransactionHash,
			InitiatedL1EventGUID: depositTx.GUID,
			Tx: l1_l2.Transaction{
				FromAddress: depositErc20Event.From,
				ToAddress:   depositErc20Event.To,
				ETHAmount:   big.NewInt(0),
				ERC20Amount: depositErc20Event.Amount,
				Data:        depositErc20Event.Data,
				Timestamp:   depositTx.Timestamp,
			},
			GasLimit: big.NewInt(0),
		}
		unpackErr := utils.UnpackLog(abi.L1StandardERC20GatewayABI, &depositErc20Event, "DepositERC20", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		l1ToL2s[i] = worker.L1ToL2{
			GUID:              uuid.New(),
			L1BlockHash:       rlpLog.BlockHash,
			QueueIndex:        nil,
			L1TransactionHash: rlpLog.TxHash,
			L2TransactionHash: common.Hash{},
			L1TxOrigin:        common.Hash{},
			Status:            0,
			FromAddress:       depositErc20Event.From,
			ToAddress:         depositErc20Event.To,
			L1TokenAddress:    depositErc20Event.L1Token,
			L2TokenAddress:    depositErc20Event.L2Token,
			ETHAmount:         big.NewInt(0),
			ERC20Amount:       depositErc20Event.Amount,
			GasLimit:          big.NewInt(0),
			Timestamp:         int64(depositTx.Timestamp),
			AssetType:         int64(common2.ERC20),
		}
	}
	if len(l1ToL2s) > 0 {
		if err := db.BridgeTransactions.StoreL1TransactionDeposits(transactionDeposits); err != nil {
			return err
		}
		if err := db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s); err != nil {
			return err
		}
	}
	return nil
}

func L1DepositERC721(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1DepositERC721Sig := abi.L1DepositERC721Sig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1DepositERC721Sig}
	depositErc721Events, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l1ToL2s := make([]worker.L1ToL2, len(depositErc721Events))
	for i := range depositErc721Events {
		depositTx := depositErc721Events[i]
		rlpLog := *depositTx.RLPLog
		depositErc721Event := abi.ERC721MessageEvent{}

		unpackErr := utils.UnpackLog(abi.L1ERC721GatewayABI, &depositErc721Event, "DepositERC721", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		l1ToL2s[i] = worker.L1ToL2{
			GUID:              uuid.New(),
			L1BlockHash:       rlpLog.BlockHash,
			QueueIndex:        nil,
			L1TransactionHash: rlpLog.TxHash,
			L2TransactionHash: common.Hash{},
			L1TxOrigin:        common.Hash{},
			Status:            0,
			FromAddress:       depositErc721Event.From,
			ToAddress:         depositErc721Event.To,
			L1TokenAddress:    depositErc721Event.L1Token,
			L2TokenAddress:    depositErc721Event.L2Token,
			ETHAmount:         depositErc721Event.Amount,
			ERC20Amount:       big.NewInt(0),
			GasLimit:          big.NewInt(0),
			Timestamp:         int64(depositTx.Timestamp),
			AssetType:         int64(common2.ERC721),
			TokenIds:          depositErc721Event.TokenID.String(),
		}
	}
	if len(l1ToL2s) > 0 {
		if err := db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s); err != nil {
			return err
		}
	}
	return nil
}

func L1DepositERC1155(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1DepositERC1155Sig := abi.L1DepositERC1155Sig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1DepositERC1155Sig}
	depositErc1155Events, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l1ToL2s := make([]worker.L1ToL2, len(depositErc1155Events))
	for i := range depositErc1155Events {
		depositTx := depositErc1155Events[i]
		rlpLog := *depositTx.RLPLog
		depositErc1155Event := abi.ERC1155MessageEvent{}

		unpackErr := utils.UnpackLog(abi.L1ERC1155GatewayABI, &depositErc1155Event, "DepositERC1155", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		l1ToL2s[i] = worker.L1ToL2{
			GUID:              uuid.New(),
			L1BlockHash:       rlpLog.BlockHash,
			QueueIndex:        nil,
			L1TransactionHash: rlpLog.TxHash,
			L2TransactionHash: common.Hash{},
			L1TxOrigin:        common.Hash{},
			Status:            0,
			FromAddress:       depositErc1155Event.From,
			ToAddress:         depositErc1155Event.To,
			L1TokenAddress:    depositErc1155Event.L1Token,
			L2TokenAddress:    depositErc1155Event.L2Token,
			ETHAmount:         depositErc1155Event.Amount,
			ERC20Amount:       big.NewInt(0),
			GasLimit:          big.NewInt(0),
			Timestamp:         int64(depositTx.Timestamp),
			AssetType:         int64(common2.ERC1155),
			TokenIds:          depositErc1155Event.TokenID.String(),
		}
	}
	if len(l1ToL2s) > 0 {
		if err := db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s); err != nil {
			return err
		}
	}
	return nil

}

func L1SentMessageEvent(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1SentMessageEventSignature := abi.L1SentMessageEventSignature
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1SentMessageEventSignature}
	sentMessageEvents, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l1BridgeMsgs := make([]l1_l2.L1BridgeMessage, len(sentMessageEvents))
	for i := range sentMessageEvents {
		depositTx := sentMessageEvents[i]
		rlpLog := *depositTx.RLPLog
		sentMessageEvent := abi.L1SentMessageEvent{}
		unpackErr := utils.UnpackLog(abi.L1ScrollMessengerABI, &sentMessageEvent, "SentMessage", rlpLog)
		if unpackErr != nil {
			log.Warn("Failed to unpack SentMessage event", "err", err)
			return unpackErr
		}
		// compute msgHash
		msgHash := utils.ComputeMessageHash(sentMessageEvent.Sender, sentMessageEvent.Target,
			sentMessageEvent.Value, sentMessageEvent.MessageNonce, sentMessageEvent.Message)
		// update l1tol2  set msgHash by txHash
		if err := db.L1ToL2.UpdateL1ToL2MsgHashByL1TxHash(worker.L1ToL2{L1TransactionHash: rlpLog.TxHash, MsgHash: msgHash}); err != nil {
			return err
		}
		relayedMessageEventGUID := uuid.New()
		l1BridgeMsgs[i] = l1_l2.L1BridgeMessage{
			BridgeMessage: l1_l2.BridgeMessage{
				MessageHash:             msgHash,
				Nonce:                   sentMessageEvent.MessageNonce,
				GasLimit:                sentMessageEvent.GasLimit,
				SentMessageEventGUID:    uuid.New(),
				RelayedMessageEventGUID: &relayedMessageEventGUID,
				Tx: l1_l2.Transaction{
					FromAddress: sentMessageEvent.Sender,
					ToAddress:   sentMessageEvent.Target,
					ETHAmount:   sentMessageEvent.Value,
					ERC20Amount: big.NewInt(0),
					Data:        sentMessageEvent.Message,
					Timestamp:   depositTx.Timestamp,
				},
			},
			TransactionSourceHash: rlpLog.TxHash,
		}
	}
	if len(l1BridgeMsgs) > 0 {
		if err := db.BridgeMessages.StoreL1BridgeMessages(l1BridgeMsgs); err != nil {
			return err
		}
	}
	return nil
}

func L1BatchDepositERC721(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1BatchDepositERC721Sig := abi.L1BatchDepositERC721Sig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1BatchDepositERC721Sig}
	batchDepositERC721Events, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l1ToL2s := make([]worker.L1ToL2, len(batchDepositERC721Events))
	for i := range batchDepositERC721Events {
		depositTx := batchDepositERC721Events[i]
		rlpLog := *depositTx.RLPLog
		batchDepositErc721Event := abi.BatchERC721MessageEvent{}

		unpackErr := utils.UnpackLog(abi.L1ERC721GatewayABI, &batchDepositErc721Event, "BatchDepositERC721", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		l1ToL2s[i] = worker.L1ToL2{
			GUID:              uuid.New(),
			L1BlockHash:       rlpLog.BlockHash,
			QueueIndex:        nil,
			L1TransactionHash: rlpLog.TxHash,
			L2TransactionHash: common.Hash{},
			L1TxOrigin:        common.Hash{},
			Status:            0,
			FromAddress:       batchDepositErc721Event.From,
			ToAddress:         batchDepositErc721Event.To,
			L1TokenAddress:    batchDepositErc721Event.L1Token,
			L2TokenAddress:    batchDepositErc721Event.L2Token,
			ETHAmount:         big.NewInt(0),
			ERC20Amount:       big.NewInt(0),
			GasLimit:          big.NewInt(0),
			Timestamp:         int64(depositTx.Timestamp),
			AssetType:         int64(common2.ERC721),
			TokenIds:          utils.ConvertBigIntArrayToString(batchDepositErc721Event.TokenIDs),
		}
	}
	if len(l1ToL2s) > 0 {
		if err := db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s); err != nil {
			return err
		}
	}
	return nil
}

func L1BatchDepositERC1155(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1BatchDepositERC1155Sig := abi.L1BatchDepositERC1155Sig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1BatchDepositERC1155Sig}
	batchDepositERC1155Events, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l1ToL2s := make([]worker.L1ToL2, len(batchDepositERC1155Events))
	for i := range batchDepositERC1155Events {
		depositTx := batchDepositERC1155Events[i]
		rlpLog := *depositTx.RLPLog
		batchDepositERC1155Event := abi.BatchERC1155MessageEvent{}

		unpackErr := utils.UnpackLog(abi.L1ERC1155GatewayABI, &batchDepositERC1155Event, "BatchWithdrawERC1155", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		l1ToL2s[i] = worker.L1ToL2{
			GUID:              uuid.New(),
			L1BlockHash:       rlpLog.BlockHash,
			QueueIndex:        nil,
			L1TransactionHash: rlpLog.TxHash,
			L2TransactionHash: common.Hash{},
			L1TxOrigin:        common.Hash{},
			Status:            0,
			FromAddress:       batchDepositERC1155Event.From,
			ToAddress:         batchDepositERC1155Event.To,
			L1TokenAddress:    batchDepositERC1155Event.L1Token,
			L2TokenAddress:    batchDepositERC1155Event.L2Token,
			ETHAmount:         big.NewInt(0),
			ERC20Amount:       big.NewInt(0),
			GasLimit:          big.NewInt(0),
			Timestamp:         int64(depositTx.Timestamp),
			AssetType:         int64(common2.ERC1155),
			TokenIds:          utils.ConvertBigIntArrayToString(batchDepositERC1155Event.TokenIDs),
			TokenAmounts:      utils.ConvertBigIntArrayToString(batchDepositERC1155Event.TokenAmounts),
		}
	}
	if len(l1ToL2s) > 0 {
		if err := db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s); err != nil {
			return err
		}
	}
	return nil
}

func L1RelayedMessageEvent(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1RelayedMessageSig := abi.L1RelayedMessageEventSignature
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1RelayedMessageSig}
	l1RelayedMessageEvents, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l1BridgeMsgs := make([]l1_l2.L1BridgeMessage, len(l1RelayedMessageEvents))

	for i := range l1RelayedMessageEvents {
		depositTx := l1RelayedMessageEvents[i]
		rlpLog := *depositTx.RLPLog
		l1RelayedMessageEvent := abi.L1RelayedMessageEvent{}
		unpackErr := utils.UnpackLog(abi.L1ScrollMessengerABI, &l1RelayedMessageEvent, "RelayedMessage", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		// update l2 to l1 Set l1_tx_hash by msg_hash
		if err := db.L2ToL1.UpdateL2ToL1L1TxHashByMsgHash(
			worker.L2ToL1{
				MsgHash:          l1RelayedMessageEvent.MessageHash,
				L1FinalizeTxHash: rlpLog.TxHash}); err != nil {
			return err
		}

		l1BridgeMsgs[i] = l1_l2.L1BridgeMessage{
			BridgeMessage: l1_l2.BridgeMessage{
				MessageHash: l1RelayedMessageEvent.MessageHash,
				Tx: l1_l2.Transaction{
					FromAddress: common.Address{},
					ToAddress:   common.Address{},
					ETHAmount:   big.NewInt(0),
					ERC20Amount: big.NewInt(0),
					Timestamp:   depositTx.Timestamp,
				},
			},
			TransactionSourceHash: l1RelayedMessageEvent.MessageHash,
		}
	}
	if len(l1BridgeMsgs) > 0 {
		if err := db.BridgeMessages.StoreL1BridgeMessages(l1BridgeMsgs); err != nil {
			return err
		}
	}
	return nil
}
