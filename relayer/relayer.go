package relayer

import (
	"context"
	"fmt"
	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/relayer"
	"github.com/cornerstone-labs/acorus/relayer/bindings"
	"github.com/cornerstone-labs/acorus/relayer/unpack"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"time"
)

type RelayerListener struct {
	db                *database.DB
	chainId           string
	l1Contracts       []string
	l2Contracts       []string
	l1EventStartBlock uint64
	l2EventStartBlock uint64
	l1StartHeight     *big.Int
	l2StartHeight     *big.Int
	layerType         int
	resourceCtx       context.Context
	resourceCancel    context.CancelFunc
	tasks             tasks.Group
	loopInterval      time.Duration
	epoch             uint64
}

func NewRelayerListener(
	db *database.DB,
	chainId string,
	l1Contracts []string,
	l2Contracts []string,
	l1EventStartBlock uint64,
	l2EventStartBlock uint64,
	layerType int,
	shutdown context.CancelCauseFunc,
	loopInterval time.Duration, epoch uint64) (*RelayerListener, error) {
	resCtx, resCancel := context.WithCancel(context.Background())
	return &RelayerListener{
		db:                db,
		chainId:           chainId,
		l1EventStartBlock: l1EventStartBlock,
		l2EventStartBlock: l2EventStartBlock,
		l1Contracts:       l1Contracts,
		l2Contracts:       l2Contracts,
		layerType:         layerType,
		resourceCtx:       resCtx,
		resourceCancel:    resCancel,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
		loopInterval: loopInterval,
		epoch:        epoch,
	}, nil
}

func (rl *RelayerListener) Start() error {
	if rl.layerType == global_const.LayerTypeOne {
		relayerEventOn1 := time.NewTicker(rl.loopInterval)
		rl.tasks.Go(func() error {
			for range relayerEventOn1.C {
				err := rl.onL1Data()
				if err != nil {
					log.Error("no more l1 etl updates. shutting down l1 task")
					continue
				}
			}
			return nil
		})
	} else {
		relayerEventOn2 := time.NewTicker(rl.loopInterval)
		rl.tasks.Go(func() error {
			for range relayerEventOn2.C {
				err := rl.onL2Data()
				if err != nil {
					log.Error("no more l1 etl updates. shutting down l1 task")
					continue
				}
			}
			return nil
		})
	}
	return nil
}

func (rl *RelayerListener) onL1Data() error {
	if rl.l1StartHeight == nil {
		lastListenBlock, err := rl.db.BridgeBlockListener.GetLastBlockNumber(rl.chainId)
		if err != nil {
			log.Error("l1 failed to get last block heard", "err", err)
			return err
		}
		if lastListenBlock == nil {
			lastListenBlock = &relayer.BridgeBlockListener{
				BlockNumber: big.NewInt(0),
			}
		}
		rl.l1StartHeight = lastListenBlock.BlockNumber
		if rl.l1StartHeight.Cmp(big.NewInt(int64(rl.l1EventStartBlock))) == -1 {
			rl.l1StartHeight = big.NewInt(int64(rl.l1EventStartBlock))
		}
	} else {

		rl.l1StartHeight = new(big.Int).Add(rl.l1StartHeight, bigint.One)
	}

	fromL1Height := rl.l1StartHeight
	toL1Height := new(big.Int).Add(fromL1Height, big.NewInt(int64(rl.epoch)))
	chainLatestBlockHeader, err := rl.db.Blocks.ChainLatestBlockHeader(rl.chainId)
	if err != nil {
		rl.l1StartHeight = new(big.Int).Sub(rl.l1StartHeight, bigint.One)
		return err
	}
	if chainLatestBlockHeader == nil {
		rl.l1StartHeight = new(big.Int).Sub(rl.l1StartHeight, bigint.One)
		return nil
	}
	if chainLatestBlockHeader.Number.Cmp(fromL1Height) == -1 {
		fromL1Height = chainLatestBlockHeader.Number
		toL1Height = chainLatestBlockHeader.Number
	} else {
		if chainLatestBlockHeader.Number.Cmp(toL1Height) == -1 {
			toL1Height = chainLatestBlockHeader.Number
		}
	}
	if err := rl.db.Transaction(func(tx *database.DB) error {
		l1EventsFetchErr := rl.l1EventsFetch(fromL1Height, toL1Height)
		if l1EventsFetchErr != nil {
			return l1EventsFetchErr
		}
		lastBlock := relayer.BridgeBlockListener{
			ChainId:     rl.chainId,
			BlockNumber: toL1Height,
		}
		updateErr := rl.db.BridgeBlockListener.SaveOrUpdateLastBlockNumber(lastBlock)
		if updateErr != nil {
			log.Error("update last block err :", updateErr)
			return updateErr
		}
		return nil
	}); err != nil {
		rl.l1StartHeight = new(big.Int).Sub(rl.l1StartHeight, bigint.One)
		return err
	}
	rl.l1StartHeight = toL1Height
	return nil
}

func (rl *RelayerListener) onL2Data() error {
	if rl.l2StartHeight == nil {
		lastListenBlock, err := rl.db.BridgeBlockListener.GetLastBlockNumber(rl.chainId)
		if err != nil {
			log.Error("l2 failed to get last block heard", "err", err)
			return err
		}
		if lastListenBlock == nil {
			if rl.l2EventStartBlock > 0 {
				lastListenBlock = &relayer.BridgeBlockListener{
					BlockNumber: big.NewInt(int64(rl.l2EventStartBlock)),
				}
			} else {
				lastListenBlock = &relayer.BridgeBlockListener{
					BlockNumber: big.NewInt(0),
				}
			}
		}
		rl.l2StartHeight = lastListenBlock.BlockNumber
	} else {
		rl.l2StartHeight = new(big.Int).Add(rl.l2StartHeight, bigint.One)
	}
	fromL2Height := rl.l2StartHeight
	toL2Height := new(big.Int).Add(fromL2Height, big.NewInt(int64(rl.epoch)))
	chainLatestBlockHeader, err := rl.db.Blocks.ChainLatestBlockHeader(rl.chainId)
	if err != nil {
		rl.l2StartHeight = new(big.Int).Sub(rl.l2StartHeight, bigint.One)
		return err
	}
	if chainLatestBlockHeader == nil {
		rl.l2StartHeight = new(big.Int).Sub(rl.l2StartHeight, bigint.One)
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
	if err := rl.db.Transaction(func(tx *database.DB) error {
		l2EventsFetchErr := rl.l2EventsFetch(fromL2Height, toL2Height)
		if l2EventsFetchErr != nil {
			return l2EventsFetchErr
		}
		lastBlock := relayer.BridgeBlockListener{
			ChainId:     rl.chainId,
			BlockNumber: toL2Height,
		}
		updateErr := rl.db.BridgeBlockListener.SaveOrUpdateLastBlockNumber(lastBlock)
		if updateErr != nil {
			return updateErr
		}
		return nil
	}); err != nil {
		rl.l2StartHeight = new(big.Int).Sub(rl.l2StartHeight, bigint.One)
		return err
	}
	rl.l2StartHeight = toL2Height
	return nil
}

func (rl *RelayerListener) l1EventsFetch(fromL1Height, toL1Height *big.Int) error {
	log.Info("relayer l1EventsFetch", "fromL1Height", fromL1Height, "toL1Height", toL1Height)
	l1Contracts := rl.l1Contracts
	for _, l1contract := range l1Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l1contract)}
		events, err := rl.db.ContractEvents.ContractEventsWithFilter(rl.chainId, contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			log.Error("failed to index L1ContractEventsWithFilter ", "err", err)
			return err
		}
		for _, contractEvent := range events {
			unpackErr := rl.eventUnpack(contractEvent)
			if unpackErr != nil {
				log.Error("failed to index L1 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
		}
	}
	return nil
}

func (rl *RelayerListener) l2EventsFetch(fromL1Height, toL1Height *big.Int) error {
	chainIdStr := rl.chainId
	l2Contracts := rl.l2Contracts
	for _, l2contract := range l2Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l2contract)}
		log.Info("chainId", chainIdStr, "contract", common.HexToAddress(l2contract).String(),
			"relayer l2EventsFetch", "fromL1Height", fromL1Height, "toL1Height", toL1Height)

		events, err := rl.db.ContractEvents.ContractEventsWithFilter(chainIdStr, contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			log.Error("failed to index L2ContractEventsWithFilter ", "err", err)
			return err
		}
		for _, contractEvent := range events {
			unpackErr := rl.eventUnpack(contractEvent)
			if unpackErr != nil {
				log.Error("failed to index L2 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
		}
	}
	return nil
}

func (rl *RelayerListener) eventUnpack(event event.ContractEvent) error {
	switch event.EventSignature.String() {
	case bindings.MsgAbi.Events["MessageSent"].ID.String():
		err := unpack.MessageSent(event, rl.db)
		return err
	case bindings.MsgAbi.Events["MessageClaimed"].ID.String():
		err := unpack.MessageClaimed(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["FinalizeETH"].ID.String():
		err := unpack.FinalizeETH(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["FinalizeWETH"].ID.String():
		err := unpack.FinalizeWETH(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["FinalizeERC20"].ID.String():
		err := unpack.FinalizeERC20(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["InitiateWETH"].ID.String():
		err := unpack.InitiateWETH(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["InitiateETH"].ID.String():
		err := unpack.InitiateETH(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["InitiateERC20"].ID.String():
		err := unpack.InitiateERC20(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["StarkingERC20Event"].ID.String():
		err := unpack.StarkingERC20Event(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["StakingETHEvent"].ID.String():
		err := unpack.StakingETHEvent(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["StakingWETHEvent"].ID.String():
		err := unpack.StakingWETHEvent(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["ClaimEvent"].ID.String():
		err := unpack.ClaimEvent(event, rl.db)
		return err
	case bindings.IL1PoolAbi.Events["Withdraw"].ID.String():
		err := unpack.Withdraw(event, rl.db)
		return err
	case bindings.IL1PoolAbi.Events["ClaimReward"].ID.String():
		err := unpack.ClaimReward(event, rl.db)
		return err
	}
	return nil
}
