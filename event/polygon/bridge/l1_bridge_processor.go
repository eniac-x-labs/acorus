package bridge

import (
	"math/big"

	"github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/oldpolygonzkevmbridge"
	"github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/polygonzkevmbridge"
	"github.com/ethereum/go-ethereum/common"

	common3 "github.com/cornerstone-labs/acorus/common"
	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/relation"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/polygon/utils"
)

func L1Deposit(chainId string, polygonBridge *polygonzkevmbridge.Polygonzkevmbridge,
	event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	d, unpackErr := polygonBridge.ParseBridgeEvent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	l1ToL2 := &worker.L1ToL2{
		QueueIndex:        nil,
		L1TransactionHash: rlpLog.TxHash,
		L2TransactionHash: common.Hash{},
		L1TxOrigin:        common.Address{},
		L1BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		Status:            common3.L1ToL2Pending,
		FromAddress:       d.DestinationAddress,
		ToAddress:         d.DestinationAddress,
		L1TokenAddress:    d.OriginAddress,
		L2TokenAddress:    common.Address{},
		ETHAmount:         big.NewInt(0),
		GasLimit:          big.NewInt(0),
		Timestamp:         int64(event.Timestamp),
		MessageHash:       common.BigToHash(big.NewInt(int64(d.DepositCount))),
	}

	if d.OriginAddress.String() == utils.L1Eth {
		l1ToL2.ETHAmount = d.Amount
		l1ToL2.AssetType = int64(common2.ETH)
	} else {
		l1ToL2.TokenAmounts = d.Amount.String()
		l1ToL2.AssetType = int64(common2.ERC20)
	}
	return nil
}

func L1WithdrawClaimedOld(chainId string, polygonBridge *oldpolygonzkevmbridge.Oldpolygonzkevmbridge,
	event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	c, unpackErr := polygonBridge.ParseClaimEvent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	index := c.Index
	indexBig := new(big.Int).SetUint64(uint64(index))
	relayRelation := relation.RelayRelation{
		TxHash:      rlpLog.TxHash,
		LayerType:   global_const.LayerTypeOne,
		BlockNumber: big.NewInt(int64(rlpLog.BlockNumber)),
		MsgHash:     common.BigToHash(indexBig),
	}
	err := db.RelayRelationD.RelayRelationStore(relayRelation, chainId)
	return err
}

func L1WithdrawClaimed(chainId string, polygonBridge *polygonzkevmbridge.Polygonzkevmbridge,
	event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	c, unpackErr := polygonBridge.ParseClaimEvent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	index := c.GlobalIndex
	_, _, localRootIndex, decodeErr := utils.DecodeGlobalIndex(index)

	if decodeErr != nil {
		return decodeErr
	}
	relayRelation := relation.RelayRelation{
		TxHash:      rlpLog.TxHash,
		LayerType:   global_const.LayerTypeOne,
		BlockNumber: big.NewInt(int64(rlpLog.BlockNumber)),
		MsgHash:     common.BigToHash(localRootIndex),
	}
	err := db.RelayRelationD.RelayRelationStore(relayRelation, chainId)
	return err
}
