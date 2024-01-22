package bridge

import (
	"fmt"
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
	decodeLog, decodeErr := utils.DecodeLog(abi.L1MessageEventABI, "MessageSent", rlpLog)
	if decodeErr != nil {
		log.Println("Failed to unpack SentMessage event", "err", decodeErr)
		return nil, decodeErr
	}
	fmt.Println(decodeLog)

	return &worker.L1ToL2{
		QueueIndex:        nil,
		L1BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		L1TransactionHash: rlpLog.TxHash,
		L2TransactionHash: common.Hash{},
		L1TxOrigin:        common.Address{},
		Status:            0,
		FromAddress:       decodeLog.From,
		ToAddress:         decodeLog.To,
		L1TokenAddress:    common.Address{},
		L2TokenAddress:    common.Address{},
		ETHAmount:         decodeLog.Amount,
		GasLimit:          big.NewInt(0),
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ETH),
		MessageHash:       common.Hash{},
	}, nil
}

func L1ClaimedMessageEvent(event event.ContractEvent) (*worker.L1ToL2, error) {

	rlpLog := *event.RLPLog
	decodeLog, unpackErr := utils.DecodeLog(abi.L1MessageEventABI, "MessageClaimed", rlpLog)
	if unpackErr != nil {
		log.Println("Failed to unpack SentMessage event", "err", unpackErr)
		return nil, unpackErr
	}
	fmt.Println(decodeLog)

	// update claimed message
	return nil, nil
}
