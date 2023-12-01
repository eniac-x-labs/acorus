package event

import (
	"context"
	op_stack "github.com/cornerstone-labs/acorus/event/processors/op-stack"
	"github.com/cornerstone-labs/acorus/event/processors/scroll"

	"github.com/ethereum/go-ethereum/log"

	common2 "github.com/cornerstone-labs/acorus/common"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/synchronizer"
)

type EventDispatcher struct {
	log               log.Logger
	opBridgeProcessor interface{}
	chainBridge       string
}

func NewEventDispatcher(log log.Logger, db *database.DB, L1Syncer *synchronizer.L1Sync, chainConfig config.ChainConfig, chainBridge string, contracts interface{}) (*EventDispatcher, error) {
	var bridgeProcessor interface{}
	var err error
	if chainBridge == common2.Op {
		bridgeProcessor, err = op_stack.NewOpBridgeProcessor(log, db, L1Syncer, chainConfig, contracts)
		if err != nil {
			return nil, err
		}
	} else if chainBridge == common2.Polygon {
		// todo: handle polygon
	} else if chainBridge == common2.Scroll {
		bridgeProcessor, err = scroll.NewScrollBridgeProcessor(log, db, L1Syncer, chainConfig, contracts)
		if err != nil {
			return nil, err
		}
	} else if chainBridge == common2.Linea {
		bridgeProcessor, err = linea.NewLineaBridgeProcessor(log, db, L1Syncer, chainConfig, contracts)

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
		// todo: handle polygon
	} else if dt.chainBridge == common2.Scroll {
		processor := dt.opBridgeProcessor.(*scroll.ScBridgeProcessor)
		err := processor.Start(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
