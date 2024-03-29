package bridge

import (
	"encoding/json"
	"fmt"
	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/relation"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/linea/abi"
	"github.com/cornerstone-labs/acorus/event/linea/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

func L1SentMessageEvent(chainId string, event event.ContractEvent, db *database.DB) error {

	rlpLog := *event.RLPLog
	eventABI := abi.L1MessageEventABI
	callDataABI := abi.L1MessageCallDataABI
	callDataFuncName := abi.L1MessageCallDataFuncName
	decodeLog, decodeErr := utils.DecodeLog(eventABI, callDataABI, "MessageSent", callDataFuncName, rlpLog)
	if decodeErr != nil {
		log.Error("Failed to unpack SentMessage event", "err", decodeErr)
		return decodeErr
	}
	fmt.Println(decodeLog)
	ethAmount := decodeLog["_value"].(*big.Int)
	l1ToL2 := worker.L1ToL2{
		QueueIndex:        nil,
		L1BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		L1TransactionHash: rlpLog.TxHash,

		L2TransactionHash: common.Hash{},
		L1TxOrigin:        common.Address{},
		Status:            0,
		FromAddress:       decodeLog["_from"].(common.Address),
		ToAddress:         decodeLog["_to"].(common.Address),
		L1TokenAddress:    common.Address{},
		L2TokenAddress:    common.Address{},
		ETHAmount:         ethAmount,
		GasLimit:          decodeLog["relayerFee"].(*big.Int),
		Timestamp:         int64(event.Timestamp),
		MessageHash:       decodeLog["_messageHash"].(common.Hash),
	}
	if ethAmount.Cmp(big.NewInt(0)) == 0 {
		l1ToL2.AssetType = int64(common2.ETH)
		l1ToL2.GasLimit = decodeLog["_fee"].(*big.Int)
	} else {
		l1ToL2.AssetType = int64(common2.ERC20)
		l1ToL2.TokenAmounts = decodeLog["amount"].(*big.Int).String()
		l1ToL2.L1TokenAddress = common.Address{}
		l1ToL2.L2TokenAddress = l1ToL2.L1TokenAddress
	}

	marshal, unpackErr := json.Marshal(l1ToL2)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelationStruct{
		TxHash:          rlpLog.TxHash,
		LayerType:       global_const.LayerTypeOne,
		Data:            string(marshal),
		MsgHash:         l1ToL2.MessageHash,
		MsgHashRelation: true,
	}
	saveErr := db.MsgSentRelationD.MsgSentRelationStore(msgSent, chainId)
	return saveErr
}

func L1ClaimedMessageEvent(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := *event.RLPLog
	decodeLog, unpackErr := utils.DecodeLog(abi.L1MessageEventABI, nil,
		"MessageClaimed", "", rlpLog)
	if unpackErr != nil {
		log.Error("Failed to unpack SentMessage event", "err", unpackErr)
		return unpackErr
	}
	messageHash := decodeLog["_messageHash"].(common.Hash)
	txHash := rlpLog.TxHash
	relayRelation := relation.RelayRelation{
		TxHash:      txHash,
		BlockNumber: big.NewInt(int64(rlpLog.BlockNumber)),
		MsgHash:     messageHash,
	}
	err := db.RelayRelationD.RelayRelationStore(relayRelation, chainId)
	return err
}
