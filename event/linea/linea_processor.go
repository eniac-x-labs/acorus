package linea

import (
	"context"
	"fmt"

	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/linea/abi"
	"github.com/cornerstone-labs/acorus/event/linea/bridge"
	"github.com/ethereum/go-ethereum/common"
)

type LineaEventProcessor struct {
	db             *database.DB
	cfgRpc         *config.RPC
	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group
	loopInterval   time.Duration
	l1StartHeight  *big.Int
	l2StartHeight  *big.Int
	epoch          uint64
}

func NewBridgeProcessor(db *database.DB,
	cfg *config.RPC, shutdown context.CancelCauseFunc,
	loopInterval time.Duration, epoch uint64) (*LineaEventProcessor, error) {
	resCtx, resCancel := context.WithCancel(context.Background())
	return &LineaEventProcessor{
		db:             db,
		cfgRpc:         cfg,
		resourceCtx:    resCtx,
		resourceCancel: resCancel,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
		loopInterval: loopInterval,
		epoch:        epoch,
	}, nil
}

func (lp *LineaEventProcessor) StartUnpack() error {
	tickerSyncer := time.NewTicker(lp.loopInterval)
	log.Println("starting scroll bridge processor...")
	lp.tasks.Go(func() error {
		for range tickerSyncer.C {
			err := lp.onL1Data()
			if err != nil {
				log.Println("no more l1 etl updates. shutting down l1 task")
				return err
			}
		}
		return nil
	})
	// start L2 worker
	lp.tasks.Go(func() error {
		for range tickerSyncer.C {
			err := lp.onL2Data()
			if err != nil {
				log.Println("no more l2 etl updates. shutting down l2 task")
				return err
			}
		}
		return nil
	})
	return nil
}

func (lp *LineaEventProcessor) ClosetUnpack() error {
	lp.resourceCancel()
	return lp.tasks.Wait()
}

func (lp *LineaEventProcessor) onL1Data() error {
	chainId := lp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))
	if lp.l1StartHeight == nil {
		lastBlockHeard, err := lp.db.L1ToL2.L1LatestBlockHeader(chainIdStr)
		if err != nil {
			log.Println("l1 failed to get last block heard", "err", err)
			return err
		}
		if lastBlockHeard == nil {
			lastBlockHeard = &common2.BlockHeader{
				Number: big.NewInt(0),
			}
		}
		lp.l1StartHeight = lastBlockHeard.Number
		if lp.l1StartHeight.Cmp(big.NewInt(int64(lp.cfgRpc.L1EventUnpackBlock))) == -1 {
			lp.l1StartHeight = big.NewInt(int64(lp.cfgRpc.L1EventUnpackBlock))
		}
	} else {
		lp.l1StartHeight = new(big.Int).Add(lp.l1StartHeight, bigint.One)
	}

	fromL1Height := lp.l1StartHeight
	toL1Height := new(big.Int).Add(fromL1Height, big.NewInt(int64(lp.epoch)))
	if err := lp.db.Transaction(func(tx *database.DB) error {
		l1EventsFetchErr := lp.l1EventsFetch(fromL1Height, toL1Height)
		if l1EventsFetchErr != nil {
			return l1EventsFetchErr
		}
		return nil
	}); err != nil {
		return err
	}
	lp.l1StartHeight = toL1Height
	return nil
}

func (lp *LineaEventProcessor) onL2Data() error {
	chainId := lp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))
	if lp.l2StartHeight == nil {
		lastBlockHeard, err := lp.db.L2ToL1.L2LatestBlockHeader(chainIdStr)
		if err != nil {
			log.Println("l2 failed to get last block heard", "err", err)
			return err
		}
		if lastBlockHeard == nil {
			lastBlockHeard = &common2.BlockHeader{
				Number: big.NewInt(0),
			}
		}
		lp.l2StartHeight = lastBlockHeard.Number
	} else {
		lp.l2StartHeight = new(big.Int).Add(lp.l2StartHeight, bigint.One)
	}
	fromL2Height := lp.l2StartHeight
	toL2Height := new(big.Int).Add(fromL2Height, big.NewInt(int64(lp.epoch)))
	if err := lp.db.Transaction(func(tx *database.DB) error {
		l2EventsFetchErr := lp.l2EventsFetch(fromL2Height, toL2Height)
		if l2EventsFetchErr != nil {
			return l2EventsFetchErr
		}
		return nil
	}); err != nil {
		return err
	}
	lp.l2StartHeight = toL2Height
	return nil
}

func (lp *LineaEventProcessor) l1EventsFetch(fromL1Height, toL1Height *big.Int) error {
	chainId := lp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))
	l1Contracts := lp.cfgRpc.Contracts
	for _, l1contract := range l1Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l1contract)}
		events, err := lp.db.ContractEvents.ContractEventsWithFilter("1", contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			log.Println("failed to index L1ContractEventsWithFilter ", "err", err)
			return err
		}
		l1ToL2s := make([]worker.L1ToL2, 0)
		for _, contractEvent := range events {
			unpackEvent, unpackErr := lp.l1EventUnpack(contractEvent)
			if unpackErr != nil {
				log.Println("failed to index L1 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
			if unpackEvent != nil {
				l1ToL2s = append(l1ToL2s, *unpackEvent)
			}
		}
		if len(l1ToL2s) > 0 {
			saveErr := lp.db.L1ToL2.StoreL1ToL2Transactions(chainIdStr, l1ToL2s)
			if saveErr != nil {
				log.Println("failed to StoreL1ToL2Transactions", "saveErr", saveErr)
				return saveErr
			}
		}
	}
	return nil
}

func (lp *LineaEventProcessor) l2EventsFetch(fromL1Height, toL1Height *big.Int) error {
	chainId := lp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))
	l2Contracts := lp.cfgRpc.EventContracts
	for _, l2contract := range l2Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l2contract)}
		events, err := lp.db.ContractEvents.ContractEventsWithFilter(chainIdStr, contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			log.Println("failed to index L2ContractEventsWithFilter ", "err", err)
			return err
		}
		l2ToL1s := make([]worker.L2ToL1, 0)
		for _, contractEvent := range events {
			unpackEvent, unpackErr := lp.l2EventUnpack(contractEvent)
			if unpackErr != nil {
				log.Println("failed to index L2 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
			if unpackEvent != nil {
				l2ToL1s = append(l2ToL1s, *unpackEvent)
			}
		}
		if len(l2ToL1s) > 0 {
			saveErr := lp.db.L2ToL1.StoreL2ToL1Transactions(chainIdStr, l2ToL1s)
			if saveErr != nil {
				log.Println("failed to StoreL2ToL1Transactions", "saveErr", saveErr)
				return saveErr
			}
		}
	}
	return nil
}

func (lp *LineaEventProcessor) l1EventUnpack(event event.ContractEvent) (*worker.L1ToL2, error) {
	chainId := lp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))
	switch event.EventSignature.String() {
	case abi.L1SentMessageSignature.String():
		l1SentMsg, err := bridge.L1SentMessageEvent(event)
		if err != nil {
			return nil, err
		}
		return l1SentMsg, nil
	case abi.L1ClaimedMessageSignature.String():
		err := bridge.L1ClaimedMessageEvent(chainIdStr, event, lp.db)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	return nil, nil
}

func (lp *LineaEventProcessor) l2EventUnpack(event event.ContractEvent) (*worker.L2ToL1, error) {
	chainId := lp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))
	switch event.EventSignature.String() {
	case abi.L2SentMessageSignature.String():
		withdrawETH, err := bridge.L2SentMessageEvent(event)
		if err != nil {
			return nil, err
		}
		return withdrawETH, nil
	case abi.L2ClaimedMessageSignature.String():
		err := bridge.L2ClaimedMessageEvent(chainIdStr, event, lp.db)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	fmt.Println(chainIdStr)
	return nil, nil
}
