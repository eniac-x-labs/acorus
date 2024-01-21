package linea

import (
	"context"
	"fmt"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"log"
	"math/big"
	"time"
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
	return nil
}

func (lp *LineaEventProcessor) onL2Data() error {
	return nil
}
