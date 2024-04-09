package bindings

import (
	delegation_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/delegation/bindings"
	stake_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/staking_manager/bindings"
	stratege_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/strategy_manager/bindings"
)

var (
	StakingManagerAbi, _    = stake_bindings.StakingManagerMetaData.GetAbi()
	StrategyManagerAbi, _   = stratege_bindings.StrategyManagerMetaData.GetAbi()
	DelegationManagerAbi, _ = delegation_bindings.DelegationManagerMetaData.GetAbi()
)
