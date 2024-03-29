package zkfair

import (
	"context"
	"fmt"
	"github.com/cornerstone-labs/acorus/event/zkfair/bridge"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/event/zkfair/bindings"
	"github.com/ethereum/go-ethereum/common"
)

type ZkFairEventProcessor struct {
	db             *database.DB
	cfgRpc         *config.RPC
	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group
	polygonBridge  *bindings.Polygonzkevmbridge
	loopInterval   time.Duration
	l1StartHeight  *big.Int
	l2StartHeight  *big.Int
	epoch          uint64
	l1ChainId      string
	l2ChainId      string
}

func NewBridgeProcessor(db *database.DB,
	l2cfg *config.RPC, shutdown context.CancelCauseFunc,
	loopInterval time.Duration, epoch uint64, l1ChainId, l2ChainId string) (*ZkFairEventProcessor, error) {
	resCtx, resCancel := context.WithCancel(context.Background())

	polygonBridge, err := bindings.NewPolygonzkevmbridge(common.Address{}, nil)
	if err != nil {
		return nil, err
	}

	return &ZkFairEventProcessor{
		db:             db,
		cfgRpc:         l2cfg,
		resourceCtx:    resCtx,
		resourceCancel: resCancel,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
		polygonBridge: polygonBridge,
		loopInterval:  loopInterval,
		epoch:         epoch,
		l1ChainId:     l1ChainId,
		l2ChainId:     l2ChainId,
	}, nil
}

func (pp *ZkFairEventProcessor) StartUnpack() error {
	tickerEventOn1 := time.NewTicker(pp.loopInterval)
	tickerEventOn2 := time.NewTicker(pp.loopInterval)
	log.Info("starting zkfair bridge processor...")
	pp.tasks.Go(func() error {
		for range tickerEventOn1.C {
			err := pp.onL1Data()
			if err != nil {
				log.Error("fail no more l1 etl updates. shutting down l1 task")
				continue
			}
		}
		return nil
	})
	// start L2 worker
	pp.tasks.Go(func() error {
		for range tickerEventOn2.C {
			err := pp.onL2Data()
			if err != nil {
				log.Error("fail no more l2 etl updates. shutting down l2 task")
				continue
			}
		}
		return nil
	})
	return nil
}

func (pp *ZkFairEventProcessor) ClosetUnpack() error {
	pp.resourceCancel()
	return pp.tasks.Wait()
}

func (pp *ZkFairEventProcessor) onL1Data() error {
	if pp.l1StartHeight == nil {
		lastBlockHeard, err := pp.db.L1ToL2.L1LatestBlockHeader(pp.l2ChainId, pp.l1ChainId)
		if err != nil {
			log.Error("l1 failed to get last block heard", "err", err)
			return err
		}
		if lastBlockHeard == nil {
			lastBlockHeard = &common2.BlockHeader{
				Number: big.NewInt(0),
			}
		}
		pp.l1StartHeight = lastBlockHeard.Number
		if pp.l1StartHeight.Cmp(big.NewInt(int64(pp.cfgRpc.L1EventUnpackBlock))) == -1 {
			pp.l1StartHeight = big.NewInt(int64(pp.cfgRpc.L1EventUnpackBlock))
		}
	} else {
		pp.l1StartHeight = new(big.Int).Add(pp.l1StartHeight, bigint.One)
	}
	fromL1Height := pp.l1StartHeight
	toL1Height := new(big.Int).Add(fromL1Height, big.NewInt(int64(pp.epoch)))

	chainLatestBlockHeader, err := pp.db.Blocks.ChainLatestBlockHeader(pp.l1ChainId)
	if err != nil {
		pp.l1StartHeight = new(big.Int).Sub(pp.l1StartHeight, bigint.One)

		return err
	}
	if chainLatestBlockHeader == nil {
		pp.l1StartHeight = new(big.Int).Sub(pp.l1StartHeight, bigint.One)

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

	if err := pp.db.Transaction(func(tx *database.DB) error {
		l1EventsFetchErr := pp.l1EventsFetch(fromL1Height, toL1Height)
		if l1EventsFetchErr != nil {
			return l1EventsFetchErr
		}
		return nil
	}); err != nil {
		pp.l1StartHeight = new(big.Int).Sub(pp.l1StartHeight, bigint.One)

		return err
	}
	pp.l1StartHeight = toL1Height
	return nil
}

func (pp *ZkFairEventProcessor) onL2Data() error {
	if pp.l2StartHeight == nil {
		lastBlockHeard, err := pp.db.L2ToL1.L2LatestBlockHeader(pp.l2ChainId)
		if err != nil {
			log.Error("l2 failed to get last block heard", "err", err)
			return err
		}
		if lastBlockHeard == nil {
			if pp.cfgRpc.StartBlock > 0 {
				lastBlockHeard = &common2.BlockHeader{
					Number: big.NewInt(int64(pp.cfgRpc.StartBlock)),
				}
			} else {
				lastBlockHeard = &common2.BlockHeader{
					Number: big.NewInt(0),
				}
			}
		}
		pp.l2StartHeight = lastBlockHeard.Number
	} else {
		pp.l2StartHeight = new(big.Int).Add(pp.l2StartHeight, bigint.One)
	}
	fromL2Height := pp.l2StartHeight
	toL2Height := new(big.Int).Add(fromL2Height, big.NewInt(int64(pp.epoch)))
	chainLatestBlockHeader, err := pp.db.Blocks.ChainLatestBlockHeader(pp.l2ChainId)
	if err != nil {
		pp.l2StartHeight = new(big.Int).Sub(pp.l2StartHeight, bigint.One)
		return err
	}
	if chainLatestBlockHeader == nil {
		pp.l2StartHeight = new(big.Int).Sub(pp.l2StartHeight, bigint.One)
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
	if err := pp.db.Transaction(func(tx *database.DB) error {
		l2EventsFetchErr := pp.l2EventsFetch(fromL2Height, toL2Height)
		if l2EventsFetchErr != nil {
			return l2EventsFetchErr
		}
		return nil
	}); err != nil {
		pp.l2StartHeight = new(big.Int).Sub(pp.l2StartHeight, bigint.One)
		return err
	}
	pp.l2StartHeight = toL2Height
	return nil
}

func (pp *ZkFairEventProcessor) l1EventsFetch(fromL1Height, toL1Height *big.Int) error {
	l1Contracts := pp.cfgRpc.Contracts
	for _, l1contract := range l1Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l1contract)}
		events, err := pp.db.ContractEvents.ContractEventsWithFilter(pp.l1ChainId, contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			log.Error("failed to index L1ContractEventsWithFilter ", "err", err)
			return err
		}
		for _, contractEvent := range events {
			unpackErr := pp.l1EventUnpack(contractEvent)
			if unpackErr != nil {
				log.Error("failed to index L1 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
		}
	}
	return nil
}

func (pp *ZkFairEventProcessor) l2EventsFetch(fromL1Height, toL1Height *big.Int) error {
	l2chainIdStr := pp.l2ChainId
	l2Contracts := pp.cfgRpc.EventContracts
	for _, l2contract := range l2Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l2contract)}
		events, err := pp.db.ContractEvents.ContractEventsWithFilter(l2chainIdStr, contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			log.Error("failed to index L2ContractEventsWithFilter ", "err", err)
			return err
		}
		for _, contractEvent := range events {
			unpackErr := pp.l2EventUnpack(contractEvent)
			if unpackErr != nil {
				log.Error("failed to index L2 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
		}
	}
	return nil
}

func (pp *ZkFairEventProcessor) l1EventUnpack(event event.ContractEvent) error {
	l2chainId := pp.l2ChainId
	l1ChainId := pp.l1ChainId
	getAbi, err := bindings.PolygonzkevmbridgeMetaData.GetAbi()
	if err != nil {
		return err
	}
	switch event.EventSignature.String() {
	case getAbi.Events["BridgeEvent"].ID.String():
		err := bridge.L1Deposit(l1ChainId, l2chainId, pp.polygonBridge, event, pp.db)
		return err
	case getAbi.Events["ClaimEvent"].ID.String():
		err := bridge.L1Claimed(l1ChainId, l2chainId, pp.polygonBridge, event, pp.db)
		return err
	}

	return nil
}

func (pp *ZkFairEventProcessor) l2EventUnpack(event event.ContractEvent) error {
	l2chainId := pp.l2ChainId
	l1ChainId := pp.l1ChainId
	getAbi, err := bindings.PolygonzkevmbridgeMetaData.GetAbi()
	if err != nil {
		return err
	}
	switch event.EventSignature.String() {
	case getAbi.Events["BridgeEvent"].ID.String():
		err := bridge.L2Withdraw(l1ChainId, l2chainId, pp.polygonBridge, event, pp.db)
		return err
	case getAbi.Events["ClaimEvent"].ID.String():
		err := bridge.L2Claimed(l1ChainId, l2chainId, pp.polygonBridge, event, pp.db)
		return err
	}
	return nil
}
