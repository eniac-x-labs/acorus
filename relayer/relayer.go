package relayer

import (
	"context"
	"fmt"
	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/relayer/bindings"
	"github.com/cornerstone-labs/acorus/relayer/unpack"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"strconv"
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
	resourceCtx       context.Context
	resourceCancel    context.CancelFunc
	tasks             tasks.Group
	loopInterval      time.Duration
	epoch             uint64
}

func NewRelayerListener(chainId string,
	l1Contracts []string,
	l2Contracts []string,
	l1EventStartBlock uint64,
	l2EventStartBlock uint64,
	shutdown context.CancelCauseFunc,
	loopInterval time.Duration, epoch uint64) (*RelayerListener, error) {
	resCtx, resCancel := context.WithCancel(context.Background())
	return &RelayerListener{
		chainId:           chainId,
		l1EventStartBlock: l1EventStartBlock,
		l2EventStartBlock: l2EventStartBlock,
		l1Contracts:       l1Contracts,
		l2Contracts:       l2Contracts,
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
	relayerEventOn1 := time.NewTicker(rl.loopInterval)
	relayerEventOn2 := time.NewTicker(rl.loopInterval)
	rl.tasks.Go(func() error {
		for range relayerEventOn1.C {
			err := rl.onL1Data()
			if err != nil {
				log.Println("no more l1 etl updates. shutting down l1 task")
				return err
			}
		}
		return nil
	})
	rl.tasks.Go(func() error {
		for range relayerEventOn2.C {
			err := rl.onL2Data()
			if err != nil {
				log.Println("no more l1 etl updates. shutting down l1 task")
				return err
			}
		}
		return nil
	})
	return nil
}

func (rl *RelayerListener) onL1Data() error {

	if rl.l1StartHeight == nil {
		// todo 修改表名处理监听到最后一个区块好
		lastBlockHeard, err := rl.db.L1ToL2.L1LatestBlockHeader(rl.chainId)
		if err != nil {
			log.Println("l1 failed to get last block heard", "err", err)
			return err
		}
		if lastBlockHeard == nil {
			lastBlockHeard = &common2.BlockHeader{
				Number: big.NewInt(0),
			}
		}
		rl.l1StartHeight = lastBlockHeard.Number
		if rl.l1StartHeight.Cmp(big.NewInt(int64(rl.l1EventStartBlock))) == -1 {
			rl.l1StartHeight = big.NewInt(int64(rl.l1EventStartBlock))
		}
	} else {
		rl.l1StartHeight = new(big.Int).Add(rl.l1StartHeight, bigint.One)
	}

	fromL1Height := rl.l1StartHeight
	toL1Height := new(big.Int).Add(fromL1Height, big.NewInt(int64(rl.epoch)))
	// todo 处理监听到最后一个区块
	chainLatestBlockHeader, err := rl.db.Blocks.ChainLatestBlockHeader(strconv.FormatUint(global_const.EthereumChainId, 10))
	if err != nil {
		return err
	}
	if chainLatestBlockHeader == nil {
		return nil
	}
	if chainLatestBlockHeader.Number.Cmp(fromL1Height) == -1 {
		return nil
	} else {
		if chainLatestBlockHeader.Number.Cmp(toL1Height) == -1 {
			toL1Height = new(big.Int).Add(fromL1Height, chainLatestBlockHeader.Number)
		}
	}
	if err := rl.db.Transaction(func(tx *database.DB) error {
		l1EventsFetchErr := rl.l1EventsFetch(fromL1Height, toL1Height)
		if l1EventsFetchErr != nil {
			return l1EventsFetchErr
		}
		return nil
	}); err != nil {
		return err
	}
	rl.l1StartHeight = toL1Height
	return nil
	return nil
}

func (rl *RelayerListener) onL2Data() error {
	if rl.l2StartHeight == nil {
		// todo 处理监听到最后一个区块
		lastBlockHeard, err := rl.db.L2ToL1.L2LatestBlockHeader(rl.chainId)
		if err != nil {
			log.Println("l2 failed to get last block heard", "err", err)
			return err
		}
		if lastBlockHeard == nil {
			if rl.l2EventStartBlock > 0 {
				lastBlockHeard = &common2.BlockHeader{
					Number: big.NewInt(int64(rl.l2EventStartBlock)),
				}
			} else {
				lastBlockHeard = &common2.BlockHeader{
					Number: big.NewInt(0),
				}
			}
		}
		rl.l2StartHeight = lastBlockHeard.Number
	} else {
		rl.l2StartHeight = new(big.Int).Add(rl.l2StartHeight, bigint.One)
	}
	fromL2Height := rl.l2StartHeight
	toL2Height := new(big.Int).Add(fromL2Height, big.NewInt(int64(rl.epoch)))
	// todo 处理监听到最后一个区块
	chainLatestBlockHeader, err := rl.db.Blocks.ChainLatestBlockHeader(rl.chainId)
	if err != nil {
		return err
	}
	if chainLatestBlockHeader == nil {
		return nil
	}
	if chainLatestBlockHeader.Number.Cmp(fromL2Height) == -1 {
		return nil
	} else {
		if chainLatestBlockHeader.Number.Cmp(toL2Height) == -1 {
			toL2Height = new(big.Int).Add(fromL2Height, chainLatestBlockHeader.Number)
		}
	}
	if err := rl.db.Transaction(func(tx *database.DB) error {
		l2EventsFetchErr := rl.l2EventsFetch(fromL2Height, toL2Height)
		if l2EventsFetchErr != nil {
			return l2EventsFetchErr
		}
		return nil
	}); err != nil {
		return err
	}
	rl.l2StartHeight = toL2Height
	return nil
}

func (rl *RelayerListener) l1EventsFetch(fromL1Height, toL1Height *big.Int) error {
	l1Contracts := rl.l1Contracts
	for _, l1contract := range l1Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l1contract)}
		events, err := rl.db.ContractEvents.ContractEventsWithFilter("1", contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			log.Println("failed to index L1ContractEventsWithFilter ", "err", err)
			return err
		}
		for _, contractEvent := range events {
			unpackErr := rl.eventUnpack(contractEvent)
			if unpackErr != nil {
				log.Println("failed to index L1 bridge events", "unpackErr", unpackErr)
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
		events, err := rl.db.ContractEvents.ContractEventsWithFilter(chainIdStr, contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			log.Println("failed to index L2ContractEventsWithFilter ", "err", err)
			return err
		}
		for _, contractEvent := range events {
			unpackErr := rl.eventUnpack(contractEvent)
			if unpackErr != nil {
				log.Println("failed to index L2 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
		}
	}
	return nil
}

func (rl *RelayerListener) eventUnpack(event event.ContractEvent) error {
	chainIdStr := rl.chainId
	switch event.EventSignature.String() {
	case bindings.MsgAbi.Events["MessageSent"].ID.String():
		err := unpack.MessageSent(chainIdStr, event, rl.db)
		return err
	case bindings.MsgAbi.Events["MessageClaimed"].ID.String():
		err := unpack.MessageClaimed(chainIdStr, event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["FinalizeETH"].ID.String():
		err := unpack.FinalizeETH(chainIdStr, event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["FinalizeWETH"].ID.String():
		err := unpack.FinalizeWETH(chainIdStr, event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["FinalizeERC20"].ID.String():
		err := unpack.FinalizeERC20(chainIdStr, event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["InitiateWETH"].ID.String():
		err := unpack.InitiateWETH(chainIdStr, event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["InitiateETH"].ID.String():
		err := unpack.InitiateETH(chainIdStr, event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["InitiateERC20"].ID.String():
		err := unpack.InitiateERC20(chainIdStr, event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["StarkingERC20Event"].ID.String():
		err := unpack.StarkingERC20Event(chainIdStr, event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["StakingETHEvent"].ID.String():
		err := unpack.StakingETHEvent(chainIdStr, event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["StakingWETHEvent"].ID.String():
		err := unpack.StakingWETHEvent(chainIdStr, event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["ClaimEvent"].ID.String():
		err := unpack.ClaimEvent(chainIdStr, event, rl.db)
		return err
	}
	return nil
}
