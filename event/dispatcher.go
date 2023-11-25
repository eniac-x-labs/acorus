package event

import (
	"context"

	"github.com/ethereum/go-ethereum/log"

	common2 "github.com/cornerstone-labs/acorus/common"
	"github.com/cornerstone-labs/acorus/config"
	op_stack2 "github.com/cornerstone-labs/acorus/config/op-stack"
	"github.com/cornerstone-labs/acorus/database"
	op_stack "github.com/cornerstone-labs/acorus/event/processors/op-stack"
	"github.com/cornerstone-labs/acorus/synchronizer"
)

type EventDispatcher struct {
	log               log.Logger
	opBridgeProcessor *op_stack.OpBridgeProcessor
	chainBridge       string
}

func NewEventDispatcher(log log.Logger, db *database.DB, L1Syncer *synchronizer.L1Sync, chainConfig config.ChainConfig, chainBridge string, contracts interface{}) (*EventDispatcher, error) {
	opBridgeProcessor, err := op_stack.NewOpBridgeProcessor(log, db, L1Syncer, chainConfig, contracts.(op_stack2.OpContracts))
	if err != nil {
		return nil, err
	}
	return &EventDispatcher{
		log:               log,
		opBridgeProcessor: opBridgeProcessor,
		chainBridge:       chainBridge,
	}, nil
}

func (dt *EventDispatcher) Start(ctx context.Context) error {
	if dt.chainBridge == common2.Op {
		err := dt.opBridgeProcessor.Start(ctx)
		if err != nil {
			return err
		}
	} else if dt.chainBridge == common2.Polygon {
		// todo: handle polygon
	} else if dt.chainBridge == common2.Scroll {
		// todo: handle scroll
	}
	return nil
}
