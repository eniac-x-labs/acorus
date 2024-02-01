package bridge

import (
	"encoding/json"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/relation"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/scroll/abi"
	"github.com/cornerstone-labs/acorus/event/scroll/utils"
)

func L1DepositETH(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	depositEvent := abi.DepositETH{}
	unpackErr := utils.UnpackLog(abi.L1ETHGatewayABI, &depositEvent, "DepositETH", *rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	l1L2Data := &worker.L1ToL2{

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
	}
	marshal, unpackErr := json.Marshal(l1L2Data)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelation{
		TxHash:    rlpLog.TxHash,
		LayerType: global_const.LayerTypeOne,
		Data:      string(marshal),
	}
	saveErr := db.MsgSentRelation.MsgSentRelationStore(msgSent, chainId)
	return saveErr

}

func L1DepositERC20(chainId string, event event.ContractEvent, db *database.DB) error {

	rlpLog := event.RLPLog
	depositErc20Event := abi.ERC20MessageEvent{}
	unpackErr := utils.UnpackLog(abi.L1StandardERC20GatewayABI, &depositErc20Event, "DepositERC20", *rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	amounts := make([]*big.Int, 0)
	amounts = append(amounts, depositErc20Event.Amount)
	l1L2Data := &worker.L1ToL2{

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
	}

	marshal, unpackErr := json.Marshal(l1L2Data)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelation{
		TxHash:    rlpLog.TxHash,
		LayerType: global_const.LayerTypeOne,
		Data:      string(marshal),
	}
	saveErr := db.MsgSentRelation.MsgSentRelationStore(msgSent, chainId)
	return saveErr

}

func L1DepositERC721(chainId string, event event.ContractEvent, db *database.DB) error {

	rlpLog := event.RLPLog
	depositErc721Event := abi.ERC721MessageEvent{}

	unpackErr := utils.UnpackLog(abi.L1ERC721GatewayABI, &depositErc721Event, "DepositERC721", *rlpLog)
	if unpackErr != nil {
		return unpackErr
	}

	l1L2Data := &worker.L1ToL2{

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
	}
	marshal, unpackErr := json.Marshal(l1L2Data)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelation{
		TxHash:    rlpLog.TxHash,
		LayerType: global_const.LayerTypeOne,
		Data:      string(marshal),
	}
	saveErr := db.MsgSentRelation.MsgSentRelationStore(msgSent, chainId)
	return saveErr

}

func L1DepositERC1155(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	depositErc1155Event := abi.ERC1155MessageEvent{}
	unpackErr := utils.UnpackLog(abi.L1ERC1155GatewayABI, &depositErc1155Event, "DepositERC1155", *rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	l1L2Data := &worker.L1ToL2{

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
	}
	marshal, unpackErr := json.Marshal(l1L2Data)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelation{
		TxHash:    rlpLog.TxHash,
		LayerType: global_const.LayerTypeOne,
		Data:      string(marshal),
	}
	saveErr := db.MsgSentRelation.MsgSentRelationStore(msgSent, chainId)
	return saveErr
}

func L1BatchDepositERC721(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	batchDepositErc721Event := abi.BatchERC721MessageEvent{}

	unpackErr := utils.UnpackLog(abi.L1ERC721GatewayABI, &batchDepositErc721Event, "BatchDepositERC721", *rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	l1L2Data := &worker.L1ToL2{

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
	}
	marshal, unpackErr := json.Marshal(l1L2Data)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelation{
		TxHash:    rlpLog.TxHash,
		LayerType: global_const.LayerTypeOne,
		Data:      string(marshal),
	}
	saveErr := db.MsgSentRelation.MsgSentRelationStore(msgSent, chainId)
	return saveErr

}

func L1BatchDepositERC1155(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	batchDepositERC1155Event := abi.BatchERC1155MessageEvent{}

	unpackErr := utils.UnpackLog(abi.L1ERC1155GatewayABI, &batchDepositERC1155Event, "BatchWithdrawERC1155", *rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	l1L2Data := &worker.L1ToL2{
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
	}
	marshal, unpackErr := json.Marshal(l1L2Data)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelation{
		TxHash:    rlpLog.TxHash,
		LayerType: global_const.LayerTypeOne,
		Data:      string(marshal),
	}
	saveErr := db.MsgSentRelation.MsgSentRelationStore(msgSent, chainId)
	return saveErr

}

func L1SentMessageEvent(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	sentMessageEvent := abi.L1SentMessageEvent{}
	unpackErr := utils.UnpackLog(abi.L1ScrollMessengerABI, &sentMessageEvent, "SentMessage", *rlpLog)
	if unpackErr != nil {
		log.Println("Failed to unpack SentMessage event", "err", unpackErr)
		return unpackErr
	}

	msgHash := utils.ComputeMessageHash(sentMessageEvent.Sender, sentMessageEvent.Target,
		sentMessageEvent.Value, sentMessageEvent.MessageNonce, sentMessageEvent.Message)

	relayRelation := relation.MsgHashRelation{
		TxHash:  rlpLog.TxHash,
		MsgHash: msgHash,
	}

	err := db.MsgHashRelation.MsgHashRelationStore(relayRelation, chainId)
	return err
}

func L1RelayedMessageEvent(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	l1RelayedMessageEvent := abi.L1RelayedMessageEvent{}
	unpackErr := utils.UnpackLog(abi.L1ScrollMessengerABI, &l1RelayedMessageEvent, "RelayedMessage", *rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	relayRelation := relation.RelayRelation{
		TxHash:      rlpLog.TxHash,
		BlockNumber: big.NewInt(int64(rlpLog.BlockNumber)),
		MsgHash:     l1RelayedMessageEvent.MessageHash,
	}
	err := db.RelayRelation.RelayRelationStore(relayRelation, chainId)
	return err
}
