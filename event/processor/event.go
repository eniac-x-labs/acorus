package processor

import (
	"context"
	"fmt"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	op_stack "github.com/cornerstone-labs/acorus/event/processor/op-stack"
	"github.com/cornerstone-labs/acorus/synchronizer"
	"github.com/ethereum/go-ethereum/log"
	"time"
)

type EventProcessor struct {
	log              log.Logger
	tasks            tasks.Group
	OpEventProcessor *op_stack.OpProcessor
}

func NewEventProcessor(
	log log.Logger,
	db *database.DB,
	l1Sync *synchronizer.L1Sync,
	OpSync *synchronizer.L2Sync,
	chainConfig config.ChainConfig,
	shutdown context.CancelCauseFunc,
) (*EventProcessor, error) {
	opProcessor, err := op_stack.NewBridgeProcessor(log, db, l1Sync, OpSync, chainConfig, shutdown)
	if err != nil {
		return nil, err
	}
	return &EventProcessor{
		log: log,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error: %w", err))
		}},
		OpEventProcessor: opProcessor,
	}, nil
}

func (ep *EventProcessor) Start() error {
	tickerRun := time.NewTicker(time.Second * 5)
	ep.tasks.Go(func() error {
		for range tickerRun.C {
			err := ep.Start()
			if err != nil {
				return err
			}
		}
		ep.log.Info("no event updates. shut down business task")
		return nil
	})
	return nil
}

func (ep *EventProcessor) Stop() error {
	err := ep.Stop()
	if err != nil {
		return err
	}
	return nil
}
