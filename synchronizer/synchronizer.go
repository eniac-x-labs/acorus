package synchronizer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/synchronizer/node"
	"github.com/cornerstone-labs/acorus/synchronizer/retry"
)

type Config struct {
	LoopIntervalMsec  uint
	HeaderBufferSize  uint
	StartHeight       *big.Int
	ConfirmationDepth *big.Int
	ChainId           uint
}

type Synchronizer struct {
	log              log.Logger
	loopInterval     time.Duration
	headerBufferSize uint64
	headerTraversal  *node.HeaderTraversal
	ethClient        node.EthClient
	headers          []types.Header
	LatestHeader     *types.Header
	resourceCtx      context.Context
	resourceCancel   context.CancelFunc
	tasks            tasks.Group
	db               *database.DB
	chainId          string
}

func NewSynchronizer(cfg *Config, log log.Logger, db *database.DB, client node.EthClient, shutdown context.CancelCauseFunc) (*Synchronizer, error) {
	log = log.New("synchronizer")
	latestHeader, err := db.Blocks.ChainLatestBlockHeader(strconv.Itoa(int(cfg.ChainId)))
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
	chainIdInt, _ := strconv.Atoi(strconv.Itoa(int(cfg.ChainId)))
	resCtx, resCancel := context.WithCancel(context.Background())
	return &Synchronizer{
		loopInterval:     time.Duration(cfg.LoopIntervalMsec) * time.Millisecond,
		headerBufferSize: uint64(cfg.HeaderBufferSize),
		log:              log,
		headerTraversal:  node.NewHeaderTraversal(client, fromHeader, cfg.ConfirmationDepth, uint(chainIdInt)),
		ethClient:        client,
		LatestHeader:     fromHeader,
		db:               db,
		resourceCtx:      resCtx,
		resourceCancel:   resCancel,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in L1 Synchronizer: %w", err))
		}},
		chainId: strconv.Itoa(int(cfg.ChainId)),
	}, nil
}

func (syncer *Synchronizer) Start() error {
	tickerSyncer := time.NewTicker(syncer.loopInterval)
	syncer.tasks.Go(func() error {
		for range tickerSyncer.C {
			if len(syncer.headers) > 0 {
				syncer.log.Info("retrying previous batch")
			} else {
				newHeaders, err := syncer.headerTraversal.NextHeaders(syncer.headerBufferSize)
				if err != nil {
					syncer.log.Error("error querying for headers", "err", err)
				} else if len(newHeaders) == 0 {
					syncer.log.Warn("no new headers. syncer at head?")
				} else {
					syncer.headers = newHeaders
				}
				latestHeader := syncer.headerTraversal.LatestHeader()
				if latestHeader != nil {
					syncer.log.Info("Latest header", "latestHeader Number", latestHeader.Number)
				}
			}
			err := syncer.processBatch(syncer.headers)
			if err == nil {
				syncer.headers = nil
			}
		}
		return nil
	})
	return nil
}

func (syncer *Synchronizer) processBatch(headers []types.Header) error {
	if len(headers) == 0 {
		return nil
	}
	firstHeader, lastHeader := headers[0], headers[len(headers)-1]
	batchLog := syncer.log.New("batch_start_block_number", firstHeader.Number, "batch_end_block_number", lastHeader.Number)
	batchLog.Info("extracting batch", "size", len(headers))

	headerMap := make(map[common.Hash]*types.Header, len(headers))
	for i := range headers {
		header := headers[i]
		headerMap[header.Hash()] = &header
	}
	headersWithLog := make(map[common.Hash]bool, len(headers))
	filterQuery := ethereum.FilterQuery{FromBlock: firstHeader.Number, ToBlock: lastHeader.Number}
	logs, err := syncer.ethClient.FilterLogs(filterQuery)
	if err != nil {
		batchLog.Info("failed to extract logs", "err", err)
		return err
	}

	if logs.ToBlockHeader.Number.Cmp(lastHeader.Number) != 0 {
		return fmt.Errorf("mismatch in FilterLog#ToBlock number")
	} else if logs.ToBlockHeader.Hash() != lastHeader.Hash() {
		return fmt.Errorf("mismatch in FitlerLog#ToBlock block hash!!!")
	}

	if len(logs.Logs) > 0 {
		batchLog.Info("detected logs", "size", len(logs.Logs))
	}

	for i := range logs.Logs {
		logEvent := logs.Logs[i]
		headersWithLog[logEvent.BlockHash] = true
		if _, ok := headerMap[logEvent.BlockHash]; !ok {
			batchLog.Error("log found with block hash not in the batch", "block_hash", logs.Logs[i].BlockHash, "log_index", logs.Logs[i].Index)
			return errors.New("parsed log with a block hash not in the batch")
		}
	}
	chainBlockHeaders := make([]common2.ChainBlockHeader, 0, len(headers))
	for i := range headers {
		chainBlockHeaders = append(chainBlockHeaders, common2.ChainBlockHeader{BlockHeader: common2.BlockHeaderFromHeader(&headers[i])})
	}

	chainContractEvent := make([]event.ChainContractEvent, len(logs.Logs))
	for i := range logs.Logs {
		timestamp := headerMap[logs.Logs[i].BlockHash].Time
		chainContractEvent[i] = event.ChainContractEvent{ContractEvent: event.ContractEventFromLog(&logs.Logs[i], timestamp)}
	}

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](syncer.resourceCtx, 10, retryStrategy, func() (interface{}, error) {
		if err := syncer.db.Transaction(func(tx *database.DB) error {
			if err := tx.Blocks.StoreBlockHeaders(syncer.chainId, chainBlockHeaders); err != nil {
				return err
			}
			if err := tx.ContractEvents.StoreChainContractEvents(syncer.chainId, chainContractEvent); err != nil {
				return err
			}
			return nil
		}); err != nil {
			syncer.log.Error("unable to persist batch", "err", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return err
	}
	return nil
}

func (syncer *Synchronizer) Close() error {
	return nil
}
