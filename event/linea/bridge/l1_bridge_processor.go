package bridge

import (
	"fmt"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/linea/abi"
	"github.com/cornerstone-labs/acorus/event/linea/utils"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

func L1SentMessageEvent(event event.ContractEvent) (*worker.L1ToL2, error) {

	rlpLog := *event.RLPLog
	eventABI := abi.L1MessageEventABI
	callDataABI := abi.L1MessageCallDataABI
	callDataFuncName := abi.L1MessageCallDataFuncName
	decodeLog, decodeErr := utils.DecodeLog(eventABI, callDataABI, "MessageSent", callDataFuncName, rlpLog)
	if decodeErr != nil {
		log.Println("Failed to unpack SentMessage event", "err", decodeErr)
		return nil, decodeErr
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

	return &l1ToL2, nil
}

func L1ClaimedMessageEvent(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := *event.RLPLog
	decodeLog, unpackErr := utils.DecodeLog(abi.L1MessageEventABI, nil,
		"MessageClaimed", "", rlpLog)
	if unpackErr != nil {
		log.Println("Failed to unpack SentMessage event", "err", unpackErr)
		return unpackErr
	}
	fmt.Println(decodeLog)
	messageHash := decodeLog["_messageHash"].(common.Hash)
	txHash := rlpLog.TxHash
	if err := db.L2ToL1.UpdateL2ToL1L1TxHashByMsgHash(
		chainId,
		worker.L2ToL1{
			L1BlockNumber:    big.NewInt(int64(rlpLog.BlockNumber)),
			MessageHash:      messageHash,
			L1FinalizeTxHash: txHash}); err != nil {
		return err
	}
	return nil
}
