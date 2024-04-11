package appchain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"time"

	"github.com/cornerstone-labs/acorus/appchain/bindings"
	"github.com/cornerstone-labs/acorus/appchain/unpack"
	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/appchain"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/rpc/bridge"
)

var (
	StakeAmount = big.NewInt(0).Mul(big.NewInt(32), big.NewInt(1e18))
)

type L2AppChainListener struct {
	chainId          string
	contracts        []string
	tasks            tasks.Group
	loopInterval     time.Duration
	startHeight      uint64
	epoch            uint64
	db               *database.DB
	bridgeRpcService bridge.BridgeRpcService
}

func NewL2AppChainListener(chainId string,
	contracts []string,
	shutdown context.CancelCauseFunc,
	loopInterval time.Duration,
	startHeight uint64,
	epoch uint64,
	bridgeRpcService bridge.BridgeRpcService,
	db *database.DB) (*L2AppChainListener, error) {
	return &L2AppChainListener{
		chainId:   chainId,
		contracts: contracts,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
		loopInterval:     loopInterval,
		startHeight:      startHeight,
		epoch:            epoch,
		bridgeRpcService: bridgeRpcService,
		db:               db,
	}, nil
}

func (l *L2AppChainListener) Start() error {
	l2ListenerEventOn1 := time.NewTicker(l.loopInterval)
	l.tasks.Go(func() error {
		for range l2ListenerEventOn1.C {
			err := l.onL2Data()
			if err != nil {
				log.Error("L2AppChainListener", "err", err)
				continue
			}
		}
		return nil
	})
	l.tasks.Go(func() error {
		for range l2ListenerEventOn1.C {
			err := l.operatorSharesTask()
			if err != nil {
				log.Error("L2AppChainListener", "err", err)
				continue
			}
		}
		return nil
	})
	return nil
}

func (l *L2AppChainListener) onL2Data() error {
	lastListenBlock, err := l.db.AppChainLastBlock.GetLastBlockNumber(l.chainId)
	if err != nil {
		log.Error("appchain L2  failed to get last block heard", "err", err)
		return err
	}
	fromL2Height := bigint.Zero
	if lastListenBlock != nil {
		fromL2Height = new(big.Int).Add(lastListenBlock.BlockNumber, bigint.One)
	}
	if l.startHeight > 0 {
		if fromL2Height.Cmp(big.NewInt(int64(l.startHeight))) == -1 {
			fromL2Height = big.NewInt(int64(l.startHeight))
		}
	}
	toL2Height := new(big.Int).Add(fromL2Height, big.NewInt(int64(l.epoch)))
	chainLatestBlockHeader, err := l.db.Blocks.ChainLatestBlockHeader(l.chainId)
	if err != nil {
		return err
	}
	if chainLatestBlockHeader == nil {
		return nil
	}
	if chainLatestBlockHeader.Number.Cmp(fromL2Height) == -1 {
		fromL2Height = chainLatestBlockHeader.Number
		toL2Height = chainLatestBlockHeader.Number
	} else {
		if chainLatestBlockHeader.Number.Cmp(toL2Height) == -1 {
			toL2Height = chainLatestBlockHeader.Number
		}
	}
	if err := l.db.Transaction(func(tx *database.DB) error {
		eventsFetchErr := l.EventsFetch(fromL2Height, toL2Height)
		if eventsFetchErr != nil {
			return eventsFetchErr
		}
		lastBlock := appchain.AppChainLastBlock{
			ChainId:     l.chainId,
			BlockNumber: toL2Height,
		}
		updateErr := l.db.AppChainLastBlock.SaveOrUpdateLastBlockNumber(lastBlock)
		if updateErr != nil {
			return updateErr
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (l *L2AppChainListener) EventsFetch(fromL1Height, toL1Height *big.Int) error {
	chainIdStr := l.chainId
	contracts := l.contracts
	for _, contract := range contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(contract)}
		log.Info("appchain L2 Fetch", "chainId", chainIdStr, "contract", common.HexToAddress(contract).String(),
			"fromL1Height", fromL1Height, "toL1Height", toL1Height)

		events, err := l.db.ContractEvents.ContractEventsWithFilter(chainIdStr, contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			log.Error("failed to index L2 ContractEventsWithFilter ", "err", err)
			return err
		}
		for _, contractEvent := range events {
			unpackErr := l.eventUnpack(contractEvent)
			if unpackErr != nil {
				log.Error("failed to index L2 events", "unpackErr", unpackErr)
				return unpackErr
			}
		}
	}
	return nil
}

func (l *L2AppChainListener) eventUnpack(event event.ContractEvent) error {
	switch event.EventSignature.String() {
	case bindings.StrategyManagerAbi.Events["Deposit"].ID.String():
		err := unpack.StakeRecord(l.chainId, event, l.db)
		return err
	case bindings.DelegationManagerAbi.Events["OperatorSharesIncreased"].ID.String():
		err := unpack.OperatorSharesIncreased(l.chainId, event, l.db)
		return err
	}
	return nil
}

func (l *L2AppChainListener) operatorSharesTask() error {
	log.Info("operatorSharesTask start")
	if err := l.db.Transaction(func(tx *database.DB) error {
		err := l.operatorSharesIncreased()
		return err
	}); err != nil {
		return err
	}
	return nil
}

type ToStakeShares struct {
	TotalShares      *big.Int
	UserInfoShares   []*UserInfoShares
	NeedUpdateShares []*appchain.AppChainOperatorSharesIncreased
}
type UserInfoShares struct {
	UserShares      *big.Int
	Staker          common.Address
	Operator        common.Address
	StrategyAddress common.Address
}

func (l *L2AppChainListener) operatorSharesIncreased() error {
	toStakeShares := new(ToStakeShares)
	toStakeShares.TotalShares = big.NewInt(0)
	needStakeShares := l.db.AppChainOperatorSharesIncreased.GetNeedStakeShares(l.chainId)
	log.Info("operatorSharesIncreased", "needStakeSharesSize", len(needStakeShares))
	for _, needStakeShare := range needStakeShares {
		operatorAddress := needStakeShare.Operator
		shares := needStakeShare.Shares
		useShares := needStakeShare.UseShares
		strategyAddress := needStakeShare.StrategyAddress
		totalShares := toStakeShares.TotalShares
		if totalShares.Cmp(StakeAmount) == 0 {
			break
		}
		leftAddShares := big.NewInt(0).Sub(StakeAmount, totalShares)
		thisUserLeftShares := big.NewInt(0).Sub(shares, useShares)

		if thisUserLeftShares.Cmp(leftAddShares) == -1 {
			totalShares = big.NewInt(0).Add(totalShares, thisUserLeftShares)
			needStakeShare.Status = 1
			needStakeShare.UseShares = shares
		}
		if thisUserLeftShares.Cmp(leftAddShares) == 0 {
			totalShares = big.NewInt(0).Add(totalShares, leftAddShares)
			needStakeShare.Status = 1
			needStakeShare.UseShares = shares
		}
		if thisUserLeftShares.Cmp(leftAddShares) == 1 {
			totalShares = big.NewInt(0).Add(totalShares, leftAddShares)
			needStakeShare.UseShares = big.NewInt(0).Add(useShares, leftAddShares)
		}

		userInfoShares := toStakeShares.UserInfoShares
		userInfoShares = append(userInfoShares, &UserInfoShares{
			UserShares:      needStakeShare.UseShares,
			Staker:          needStakeShare.Staker,
			StrategyAddress: strategyAddress,
			Operator:        operatorAddress,
		})
		needUpdateShares := toStakeShares.NeedUpdateShares
		needUpdateShares = append(needUpdateShares, needStakeShare)
		toStakeShares.TotalShares = totalShares
		toStakeShares.UserInfoShares = userInfoShares
		toStakeShares.NeedUpdateShares = needUpdateShares
	}
	if toStakeShares.TotalShares.Cmp(StakeAmount) == 0 {
		batchIdInt := time.Now().Unix()
		batchId := big.NewInt(batchIdInt)
		increaseBatches := make([]appchain.AppChainIncreaseBatch, 0)
		// create batchMint param
		mintMap := make(map[string]string)
		for _, userInfoShares := range toStakeShares.UserInfoShares {
			increaseBatches = append(increaseBatches, appchain.AppChainIncreaseBatch{
				Staker:          userInfoShares.Staker,
				Shares:          userInfoShares.UserShares,
				ChainId:         l.chainId,
				BatchId:         batchId.String(),
				StrategyAddress: userInfoShares.StrategyAddress,
				Operator:        userInfoShares.Operator,
				Created:         uint64(time.Now().Unix()),
			})
			staker := userInfoShares.Staker.String()
			exitsSharesStr := mintMap[staker]
			if exitsSharesStr == "" {
				exitsSharesStr = "0"
			}
			exitsShares, _ := big.NewInt(0).SetString(exitsSharesStr, 10)
			mintMap[staker] = big.NewInt(0).Add(exitsShares, userInfoShares.UserShares).String()
		}

		// update shares
		log.Info("operatorSharesIncreased", "needStakeSharesSize", len(toStakeShares.NeedUpdateShares))
		for _, needStakeShare := range toStakeShares.NeedUpdateShares {
			err := l.db.AppChainOperatorSharesIncreased.UpdateOperatorUseShares(*needStakeShare)
			if err != nil {
				return err
			}
		}

		log.Info("operatorSharesIncreased", "increaseBatchesSize", len(increaseBatches))
		// save batch
		err := l.db.AppChainIncreaseBatch.StoreAppChainIncreasedBatch(increaseBatches)
		if err != nil {
			return err
		}
		// call rpc
		batchMintResp, err := l.bridgeRpcService.BatchMint(batchId.Uint64(), mintMap)
		if err != nil {
			return err
		}
		log.Info("BatchMint", "batchId", batchId, "rpcResp", batchMintResp)
		success := batchMintResp.Success
		if success {
			log.Info("BatchMint", "batchId", batchId, "success", success)
		} else {
			return fmt.Errorf("call rpc BatchMint failed")
		}
	}
	return nil
}
