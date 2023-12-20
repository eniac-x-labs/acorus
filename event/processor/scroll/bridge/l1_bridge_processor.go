package bridge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"

	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/processor/scroll/abi"
	"github.com/cornerstone-labs/acorus/event/processor/scroll/utils"
)

func L1DepositETH(event event.L1ContractEvent) (*worker.L1ToL2, error) {
	rlpLog := event.RLPLog
	depositEvent := abi.DepositETH{}
	unpackErr := utils.UnpackLog(abi.L1ETHGatewayABI, &depositEvent, "DepositETH", *rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L1ToL2{
		GUID:              uuid.New(),
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
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ETH),
		MessageHash:       common.Hash{},
	}, nil

}

func L1DepositERC20(event event.L1ContractEvent) (*worker.L1ToL2, error) {

	rlpLog := event.RLPLog
	depositErc20Event := abi.ERC20MessageEvent{}
	unpackErr := utils.UnpackLog(abi.L1StandardERC20GatewayABI, &depositErc20Event, "DepositERC20", *rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L1ToL2{
		GUID:              uuid.New(),
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
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ERC20),
	}, nil

}

func L1DepositERC721(event event.L1ContractEvent) (*worker.L1ToL2, error) {

	rlpLog := event.RLPLog
	depositErc721Event := abi.ERC721MessageEvent{}

	unpackErr := utils.UnpackLog(abi.L1ERC721GatewayABI, &depositErc721Event, "DepositERC721", *rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L1ToL2{
		GUID:              uuid.New(),
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
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ERC721),
		TokenIds:          depositErc721Event.TokenID.String(),
	}, nil

}

func L1DepositERC1155(event event.L1ContractEvent) (*worker.L1ToL2, error) {
	rlpLog := event.RLPLog
	depositErc1155Event := abi.ERC1155MessageEvent{}
	unpackErr := utils.UnpackLog(abi.L1ERC1155GatewayABI, &depositErc1155Event, "DepositERC1155", *rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L1ToL2{
		GUID:              uuid.New(),
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
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ERC1155),
		TokenIds:          depositErc1155Event.TokenID.String(),
	}, nil
}

func L1BatchDepositERC721(event event.L1ContractEvent) (*worker.L1ToL2, error) {
	rlpLog := event.RLPLog
	batchDepositErc721Event := abi.BatchERC721MessageEvent{}

	unpackErr := utils.UnpackLog(abi.L1ERC721GatewayABI, &batchDepositErc721Event, "BatchDepositERC721", *rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L1ToL2{
		GUID:              uuid.New(),
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
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ERC721),
		TokenIds:          utils.ConvertBigIntArrayToString(batchDepositErc721Event.TokenIDs),
	}, nil

}

func L1BatchDepositERC1155(event event.L1ContractEvent) (*worker.L1ToL2, error) {
	rlpLog := event.RLPLog
	batchDepositERC1155Event := abi.BatchERC1155MessageEvent{}

	unpackErr := utils.UnpackLog(abi.L1ERC1155GatewayABI, &batchDepositERC1155Event, "BatchWithdrawERC1155", *rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L1ToL2{
		GUID:              uuid.New(),
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
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ERC1155),
		TokenIds:          utils.ConvertBigIntArrayToString(batchDepositERC1155Event.TokenIDs),
		TokenAmounts:      utils.ConvertBigIntArrayToString(batchDepositERC1155Event.TokenAmounts),
	}, nil

}

func L1SentMessageEvent(event event.L1ContractEvent, db *database.DB) (*worker.L1ToL2, error) {
	rlpLog := event.RLPLog
	sentMessageEvent := abi.L1SentMessageEvent{}
	unpackErr := utils.UnpackLog(abi.L1ScrollMessengerABI, &sentMessageEvent, "SentMessage", *rlpLog)
	if unpackErr != nil {
		log.Warn("Failed to unpack SentMessage event", "err", unpackErr)
		return nil, unpackErr
	}
	// compute msgHash
	msgHash := utils.ComputeMessageHash(sentMessageEvent.Sender, sentMessageEvent.Target,
		sentMessageEvent.Value, sentMessageEvent.MessageNonce, sentMessageEvent.Message)
	// update l1tol2  set msgHash by txHash
	if err := db.L1ToL2.UpdateL1ToL2MsgHashByL1TxHash(worker.L1ToL2{L1TransactionHash: rlpLog.TxHash, MessageHash: msgHash}); err != nil {
		return nil, err
	}
	return nil, nil
}

func L1RelayedMessageEvent(event event.L1ContractEvent, db *database.DB) (*worker.L1ToL2, error) {
	rlpLog := event.RLPLog
	l1RelayedMessageEvent := abi.L1RelayedMessageEvent{}
	unpackErr := utils.UnpackLog(abi.L1ScrollMessengerABI, &l1RelayedMessageEvent, "RelayedMessage", *rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	// update l2 to l1 Set l1_tx_hash by msg_hash
	if err := db.L2ToL1.UpdateL2ToL1L1TxHashByMsgHash(
		worker.L2ToL1{
			MessageHash:      l1RelayedMessageEvent.MessageHash,
			L1FinalizeTxHash: rlpLog.TxHash}); err != nil {
		return nil, err
	}
	return nil, nil
}
