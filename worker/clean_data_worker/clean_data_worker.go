package clean_data_worker

import (
	"context"
	"fmt"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/ethereum/go-ethereum/log"
	"strconv"
	"time"
)

type WorkerProcessor struct {
	db     *database.DB
	cfgRpc config.Config
	tasks  tasks.Group
}

func NewWorkerProcessor(db *database.DB, cfgRpc config.Config, shutdown context.CancelCauseFunc) (*WorkerProcessor, error) {
	workerProcessor := WorkerProcessor{
		db:     db,
		cfgRpc: cfgRpc,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in worker processor: %w", err))
		}},
	}
	return &workerProcessor, nil
}

func (b *WorkerProcessor) WorkerStart() error {
	tickerRun := time.NewTicker(time.Second * 5)
	b.tasks.Go(func() error {
		for range tickerRun.C {
			rpCs := b.cfgRpc.RPCs
			for i := range rpCs {
				log.Info("clean table by chainId", "chainId", rpCs[i].ChainId)
				err := b.db.Blocks.CleanBlockHerders(strconv.Itoa(int(rpCs[i].ChainId)))
				if err != nil {
					log.Error("clean table by chainId", "chainId", rpCs[i].ChainId, "err ", err)
					continue
				}
			}
		}
		return nil
	})
	tickerRunEvent := time.NewTicker(time.Second * 5)
	b.tasks.Go(func() error {
		for range tickerRunEvent.C {
			rpCs := b.cfgRpc.RPCs
			for i := range rpCs {
				log.Info("clean events by chainId", "chainId", rpCs[i].ChainId)
				err := b.db.ContractEvents.CleanEventsByChainId(strconv.Itoa(int(rpCs[i].ChainId)))
				if err != nil {
					log.Error("clean events by chainId", "chainId", rpCs[i].ChainId, "err ", err)
					continue
				}
			}
		}
		return nil
	})
	return nil
}
