package worker

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/worker/l1l2"
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
	done := ctx.Done()

	startup := make(chan interface{}, 1)
	startup <- nil

	b.log.Info("starting worker processor...")

	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			if err := b.db.L2ToL1.UpdateTimeLeft(); err != nil {
				b.log.Error(err.Error())
			}
		}
	}()
	defer ticker.Stop()

	for {
		select {
		case <-done:
			b.log.Info("stopping bridge processor")
			return nil

		// Tickers
		case <-startup:
		case <-b.listeners:
		}

		b.run()
	}
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

	err = b.db.L2ToL1.UpdateReadyForProvedStatus(blockTimestamp)
	if err != nil {
		b.log.Error(err.Error())
		return err
	}

	b.log.Info("update ready for proven status success")
	return nil
}

func (b *WorkerProcessor) syncL1ToL2Worker() error {
	startTimestamp := uint64(1)

	val, err := b.redis.Get(context.Background(), "l1StartTimestamp").Result()
	if err != nil && err.Error() != "redis: nil" {
		b.log.Error("fail to get key value")
		return err
	}

	if val == "" {
		startTimestamp = b.db.BridgeTransfers.GetL1FirstDepositTimestamp()
		if startTimestamp == 1 {
			log.Info("wait the first deposit transaction")
			return nil
		}
	}
	num, _ := strconv.ParseUint(val, 10, 64)
	startTimestamp = num

	endTimestamp := b.db.BridgeTransfers.GetL1EndDepositTimestamp(startTimestamp, insertLoop)

	if endTimestamp+1 == startTimestamp {
		log.Info("wait the next deposit transaction")
		return nil
	}

	res, err := b.db.BridgeTransfers.StoreL1BridgeDepositsByTimestamp(startTimestamp, endTimestamp)
	l1ToL2 := l1l2.ConstructL1ToL2(res)
	if err := b.db.L1ToL2.StoreL1ToL2Transactions(l1ToL2); err != nil {
		b.log.Error(err.Error())
		return err
	}

	b.log.Info("store l1 to l2 success")
	startTimestamp = endTimestamp + 1
	err = b.redis.Set(context.Background(), "l1StartTimestamp", startTimestamp, 0).Err()
	if err != nil {
		b.log.Error("fail to set key value")
		return err
	}

	return nil
}

func (b *WorkerProcessor) syncL2ToL1Worker() error {
	startTimestamp := uint64(1)

	val, err := b.redis.Get(context.Background(), "l2StartTimestamp").Result()
	if err != nil && err.Error() != "redis: nil" {
		b.log.Error("fail to get key value")
		return err
	}

	if val == "" {
		startTimestamp = b.db.BridgeTransfers.GetL2FirstWithdrawTimestamp()
		if startTimestamp == 1 {
			log.Info("wait the first withdraw transaction")
			return nil
		}
	}
	num, _ := strconv.ParseUint(val, 10, 64)
	startTimestamp = num
	endTimestamp := b.db.BridgeTransfers.GetL2EndWithdrawTimestamp(startTimestamp, insertLoop)

	if endTimestamp+1 == startTimestamp {
		log.Info("wait the next withdraw transaction")
		return nil
	}

	res, err := b.db.BridgeTransfers.StoreL2BridgeWithdrawsByTimestamp(startTimestamp, endTimestamp)
	l2ToL1 := l1l2.ConstructL2ToL1(res)
	if err := b.db.L2ToL1.StoreL2ToL1Transactions(l2ToL1); err != nil {
		b.log.Error(err.Error())
		return err
	}

	startTimestamp = endTimestamp + 1
	err = b.redis.Set(context.Background(), "l2StartTimestamp", startTimestamp, 0).Err()
	if err != nil {
		b.log.Error("fail to set key value")
		return err
	}

	b.log.Info("store l2 to l1 success")
	return nil
}
