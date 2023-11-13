package synchronizer

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/synchronizer/node"
	"github.com/cornerstone-labs/acorus/synchronizer/retry"
)

type L2Sync struct {
	Synchronizer

	db *database.DB
}

func NewL2Sync(cfg Config, log log.Logger, db *database.DB, client node.EthClient, contracts config.L2Contracts) (*L2Sync, error) {
	log = log.New("syncer", "l2")

	zeroAddr := common.Address{}
	l2Contracts := []common.Address{}
	if err := contracts.ForEach(func(name string, addr common.Address) error {
		// Since we dont have backfill support yet, we want to make sure all expected
		// contracts are specified to ensure consistent behavior. Once backfill support
		// is ready, we can relax this requirement.
		if addr == zeroAddr {
			log.Error("address not configured", "name", name)
			return errors.New("all L2Contracts must be configured")
		}

		log.Info("configured contract", "name", name, "addr", addr)
		l2Contracts = append(l2Contracts, addr)
		return nil
	}); err != nil {
		return nil, err
	}

	latestHeader, err := db.Blocks.L2LatestBlockHeader()
	if err != nil {
		return nil, err
	}

	var fromHeader *types.Header
	if latestHeader != nil {
		log.Info("detected last indexed block", "number", latestHeader.Number, "hash", latestHeader.Hash)
		fromHeader = latestHeader.RLPHeader.Header()
	} else {
		log.Info("no indexed state, starting from genesis")
	}

	syncerBatches := make(chan SynchronizerBatch)
	syncer := Synchronizer{
		loopInterval:     time.Duration(cfg.LoopIntervalMsec) * time.Millisecond,
		headerBufferSize: uint64(cfg.HeaderBufferSize),

		log:             log,
		headerTraversal: node.NewHeaderTraversal(client, fromHeader, cfg.ConfirmationDepth),
		contracts:       l2Contracts,
		syncerBatches:   syncerBatches,

		EthClient: client,
	}

	return &L2Sync{Synchronizer: syncer, db: db}, nil
}

func (l2Sync *L2Sync) Start(ctx context.Context) error {
	errCh := make(chan error, 1)
	go func() {
		errCh <- l2Sync.Synchronizer.Start(ctx)
	}()

	for {
		select {
		case err := <-errCh:
			return err

		// Index incoming batches (all L2 Blocks)
		case batch := <-l2Sync.syncerBatches:
			l2BlockHeaders := make([]common2.L2BlockHeader, len(batch.Headers))
			for i := range batch.Headers {
				l2BlockHeaders[i] = common2.L2BlockHeader{BlockHeader: common2.BlockHeaderFromHeader(&batch.Headers[i])}
			}

			l2ContractEvents := make([]event.L2ContractEvent, len(batch.Logs))
			for i := range batch.Logs {
				timestamp := batch.HeaderMap[batch.Logs[i].BlockHash].Time
				l2ContractEvents[i] = event.L2ContractEvent{ContractEvent: event.ContractEventFromLog(&batch.Logs[i], timestamp)}
			}

			// Continually try to persist this batch. If it fails after 10 attempts, we simply error out
			retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
			if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
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

			batch.Logger.Info("synced batch")
		}
	}
}
