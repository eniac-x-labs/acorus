package synchronizer

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/common/bigint"
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

func NewSynchronizer(cfg *Config, db *database.DB, client node.EthClient, shutdown context.CancelCauseFunc) (*Synchronizer, error) {
	latestHeader, err := db.Blocks.ChainLatestBlockHeader(strconv.Itoa(int(cfg.ChainId)))
	if err != nil {
		return nil, err
	}
	var fromHeader *types.Header
	if latestHeader != nil {
		log.Info("detected last indexed block", "number", latestHeader.Number, "hash", latestHeader.Hash)
		fromHeader = latestHeader.RLPHeader.Header()
	} else if cfg.StartHeight.BitLen() > 0 {
		log.Info("no indexed state starting from supplied L1 height;", "height =", cfg.StartHeight.String())
		header, err := client.BlockHeaderByNumber(new(big.Int).Sub(cfg.StartHeight, bigint.One))
		if err != nil {
			log.Error("Fetch block header by number fail", "err", err)
			return nil, fmt.Errorf("could not fetch starting block header: %w", err)
		}
		fromHeader = header
	} else {
		log.Info("no indexed state, starting from genesis")
	}
	chainIdInt, _ := strconv.Atoi(strconv.Itoa(int(cfg.ChainId)))
	log.Info("Support chain", "chainId", chainIdInt)
	resCtx, resCancel := context.WithCancel(context.Background())
	return &Synchronizer{
		loopInterval:     time.Duration(cfg.LoopIntervalMsec) * time.Second,
		headerBufferSize: uint64(cfg.HeaderBufferSize),
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
				log.Info("sync", "chain ", syncer.chainId, "retrying previous batch")
			} else {
				newHeaders, err := syncer.headerTraversal.NextHeaders(syncer.headerBufferSize)
				if err != nil {
					log.Error("sync", "chain ", syncer.chainId, "error querying for headers err", err)
					continue
				} else if len(newHeaders) == 0 {
					log.Warn("sync", "chain ", syncer.chainId, "no new headers. syncer at head?", 0)
				} else {
					syncer.headers = newHeaders
				}
				latestHeader := syncer.headerTraversal.LatestHeader()
				if latestHeader != nil {
					log.Warn("sync", "chain ", syncer.chainId, "latestHeader Number", latestHeader.Number)
				}
			}
			err := syncer.processBatch(syncer.headers)
			if err != nil {
				continue
			} else {
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
	log.Info("processBatch", "chain ", syncer.chainId, "extracting batch size", len(headers),
		"firstHeader", firstHeader.Number, "lastHeader", lastHeader.Number)

	headerMap := make(map[common.Hash]*types.Header, len(headers))
	for i := range headers {
		header := headers[i]
		headerMap[header.Hash()] = &header
	}
	chainIdInt, _ := strconv.Atoi(syncer.chainId)

	filterQuery := ethereum.FilterQuery{FromBlock: firstHeader.Number, ToBlock: lastHeader.Number}
	logs, err := syncer.ethClient.FilterLogs(filterQuery, uint(chainIdInt))
	if err != nil {
		log.Error("processBatch", "chain ", syncer.chainId, "failed to extract logs err", err)
		return err
	}

	if logs.ToBlockHeader.Number.Cmp(lastHeader.Number) != 0 {
		return fmt.Errorf("mismatch in FilterLog#ToBlock number")
	} else if logs.ToBlockHeader.Hash() != lastHeader.Hash() {
		return fmt.Errorf("mismatch in FitlerLog#ToBlock block hash!!!")
	}

	if len(logs.Logs) > 0 {
		log.Info("processBatch", "chain ", syncer.chainId, "detected logs size", len(logs.Logs))
	}

	chainBlockHeaders := make([]common2.ChainBlockHeader, 0, len(headers))
	for i := range headers {
		if headers[i].Number == nil {
			continue
		}
		chainBlockHeaders = append(chainBlockHeaders, common2.ChainBlockHeader{BlockHeader: common2.BlockHeaderFromHeader(&headers[i])})
	}

	chainContractEvents := make([]event.ChainContractEvent, 0)
	for i := range logs.Logs {
		logEvent := logs.Logs[i]
		if _, ok := headerMap[logEvent.BlockHash]; !ok {
			continue
		}
		timestamp := headerMap[logEvent.BlockHash].Time
		blockNumber := headerMap[logEvent.BlockHash].Number
		if blockNumber != nil {
			chainContractEvent := event.ChainContractEvent{ContractEvent: event.ContractEventFromLog(&logs.Logs[i], timestamp, blockNumber)}
			chainContractEvents = append(chainContractEvents, chainContractEvent)
		}
	}

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](syncer.resourceCtx, 10, retryStrategy, func() (interface{}, error) {
		if err := syncer.db.Transaction(func(tx *database.DB) error {
			if err := tx.Blocks.StoreBlockHeaders(syncer.chainId, chainBlockHeaders); err != nil {
				return err
			}
			if err := tx.ContractEvents.StoreChainContractEvents(syncer.chainId, chainContractEvents); err != nil {
				return err
			}
			return nil
		}); err != nil {
			log.Error("unable to persist batch", "err", err)
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
