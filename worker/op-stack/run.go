package op_stack

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/cornerstone-labs/acorus/database"
	"github.com/ethereum/go-ethereum/log"
)

const insertLoop = 100

type WorkerProcessor struct {
	log       log.Logger
	db        *database.DB
	redis     *redis.Client
	listeners chan interface{}
}

func NewWorkerProcessor(logger log.Logger, db *database.DB, redis *redis.Client) *WorkerProcessor {
	workerProcessor := WorkerProcessor{
		log:       logger,
		db:        db,
		redis:     redis,
		listeners: make(chan interface{}, 1),
	}
	return &workerProcessor
}

func (b *WorkerProcessor) Start(ctx context.Context) error {
	return nil
}

func (b *WorkerProcessor) run() error {
	if err := b.syncL2ToL1StateRoot(); err != nil {
		log.Info("err info", "err", err)
		return err
	}

	time.Sleep(time.Second * 5)
	b.listeners <- struct{}{}
	return nil
}

func (b *WorkerProcessor) syncL2ToL1StateRoot() error {
	blockNumber, err := b.db.StateRoots.GetLatestStateRootL2BlockNumber()
	if err != nil {
		b.log.Error(err.Error())
		return err
	}
	if blockNumber == 0 {
		return nil
	}

	blockTimestamp, err := b.db.Blocks.L2BlockTimeStampByNum(blockNumber)
	if err != nil {
		b.log.Error(err.Error())
		return err
	}

	err = b.db.L2ToL1.UpdateReadyForProvedStatus(blockTimestamp, 1)
	if err != nil {
		b.log.Error(err.Error())
		return err
	}

	b.log.Info("update ready for proven status success")
	return nil
}
