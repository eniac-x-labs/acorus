package bridge

import (
	"github.com/cornerstone-labs/acorus/event/scroll/abi"
	"github.com/cornerstone-labs/acorus/event/scroll/utils"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"

	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
)

func L1DepositETH(event event.ContractEvent) (*worker.L1ToL2, error) {
	rlpLog := event.RLPLog
	depositEvent := abi.DepositETH{}
	unpackErr := utils.UnpackLog(abi.L1ETHGatewayABI, &depositEvent, "DepositETH", *rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L1ToL2{
		GUID:              uuid.New(),
		QueueIndex:        nil,
		L1BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		L1TransactionHash: rlpLog.TxHash,
		L2TransactionHash: common.Hash{},
		L1TxOrigin:        common.Address{},
		Status:            0,
		FromAddress:       depositEvent.From,
		ToAddress:         depositEvent.To,
		L1TokenAddress:    common.Address{},
		L2TokenAddress:    common.Address{},
		ETHAmount:         depositEvent.Amount,
		GasLimit:          big.NewInt(0),
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ETH),
		MessageHash:       common.Hash{},
	}, nil

}

func L1DepositERC20(event event.ContractEvent) (*worker.L1ToL2, error) {

	rlpLog := event.RLPLog
	depositErc20Event := abi.ERC20MessageEvent{}
	unpackErr := utils.UnpackLog(abi.L1StandardERC20GatewayABI, &depositErc20Event, "DepositERC20", *rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	amounts := make([]*big.Int, 0)
	amounts = append(amounts, depositErc20Event.Amount)
	return &worker.L1ToL2{
		GUID:              uuid.New(),
		QueueIndex:        nil,
		L1BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		L1TransactionHash: rlpLog.TxHash,
		L2TransactionHash: common.Hash{},
		L1TxOrigin:        common.Address{},
		Status:            0,
		FromAddress:       depositErc20Event.From,
		ToAddress:         depositErc20Event.To,
		L1TokenAddress:    depositErc20Event.L1Token,
		L2TokenAddress:    depositErc20Event.L2Token,
		ETHAmount:         big.NewInt(0),
		GasLimit:          big.NewInt(0),
		Timestamp:         int64(event.Timestamp),
		TokenAmounts:      utils.ConvertBigIntArrayToString(amounts),
		AssetType:         int64(common2.ERC20),
	}, nil

}

func L1DepositERC721(event event.ContractEvent) (*worker.L1ToL2, error) {

	rlpLog := event.RLPLog
	depositErc721Event := abi.ERC721MessageEvent{}

	unpackErr := utils.UnpackLog(abi.L1ERC721GatewayABI, &depositErc721Event, "DepositERC721", *rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}

	return &worker.L1ToL2{
		GUID:              uuid.New(),
		QueueIndex:        nil,
		L1BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		L1TransactionHash: rlpLog.TxHash,
		L2TransactionHash: common.Hash{},
		L1TxOrigin:        common.Address{},
		Status:            0,
		FromAddress:       depositErc721Event.From,
		ToAddress:         depositErc721Event.To,
		L1TokenAddress:    depositErc721Event.L1Token,
		L2TokenAddress:    depositErc721Event.L2Token,
		ETHAmount:         depositErc721Event.Amount,
		GasLimit:          big.NewInt(0),
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ERC721),
		TokenIds:          depositErc721Event.TokenID.String(),
		TokenAmounts:      depositErc721Event.Amount.String(),
	}, nil

}

func L1DepositERC1155(event event.ContractEvent) (*worker.L1ToL2, error) {
	rlpLog := event.RLPLog
	depositErc1155Event := abi.ERC1155MessageEvent{}
	unpackErr := utils.UnpackLog(abi.L1ERC1155GatewayABI, &depositErc1155Event, "DepositERC1155", *rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L1ToL2{
		GUID:              uuid.New(),
		QueueIndex:        nil,
		L1BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		L1TransactionHash: rlpLog.TxHash,
		L2TransactionHash: common.Hash{},
		L1TxOrigin:        common.Address{},
		Status:            0,
		FromAddress:       depositErc1155Event.From,
		ToAddress:         depositErc1155Event.To,
		L1TokenAddress:    depositErc1155Event.L1Token,
		L2TokenAddress:    depositErc1155Event.L2Token,
		ETHAmount:         depositErc1155Event.Amount,
		GasLimit:          big.NewInt(0),
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ERC1155),
		TokenIds:          depositErc1155Event.TokenID.String(),
		TokenAmounts:      depositErc1155Event.Amount.String(),
	}, nil
}

func L1BatchDepositERC721(event event.ContractEvent) (*worker.L1ToL2, error) {
	rlpLog := event.RLPLog
	batchDepositErc721Event := abi.BatchERC721MessageEvent{}

	unpackErr := utils.UnpackLog(abi.L1ERC721GatewayABI, &batchDepositErc721Event, "BatchDepositERC721", *rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L1ToL2{
		GUID:              uuid.New(),
		QueueIndex:        nil,
		L1BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		L1TransactionHash: rlpLog.TxHash,
		L2TransactionHash: common.Hash{},
		L1TxOrigin:        common.Address{},
		Status:            0,
		FromAddress:       batchDepositErc721Event.From,
		ToAddress:         batchDepositErc721Event.To,
		L1TokenAddress:    batchDepositErc721Event.L1Token,
		L2TokenAddress:    batchDepositErc721Event.L2Token,
		ETHAmount:         big.NewInt(0),
		GasLimit:          big.NewInt(0),
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ERC721),
		TokenAmounts:      "1",
		TokenIds:          utils.ConvertBigIntArrayToString(batchDepositErc721Event.TokenIDs),
	}, nil

}

func L1BatchDepositERC1155(event event.ContractEvent) (*worker.L1ToL2, error) {
	rlpLog := event.RLPLog
	batchDepositERC1155Event := abi.BatchERC1155MessageEvent{}

	unpackErr := utils.UnpackLog(abi.L1ERC1155GatewayABI, &batchDepositERC1155Event, "BatchWithdrawERC1155", *rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	return &worker.L1ToL2{
		GUID:              uuid.New(),
		QueueIndex:        nil,
		L1BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		L1TransactionHash: rlpLog.TxHash,
		L2TransactionHash: common.Hash{},
		L1TxOrigin:        common.Address{},
		Status:            0,
		FromAddress:       batchDepositERC1155Event.From,
		ToAddress:         batchDepositERC1155Event.To,
		L1TokenAddress:    batchDepositERC1155Event.L1Token,
		L2TokenAddress:    batchDepositERC1155Event.L2Token,
		ETHAmount:         big.NewInt(0),
		GasLimit:          big.NewInt(0),
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ERC1155),
		TokenIds:          utils.ConvertBigIntArrayToString(batchDepositERC1155Event.TokenIDs),
		TokenAmounts:      utils.ConvertBigIntArrayToString(batchDepositERC1155Event.TokenAmounts),
	}, nil

}

func L1SentMessageEvent(chainId string, event event.ContractEvent, db *database.DB) (*worker.L1ToL2, error) {
	rlpLog := event.RLPLog
	sentMessageEvent := abi.L1SentMessageEvent{}
	unpackErr := utils.UnpackLog(abi.L1ScrollMessengerABI, &sentMessageEvent, "SentMessage", *rlpLog)
	if unpackErr != nil {
		log.Println("Failed to unpack SentMessage event", "err", unpackErr)
		return nil, unpackErr
	}
	// compute msgHash
	msgHash := utils.ComputeMessageHash(sentMessageEvent.Sender, sentMessageEvent.Target,
		sentMessageEvent.Value, sentMessageEvent.MessageNonce, sentMessageEvent.Message)
	// update l1tol2  set msgHash by txHash
	if err := db.L1ToL2.UpdateL1ToL2MsgHashByL1TxHash(chainId, worker.L1ToL2{L1TransactionHash: rlpLog.TxHash, MessageHash: msgHash}); err != nil {
		return nil, err
	}
	return nil, nil
}

func L1RelayedMessageEvent(chainId string, event event.ContractEvent, db *database.DB) (*worker.L1ToL2, error) {
	rlpLog := event.RLPLog
	l1RelayedMessageEvent := abi.L1RelayedMessageEvent{}
	unpackErr := utils.UnpackLog(abi.L1ScrollMessengerABI, &l1RelayedMessageEvent, "RelayedMessage", *rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	// update l2 to l1 Set l1_tx_hash by msg_hash
	if err := db.L2ToL1.UpdateL2ToL1L1TxHashByMsgHash(
		chainId,
		worker.L2ToL1{
			L1BlockNumber:    big.NewInt(int64(rlpLog.BlockNumber)),
			MessageHash:      l1RelayedMessageEvent.MessageHash,
			L1FinalizeTxHash: rlpLog.TxHash}); err != nil {
		return nil, err
	}
	return nil, nil
}
