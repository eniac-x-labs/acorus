package bridge

import (
	"fmt"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/event/arb/utils"
)

func MessageDeliveredUnpack(l1ChainId, l2ChainId string, eventInfo event.ContractEvent, db *database.DB) error {

	rlpLog := eventInfo.RLPLog
	d, unpackErr := utils.IBridge.ParseMessageDelivered(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	fmt.Println(d)

	//l1ToL2 := worker.L1ToL2{
	//	QueueIndex:        nil,
	//	L1TransactionHash: rlpLog.TxHash,
	//	L2TransactionHash: common.Hash{},
	//	L1TxOrigin:        common.Address{},
	//	L1BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
	//	Status:            common3.L1ToL2Pending,
	//	FromAddress:       common.Address{},
	//	ToAddress:         d.DestinationAddress,
	//	L1TokenAddress:    d.OriginAddress,
	//	L2TokenAddress:    common.Address{},
	//	GasLimit:          big.NewInt(0),
	//	Timestamp:         int64(eventInfo.Timestamp),
	//	MessageHash:       common.BigToHash(big.NewInt(int64(d.DepositCount))),
	//	ETHAmount:         big.NewInt(0),
	//	TokenAmounts:      "0",
	//}

	return nil
}
