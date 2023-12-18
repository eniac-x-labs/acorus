package polygon

import (
	"context"
	"errors"
	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/config/polygon"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/event/processors/polygon/abi"
	"github.com/cornerstone-labs/acorus/event/processors/polygon/bridge"
	"github.com/cornerstone-labs/acorus/metrics"
	"github.com/cornerstone-labs/acorus/synchronizer"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

type PlBridgeProcessor struct {
	log log.Logger
	db  *database.DB

	metrics     metrics.Metricer
	l1Etl       *synchronizer.L1Sync
	chainConfig config.ChainConfig
	plContracts interface{}

	LatestL1Header *types.Header
	LatestL2Header *types.Header

	L1PolygonBridge            *abi.Polygonzkevmbridge
	L2PolygonBridge            *abi.Polygonzkevmbridge
	PolygonZkEVMGlobalExitRoot *abi.Polygonzkevmglobalexitroot
	PolygonRollupManager       *abi.Polygonrollupmanager
}

func NewPolygonBridgeProcessor(log log.Logger, db *database.DB, metrics metrics.Metricer, L1Syncer *synchronizer.L1Sync, chainConfig config.ChainConfig, plContracts interface{}, L1URL string, L2URL string) (*PlBridgeProcessor, error) {
	ethClient, _ := ethclient.Dial(L1URL)
	zkEVMClient, _ := ethclient.Dial(L2URL)
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
	l1PolygonBridge, err := abi.NewPolygonzkevmbridge(plContracts.(polygon.PlContracts).L1Contracts.PolygonZKEVMBridgeAddr, ethClient)
	if err != nil {
		return nil, err
	}
	l2PolygonBridge, err := abi.NewPolygonzkevmbridge(plContracts.(polygon.PlContracts).L1Contracts.PolygonZKEVMBridgeAddr, zkEVMClient)
	if err != nil {
		return nil, err
	}
	//polygonZkEVMGlobalExitRoot, err := abi.NewPolygonzkevmglobalexitroot(plContracts.(polygon.PlContracts).L1Contracts.PolygonZkEVMGlobalExitRootAddr, ethClient)
	//if err != nil {
	//	return nil, err
	//}
	//polygonRollupManager, err := abi.NewPolygonrollupmanager(plContracts.(polygon.PlContracts).L1Contracts.PolygonRollupManagerAddr, ethClient)
	//if err != nil {
	//	return nil, err
	//}
	return &PlBridgeProcessor{log, db, metrics, L1Syncer, chainConfig, plContracts, l1Header, l2Header, l1PolygonBridge, l2PolygonBridge, nil, nil}, nil
}

func (b *PlBridgeProcessor) Start(ctx context.Context) error {
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

		done := b.metrics.RecordInterval()
		done(b.run())
	}
}

func (b *PlBridgeProcessor) run() error {
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

	batchLog := b.log.New("epoch_start_number", fromL1Height, "epoch_end_number", toL1Height)
	batchLog.Info("unobserved epochs", "latest_l1_block_number", fromL1Height, "latest_l2_block_number", fromL2Height)
	if err := b.db.Transaction(func(tx *database.DB) error {
		l1BridgeLog := b.log.New("bridge", "l1")
		l2BridgeLog := b.log.New("bridge", "l2")

		l1BridgeLog.Info("scanning for bridge events")
		l2BridgeLog.Info("scanning for bridge events")
		l1Contracts := b.plContracts.(polygon.PlContracts).L1Contracts
		// l1 handle start
		// eth deposit handle
		if err := bridge.L1Deposit(b.L1PolygonBridge, l1Contracts.PolygonZKEVMBridgeAddr, tx, fromL1Height, toL1Height); err != nil {
			batchLog.Error("failed to index L1Deposit bridge events", "err", err)
			return err
		}
		if err := bridge.L1WithdrawClaimed(b.L1PolygonBridge, l1Contracts.PolygonZKEVMBridgeAddr, tx, fromL1Height, toL1Height); err != nil {
			batchLog.Error("failed to index L1WithdrawClaimed bridge events", "err", err)
			return err
		}
		// l1 handle end

		// l2 handle start
		l2Contracts := b.plContracts.(polygon.PlContracts).L2Contracts

		if err := bridge.L2Withdraw(b.L2PolygonBridge, l2Contracts.PolygonZKEVMBridgeAddr, tx, fromL2Height, toL2Height); err != nil {
			batchLog.Error("failed to index L2Withdraw bridge events", "err", err)
			return err
		}

		// l2 handle end

		return nil
	}); err != nil {
		return err
	}
	batchLog.Info("indexed bridge events", "latest_l1_block_number", toL1Height, "latest_l2_block_number", toL2Height)
	b.LatestL1Header = latestEpoch.L1BlockHeader.RLPHeader.Header()
	b.LatestL2Header = latestEpoch.L2BlockHeader.RLPHeader.Header()
	return nil

}
