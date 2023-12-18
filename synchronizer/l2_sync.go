package synchronizer

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	common1 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/synchronizer/node"
	"github.com/cornerstone-labs/acorus/synchronizer/retry"
)

type L2Sync struct {
	Synchronizer
	LatestHeader   *types.Header
	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group
	db             *database.DB
	mu             sync.Mutex
	listeners      []chan interface{}
}

func NewL2Sync(cfg Config, log log.Logger, db *database.DB, client node.EthClient, contracts config.ChainConfig, shutdown context.CancelCauseFunc) (*L2Sync, error) {
	log = log.New("syncer", "l2")

	l2Contracts := []common.Address{}

	latestHeader, err := db.Blocks.L2LatestBlockHeader()
	if err != nil {
		return nil, err
	}

	var fromHeader *types.Header
	if latestHeader != nil {
		log.Info("detected last indexed block", "number", latestHeader.Number, "hash", latestHeader.Hash)
		fromHeader = latestHeader.RLPHeader.Header()
	} else if cfg.StartHeight.BitLen() > 0 {
		log.Info("no indexed state starting from supplied L2 height", "height", cfg.StartHeight.String())
		header, err := client.BlockHeaderByNumber(cfg.StartHeight)
		if err != nil {
			return nil, fmt.Errorf("could not fetch starting block header: %w", err)
		}
		fromHeader = header
	} else {
		log.Info("no indexed state")
	}

	syncerBatches := make(chan *SynchronizerBatch)
	syncer := Synchronizer{
		loopInterval:     time.Duration(cfg.LoopIntervalMsec) * time.Millisecond,
		headerBufferSize: uint64(cfg.HeaderBufferSize),
		log:              log,
		headerTraversal:  node.NewHeaderTraversal(client, fromHeader, cfg.ConfirmationDepth),
		contracts:        l2Contracts,
		syncerBatches:    syncerBatches,
		EthClient:        client,
	}

	resCtx, resCancel := context.WithCancel(context.Background())
	return &L2Sync{
		Synchronizer:   syncer,
		LatestHeader:   fromHeader,
		resourceCtx:    resCtx,
		resourceCancel: resCancel,
		db:             db,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in L2 Synchronizer: %w", err))
		}},
	}, nil
}

func (l2Sync *L2Sync) Close() error {
	var result error
	if err := l2Sync.Synchronizer.Close(); err != nil {
		result = errors.Join(result, fmt.Errorf("failed to close internal Synchronizer: %w", err))
	}
	l2Sync.resourceCancel()
	if err := l2Sync.tasks.Wait(); err != nil {
		result = errors.Join(result, fmt.Errorf("failed to await batch handler completion: %w", err))
	}
	for i := range l2Sync.listeners {
		close(l2Sync.listeners[i])
	}
	return result
}

func (l2Sync *L2Sync) Start() error {
	l2Sync.log.Info("starting syncer...")

	if err := l2Sync.Synchronizer.Start(); err != nil {
		return fmt.Errorf("failed to start internal Synchronizer: %w", err)
	}

	l2Sync.tasks.Go(func() error {
		for batch := range l2Sync.syncerBatches {
			if err := l2Sync.handleBatch(batch); err != nil {
				return fmt.Errorf("failed to handle batch, stopping L2 Synchronizer: %w", err)
			}
		}
		l2Sync.log.Info("no more batches, shutting down batch handler")
		return nil
	})
	return nil
}

func (l2Sync *L2Sync) handleBatch(batch *SynchronizerBatch) error {
	l2BlockHeaders := make([]common1.L2BlockHeader, len(batch.Headers))
	for i := range batch.Headers {
		l2BlockHeaders[i] = common1.L2BlockHeader{BlockHeader: common1.BlockHeaderFromHeader(&batch.Headers[i])}
	}

	l2ContractEvents := make([]event.L2ContractEvent, len(batch.Logs))
	for i := range batch.Logs {
		timestamp := batch.HeaderMap[batch.Logs[i].BlockHash].Time
		l2ContractEvents[i] = event.L2ContractEvent{ContractEvent: event.ContractEventFromLog(&batch.Logs[i], timestamp, 1)}
	}

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](l2Sync.resourceCtx, 10, retryStrategy, func() (interface{}, error) {
		if err := l2Sync.db.Transaction(func(tx *database.DB) error {
			if err := tx.Blocks.StoreL2BlockHeaders(l2BlockHeaders); err != nil {
				return err
			}
			if len(l2ContractEvents) > 0 {
				if err := tx.ContractEvents.StoreL2ContractEvents(l2ContractEvents); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			batch.Logger.Error("unable to persist batch", "err", err)
			return nil, err
		}

		return nil, nil
	}); err != nil {
		return err
	}

	batch.Logger.Info("indexed batch")
	l2Sync.LatestHeader = &batch.Headers[len(batch.Headers)-1]

	l2Sync.mu.Lock()
	defer l2Sync.mu.Unlock()
	for i := range l2Sync.listeners {
		select {
		case l2Sync.listeners[i] <- struct{}{}:
		default:
		}
	}
	return nil
}

func (l2Sync *L2Sync) Notify() <-chan interface{} {
	receiver := make(chan interface{})
	l2Sync.mu.Lock()
	defer l2Sync.mu.Unlock()
	l2Sync.listeners = append(l2Sync.listeners, receiver)
	return receiver
}
