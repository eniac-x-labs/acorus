package unpack

import (
	"github.com/cornerstone-labs/acorus/appchain/bindings"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/appchain"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/ethereum/go-ethereum/common"
)

var (
	StakeUnpack, _ = bindings.NewStakingManager(common.Address{}, nil)
)

const (
	Pending = 0
	Claim   = 1
)

func UnstakeRequestClaimed(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := StakeUnpack.ParseUnstakeRequestClaimed(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	unStakeSingle := appchain.AppChainUnStakeBatch{
		ClaimTxHash:   rlpLog.TxHash,
		Bridge:        uEvent.Bridge,
		Staker:        uEvent.Staker,
		BatchId:       uEvent.Id,
		SourceChainId: uEvent.SourceChainId.String(),
		DestChainId:   uEvent.DestChainId.String(),
		Updated:       event.Timestamp,
		Status:        Claim,
	}
	return db.AppChainUnStakeBatch.ClaimAppChainUnStakeBatch(unStakeSingle)
	return nil
}

func UnstakeSingle(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := StakeUnpack.ParseUnstakeSingle(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}

	unStakeSingle := appchain.AppChainUnStakeSingle{
		TxHash:       rlpLog.TxHash,
		BlockNumber:  event.BlockNumber,
		LockedAmount: uEvent.DETHLocked,
		L2Strategy:   uEvent.L2Strategy,
		Created:      event.Timestamp,
	}
	return db.AppChainUnStakeSingle.StoreAppChainUnStakeSingle(unStakeSingle)

}

func UnstakeBatchRequest(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := StakeUnpack.ParseUnstakeBatchRequest(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	unStakeBatch := appchain.AppChainUnStakeBatch{
		TxHash:          rlpLog.TxHash,
		BlockNumber:     event.BlockNumber,
		BatchEthAmount:  uEvent.BatchEthAmount,
		BatchDETHLocked: uEvent.BatchDETHLocked,
		Status:          Pending,
		Created:         event.Timestamp,
	}
	return db.AppChainUnStakeBatch.StoreAppChainUnStakeBatch(unStakeBatch)
}
