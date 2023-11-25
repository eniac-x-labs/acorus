package op_stack

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/config"
	op_stack "github.com/cornerstone-labs/acorus/config/op-stack"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/event/processors/op-stack/bridge"
	"github.com/cornerstone-labs/acorus/synchronizer"
)

type OpBridgeProcessor struct {
	log log.Logger
	db  *database.DB

	l1Etl       *synchronizer.L1Sync
	chainConfig config.ChainConfig
	opContracts interface{}

	LatestL1Header *types.Header
	LatestL2Header *types.Header
}

func NewOpBridgeProcessor(log log.Logger, db *database.DB, L1Syncer *synchronizer.L1Sync, chainConfig config.ChainConfig, opContracts interface{}) (*OpBridgeProcessor, error) {
	log = log.New("processor", "bridge")

	latestL1Header, err := db.BridgeTransactions.L1LatestBlockHeader()
	if err != nil {
		return nil, err
	}
	latestL2Header, err := db.BridgeTransactions.L2LatestBlockHeader()
	if err != nil {
		return nil, err
	}

	var l1Header, l2Header *types.Header
	if latestL1Header == nil && latestL2Header == nil {
		log.Info("no indexed state, starting from rollup genesis")
	} else {
		l1Height, l2Height := bigint.Zero, bigint.Zero
		if latestL1Header != nil {
			l1Height = latestL1Header.Number
			l1Header = latestL1Header.RLPHeader.Header()
		}
		if latestL2Header != nil {
			l2Height = latestL2Header.Number
			l2Header = latestL2Header.RLPHeader.Header()
		}
		log.Info("detected latest indexed bridge state", "l1_block_number", l1Height, "l2_block_number", l2Height)
	}
	return &OpBridgeProcessor{log, db, L1Syncer, chainConfig, opContracts, l1Header, l2Header}, nil
}

func (b *OpBridgeProcessor) Start(ctx context.Context) error {
	done := ctx.Done()
	l1EtlUpdates := b.l1Etl.Notify()
	startup := make(chan interface{}, 1)
	startup <- nil

	b.log.Info("starting bridge processor...")
	for {
		select {
		case <-done:
			b.log.Info("stopping bridge processor")
			return nil

		// Tickers
		case <-startup:
		case <-l1EtlUpdates:
		}

		b.run()
	}
}

func (b *OpBridgeProcessor) run() error {
	maxEpochRange := uint64(10_000)
	var lastEpoch *big.Int
	if b.LatestL1Header != nil {
		lastEpoch = b.LatestL1Header.Number
	}

	latestEpoch, err := b.db.Blocks.LatestObservedEpoch(lastEpoch, maxEpochRange)
	if err != nil {
		return err
	} else if latestEpoch == nil {
		if b.LatestL1Header != nil || b.LatestL2Header != nil {
			b.log.Error("bridge events indexed, but no observed epoch returned", "latest_bridge_l1_block_number", b.LatestL1Header.Number)
			return errors.New("bridge events indexed, but no observed epoch returned")
		}
		b.log.Warn("no observed epochs available. waiting...")
		return nil
	}

	if b.LatestL1Header != nil && latestEpoch.L1BlockHeader.Hash == b.LatestL1Header.Hash() {
		b.log.Warn("all available epochs indexed", "latest_bridge_l1_block_number", b.LatestL1Header.Number)
		return nil
	}

	genesisL1Height := big.NewInt(int64(b.chainConfig.L2StartHeight))
	if latestEpoch.L1BlockHeader.Number.Cmp(genesisL1Height) < 0 {
		b.log.Error("L1 epoch less than starting L1 height observed", "l1_starting_number", genesisL1Height, "latest_epoch_number", latestEpoch.L1BlockHeader.Number)
		return errors.New("L1 epoch less than starting L1 height observed")
	}
	if b.LatestL1Header != nil && latestEpoch.L1BlockHeader.Number.Cmp(b.LatestL1Header.Number) <= 0 {
		b.log.Error("non-increasing l1 block height observed", "latest_bridge_l1_block_number", b.LatestL1Header.Number, "latest_epoch_l1_block_number", latestEpoch.L1BlockHeader.Number)
		return errors.New("non-increasing l1 block height observed")
	}
	if b.LatestL2Header != nil && latestEpoch.L2BlockHeader.Number.Cmp(b.LatestL2Header.Number) <= 0 {
		b.log.Error("non-increasing l2 block height observed", "latest_bridge_l2_block_number", b.LatestL2Header.Number, "latest_epoch_l2_block_number", latestEpoch.L2BlockHeader.Number)
		return errors.New("non-increasing l2 block height observed")
	}

	toL1Height, toL2Height := latestEpoch.L1BlockHeader.Number, latestEpoch.L2BlockHeader.Number
	fromL1Height, fromL2Height := genesisL1Height, bigint.Zero
	if b.LatestL1Header != nil {
		fromL1Height = new(big.Int).Add(b.LatestL1Header.Number, bigint.One)
	}
	if b.LatestL2Header != nil {
		fromL2Height = new(big.Int).Add(b.LatestL2Header.Number, bigint.One)
	}

	l1BedrockStartingHeight := big.NewInt(int64(b.chainConfig.L1StartHeight))
	l2BedrockStartingHeight := big.NewInt(int64(b.chainConfig.L2StartHeight))

	batchLog := b.log.New("epoch_start_number", fromL1Height, "epoch_end_number", toL1Height)
	batchLog.Info("unobserved epochs", "latest_l1_block_number", fromL1Height, "latest_l2_block_number", fromL2Height)
	if err := b.db.Transaction(func(tx *database.DB) error {
		l1BridgeLog := b.log.New("bridge", "l1")
		l2BridgeLog := b.log.New("bridge", "l2")

		if l1BedrockStartingHeight.Cmp(fromL1Height) > 0 {
			legacyFromL1Height, legacyToL1Height := fromL1Height, toL1Height
			legacyFromL2Height, legacyToL2Height := fromL2Height, toL2Height
			if l1BedrockStartingHeight.Cmp(toL1Height) <= 0 {
				legacyToL1Height = new(big.Int).Sub(l1BedrockStartingHeight, bigint.One)
				legacyToL2Height = new(big.Int).Sub(l2BedrockStartingHeight, bigint.One)
			}

			l1BridgeLog = l1BridgeLog.New("mode", "legacy", "from_l1_block_number", legacyFromL1Height, "to_l1_block_number", legacyToL1Height)
			l1BridgeLog.Info("scanning for bridge events")

			l2BridgeLog = l2BridgeLog.New("mode", "legacy", "from_l2_block_number", legacyFromL2Height, "to_l2_block_number", legacyToL2Height)
			l2BridgeLog.Info("scanning for bridge events")

			if err := bridge.LegacyL1ProcessInitiatedBridgeEvents(l1BridgeLog, tx, b.opContracts.(op_stack.OpContracts).L1Contracts, legacyFromL1Height, legacyToL1Height); err != nil {
				batchLog.Error("failed to index legacy l1 initiated bridge events", "err", err)
				return err
			}
			if err := bridge.LegacyL2ProcessInitiatedBridgeEvents(l2BridgeLog, tx, b.opContracts.(op_stack.OpContracts).L2Contracts, legacyFromL2Height, legacyToL2Height); err != nil {
				batchLog.Error("failed to index legacy l2 initiated bridge events", "err", err)
				return err
			}

			if err := bridge.LegacyL1ProcessFinalizedBridgeEvents(l1BridgeLog, tx, b.l1Etl.EthClient, b.opContracts.(op_stack.OpContracts).L1Contracts, legacyFromL1Height, legacyToL1Height); err != nil {
				batchLog.Error("failed to index legacy l1 finalized bridge events", "err", err)
				return err
			}
			if err := bridge.LegacyL2ProcessFinalizedBridgeEvents(l2BridgeLog, tx, b.opContracts.(op_stack.OpContracts).L2Contracts, legacyFromL2Height, legacyToL2Height); err != nil {
				batchLog.Error("failed to index legacy l2l finalized bridge events", "err", err)
				return err
			}

			if legacyToL1Height.Cmp(toL1Height) == 0 {
				return nil
			}

			batchLog.Info("detected switch to bedrock", "l1_bedrock_starting_height", l1BedrockStartingHeight, "l2_bedrock_starting_height", l2BedrockStartingHeight)
			fromL1Height = l1BedrockStartingHeight
			fromL2Height = l2BedrockStartingHeight
		}

		l1BridgeLog = l1BridgeLog.New("from_l1_block_number", fromL1Height, "to_l1_block_number", toL1Height)
		l1BridgeLog.Info("scanning for bridge events")

		l2BridgeLog = l2BridgeLog.New("from_l2_block_number", fromL2Height, "to_l2_block_number", toL2Height)
		l2BridgeLog.Info("scanning for bridge events")

		if err := bridge.L1ProcessInitiatedBridgeEvents(l1BridgeLog, tx, b.opContracts.(op_stack.OpContracts).L1Contracts, fromL1Height, toL1Height); err != nil {
			batchLog.Error("failed to index l1 initiated bridge events", "err", err)
			return err
		}
		if err := bridge.L2ProcessInitiatedBridgeEvents(l2BridgeLog, tx, b.opContracts.(op_stack.OpContracts).L2Contracts, fromL2Height, toL2Height); err != nil {
			batchLog.Error("failed to index l2 initiated bridge events", "err", err)
			return err
		}

		if err := bridge.L1ProcessFinalizedBridgeEvents(l1BridgeLog, tx, b.opContracts.(op_stack.OpContracts).L1Contracts, fromL1Height, toL1Height); err != nil {
			batchLog.Error("failed to index l1 finalized bridge events", "err", err)
			return err
		}
		if err := bridge.L2ProcessFinalizedBridgeEvents(l2BridgeLog, tx, b.opContracts.(op_stack.OpContracts).L2Contracts, fromL2Height, toL2Height); err != nil {
			batchLog.Error("failed to index l2 finalized bridge events", "err", err)
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	batchLog.Info("indexed bridge events", "latest_l1_block_number", toL1Height, "latest_l2_block_number", toL2Height)
	b.LatestL1Header = latestEpoch.L1BlockHeader.RLPHeader.Header()
	b.LatestL2Header = latestEpoch.L2BlockHeader.RLPHeader.Header()
	return nil
}
