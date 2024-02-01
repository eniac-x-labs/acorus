package bridge

import (
	"encoding/json"
	"fmt"
	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/database/relation"

	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/linea/abi"
	"github.com/cornerstone-labs/acorus/event/linea/utils"
)

func L2SentMessageEvent(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := *event.RLPLog
	eventABI := abi.L2MessageEventABI
	callDataABI := abi.L2MessageCallDataABI
	callDataFuncName := abi.L2MessageCallDataFuncName
	decodeLog, decodeErr := utils.DecodeLog(eventABI, callDataABI, "MessageSent", callDataFuncName, rlpLog)
	if decodeErr != nil {
		log.Println("Failed to unpack SentMessage event", "err", decodeErr)
		return decodeErr
	}
	ethAmount := decodeLog["_value"].(*big.Int)
	l2ToL1 := worker.L2ToL1{
		L2BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		L2TransactionHash: rlpLog.TxHash,
		Status:            0,
		FromAddress:       decodeLog["_from"].(common.Address),
		ToAddress:         decodeLog["_to"].(common.Address),
		L1TokenAddress:    common.Address{},
		L2TokenAddress:    common.Address{},
		ETHAmount:         ethAmount,
		Timestamp:         int64(event.Timestamp),
		MessageHash:       decodeLog["_messageHash"].(common.Hash),
	}
	if ethAmount.Cmp(big.NewInt(0)) == 0 {
		l2ToL1.AssetType = int64(common2.ETH)
		l2ToL1.GasLimit = decodeLog["_fee"].(*big.Int)
	} else {
		l2ToL1.AssetType = int64(common2.ERC20)
		l2ToL1.TokenAmounts = decodeLog["_amount"].(*big.Int).String()
		l2ToL1.GasLimit = big.NewInt(0)
		l2ToL1.L1TokenAddress = decodeLog["_nativeToken"].(common.Address)
		l2ToL1.L2TokenAddress = l2ToL1.L1TokenAddress
	}

	marshal, unpackErr := json.Marshal(l2ToL1)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelation{
		TxHash:          rlpLog.TxHash,
		LayerType:       global_const.LayerTypeTwo,
		MsgHash:         l2ToL1.MessageHash,
		Data:            string(marshal),
		MsgHashRelation: true,
	}
	saveErr := db.MsgSentRelation.MsgSentRelationStore(msgSent, chainId)
	return saveErr
}

func L2ClaimedMessageEvent(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := *event.RLPLog
	decodeLog, unpackErr := utils.DecodeLog(abi.L2MessageEventABI, nil,
		"MessageClaimed", "", rlpLog)
	if unpackErr != nil {
		log.Println("Failed to unpack SentMessage event", "err", unpackErr)
		return unpackErr
	}
	fmt.Println(decodeLog)
	messageHash := decodeLog["_messageHash"].(common.Hash)
	txHash := rlpLog.TxHash
	relayRelation := relation.RelayRelation{
		TxHash:      txHash,
		BlockNumber: big.NewInt(int64(rlpLog.BlockNumber)),
		MsgHash:     messageHash,
	}
	err := db.RelayRelation.RelayRelationStore(relayRelation, chainId)
	return err
}
