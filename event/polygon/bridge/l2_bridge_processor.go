package bridge

import (
	"encoding/json"

	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/relation"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/polygon/abi"
	"github.com/cornerstone-labs/acorus/event/polygon/utils"
)

func L2Withdraw(chainId string, polygonBridge *abi.Polygonzkevmbridge,
	event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	w, unpackErr := polygonBridge.ParseBridgeEvent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	l2ToL1 := worker.L2ToL1{
		L1FinalizeTxHash:  common.Hash{},
		L2TransactionHash: rlpLog.TxHash,
		L1ProveTxHash:     common.Hash{},
		Status:            0,
		TimeLeft:          big.NewInt(0),
		FromAddress:       w.DestinationAddress,
		ToAddress:         w.DestinationAddress,
		L2BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		L1TokenAddress:    common.Address{},
		L2TokenAddress:    common.Address{},
		ETHAmount:         big.NewInt(0),
		GasLimit:          big.NewInt(0),
		Timestamp:         int64(event.Timestamp),
		MsgNonce:          big.NewInt(0),
		MessageHash:       common.BigToHash(big.NewInt(int64(w.DepositCount))),
		ClaimedIndex:      int64(w.DepositCount),
	}

	if w.OriginAddress.String() == utils.L1_ETH {
		l2ToL1.ETHAmount = w.Amount
		l2ToL1.AssetType = int64(common2.ETH)
	} else {
		l2ToL1.TokenAmounts = w.Amount.String()
		l2ToL1.AssetType = int64(common2.ERC20)
	}

	marshal, unpackErr := json.Marshal(l2ToL1)
	if unpackErr != nil {
		return unpackErr
	}
	msgSent := relation.MsgSentRelationStruct{
		TxHash:          rlpLog.TxHash,
		LayerType:       global_const.LayerTypeTwo,
		Data:            string(marshal),
		MsgHash:         l2ToL1.MessageHash,
		MsgHashRelation: true,
	}
	saveErr := db.MsgSentRelationD.MsgSentRelationStore(msgSent, chainId)
	return saveErr
}

func L2WithdrawClaimed(chainId string, polygonBridge *abi.Polygonzkevmbridge,
	event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	c, unpackErr := polygonBridge.ParseClaimEvent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	index := c.GlobalIndex
	relayRelation := relation.RelayRelation{
		TxHash:      rlpLog.TxHash,
		LayerType:   global_const.LayerTypeTwo,
		BlockNumber: big.NewInt(int64(rlpLog.BlockNumber)),
		MsgHash:     common.BigToHash(index),
	}
	err := db.RelayRelationD.RelayRelationStore(relayRelation, chainId)
	return err
}
