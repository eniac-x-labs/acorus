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

func L2WithdrawETH(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := *event.RLPLog
	depositEvent := abi.DepositETH{}
	unpackErr := utils.UnpackLog(abi.L2ETHGatewayABI, &depositEvent, "WithdrawETH", rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	l1L2Data := &worker.L2ToL1{
		L2BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		L1FinalizeTxHash:  common.Hash{},
		L2TransactionHash: rlpLog.TxHash,
		L1ProveTxHash:     common.Hash{},
		Status:            0,
		TimeLeft:          big.NewInt(0),
		FromAddress:       depositEvent.From,
		ToAddress:         depositEvent.To,
		L1TokenAddress:    common.Address{},
		L2TokenAddress:    common.Address{},
		ETHAmount:         depositEvent.Amount,
		GasLimit:          big.NewInt(0),
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ETH),
		MsgNonce:          big.NewInt(0),
		MessageHash:       common.Hash{},
	}

	marshal, unpackErr := json.Marshal(l1L2Data)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelationStruct{
		TxHash:    rlpLog.TxHash,
		LayerType: global_const.LayerTypeTwo,
		Data:      string(marshal),
	}
	saveErr := db.MsgSentRelationD.MsgSentRelationStore(msgSent, chainId)
	return saveErr

}

func L2WithdrawERC20(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := *event.RLPLog
	withdrawErc20Event := abi.ERC20MessageEvent{}

	unpackErr := utils.UnpackLog(abi.L2StandardERC20GatewayABI, &withdrawErc20Event, "WithdrawERC20", rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	amounts := make([]*big.Int, 0)
	amounts = append(amounts, withdrawErc20Event.Amount)
	l1L2Data := &worker.L2ToL1{

		L1FinalizeTxHash:  common.Hash{},
		L2BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		L2TransactionHash: rlpLog.TxHash,
		L1ProveTxHash:     common.Hash{},
		Status:            0,
		TimeLeft:          big.NewInt(0),
		FromAddress:       withdrawErc20Event.From,
		ToAddress:         withdrawErc20Event.To,
		L1TokenAddress:    withdrawErc20Event.L1Token,
		L2TokenAddress:    withdrawErc20Event.L2Token,
		GasLimit:          big.NewInt(0),
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ERC20),
		MsgNonce:          big.NewInt(0),
		MessageHash:       common.Hash{},
		TokenAmounts:      utils.ConvertBigIntArrayToString(amounts),
	}
	marshal, unpackErr := json.Marshal(l1L2Data)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelationStruct{
		TxHash:    rlpLog.TxHash,
		LayerType: global_const.LayerTypeTwo,
		Data:      string(marshal),
	}
	saveErr := db.MsgSentRelationD.MsgSentRelationStore(msgSent, chainId)
	return saveErr
}

func L2WithdrawERC721(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := *event.RLPLog
	withdrawERC721Event := abi.ERC721MessageEvent{}
	unpackErr := utils.UnpackLog(abi.L2ERC721GatewayABI, &withdrawERC721Event, "WithdrawERC721", rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	l1L2Data := &worker.L2ToL1{

		L1FinalizeTxHash:  common.Hash{},
		L2TransactionHash: rlpLog.TxHash,
		L2BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		L1ProveTxHash:     common.Hash{},
		Status:            0,
		TimeLeft:          big.NewInt(0),
		FromAddress:       withdrawERC721Event.From,
		ToAddress:         withdrawERC721Event.To,
		L1TokenAddress:    withdrawERC721Event.L1Token,
		L2TokenAddress:    withdrawERC721Event.L2Token,
		ETHAmount:         big.NewInt(0),
		GasLimit:          big.NewInt(0),
		TokenAmounts:      withdrawERC721Event.Amount.String(),
		TokenIds:          withdrawERC721Event.TokenID.String(),
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ERC721),
		MsgNonce:          big.NewInt(0),
		MessageHash:       common.Hash{},
	}

	marshal, unpackErr := json.Marshal(l1L2Data)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelationStruct{
		TxHash:    rlpLog.TxHash,
		LayerType: global_const.LayerTypeTwo,
		Data:      string(marshal),
	}
	saveErr := db.MsgSentRelationD.MsgSentRelationStore(msgSent, chainId)
	return saveErr

}

func L2WithdrawERC1155(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := *event.RLPLog
	withdrawERC1155Event := abi.ERC1155MessageEvent{}
	unpackErr := utils.UnpackLog(abi.L2ERC1155GatewayABI, &withdrawERC1155Event, "WithdrawERC1155", rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	l1L2Data := &worker.L2ToL1{

		L1FinalizeTxHash:  common.Hash{},
		L2TransactionHash: rlpLog.TxHash,
		L2BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		L1ProveTxHash:     common.Hash{},
		Status:            0,
		TimeLeft:          big.NewInt(0),
		FromAddress:       withdrawERC1155Event.From,
		ToAddress:         withdrawERC1155Event.To,
		L1TokenAddress:    withdrawERC1155Event.L1Token,
		L2TokenAddress:    withdrawERC1155Event.L2Token,
		ETHAmount:         big.NewInt(0),
		GasLimit:          big.NewInt(0),
		TokenAmounts:      withdrawERC1155Event.Amount.String(),
		TokenIds:          withdrawERC1155Event.TokenID.String(),
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ERC1155),
		MsgNonce:          big.NewInt(0),
		MessageHash:       common.Hash{},
	}
	marshal, unpackErr := json.Marshal(l1L2Data)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelationStruct{
		TxHash:    rlpLog.TxHash,
		LayerType: global_const.LayerTypeTwo,
		Data:      string(marshal),
	}
	saveErr := db.MsgSentRelationD.MsgSentRelationStore(msgSent, chainId)
	return saveErr
}

func L2BatchWithdrawERC721(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := *event.RLPLog
	batchWithdrawERC721Event := abi.BatchERC721MessageEvent{}
	unpackErr := utils.UnpackLog(abi.L2ERC721GatewayABI, &batchWithdrawERC721Event, "BatchWithdrawERC721", rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	l1L2Data := &worker.L2ToL1{

		L1FinalizeTxHash:  common.Hash{},
		L2TransactionHash: rlpLog.TxHash,
		L2BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		L1ProveTxHash:     common.Hash{},
		Status:            0,
		TimeLeft:          big.NewInt(0),
		FromAddress:       batchWithdrawERC721Event.From,
		ToAddress:         batchWithdrawERC721Event.To,
		L1TokenAddress:    batchWithdrawERC721Event.L1Token,
		L2TokenAddress:    batchWithdrawERC721Event.L2Token,
		ETHAmount:         big.NewInt(0),
		GasLimit:          big.NewInt(0),
		TokenAmounts:      "1",
		TokenIds:          utils.ConvertBigIntArrayToString(batchWithdrawERC721Event.TokenIDs),
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ERC721),
		MsgNonce:          big.NewInt(0),
		MessageHash:       common.Hash{},
	}
	marshal, unpackErr := json.Marshal(l1L2Data)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelationStruct{
		TxHash:    rlpLog.TxHash,
		LayerType: global_const.LayerTypeTwo,
		Data:      string(marshal),
	}
	saveErr := db.MsgSentRelationD.MsgSentRelationStore(msgSent, chainId)
	return saveErr
}

func L2BatchWithdrawERC1155(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := *event.RLPLog
	batchWithdrawERC1155Event := abi.BatchERC1155MessageEvent{}
	unpackErr := utils.UnpackLog(abi.L2ERC1155GatewayABI, &batchWithdrawERC1155Event, "BatchWithdrawERC1155", rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	l1L2Data := &worker.L2ToL1{

		L1FinalizeTxHash:  common.Hash{},
		L2TransactionHash: rlpLog.TxHash,
		L1ProveTxHash:     common.Hash{},
		L2BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		Status:            0,
		TimeLeft:          big.NewInt(0),
		FromAddress:       batchWithdrawERC1155Event.From,
		ToAddress:         batchWithdrawERC1155Event.To,
		L1TokenAddress:    batchWithdrawERC1155Event.L1Token,
		L2TokenAddress:    batchWithdrawERC1155Event.L2Token,
		ETHAmount:         big.NewInt(0),
		GasLimit:          big.NewInt(0),
		TokenAmounts:      utils.ConvertBigIntArrayToString(batchWithdrawERC1155Event.TokenAmounts),
		TokenIds:          utils.ConvertBigIntArrayToString(batchWithdrawERC1155Event.TokenIDs),
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ERC1155),
		MsgNonce:          big.NewInt(0),
		MessageHash:       common.Hash{},
	}
	marshal, unpackErr := json.Marshal(l1L2Data)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelationStruct{
		TxHash:    rlpLog.TxHash,
		LayerType: global_const.LayerTypeTwo,
		Data:      string(marshal),
	}
	saveErr := db.MsgSentRelationD.MsgSentRelationStore(msgSent, chainId)
	return saveErr

}

func L2SentMessageEvent(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := *event.RLPLog
	sentMessageEvent := abi.L2SentMessageEvent{}
	unpackErr := utils.UnpackLog(abi.L2ScrollMessengerABI, &sentMessageEvent, "SentMessage", rlpLog)
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

	err := db.MsgHashRelationD.MsgHashRelationStore(relayRelation, chainId)
	return err
}

func L2RelayedMessageEvent(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := *event.RLPLog
	l2RelayedMessageEvent := abi.L2RelayedMessageEvent{}
	unpackErr := utils.UnpackLog(abi.L2ScrollMessengerABI, &l2RelayedMessageEvent, "RelayedMessage", rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	relayRelation := relation.RelayRelation{
		TxHash:      rlpLog.TxHash,
		BlockNumber: big.NewInt(int64(rlpLog.BlockNumber)),
		MsgHash:     l2RelayedMessageEvent.MessageHash,
	}
	err := db.RelayRelationD.RelayRelationStore(relayRelation, chainId)
	return err
}
