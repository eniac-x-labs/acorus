package unpack

import (
	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/rpc/bridge"
	"github.com/cornerstone-labs/acorus/rpc/bridge/protobuf/pb"
	"github.com/ethereum/go-ethereum/log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	delegation_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/delegation/bindings"
	stake_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/staking_manager/bindings"
	strategybasebindings "github.com/cornerstone-labs/acorus/appchain/bindings/strategy_base/bindings"
	stratege_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/strategy_manager/bindings"
	unstake_manager_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/unstake_requests_manager/bindings"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/appchain"
	"github.com/cornerstone-labs/acorus/database/event"
)

var (
	StakeUnpack, _          = stake_bindings.NewStakingManager(common.Address{}, nil)
	StrategyUnpack, _       = stratege_bindings.NewStrategyManager(common.Address{}, nil)
	DelegateUnpack, _       = delegation_bindings.NewDelegationManager(common.Address{}, nil)
	StrategyBaseUnpack, _   = strategybasebindings.NewStrategyBase(common.Address{}, nil)
	UnstakeManagerUnpack, _ = unstake_manager_bindings.NewUnstakeRequestsManager(common.Address{}, nil)
)

const (
	Pending = 0
	Claim   = 1
)

func UnstakeRequestClaimed(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := UnstakeManagerUnpack.ParseUnstakeRequestClaimed(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	unStakeSingle := appchain.AppChainUnStake{
		ClaimTxHash:  rlpLog.TxHash,
		Bridge:       uEvent.BridgeAddress,
		UnstakeNonce: uEvent.UnStakeMessageNonce,
		Updated:      event.Timestamp,
		Status:       Claim,
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
		UnstakeNonce:  uEvent.UnStakeMessageNonce,
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

func TransferETHToL2DappLinkBridge(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := StrategyBaseUnpack.ParseTransferETHToL2DappLinkBridge(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	dapplinkBridgeRecord := appchain.AppChainDappLinkBridge{
		TxHash:                rlpLog.TxHash,
		SourceChainId:         uEvent.SourceChainId.String(),
		DestChainId:           uEvent.DestChainId.String(),
		ChainId:               chainId,
		Bridge:                uEvent.Bridge,
		StakingManagerAddress: uEvent.L1StakingManagerAddr,
		TokenAddress:          uEvent.TokenAddress,
		Amount:                uEvent.BridgeEthAmount,
		BatchId:               uEvent.BatchId,
		Nonce:                 uEvent.Nonce,
		NotifyRelayer:         false,
		Created:               event.Timestamp,
	}

	return db.AppChainDappLinkBridge.StroeAppChainDappLinkBridge(dapplinkBridgeRecord)
}

func MigrateRelatedL1StakerShares(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := StrategyUnpack.ParseMigrateRelatedL1StakerShares(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	appChainMigrateShares := appchain.AppChainMigrateShares{
		TxHash:        rlpLog.TxHash,
		ChainId:       chainId,
		BlockNumber:   event.BlockNumber,
		UnstakeNonce:  uEvent.L1UnStakeMessageNonce,
		Shares:        uEvent.Shares,
		Staker:        uEvent.Staker,
		Strategy:      uEvent.Strategy,
		NotifyRelayer: false,
		Created:       event.Timestamp,
	}

	return db.AppChainMigrateShares.StoreMigrateShares(appChainMigrateShares)
}

func WithdrawalCompleted(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := DelegateUnpack.ParseWithdrawalCompleted(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	appChainMigrateShares := appchain.AppChainWithdraw{
		TxHash:      rlpLog.TxHash,
		ChainId:     chainId,
		BlockNumber: event.BlockNumber,
		Shares:      uEvent.Shares,
		Staker:      uEvent.Staker,
		Strategy:    uEvent.Strategy,
		Operator:    uEvent.Operator,
		Created:     event.Timestamp,
	}
	return db.AppChainWithdraw.StoreAppchainWithdraw(appChainMigrateShares)

}

func ComputeMsgHash(sourceChainId, destChainId, nonce *big.Int) common.Hash {
	return crypto.Keccak256Hash(common.LeftPadBytes(sourceChainId.Bytes(), 32),
		common.LeftPadBytes(destChainId.Bytes(), 32),
		common.LeftPadBytes(nonce.Bytes(), 32))
}

func TransferL2Shares(bridgeRpcService bridge.BridgeRpcService, fromAddress, toAddress common.Address,
	shares *big.Int, db *database.DB) error {
	fromAddressDethList := db.AppChainDethShares.ListDethSharesByAddress(fromAddress)
	log.Info("TransferL2Shares", "fromAddressDethList", len(fromAddressDethList))
	if len(fromAddressDethList) == 0 {
		return nil
	}
	sharesMap := make(map[uint64]*pb.ShareMap)
	for i := 0; i < len(fromAddressDethList); i++ {
		strategyAmountMap := make(map[string]string)
		dethSharesInfo := fromAddressDethList[i]
		thisStrategyShares := dethSharesInfo.Shares
		chainIdStr := dethSharesInfo.ChainId
		chainId, _ := big.NewInt(0).SetString(chainIdStr, 10)
		chainIdUint := chainId.Uint64()
		totalShares := getTotalShares(sharesMap)
		needShares := big.NewInt(0).Sub(shares, totalShares)
		if needShares.Cmp(shares) == 0 {
			break
		}
		strategyAddress := dethSharesInfo.StrategyAddress
		if needShares.Cmp(thisStrategyShares) == -1 {
			strategyAmountMap[strategyAddress.String()] = needShares.String()
			sharesMap[chainIdUint] = &pb.ShareMap{
				ShareMap: strategyAmountMap,
			}
			break
		}
		if needShares.Cmp(thisStrategyShares) == 0 {
			strategyAmountMap[strategyAddress.String()] = needShares.String()
			sharesMap[chainIdUint] = &pb.ShareMap{
				ShareMap: strategyAmountMap,
			}
			break
		}
		if needShares.Cmp(thisStrategyShares) == 1 {
			strategyAmountMap[strategyAddress.String()] = thisStrategyShares.String()
			sharesMap[chainIdUint] = &pb.ShareMap{
				ShareMap: strategyAmountMap,
			}
		}
	}
	l2ShareReq := &pb.TransferL2ShareRequest{}
	bridgeRpcService.TransferL2Share(l2ShareReq)
	return nil
}

func getTotalShares(sharesMap map[uint64]*pb.ShareMap) *big.Int {
	totalShares := bigint.Zero
	for _, shareMapInfo := range sharesMap {
		shareMap := shareMapInfo.GetShareMap()
		for _, shares := range shareMap {
			sharesBig, _ := big.NewInt(0).SetString(shares, 10)
			totalShares = big.NewInt(0).Add(totalShares, sharesBig)
		}
	}
	return totalShares
}
