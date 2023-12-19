package op_stack

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math/big"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	common3 "github.com/cornerstone-labs/acorus/event/processor/common"
	"github.com/cornerstone-labs/acorus/event/processor/op-stack/bridge"
	"github.com/cornerstone-labs/acorus/event/processor/op-stack/stateroot"
	"github.com/cornerstone-labs/acorus/synchronizer"
)

var blocksLimit = 10_000

type OpProcessor struct {
	common3.IProcessor
}

func NewBridgeProcessor(log log.Logger, db *database.DB, l1Sync *synchronizer.L1Sync, l2Sync *synchronizer.L2Sync,
	chainConfig config.ChainConfig, shutdown context.CancelCauseFunc) (*common3.IProcessor, error) {
	log = log.New("processor", "bridge")
	latestL1Header, latestL2Header,
		latestFinalizedL1Header, latestFinalizedL2Header,
		err := common3.GetLastBlockHeardByChain(log, db, chainConfig)
	if err != nil {
		return nil, err
	}
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

func (ep *OpProcessor) Start() error {
	ep.Log.Info("starting bridge processor...")
	ep.Tasks.Go(func() error {
		l1SyncUpdates := ep.L1Sync.Notify()
		for range l1SyncUpdates {
			err := ep.onL1Data()
			if err != nil {
				return err
			}
		}
		ep.Log.Info("no more l1 etl updates. shutting down l1 task")
		return nil
	})
	// start L2 worker
	ep.Tasks.Go(func() error {
		l2SyncUpdates := ep.L2Sync.Notify()
		for range l2SyncUpdates {
			err := ep.onL2Data()
			if err != nil {
				return err
			}
		}
		ep.Log.Info("no more l2 etl updates. shutting down l2 task")
		return nil
	})
	return nil
}

func (ep *OpProcessor) Close() error {
	ep.ResourceCancel()
	return ep.Tasks.Wait()
}

func (ep *OpProcessor) processRollup() error {
	rollupLog := ep.Log.New("rollup", "l1", "kind", "mantleDa and state root")
	lastRollupL1BlockNumber := big.NewInt(int64(ep.ChainConfig.L1StartingHeight) - 1)
	if ep.LastRollupL1Header != nil {
		lastRollupL1BlockNumber = ep.LastRollupL1Header.Number
	}

	latestRollupL1HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(common2.L1BlockHeader{}).Where("number > ?", lastRollupL1BlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(blocksLimit)).Select("MAX(number)"))
	}
	latestL1RollupHeader, err := ep.Db.Blocks.L1BlockHeaderWithScope(latestRollupL1HeaderScope)
	if err != nil {
		return fmt.Errorf("failed to query new L1 state: %w", err)
	} else if latestL1RollupHeader == nil {
		return fmt.Errorf("no new L1 state found")
	}

	fromL1Height, toL1Height := new(big.Int).Add(lastRollupL1BlockNumber, bigint.One), latestL1RollupHeader.Number
	if err := ep.Db.Transaction(func(tx *database.DB) error {
		rollupLog = rollupLog.New("from_block_number", fromL1Height, "to_block_number", toL1Height)
		rollupLog.Info("scanning for mantle da and state root events")

		if err := stateroot.L2OutputEvent(rollupLog, tx, fromL1Height, toL1Height); err != nil {
			ep.Log.Error("failed to index l1 l2output proposed events", "err", err)
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	ep.LastRollupL1Header = latestL1RollupHeader
	return nil
}

func (ep *OpProcessor) onL1Data() error {
	latestL1Header := ep.L1Sync.LatestHeader
	ep.Log.Info("notified of new L1 state", "l1_etl_block_number", latestL1Header.Number)

	var errs error
	if err := ep.processInitiatedL1Ops(); err != nil {
		ep.Log.Error("failed to process initiated L1 events", "err", err)
		errs = errors.Join(errs, err)
	}

	if ep.LastFinalizedL2Header == nil || ep.LastFinalizedL2Header.Timestamp < ep.LastL1Header.Timestamp {
		if err := ep.processFinalizedL2Ops(); err != nil {
			ep.Log.Error("failed to process finalized L2 events", "err", err)
			errs = errors.Join(errs, err)
		}
	}

	if err := ep.processRollup(); err != nil {
		ep.Log.Error("failed to process rollup events", "err", err)
		errs = errors.Join(errs, err)
	}

	return errs
}

func (ep *OpProcessor) onL2Data() error {
	if ep.L2Sync.LatestHeader.Number.Cmp(bigint.Zero) == 0 {
		return nil
	}
	ep.Log.Info("notified of new L2 state", "l2_etl_block_number", ep.L2Sync.LatestHeader.Number)

	var errs error
	if err := ep.processInitiatedL2Ops(); err != nil {
		ep.Log.Error("failed to process initiated L2 events", "err", err)
		errs = errors.Join(errs, err)
	}

	if ep.LastFinalizedL1Header == nil || ep.LastFinalizedL1Header.Timestamp < ep.LastL2Header.Timestamp {
		if err := ep.processFinalizedL1Ops(); err != nil {
			ep.Log.Error("failed to process finalized L1 events", "err", err)
			errs = errors.Join(errs, err)
		}
	}
	return errs
}

func (ep *OpProcessor) processInitiatedL1Ops() error {
	l1BridgeLog := ep.Log.New("bridge", "l1", "kind", "initiated")
	lastL1BlockNumber := big.NewInt(int64(ep.ChainConfig.L1StartingHeight) - 1)
	if ep.LastL1Header != nil {
		lastL1BlockNumber = ep.LastL1Header.Number
	}

	latestL1HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(common2.L1BlockHeader{}).Where("number > ?", lastL1BlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(blocksLimit)).Select("MAX(number)"))
	}
	latestL1Header, err := ep.Db.Blocks.L1BlockHeaderWithScope(latestL1HeaderScope)
	if err != nil {
		return fmt.Errorf("failed to query new L1 state: %w", err)
	} else if latestL1Header == nil {
		return fmt.Errorf("no new L1 state found")
	}

	fromL1Height, toL1Height := new(big.Int).Add(lastL1BlockNumber, bigint.One), latestL1Header.Number
	if err := ep.Db.Transaction(func(tx *database.DB) error {
		l1BedrockStartingHeight := big.NewInt(int64(ep.ChainConfig.L1BedrockStartingHeight))
		if l1BedrockStartingHeight.Cmp(fromL1Height) > 0 {
			legacyFromL1Height, legacyToL1Height := fromL1Height, toL1Height
			if l1BedrockStartingHeight.Cmp(toL1Height) <= 0 {
				legacyToL1Height = new(big.Int).Sub(l1BedrockStartingHeight, bigint.One)
			}

			legacyBridgeLog := l1BridgeLog.New("mode", "legacy", "from_block_number", legacyFromL1Height, "to_block_number", legacyToL1Height)
			legacyBridgeLog.Info("scanning for initiated bridge events")
			if err := bridge.LegacyL1ProcessInitiatedBridgeEvents(legacyBridgeLog, tx, legacyFromL1Height, legacyToL1Height); err != nil {
				return err
			} else if legacyToL1Height.Cmp(toL1Height) == 0 {
				return nil
			}
			legacyBridgeLog.Info("detected switch to bedrock", "bedrock_block_number", l1BedrockStartingHeight)
			fromL1Height = l1BedrockStartingHeight
		}

		l1BridgeLog = l1BridgeLog.New("from_block_number", fromL1Height, "to_block_number", toL1Height)
		l1BridgeLog.Info("scanning for initiated bridge events")
		return bridge.L1ProcessInitiatedBridgeEvents(l1BridgeLog, tx, fromL1Height, toL1Height)
	}); err != nil {
		return err
	}
	ep.LastL1Header = latestL1Header
	return nil
}

func (ep *OpProcessor) processInitiatedL2Ops() error {
	l2BridgeLog := ep.Log.New("bridge", "l2", "kind", "initiated")
	lastL2BlockNumber := bigint.Zero
	if ep.LastL2Header != nil {
		lastL2BlockNumber = ep.LastL2Header.Number
	}

	latestL2HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(common2.L2BlockHeader{}).Where("number > ?", lastL2BlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(blocksLimit)).Select("MAX(number)"))
	}
	latestL2Header, err := ep.Db.Blocks.L2BlockHeaderWithScope(latestL2HeaderScope)
	if err != nil {
		return fmt.Errorf("failed to query new L2 state: %w", err)
	} else if latestL2Header == nil {
		return fmt.Errorf("no new L2 state found")
	}

	fromL2Height, toL2Height := new(big.Int).Add(lastL2BlockNumber, bigint.One), latestL2Header.Number
	if err := ep.Db.Transaction(func(tx *database.DB) error {
		l2BedrockStartingHeight := big.NewInt(int64(ep.ChainConfig.L2BedrockStartingHeight))
		if l2BedrockStartingHeight.Cmp(fromL2Height) > 0 { // OP Mainnet & OP Goerli Only
			legacyFromL2Height, legacyToL2Height := fromL2Height, toL2Height
			if l2BedrockStartingHeight.Cmp(toL2Height) <= 0 {
				legacyToL2Height = new(big.Int).Sub(l2BedrockStartingHeight, bigint.One)
			}

			legacyBridgeLog := l2BridgeLog.New("mode", "legacy", "from_block_number", legacyFromL2Height, "to_block_number", legacyToL2Height)
			legacyBridgeLog.Info("scanning for initiated bridge events")
			if err := bridge.LegacyL2ProcessInitiatedBridgeEvents(legacyBridgeLog, tx, legacyFromL2Height, legacyToL2Height); err != nil {
				return err
			} else if legacyToL2Height.Cmp(toL2Height) == 0 {
				return nil
			}
			legacyBridgeLog.Info("detected switch to bedrock")
			fromL2Height = l2BedrockStartingHeight
		}

		l2BridgeLog = l2BridgeLog.New("from_block_number", fromL2Height, "to_block_number", toL2Height)
		l2BridgeLog.Info("scanning for initiated bridge events")
		return bridge.L2ProcessInitiatedBridgeEvents(l2BridgeLog, tx, fromL2Height, toL2Height)
	}); err != nil {
		return err
	}

	ep.LastL2Header = latestL2Header
	return nil
}

func (ep *OpProcessor) processFinalizedL1Ops() error {
	l1BridgeLog := ep.Log.New("bridge", "l1", "kind", "finalization")
	lastFinalizedL1BlockNumber := big.NewInt(int64(ep.ChainConfig.L1StartingHeight) - 1)
	if ep.LastFinalizedL1Header != nil {
		lastFinalizedL1BlockNumber = ep.LastFinalizedL1Header.Number
	}

	latestL1HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(common2.L1BlockHeader{}).Where("number > ? AND timestamp <= ?", lastFinalizedL1BlockNumber, ep.LastL2Header.Timestamp)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(blocksLimit)).Select("MAX(number)"))
	}
	latestL1Header, err := ep.Db.Blocks.L1BlockHeaderWithScope(latestL1HeaderScope)
	if err != nil {
		return fmt.Errorf("failed to query for latest unfinalized L1 state: %w", err)
	} else if latestL1Header == nil {
		l1BridgeLog.Debug("no new l1 state to finalize", "last_finalized_block_number", lastFinalizedL1BlockNumber)
		return nil
	}

	fromL1Height, toL1Height := new(big.Int).Add(lastFinalizedL1BlockNumber, bigint.One), latestL1Header.Number
	if err := ep.Db.Transaction(func(tx *database.DB) error {
		l1BedrockStartingHeight := big.NewInt(int64(ep.ChainConfig.L1BedrockStartingHeight))
		if l1BedrockStartingHeight.Cmp(fromL1Height) > 0 {
			legacyFromL1Height, legacyToL1Height := fromL1Height, toL1Height
			if l1BedrockStartingHeight.Cmp(toL1Height) <= 0 {
				legacyToL1Height = new(big.Int).Sub(l1BedrockStartingHeight, bigint.One)
			}

			legacyBridgeLog := l1BridgeLog.New("mode", "legacy", "from_block_number", legacyFromL1Height, "to_block_number", legacyToL1Height)
			legacyBridgeLog.Info("scanning for finalized bridge events")
			if err := bridge.LegacyL1ProcessFinalizedBridgeEvents(legacyBridgeLog, tx, legacyFromL1Height, legacyToL1Height); err != nil {
				return err
			} else if legacyToL1Height.Cmp(toL1Height) == 0 {
				return nil
			}
			legacyBridgeLog.Info("detected switch to bedrock")
			fromL1Height = l1BedrockStartingHeight
		}

		l1BridgeLog = l1BridgeLog.New("from_block_number", fromL1Height, "to_block_number", toL1Height)
		l1BridgeLog.Info("scanning for finalized bridge events")
		return bridge.L1ProcessFinalizedBridgeEvents(l1BridgeLog, tx, fromL1Height, toL1Height)
	}); err != nil {
		return err
	}
	ep.LastFinalizedL1Header = latestL1Header
	return nil
}

func (ep *OpProcessor) processFinalizedL2Ops() error {
	l2BridgeLog := ep.Log.New("bridge", "l2", "kind", "finalization")
	lastFinalizedL2BlockNumber := bigint.Zero
	if ep.LastFinalizedL2Header != nil {
		lastFinalizedL2BlockNumber = ep.LastFinalizedL2Header.Number
	}

	latestL2HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(common2.L2BlockHeader{}).Where("number > ? AND timestamp <= ?", lastFinalizedL2BlockNumber, ep.LastL1Header.Timestamp)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(blocksLimit)).Select("MAX(number)"))
	}
	latestL2Header, err := ep.Db.Blocks.L2BlockHeaderWithScope(latestL2HeaderScope)
	if err != nil {
		return fmt.Errorf("failed to query for latest unfinalized L2 state: %w", err)
	} else if latestL2Header == nil {
		l2BridgeLog.Debug("no new l2 state to finalize", "last_finalized_block_number", lastFinalizedL2BlockNumber)
		return nil
	}

	fromL2Height, toL2Height := new(big.Int).Add(lastFinalizedL2BlockNumber, bigint.One), latestL2Header.Number
	if err := ep.Db.Transaction(func(tx *database.DB) error {
		l2BedrockStartingHeight := big.NewInt(int64(ep.ChainConfig.L2BedrockStartingHeight))
		if l2BedrockStartingHeight.Cmp(fromL2Height) > 0 {
			legacyFromL2Height, legacyToL2Height := fromL2Height, toL2Height
			if l2BedrockStartingHeight.Cmp(toL2Height) <= 0 {
				legacyToL2Height = new(big.Int).Sub(l2BedrockStartingHeight, bigint.One)
			}

			legacyBridgeLog := l2BridgeLog.New("mode", "legacy", "from_block_number", legacyFromL2Height, "to_block_number", legacyToL2Height)
			legacyBridgeLog.Info("scanning for finalized bridge events")
			if err := bridge.LegacyL2ProcessFinalizedBridgeEvents(legacyBridgeLog, tx, legacyFromL2Height, legacyToL2Height); err != nil {
				return err
			} else if legacyToL2Height.Cmp(toL2Height) == 0 {
				return nil
			}
			legacyBridgeLog.Info("detected switch to bedrock", "bedrock_block_number", l2BedrockStartingHeight)
			fromL2Height = l2BedrockStartingHeight
		}

		l2BridgeLog = l2BridgeLog.New("from_block_number", fromL2Height, "to_block_number", toL2Height)
		l2BridgeLog.Info("scanning for finalized bridge events")
		return bridge.L2ProcessFinalizedBridgeEvents(l2BridgeLog, tx, fromL2Height, toL2Height)
	}); err != nil {
		return err
	}
	ep.LastFinalizedL2Header = latestL2Header
	return nil
}
