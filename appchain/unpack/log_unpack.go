package unpack

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"

	delegation_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/delegation/bindings"
	stake_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/staking_manager/bindings"
	stratege_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/strategy_manager/bindings"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/appchain"
	"github.com/cornerstone-labs/acorus/database/event"
)

var (
	StakeUnpack, _    = stake_bindings.NewStakingManager(common.Address{}, nil)
	StrategyUnpack, _ = stratege_bindings.NewStrategyManager(common.Address{}, nil)
	DelegateUnpack, _ = delegation_bindings.NewDelegationManager(common.Address{}, nil)
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

func StakeRecord(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := StrategyUnpack.ParseDeposit(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	stakeSingle := appchain.AppChainStake{
		TxHash:          rlpLog.TxHash,
		BlockNumber:     event.BlockNumber,
		StrategyAddress: uEvent.Strategy,
		Staker:          uEvent.Staker,
		Shares:          uEvent.Shares,
		TokenAddress:    uEvent.Weth,
		ChainId:         chainId,
		Created:         event.Timestamp,
	}
	return db.AppChainStake.StoreAppChainStake(stakeSingle)
}

func OperatorSharesIncreased(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := DelegateUnpack.ParseOperatorSharesIncreased(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	operatorSharesIncreased := appchain.AppChainOperatorSharesIncreased{
		TxHash:          rlpLog.TxHash,
		BlockNumber:     event.BlockNumber,
		StrategyAddress: uEvent.Strategy,
		Operator:        uEvent.Operator,
		Staker:          uEvent.Staker,
		Shares:          uEvent.Shares,
		UseShares:       big.NewInt(0),
		ChainId:         chainId,
		Created:         event.Timestamp,
	}
	return db.AppChainOperatorSharesIncreased.StoreAppChainOperatorSharesIncreased(operatorSharesIncreased)
}