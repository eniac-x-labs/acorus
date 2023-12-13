package event

import (
	"context"
	"github.com/cornerstone-labs/acorus/event/processors/polygon"

	"github.com/cornerstone-labs/acorus/metrics"

	"github.com/ethereum/go-ethereum/log"

	common2 "github.com/cornerstone-labs/acorus/common"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/event/processors/linea"
	op_stack "github.com/cornerstone-labs/acorus/event/processors/op-stack"
	"github.com/cornerstone-labs/acorus/event/processors/scroll"
	"github.com/cornerstone-labs/acorus/synchronizer"
)

type EventDispatcher struct {
	log               log.Logger
	opBridgeProcessor interface{}
	chainBridge       string
}

func NewEventDispatcher(log log.Logger, db *database.DB, L1Syncer *synchronizer.L1Sync, chainConfig config.ChainConfig, chainBridge string, contracts interface{}, L1URL string, L2URL string) (*EventDispatcher, error) {
	var bridgeProcessor interface{}
	var err error
	if chainBridge == common2.Op {
		metricsRegistry := metrics.NewRegistry()
		bridgeProcessor, err = op_stack.NewOpBridgeProcessor(log, db, metrics.NewMetrics(metricsRegistry), L1Syncer, chainConfig, contracts)
		if err != nil {
			return nil, err
		}
	} else if chainBridge == common2.Polygon {
		metricsRegistry := metrics.NewRegistry()
		bridgeProcessor, err = polygon.NewPolygonBridgeProcessor(log, db, metrics.NewMetrics(metricsRegistry), L1Syncer, chainConfig, contracts, L1URL, L2URL)
	} else if chainBridge == common2.Scroll {
		metricsRegistry := metrics.NewRegistry()
		bridgeProcessor, err = scroll.NewScrollBridgeProcessor(log, db, metrics.NewMetrics(metricsRegistry), L1Syncer, chainConfig, contracts)
		if err != nil {
			return nil, err
		}
	} else if chainBridge == common2.Linea {
		metricsRegistry := metrics.NewRegistry()
		bridgeProcessor, err = linea.NewLineaBridgeProcessor(log, db, metrics.NewMetrics(metricsRegistry), L1Syncer, chainConfig, contracts)

		if err != nil {
			return nil, err
		}
	}
	return &EventDispatcher{
		log:               log,
		opBridgeProcessor: bridgeProcessor,
		chainBridge:       chainBridge,
	}, nil
}

func (dt *EventDispatcher) Start(ctx context.Context) error {
	if dt.chainBridge == common2.Op {
		processor := dt.opBridgeProcessor.(*op_stack.OpBridgeProcessor)
		err := processor.Start(ctx)
		if err != nil {
			return err
		}
	} else if dt.chainBridge == common2.Polygon {
		processor := dt.opBridgeProcessor.(*polygon.PlBridgeProcessor)
		err := processor.Start(ctx)
		if err != nil {
			return err
		}
	} else if dt.chainBridge == common2.Scroll {
		processor := dt.opBridgeProcessor.(*scroll.ScBridgeProcessor)
		err := processor.Start(ctx)
		if err != nil {
			return err
		}
	} else if dt.chainBridge == common2.Linea {
		processor := dt.opBridgeProcessor.(*linea.LineaBridgeProcessor)
		err := processor.Start(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
