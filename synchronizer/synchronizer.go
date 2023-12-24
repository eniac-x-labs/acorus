package synchronizer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/clock"
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
	ChainId           uint
}

type Synchronizer struct {
	log              log.Logger
	loopInterval     time.Duration
	headerBufferSize uint64
	headerTraversal  *node.HeaderTraversal
	contracts        []common.Address
	syncerBatches    chan *SynchronizerBatch
	EthClient        node.EthClient
	headers          []types.Header
	worker           *clock.LoopFn
}

type SynchronizerBatch struct {
	Logger         log.Logger
	Headers        []types.Header
	HeaderMap      map[common.Hash]*types.Header
	Logs           []types.Log
	HeadersWithLog map[common.Hash]bool
}

func (syncer *Synchronizer) Start() error {
	if syncer.worker != nil {
		return errors.New("already started")
	}
	syncer.worker = clock.NewLoopFn(clock.SystemClock, syncer.tick, func() error {
		syncer.log.Info("shutting down batch producer")
		close(syncer.syncerBatches)
		return nil
	}, syncer.loopInterval)
	return nil
}

func (syncer *Synchronizer) Close() error {
	if syncer.worker == nil {
		return nil
	}
	return syncer.worker.Close()
}

func (syncer *Synchronizer) tick(_ context.Context) {
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
	filterQuery := ethereum.FilterQuery{FromBlock: firstHeader.Number, ToBlock: lastHeader.Number, Addresses: syncer.contracts}
	logs, err := syncer.EthClient.FilterLogs(filterQuery)
	if err != nil {
		batchLog.Info("failed to extract logs", "err", err)
		return err
	}

	if logs.ToBlockHeader.Number.Cmp(lastHeader.Number) != 0 {
		// Warn and simply wait for the provider to synchronize state
		batchLog.Warn("mismatch in FilterLog#ToBlock number", "queried_to_block_number", lastHeader.Number, "reported_to_block_number", logs.ToBlockHeader.Number)
		return fmt.Errorf("mismatch in FilterLog#ToBlock number")
	} else if logs.ToBlockHeader.Hash() != lastHeader.Hash() {
		batchLog.Error("mismatch in FitlerLog#ToBlock block hash!!!", "queried_to_block_hash", lastHeader.Hash().String(), "reported_to_block_hash", logs.ToBlockHeader.Hash().String())
		return fmt.Errorf("mismatch in FitlerLog#ToBlock block hash!!!")
	}

	if len(logs.Logs) > 0 {
		batchLog.Info("detected logs", "size", len(logs.Logs))
	}

	for i := range logs.Logs {
		log := logs.Logs[i]
		headersWithLog[log.BlockHash] = true
		if _, ok := headerMap[log.BlockHash]; !ok {
			// NOTE. Definitely an error state if the none of the headers were re-orged out in between
			// the blocks and logs retrieval operations. Unlikely as long as the confirmation depth has
			// been appropriately set or when we get to natively handling reorgs.
			batchLog.Error("log found with block hash not in the batch", "block_hash", logs.Logs[i].BlockHash, "log_index", logs.Logs[i].Index)
			return errors.New("parsed log with a block hash not in the batch")
		}
	}

	// ensure we use unique downstream references for the syncer batch
	headersRef := headers
	syncer.syncerBatches <- &SynchronizerBatch{Logger: batchLog, Headers: headersRef, HeaderMap: headerMap, Logs: logs.Logs, HeadersWithLog: headersWithLog}
	return nil
}
