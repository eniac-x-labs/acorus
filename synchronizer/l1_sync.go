package synchronizer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/synchronizer/node"
	"github.com/cornerstone-labs/acorus/synchronizer/retry"
)

type L1Sync struct {
	Synchronizer
	preset    int
	db        *database.DB
	mu        *sync.Mutex
	listeners []chan interface{}
}

func NewL1Sync(cfg Config, log log.Logger, db *database.DB, client node.EthClient, l1Contracts []common.Address, preset int) (*L1Sync, error) {
	log = log.New("sync", "l1")

	latestHeader, err := db.Blocks.L1LatestBlockHeader()
	if err != nil {
		return nil, err
	}

	var fromHeader *types.Header
	if latestHeader != nil {
		log.Info("detected last indexed block", "number", latestHeader.Number, "hash", latestHeader.Hash)
		fromHeader = latestHeader.RLPHeader.Header()
	} else if cfg.StartHeight.BitLen() > 0 {
		log.Info("no indexed state starting from supplied L1 height", "height", cfg.StartHeight.String())
		header, err := client.BlockHeaderByNumber(cfg.StartHeight)
		if err != nil {
			return nil, fmt.Errorf("could not fetch starting block header: %w", err)
		}
		fromHeader = header
	} else {
		log.Info("no indexed state, starting from genesis")
	}

	syncerBatches := make(chan SynchronizerBatch)
	syncer := Synchronizer{
		loopInterval:     time.Duration(cfg.LoopIntervalMsec) * time.Millisecond,
		headerBufferSize: uint64(cfg.HeaderBufferSize),
		log:              log,
		headerTraversal:  node.NewHeaderTraversal(client, fromHeader, cfg.ConfirmationDepth),
		contracts:        l1Contracts,
		syncerBatches:    syncerBatches,
		EthClient:        client,
	}

	return &L1Sync{Synchronizer: syncer, db: db, mu: new(sync.Mutex), preset: preset}, nil
}

func (l1Sync *L1Sync) Start(ctx context.Context) error {
	errCh := make(chan error, 1)
	go func() {
		errCh <- l1Sync.Synchronizer.Start(ctx)
	}()

	for {
		select {
		case err := <-errCh:
			return err

		case batch := <-l1Sync.syncerBatches:
			l1BlockHeaders := make([]common2.L1BlockHeader, 0, len(batch.Headers))
			for i := range batch.Headers {
				if _, ok := batch.HeadersWithLog[batch.Headers[i].Hash()]; ok {
					l1BlockHeaders = append(l1BlockHeaders, common2.L1BlockHeader{BlockHeader: common2.BlockHeaderFromHeader(&batch.Headers[i])})
				}
			}
			if len(l1BlockHeaders) == 0 {
				batch.Logger.Info("no l1 blocks with logs in batch")
				continue
			}

			l1ContractEvents := make([]event.L1ContractEvent, len(batch.Logs))
			for i := range batch.Logs {
				timestamp := batch.HeaderMap[batch.Logs[i].BlockHash].Time
				l1ContractEvents[i] = event.L1ContractEvent{ContractEvent: event.ContractEventFromLog(&batch.Logs[i], timestamp)}
			}

			retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
			if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
				if err := l1Sync.db.Transaction(func(tx *database.DB) error {
					if err := tx.Blocks.StoreL1BlockHeaders(l1BlockHeaders); err != nil {
						return err
					}
					if err := tx.ContractEvents.StoreL1ContractEvents(l1ContractEvents); err != nil {
						return err
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

			l1Sync.mu.Lock()
			for i := range l1Sync.listeners {
				select {
				case l1Sync.listeners[i] <- struct{}{}:
				default:

				}
			}
			l1Sync.mu.Unlock()
		}
	}
}

func (l1Sync *L1Sync) Notify() <-chan interface{} {
	receiver := make(chan interface{})
	l1Sync.mu.Lock()
	defer l1Sync.mu.Unlock()

	l1Sync.listeners = append(l1Sync.listeners, receiver)
	return receiver
}
