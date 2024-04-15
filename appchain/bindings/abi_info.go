package bindings

import (
	delegation_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/delegation/bindings"
	stake_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/staking_manager/bindings"
	strategybasebindings "github.com/cornerstone-labs/acorus/appchain/bindings/strategy_base/bindings"
	stratege_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/strategy_manager/bindings"
	unstake_manager_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/unstake_requests_manager/bindings"
)

var (
	StakingManagerAbi, _         = stake_bindings.StakingManagerMetaData.GetAbi()
	StrategyManagerAbi, _        = stratege_bindings.StrategyManagerMetaData.GetAbi()
	DelegationManagerAbi, _      = delegation_bindings.DelegationManagerMetaData.GetAbi()
	StrategyBaseAbi, _           = strategybasebindings.StrategyBaseMetaData.GetAbi()
	UnstakeRequestsManagerAbi, _ = unstake_manager_bindings.UnstakeRequestsManagerMetaData.GetAbi()
)
