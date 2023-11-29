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

func L2WithdrawETH(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l2WithdrawETHSig := abi.L2WithdrawETHSig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l2WithdrawETHSig}
	withdrawEvents, err := db.ContractEvents.L2ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l2ToL1s := make([]worker.L2ToL1, len(withdrawEvents))
	transactionWithdrawals := make([]l1_l2.L2TransactionWithdrawal, len(withdrawEvents))
	if len(withdrawEvents) > 0 {
		log.Info("detected eth transaction deposits", "size", len(transactionWithdrawals))
	}
	for i := range withdrawEvents {
		ethDepositTx := withdrawEvents[i]
		rlpLog := *ethDepositTx.RLPLog
		depositEvent := abi.DepositETH{}
		unpackErr := utils.UnpackLog(abi.L2ETHGatewayABI, &depositEvent, "WithdrawETH", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		transactionWithdrawals[i] = l1_l2.L2TransactionWithdrawal{
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
		l2ToL1s[i] = worker.L2ToL1{
			GUID:                      uuid.New(),
			L2BlockHash:               rlpLog.BlockHash,
			L1FinalizeTxHash:          common.Hash{},
			L2TransactionHash:         rlpLog.TxHash,
			L2WithdrawTransactionHash: common.Hash{},
			L1ProveTxHash:             common.Hash{},
			Status:                    0,
			TimeLeft:                  big.NewInt(0),
			FromAddress:               depositEvent.From,
			ToAddress:                 depositEvent.To,
			L1TokenAddress:            common.Address{},
			L2TokenAddress:            common.Address{},
			ETHAmount:                 depositEvent.Amount,
			ERC20Amount:               big.NewInt(0),
			GasLimit:                  big.NewInt(0),
			Timestamp:                 int64(ethDepositTx.Timestamp),
			AssetType:                 int64(common2.ETH),
			MsgNonce:                  big.NewInt(0),
			MsgHash:                   common.Hash{},
		}
	}
	if len(l2ToL1s) > 0 {
		if err := db.BridgeTransactions.StoreL2TransactionWithdrawals(transactionWithdrawals); err != nil {
			return err
		}
		if err := db.L2ToL1.StoreL2ToL1Transactions(l2ToL1s); err != nil {
			return err
		}
	}
	return nil
}

func L2WithdrawERC20(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l2WithdrawERC20Sig := abi.L2WithdrawERC20Sig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l2WithdrawERC20Sig}
	l2WithdrawERC20Events, err := db.ContractEvents.L2ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l2ToL1s := make([]worker.L2ToL1, len(l2WithdrawERC20Events))
	transactionWithdraws := make([]l1_l2.L2TransactionWithdrawal, len(l2WithdrawERC20Events))

	for i := range l2WithdrawERC20Events {
		withdrawTx := l2WithdrawERC20Events[i]
		rlpLog := *withdrawTx.RLPLog
		withdrawErc20Event := abi.ERC20MessageEvent{}
		transactionWithdraws[i] = l1_l2.L2TransactionWithdrawal{
			Tx: l1_l2.Transaction{
				FromAddress: withdrawErc20Event.From,
				ToAddress:   withdrawErc20Event.To,
				ETHAmount:   big.NewInt(0),
				ERC20Amount: withdrawErc20Event.Amount,
				Data:        withdrawErc20Event.Data,
				Timestamp:   withdrawTx.Timestamp,
			},
			GasLimit: big.NewInt(0),
		}
		unpackErr := utils.UnpackLog(abi.L2StandardERC20GatewayABI, &withdrawErc20Event, "WithdrawERC20", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		l2ToL1s[i] = worker.L2ToL1{
			GUID:                      uuid.New(),
			L2BlockHash:               rlpLog.BlockHash,
			L1FinalizeTxHash:          common.Hash{},
			L2TransactionHash:         rlpLog.TxHash,
			L2WithdrawTransactionHash: common.Hash{},
			L1ProveTxHash:             common.Hash{},
			Status:                    0,
			TimeLeft:                  big.NewInt(0),
			FromAddress:               withdrawErc20Event.From,
			ToAddress:                 withdrawErc20Event.To,
			L1TokenAddress:            common.Address{},
			L2TokenAddress:            common.Address{},
			ETHAmount:                 withdrawErc20Event.Amount,
			ERC20Amount:               big.NewInt(0),
			GasLimit:                  big.NewInt(0),
			Timestamp:                 int64(withdrawTx.Timestamp),
			AssetType:                 int64(common2.ERC20),
			MsgNonce:                  big.NewInt(0),
			MsgHash:                   common.Hash{},
		}
	}
	if len(l2ToL1s) > 0 {
		if err := db.BridgeTransactions.StoreL2TransactionWithdrawals(transactionWithdraws); err != nil {
			return err
		}
		if err := db.L2ToL1.StoreL2ToL1Transactions(l2ToL1s); err != nil {
			return err
		}
	}
	return nil
}

func L2WithdrawERC721(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l2WithdrawERC721Sig := abi.L2WithdrawERC721Sig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l2WithdrawERC721Sig}
	l2WithdrawERC721Events, err := db.ContractEvents.L2ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l2ToL1s := make([]worker.L2ToL1, len(l2WithdrawERC721Events))

	for i := range l2WithdrawERC721Events {
		withdrawTx := l2WithdrawERC721Events[i]
		rlpLog := *withdrawTx.RLPLog
		withdrawERC721Event := abi.ERC721MessageEvent{}
		unpackErr := utils.UnpackLog(abi.L2ERC721GatewayABI, &withdrawERC721Event, "WithdrawERC721", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		l2ToL1s[i] = worker.L2ToL1{
			GUID:                      uuid.New(),
			L2BlockHash:               rlpLog.BlockHash,
			L1FinalizeTxHash:          common.Hash{},
			L2TransactionHash:         rlpLog.TxHash,
			L2WithdrawTransactionHash: common.Hash{},
			L1ProveTxHash:             common.Hash{},
			Status:                    0,
			TimeLeft:                  big.NewInt(0),
			FromAddress:               withdrawERC721Event.From,
			ToAddress:                 withdrawERC721Event.To,
			L1TokenAddress:            common.Address{},
			L2TokenAddress:            common.Address{},
			ETHAmount:                 big.NewInt(0),
			ERC20Amount:               big.NewInt(0),
			GasLimit:                  big.NewInt(0),
			TokenAmounts:              withdrawERC721Event.Amount.String(),
			TokenIds:                  withdrawERC721Event.TokenID.String(),
			Timestamp:                 int64(withdrawTx.Timestamp),
			AssetType:                 int64(common2.ERC721),
			MsgNonce:                  big.NewInt(0),
			MsgHash:                   common.Hash{},
		}
	}
	if len(l2ToL1s) > 0 {
		if err := db.L2ToL1.StoreL2ToL1Transactions(l2ToL1s); err != nil {
			return err
		}
	}
	return nil
}

func L2WithdrawERC1155(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l2WithdrawERC1155Sig := abi.L2WithdrawERC1155Sig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l2WithdrawERC1155Sig}
	l2WithdrawERC1155Events, err := db.ContractEvents.L2ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l2ToL1s := make([]worker.L2ToL1, len(l2WithdrawERC1155Events))

	for i := range l2WithdrawERC1155Events {
		withdrawTx := l2WithdrawERC1155Events[i]
		rlpLog := *withdrawTx.RLPLog
		withdrawERC1155Event := abi.ERC1155MessageEvent{}
		unpackErr := utils.UnpackLog(abi.L2ERC1155GatewayABI, &withdrawERC1155Event, "WithdrawERC1155", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		l2ToL1s[i] = worker.L2ToL1{
			GUID:                      uuid.New(),
			L2BlockHash:               rlpLog.BlockHash,
			L1FinalizeTxHash:          common.Hash{},
			L2TransactionHash:         rlpLog.TxHash,
			L2WithdrawTransactionHash: common.Hash{},
			L1ProveTxHash:             common.Hash{},
			Status:                    0,
			TimeLeft:                  big.NewInt(0),
			FromAddress:               withdrawERC1155Event.From,
			ToAddress:                 withdrawERC1155Event.To,
			L1TokenAddress:            common.Address{},
			L2TokenAddress:            common.Address{},
			ETHAmount:                 big.NewInt(0),
			ERC20Amount:               big.NewInt(0),
			GasLimit:                  big.NewInt(0),
			TokenAmounts:              withdrawERC1155Event.Amount.String(),
			TokenIds:                  withdrawERC1155Event.TokenID.String(),
			Timestamp:                 int64(withdrawTx.Timestamp),
			AssetType:                 int64(common2.ERC1155),
			MsgNonce:                  big.NewInt(0),
			MsgHash:                   common.Hash{},
		}
	}
	if len(l2ToL1s) > 0 {
		if err := db.L2ToL1.StoreL2ToL1Transactions(l2ToL1s); err != nil {
			return err
		}
	}
	return nil
}

func L2BatchWithdrawERC721(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l2BatchWithdrawERC721Sig := abi.L2BatchWithdrawERC721Sig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l2BatchWithdrawERC721Sig}
	l2BatchWithdrawERC721Events, err := db.ContractEvents.L2ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l2ToL1s := make([]worker.L2ToL1, len(l2BatchWithdrawERC721Events))

	for i := range l2BatchWithdrawERC721Events {
		withdrawTx := l2BatchWithdrawERC721Events[i]
		rlpLog := *withdrawTx.RLPLog
		batchWithdrawERC721Event := abi.BatchERC721MessageEvent{}
		unpackErr := utils.UnpackLog(abi.L2ERC721GatewayABI, &batchWithdrawERC721Event, "BatchWithdrawERC721", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		l2ToL1s[i] = worker.L2ToL1{
			GUID:                      uuid.New(),
			L2BlockHash:               rlpLog.BlockHash,
			L1FinalizeTxHash:          common.Hash{},
			L2TransactionHash:         rlpLog.TxHash,
			L2WithdrawTransactionHash: common.Hash{},
			L1ProveTxHash:             common.Hash{},
			Status:                    0,
			TimeLeft:                  big.NewInt(0),
			FromAddress:               batchWithdrawERC721Event.From,
			ToAddress:                 batchWithdrawERC721Event.To,
			L1TokenAddress:            common.Address{},
			L2TokenAddress:            common.Address{},
			ETHAmount:                 big.NewInt(0),
			ERC20Amount:               big.NewInt(0),
			GasLimit:                  big.NewInt(0),
			TokenAmounts:              "1",
			TokenIds:                  utils.ConvertBigIntArrayToString(batchWithdrawERC721Event.TokenIDs),
			Timestamp:                 int64(withdrawTx.Timestamp),
			AssetType:                 int64(common2.ERC721),
			MsgNonce:                  big.NewInt(0),
			MsgHash:                   common.Hash{},
		}
	}
	if len(l2ToL1s) > 0 {
		if err := db.L2ToL1.StoreL2ToL1Transactions(l2ToL1s); err != nil {
			return err
		}
	}
	return nil
}

func L2BatchWithdrawERC1155(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l2BatchWithdrawERC1155Sig := abi.L2BatchWithdrawERC1155Sig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l2BatchWithdrawERC1155Sig}
	l2BatchWithdrawERC1155Events, err := db.ContractEvents.L2ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l2ToL1s := make([]worker.L2ToL1, len(l2BatchWithdrawERC1155Events))

	for i := range l2BatchWithdrawERC1155Events {
		withdrawTx := l2BatchWithdrawERC1155Events[i]
		rlpLog := *withdrawTx.RLPLog
		batchWithdrawERC1155Event := abi.BatchERC1155MessageEvent{}
		unpackErr := utils.UnpackLog(abi.L2ERC1155GatewayABI, &batchWithdrawERC1155Event, "BatchWithdrawERC1155", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		l2ToL1s[i] = worker.L2ToL1{
			GUID:                      uuid.New(),
			L2BlockHash:               rlpLog.BlockHash,
			L1FinalizeTxHash:          common.Hash{},
			L2TransactionHash:         rlpLog.TxHash,
			L2WithdrawTransactionHash: common.Hash{},
			L1ProveTxHash:             common.Hash{},
			Status:                    0,
			TimeLeft:                  big.NewInt(0),
			FromAddress:               batchWithdrawERC1155Event.From,
			ToAddress:                 batchWithdrawERC1155Event.To,
			L1TokenAddress:            common.Address{},
			L2TokenAddress:            common.Address{},
			ETHAmount:                 big.NewInt(0),
			ERC20Amount:               big.NewInt(0),
			GasLimit:                  big.NewInt(0),
			TokenAmounts:              utils.ConvertBigIntArrayToString(batchWithdrawERC1155Event.TokenAmounts),
			TokenIds:                  utils.ConvertBigIntArrayToString(batchWithdrawERC1155Event.TokenIDs),
			Timestamp:                 int64(withdrawTx.Timestamp),
			AssetType:                 int64(common2.ERC721),
			MsgNonce:                  big.NewInt(0),
			MsgHash:                   common.Hash{},
		}
	}
	if len(l2ToL1s) > 0 {
		if err := db.L2ToL1.StoreL2ToL1Transactions(l2ToL1s); err != nil {
			return err
		}
	}
	return nil
}

func L2SentMessageEvent(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l2SentMessageEventSignature := abi.L2SentMessageEventSignature
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l2SentMessageEventSignature}
	sentMessageEvents, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l2BridgeMsgs := make([]l1_l2.L2BridgeMessage, len(sentMessageEvents))
	for i := range sentMessageEvents {
		depositTx := sentMessageEvents[i]
		rlpLog := *depositTx.RLPLog
		sentMessageEvent := abi.L2SentMessageEvent{}
		unpackErr := utils.UnpackLog(abi.L2ScrollMessengerABI, &sentMessageEvent, "SentMessage", rlpLog)
		if unpackErr != nil {
			log.Warn("Failed to unpack SentMessage event", "err", err)
			return unpackErr
		}
		// compute msgHash
		msgHash := utils.ComputeMessageHash(sentMessageEvent.Sender, sentMessageEvent.Target,
			sentMessageEvent.Value, sentMessageEvent.MessageNonce, sentMessageEvent.Message)
		// update l1tol2  set msgHash by txHash
		if err := db.L2ToL1.UpdateL2ToL1MsgHashByL2TxHash(worker.L2ToL1{L2TransactionHash: rlpLog.TxHash, MsgHash: msgHash}); err != nil {
			return err
		}
		relayedMessageEventGUID := uuid.New()
		l2BridgeMsgs[i] = l1_l2.L2BridgeMessage{
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
			TransactionWithdrawalHash: rlpLog.TxHash,
		}
	}
	if len(l2BridgeMsgs) > 0 {
		if err := db.BridgeMessages.StoreL2BridgeMessages(l2BridgeMsgs); err != nil {
			return err
		}
	}
	return nil
}

func L2RelayedMessageEvent(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l2RelayedMessageSig := abi.L2RelayedMessageEventSignature
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l2RelayedMessageSig}
	l2RelayedMessageEvents, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l2BridgeMsgs := make([]l1_l2.L2BridgeMessage, len(l2RelayedMessageEvents))

	for i := range l2RelayedMessageEvents {
		depositTx := l2RelayedMessageEvents[i]
		rlpLog := *depositTx.RLPLog
		l2RelayedMessageEvent := abi.L2RelayedMessageEvent{}
		unpackErr := utils.UnpackLog(abi.L2ScrollMessengerABI, &l2RelayedMessageEvent, "RelayedMessage", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		// update l2 to l1 Set l1_tx_hash by msg_hash
		if err := db.L1ToL2.UpdateL1ToL2L2TxHashByMsgHash(
			worker.L1ToL2{
				MsgHash:           l2RelayedMessageEvent.MessageHash,
				L2TransactionHash: rlpLog.TxHash}); err != nil {
			return err
		}

		l2BridgeMsgs[i] = l1_l2.L2BridgeMessage{
			BridgeMessage: l1_l2.BridgeMessage{
				MessageHash: l2RelayedMessageEvent.MessageHash,
				Tx: l1_l2.Transaction{
					FromAddress: common.Address{},
					ToAddress:   common.Address{},
					ETHAmount:   big.NewInt(0),
					ERC20Amount: big.NewInt(0),
					Timestamp:   depositTx.Timestamp,
				},
			},
			TransactionWithdrawalHash: l2RelayedMessageEvent.MessageHash,
		}
	}
	if len(l2BridgeMsgs) > 0 {
		if err := db.BridgeMessages.StoreL2BridgeMessages(l2BridgeMsgs); err != nil {
			return err
		}
	}
	return nil
}
