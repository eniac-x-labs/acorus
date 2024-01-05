package op_stack

import (
	"context"
	"fmt"
	"time"

	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/ethereum/go-ethereum/log"
)

type WorkerProcessor struct {
	log     log.Logger
	db      *database.DB
	chainId string
	tasks   tasks.Group
}

func NewWorkerProcessor(logger log.Logger, db *database.DB, chainId string, shutdown context.CancelCauseFunc) *WorkerProcessor {
	workerProcessor := WorkerProcessor{
		log:     logger,
		db:      db,
		chainId: chainId,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in business processor: %w", err))
		}},
	}
	return &workerProcessor
}

func (b *WorkerProcessor) Start(ctx context.Context) error {
	ticker := time.NewTicker(time.Second)
	b.tasks.Go(func() error {
		for range ticker.C {
			if err := b.db.L2ToL1.UpdateTimeLeft(b.chainId); err != nil {
				b.log.Error(err.Error())
			}
		}
		return nil
	})
	tickerRun := time.NewTicker(time.Second * 5)
	b.tasks.Go(func() error {
		for range tickerRun.C {
			err := b.syncL2ToL1StateRoot()
			if err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

func (b *WorkerProcessor) syncL2ToL1StateRoot() error {
	blockNumber, err := b.db.StateRoots.GetLatestStateRootL2BlockNumber(b.chainId)
	if err != nil {
		b.log.Error(err.Error())
		return err
	}
	if blockNumber == 0 {
		return nil
	}

	blockTimestamp, err := b.db.Blocks.BlockTimeStampByNum(b.chainId, blockNumber)
	if err != nil {
		b.log.Error(err.Error())
		return err
	}

	err = b.db.L2ToL1.UpdateReadyForProvedStatus(b.chainId, blockTimestamp, 1)
	if err != nil {
		b.log.Error(err.Error())
		return err
	}

	b.log.Info("update ready for proven status success")
	return nil
}
