package worker

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/ethereum/go-ethereum/log"

	common2 "github.com/cornerstone-labs/acorus/common"
	"github.com/cornerstone-labs/acorus/database"
	op_stack "github.com/cornerstone-labs/acorus/worker/op-stack"
)

type WorkerDispatcher struct {
	log             log.Logger
	workerProcessor *op_stack.WorkerProcessor
	chainBridge     string
}

func NewWorkerDispatcher(logger log.Logger, db *database.DB, redis *redis.Client, chainBridge string) (*WorkerDispatcher, error) {
	workerProcessor := op_stack.NewWorkerProcessor(logger, db, redis)
	return &WorkerDispatcher{
		log:             logger,
		workerProcessor: workerProcessor,
		chainBridge:     chainBridge,
	}, nil
}

func (wd *WorkerDispatcher) Start(ctx context.Context) error {
	if wd.chainBridge == common2.Op {
		err := wd.workerProcessor.Start(ctx)
		if err != nil {
			return err
		}
	} else if wd.chainBridge == common2.Polygon {
		// todo: handle polygon
	} else if wd.chainBridge == common2.Scroll {
		// todo: handle scroll
	}
	return nil
}
