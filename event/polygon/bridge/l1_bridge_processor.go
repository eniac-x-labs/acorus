package bridge

import (
	"encoding/json"
	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/database/relation"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/polygon/abi"
	"github.com/cornerstone-labs/acorus/event/polygon/utils"
)

func L1Deposit(chainId string, polygonBridge *abi.Polygonzkevmbridge,
	event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	d, unpackErr := polygonBridge.ParseBridgeEvent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	l1ToL2 := worker.L1ToL2{
		QueueIndex:        nil,
		L1TransactionHash: rlpLog.TxHash,
		L2TransactionHash: common.Hash{},
		L1TxOrigin:        common.Address{},
		Status:            0,
		FromAddress:       d.DestinationAddress,
		ToAddress:         d.DestinationAddress,
		L1TokenAddress:    d.OriginAddress,
		L2TokenAddress:    common.Address{},
		ETHAmount:         big.NewInt(0),
		GasLimit:          big.NewInt(0),
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ETH),
		TokenAmounts:      "0",
		MessageHash:       common.BigToHash(big.NewInt(int64(d.DepositCount))),
	}

	if d.OriginAddress.String() == utils.L1_ETH {
		l1ToL2.ETHAmount = d.Amount
		l1ToL2.AssetType = int64(common2.ETH)
	} else {
		l1ToL2.TokenAmounts = d.Amount.String()
		l1ToL2.AssetType = int64(common2.ERC20)
	}

	marshal, unpackErr := json.Marshal(l1ToL2)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelation{
		TxHash:          rlpLog.TxHash,
		LayerType:       global_const.LayerTypeOne,
		Data:            string(marshal),
		MsgHash:         l1ToL2.MessageHash,
		MsgHashRelation: true,
	}
	saveErr := db.MsgSentRelation.MsgSentRelationStore(msgSent, chainId)
	return saveErr
}

func L1WithdrawClaimed(chainId string, polygonBridge *abi.Polygonzkevmbridge,
	event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	c, unpackErr := polygonBridge.ParseClaimEvent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	relayRelation := relation.RelayRelation{
		TxHash:      rlpLog.TxHash,
		BlockNumber: big.NewInt(int64(rlpLog.BlockNumber)),
		MsgHash:     common.BigToHash(c.GlobalIndex),
	}
	err := db.RelayRelation.RelayRelationStore(relayRelation, chainId)
	return err
}
