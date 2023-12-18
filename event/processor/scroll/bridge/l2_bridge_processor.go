package bridge

import (
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/processor/scroll/abi"
	"github.com/cornerstone-labs/acorus/event/processor/scroll/utils"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
)

func L2WithdrawETH(event event.L2ContractEvent) (*worker.L2ToL1, error) {
	rlpLog := *event.RLPLog
	depositEvent := abi.DepositETH{}
	unpackErr := utils.UnpackLog(abi.L2ETHGatewayABI, &depositEvent, "WithdrawETH", rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L2ToL1{
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
		Timestamp:                 int64(event.Timestamp),
		AssetType:                 int64(common2.ETH),
		MsgNonce:                  big.NewInt(0),
		MsgHash:                   common.Hash{},
	}, nil

}

func L2WithdrawERC20(event event.L2ContractEvent) (*worker.L2ToL1, error) {
	rlpLog := *event.RLPLog
	withdrawErc20Event := abi.ERC20MessageEvent{}

	unpackErr := utils.UnpackLog(abi.L2StandardERC20GatewayABI, &withdrawErc20Event, "WithdrawERC20", rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L2ToL1{
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
		Timestamp:                 int64(event.Timestamp),
		AssetType:                 int64(common2.ERC20),
		MsgNonce:                  big.NewInt(0),
		MsgHash:                   common.Hash{},
	}, nil
}

func L2WithdrawERC721(event event.L2ContractEvent) (*worker.L2ToL1, error) {
	rlpLog := *event.RLPLog
	withdrawERC721Event := abi.ERC721MessageEvent{}
	unpackErr := utils.UnpackLog(abi.L2ERC721GatewayABI, &withdrawERC721Event, "WithdrawERC721", rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L2ToL1{
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
		Timestamp:                 int64(event.Timestamp),
		AssetType:                 int64(common2.ERC721),
		MsgNonce:                  big.NewInt(0),
		MsgHash:                   common.Hash{},
	}, nil

}

func L2WithdrawERC1155(event event.L2ContractEvent) (*worker.L2ToL1, error) {
	rlpLog := *event.RLPLog
	withdrawERC1155Event := abi.ERC1155MessageEvent{}
	unpackErr := utils.UnpackLog(abi.L2ERC1155GatewayABI, &withdrawERC1155Event, "WithdrawERC1155", rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L2ToL1{
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
		Timestamp:                 int64(event.Timestamp),
		AssetType:                 int64(common2.ERC1155),
		MsgNonce:                  big.NewInt(0),
		MsgHash:                   common.Hash{},
	}, nil
}

func L2BatchWithdrawERC721(event event.L2ContractEvent) (*worker.L2ToL1, error) {
	rlpLog := *event.RLPLog
	batchWithdrawERC721Event := abi.BatchERC721MessageEvent{}
	unpackErr := utils.UnpackLog(abi.L2ERC721GatewayABI, &batchWithdrawERC721Event, "BatchWithdrawERC721", rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L2ToL1{
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
		Timestamp:                 int64(event.Timestamp),
		AssetType:                 int64(common2.ERC721),
		MsgNonce:                  big.NewInt(0),
		MsgHash:                   common.Hash{},
	}, nil
}

func L2BatchWithdrawERC1155(event event.L2ContractEvent) (*worker.L2ToL1, error) {
	rlpLog := *event.RLPLog
	batchWithdrawERC1155Event := abi.BatchERC1155MessageEvent{}
	unpackErr := utils.UnpackLog(abi.L2ERC1155GatewayABI, &batchWithdrawERC1155Event, "BatchWithdrawERC1155", rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L2ToL1{
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
		Timestamp:                 int64(event.Timestamp),
		AssetType:                 int64(common2.ERC721),
		MsgNonce:                  big.NewInt(0),
		MsgHash:                   common.Hash{},
	}, nil

}

func L2SentMessageEvent(event event.L2ContractEvent, db *database.DB) (*worker.L2ToL1, error) {
	rlpLog := *event.RLPLog
	sentMessageEvent := abi.L2SentMessageEvent{}
	unpackErr := utils.UnpackLog(abi.L2ScrollMessengerABI, &sentMessageEvent, "SentMessage", rlpLog)
	if unpackErr != nil {
		log.Warn("Failed to unpack SentMessage event", "err", unpackErr)
		return nil, unpackErr
	}
	// compute msgHash
	msgHash := utils.ComputeMessageHash(sentMessageEvent.Sender, sentMessageEvent.Target,
		sentMessageEvent.Value, sentMessageEvent.MessageNonce, sentMessageEvent.Message)
	// update l1tol2  set msgHash by txHash
	if err := db.L2ToL1.UpdateL2ToL1MsgHashByL2TxHash(worker.L2ToL1{L2TransactionHash: rlpLog.TxHash, MsgHash: msgHash}); err != nil {
		return nil, err
	}
	return nil, nil
}

func L2RelayedMessageEvent(event event.L2ContractEvent, db *database.DB) (*worker.L2ToL1, error) {
	rlpLog := *event.RLPLog
	l2RelayedMessageEvent := abi.L2RelayedMessageEvent{}
	unpackErr := utils.UnpackLog(abi.L2ScrollMessengerABI, &l2RelayedMessageEvent, "RelayedMessage", rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	// update l2 to l1 Set l1_tx_hash by msg_hash
	if err := db.L1ToL2.UpdateL1ToL2L2TxHashByMsgHash(
		worker.L1ToL2{
			MsgHash:           l2RelayedMessageEvent.MessageHash,
			L2TransactionHash: rlpLog.TxHash}); err != nil {
		return nil, err
	}

	return nil, nil
}
