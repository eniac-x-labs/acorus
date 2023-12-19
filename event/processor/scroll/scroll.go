package scroll

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	common3 "github.com/cornerstone-labs/acorus/event/processor/common"
	"github.com/cornerstone-labs/acorus/synchronizer"
)

type ScrollProcessor struct {
	common3.IProcessor
}

func NewBridgeProcessor(log log.Logger, db *database.DB, l1Sync *synchronizer.L1Sync, l2Sync *synchronizer.L2Sync,
	chainConfig config.ChainConfig, shutdown context.CancelCauseFunc) (*common3.IProcessor, error) {
	latestL1Header, latestL2Header,
		latestFinalizedL1Header, latestFinalizedL2Header,
		err := common3.GetLastBlockHeardByChain(log, db, chainConfig)
	if err != nil {
		return nil, err
	}

	log.Info("detected indexed bridge state",
		"l1_block", latestL1Header, "l2_block", latestL2Header,
		"finalized_l1_block", latestFinalizedL1Header, "finalized_l2_block", latestFinalizedL2Header)

	resCtx, resCancel := context.WithCancel(context.Background())
	return &common3.IProcessor{
		Log:                   log,
		Db:                    db,
		L1Sync:                l1Sync,
		L2Sync:                l2Sync,
		ResourceCtx:           resCtx,
		ResourceCancel:        resCancel,
		ChainConfig:           chainConfig,
		LastL1Header:          latestL1Header,
		LastL2Header:          latestL2Header,
		LastRollupL1Header:    latestL1Header,
		LastFinalizedL1Header: latestFinalizedL1Header,
		LastFinalizedL2Header: latestFinalizedL2Header,
		Tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
	}, nil
}

func (sp *ScrollProcessor) Start() error {
	sp.Log.Info("starting bridge processor...")
	sp.Tasks.Go(func() error {
		l1SyncUpdates := sp.L1Sync.Notify()
		for range l1SyncUpdates {
			err := sp.onL1Data()
			if err != nil {
				return err
			}
		}
		sp.Log.Info("no more l1 etl updates. shutting down l1 task")
		return nil
	})
	// start L2 worker
	sp.Tasks.Go(func() error {
		l2SyncUpdates := sp.L2Sync.Notify()
		for range l2SyncUpdates {
			err := sp.onL2Data()
			if err != nil {
				return err
			}
		}
		sp.Log.Info("no more l2 etl updates. shutting down l2 task")
		return nil
	})
	return nil
}

func (sp *ScrollProcessor) Close() error {
	sp.ResourceCancel()
	return sp.Tasks.Wait()
}

func (sp *ScrollProcessor) onL1Data() error {
	return nil
}

func (sp *ScrollProcessor) onL2Data() error {
	return nil
}
