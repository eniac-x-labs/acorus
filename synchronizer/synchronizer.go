package synchronizer

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/synchronizer/node"
)

type Config struct {
	LoopIntervalMsec  uint
	HeaderBufferSize  uint
	StartHeight       *big.Int
	ConfirmationDepth *big.Int
}

type Synchronizer struct {
	log              log.Logger
	loopInterval     time.Duration
	headerBufferSize uint64
	headerTraversal  *node.HeaderTraversal
	contracts        []common.Address
	syncerBatches    chan SynchronizerBatch
	EthClient        node.EthClient
	ChainId          uint64
}

type SynchronizerBatch struct {
	Logger         log.Logger
	Headers        []types.Header
	HeaderMap      map[common.Hash]*types.Header
	Logs           []types.Log
	HeadersWithLog map[common.Hash]bool
}

func (syncer *Synchronizer) Start(ctx context.Context) error {
	done := ctx.Done()
	pollTicker := time.NewTicker(syncer.loopInterval)
	defer pollTicker.Stop()

	var headers []types.Header

	syncer.log.Info("starting syncer...")
	for {
		select {
		case <-done:
			syncer.log.Info("stopping syncer")
			return nil

		case <-pollTicker.C:
			if len(headers) > 0 {
				syncer.log.Info("retrying previous batch")
			} else {
				newHeaders, err := syncer.headerTraversal.NextFinalizedHeaders(syncer.headerBufferSize)
				if err != nil {
					syncer.log.Error("error querying for headers", "err", err)
				} else if len(newHeaders) == 0 {
					syncer.log.Warn("no new headers. processor unexpectedly at head...")
				} else {
					headers = newHeaders
				}
			}
			err := syncer.processBatch(headers)
			if err == nil {
				headers = nil
			}
		}
	}
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
	logs, err := syncer.EthClient.FilterLogs(ethereum.FilterQuery{FromBlock: firstHeader.Number, ToBlock: lastHeader.Number, Addresses: syncer.contracts})
	if err != nil {
		batchLog.Info("failed to extract logs", "err", err)
		return err
	}
	if len(logs) > 0 {
		batchLog.Info("detected logs", "size", len(logs))
	}

	for i := range logs {
		log := logs[i]
		if _, ok := headerMap[log.BlockHash]; !ok {
			batchLog.Error("log found with block hash not in the batch", "block_hash", logs[i].BlockHash, "log_index", logs[i].Index)
			return errors.New("parsed log with a block hash not in the batch")
		}
		headersWithLog[log.BlockHash] = true
	}

	headersRef := headers
	syncer.syncerBatches <- SynchronizerBatch{Logger: batchLog, Headers: headersRef, HeaderMap: headerMap, Logs: logs, HeadersWithLog: headersWithLog}
	return nil
}
