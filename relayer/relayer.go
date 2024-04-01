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
					log.Error("l1 relayer. shutting down l1 task", "err", err)
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
					log.Error("l2 relayer. shutting down l2 task", "err", err)
					continue
				}
			}
			return nil
		})
	}
	return nil
}

func (rl *RelayerListener) onL1Data() error {
	lastListenBlock, err := rl.db.BridgeBlockListener.GetLastBlockNumber(rl.chainId)
	if err != nil {
		log.Error("l1 failed to get last block heard", "err", err)
		return err
	}
	fromL1Height := bigint.Zero

	if lastListenBlock != nil {
		fromL1Height = new(big.Int).Add(lastListenBlock.BlockNumber, bigint.One)
	}

	if rl.l1EventStartBlock > 0 {
		if fromL1Height.Cmp(big.NewInt(int64(rl.l1EventStartBlock))) == -1 {
			fromL1Height = big.NewInt(int64(rl.l1EventStartBlock))
		}
	}

	toL1Height := new(big.Int).Add(fromL1Height, big.NewInt(int64(rl.epoch)))
	chainLatestBlockHeader, err := rl.db.Blocks.ChainLatestBlockHeader(rl.chainId)
	if err != nil {
		return err
	}
	if chainLatestBlockHeader == nil {
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
		return err
	}
	return nil
}

func (rl *RelayerListener) onL2Data() error {
	lastListenBlock, err := rl.db.BridgeBlockListener.GetLastBlockNumber(rl.chainId)
	if err != nil {
		log.Error("l2 failed to get last block heard", "err", err)
		return err
	}
	fromL2Height := bigint.Zero
	if lastListenBlock != nil {
		fromL2Height = new(big.Int).Add(lastListenBlock.BlockNumber, bigint.One)
	}
	if rl.l2EventStartBlock > 0 {
		if fromL2Height.Cmp(big.NewInt(int64(rl.l2EventStartBlock))) == -1 {
			fromL2Height = big.NewInt(int64(rl.l2EventStartBlock))
		}
	}
	toL2Height := new(big.Int).Add(fromL2Height, big.NewInt(int64(rl.epoch)))
	chainLatestBlockHeader, err := rl.db.Blocks.ChainLatestBlockHeader(rl.chainId)
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
		return err
	}
	return nil
}

func (rl *RelayerListener) l1EventsFetch(fromL1Height, toL1Height *big.Int) error {
	log.Info("Relayer l1EventsFetch", "fromL1Height", fromL1Height, "toL1Height", toL1Height)
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
		log.Info("Relayer l2EventsFetch", "chainId", chainIdStr, "contract", common.HexToAddress(l2contract).String(),
			"fromL1Height", fromL1Height, "toL1Height", toL1Height)

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
