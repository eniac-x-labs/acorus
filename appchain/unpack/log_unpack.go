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
	unStakeSingle := appchain.AppChainUnStake{
		ClaimTxHash:   rlpLog.TxHash,
		Bridge:        uEvent.Bridge,
		Staker:        uEvent.Staker,
		SourceChainId: uEvent.SourceChainId.String(),
		DestChainId:   uEvent.DestChainId.String(),
		L2Strategy:    uEvent.L2Strategy,
		Updated:       event.Timestamp,
		Status:        Claim,
	}
	return db.AppChainUnStake.ClaimAppChainUnStake(unStakeSingle, Pending)
}

func UnstakeRequested(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := StakeUnpack.ParseUnstakeRequested(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	unStakeSingle := appchain.AppChainUnStake{
		TxHash:        rlpLog.TxHash,
		BlockNumber:   event.BlockNumber,
		L2Strategy:    uEvent.L2Strategy,
		Staker:        uEvent.Staker,
		EthAmount:     uEvent.EthAmount,
		DETHLocked:    uEvent.DETHLocked,
		SourceChainId: chainId,
		DestChainId:   uEvent.DestChainId.String(),
		Created:       event.Timestamp,
		Status:        Pending,
	}
	return db.AppChainUnStake.StoreAppChainUnStake(unStakeSingle)
}
