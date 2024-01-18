package op_stack

import (
	"context"
	"errors"
	"fmt"
	"github.com/cornerstone-labs/acorus/common/bigint"
	"gorm.io/gorm"
	"math/big"
	"time"

	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/event/op_stack/bridge"
	"github.com/ethereum-optimism/optimism/indexer/config"
)

var blocksLimit = 10_000

type OpEventProcessor struct {
	db                          *database.DB
	resourceCtx                 context.Context
	resourceCancel              context.CancelFunc
	tasks                       tasks.Group
	chainConfig                 config.ChainConfig
	LatestL1L2InitL1Header      *common2.BlockHeader
	LatestL1L2L2Header          *common2.BlockHeader
	LatestL1L2FinalizedL2Header *common2.BlockHeader
	LatestStateRootL1Header     *common2.BlockHeader
	LatestL2L1InitL2Header      *common2.BlockHeader
	LatestProvenL1Header        *common2.BlockHeader
	LatestFinalizedL1Header     *common2.BlockHeader
}

func NewOpProcessor(db *database.DB, chainConfig config.ChainConfig, shutdown context.CancelCauseFunc) (*OpEventProcessor, error) {
	latestL1L2InitL1Header, err := db.L1ToL2.LatestBlockHeader("1")
	if err != nil {
		return nil, err
	}
	latestL1L2L2Header, err := db.L1ToL2.LatestBlockHeader("10")
	if err != nil {
		return nil, err
	}
	latestL1L2FinalizedL2Header, err := db.L1ToL2.L2LatestFinalizedBlockHeader("10")
	if err != nil {
		return nil, err
	}
	latestStateRootL1Header, err := db.StateRoots.StateRootL1BlockHeader()
	if err != nil {
		return nil, err
	}

	latestL2L1InitL2Header, err := db.L2ToL1.LatestBlockHeader("10")
	if err != nil {
		return nil, err
	}
	latestProvenL1Header, err := db.WithdrawProven.WithdrawProvenL1BlockHeader()
	if err != nil {
		return nil, err
	}
	latestFinalizedL1Header, err := db.WithdrawFinalized.WithdrawFinalizedL1BlockHeader()
	if err != nil {
		return nil, err
	}
	resCtx, resCancel := context.WithCancel(context.Background())
	return &OpEventProcessor{
		db:             db,
		resourceCtx:    resCtx,
		resourceCancel: resCancel,
		chainConfig:    chainConfig,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
		LatestL1L2InitL1Header:      latestL1L2InitL1Header,
		LatestL1L2L2Header:          latestL1L2L2Header,
		LatestL1L2FinalizedL2Header: latestL1L2FinalizedL2Header,
		LatestStateRootL1Header:     latestStateRootL1Header,
		LatestL2L1InitL2Header:      latestL2L1InitL2Header,
		LatestProvenL1Header:        latestProvenL1Header,
		LatestFinalizedL1Header:     latestFinalizedL1Header,
	}, nil
}

func (ep *OpEventProcessor) Start() error {
	tickerL1Worker := time.NewTicker(time.Second * 5)
	ep.tasks.Go(func() error {
		for range tickerL1Worker.C {
			err := ep.onL1Data()
			if err != nil {
				return err
			}
		}
		return nil
	})

	tickerL2Worker := time.NewTicker(time.Second * 5)
	ep.tasks.Go(func() error {
		for range tickerL2Worker.C {
			err := ep.onL2Data()
			if err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

func (ep *OpEventProcessor) Close() error {
	ep.resourceCancel()
	return ep.tasks.Wait()
}

func (ep *OpEventProcessor) onL1Data() error {
	var errs error
	if err := ep.processInitiatedL1Events(); err != nil {
		errs = errors.Join(errs, err)
	}

	if err := ep.processFinalizedL2Events(); err != nil {
		errs = errors.Join(errs, err)
	}

	if err := ep.processRollupStateRoot(); err != nil {
		errs = errors.Join(errs, err)
	}
	return errs
}

func (ep *OpEventProcessor) onL2Data() error {
	var errs error
	if err := ep.processInitiatedL2Events(); err != nil {
		errs = errors.Join(errs, err)
	}

	if err := ep.processProvenL1Events(); err != nil {
		errs = errors.Join(errs, err)
	}

	if err := ep.processFinalizedL1Events(); err != nil {
		errs = errors.Join(errs, err)
	}
	return errs
}

func (ep *OpEventProcessor) processInitiatedL1Events() error {
	lastL1BlockNumber := big.NewInt(int64(ep.chainConfig.L1StartingHeight))
	if ep.LatestL1L2InitL1Header != nil {
		lastL1BlockNumber = ep.LatestL1L2InitL1Header.Number
	}
	latestL1HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(common2.BlockHeader{}).Where("number > ?", lastL1BlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(blocksLimit)).Select("MAX(number)"))
	}
	latestL1Header, err := ep.db.Blocks.ChainBlockHeaderWithFilter("1", latestL1HeaderScope)
	if err != nil {
		return fmt.Errorf("failed to query new L1 state: %w", err)
	} else if latestL1Header == nil {
		return nil
	}
	fromL1Height, toL1Height := new(big.Int).Add(lastL1BlockNumber, bigint.One), latestL1Header.Number
	if err := ep.db.Transaction(func(tx *database.DB) error {
		l1BedrockStartingHeight := big.NewInt(int64(ep.chainConfig.L1BedrockStartingHeight))
		if l1BedrockStartingHeight.Cmp(fromL1Height) > 0 {
			legacyFromL1Height, legacyToL1Height := fromL1Height, toL1Height
			if l1BedrockStartingHeight.Cmp(toL1Height) <= 0 {
				legacyToL1Height = new(big.Int).Sub(l1BedrockStartingHeight, bigint.One)
			}

			if err := bridge.LegacyL1ProcessInitiatedBridgeEvents(tx, legacyFromL1Height, legacyToL1Height); err != nil {
				return err
			} else if legacyToL1Height.Cmp(toL1Height) == 0 {
				return nil
			}
			fromL1Height = l1BedrockStartingHeight
		}
		return bridge.L1ProcessInitiatedBridgeEvents(tx, fromL1Height, toL1Height)
	}); err != nil {
		return err
	}
	ep.LatestL1L2InitL1Header = latestL1Header
	return nil
}

func (ep *OpEventProcessor) processInitiatedL2Events() error {
	lastL2BlockNumber := big.NewInt(int64(ep.chainConfig.L2StartingHeight))
	if ep.LatestL2L1InitL2Header != nil {
		lastL2BlockNumber = ep.LatestL2L1InitL2Header.Number
	}
	latestL2HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(common2.BlockHeader{}).Where("number > ?", lastL2BlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(blocksLimit)).Select("MAX(number)"))
	}
	latestL2Header, err := ep.db.Blocks.ChainBlockHeaderWithScope(latestL2HeaderScope)
	if err != nil {
		return fmt.Errorf("failed to query new L2 state: %w", err)
	} else if latestL2Header == nil {
		return nil
	}
	fromL2Height, toL2Height := new(big.Int).Add(lastL2BlockNumber, bigint.One), latestL2Header.Number
	if err := ep.db.Transaction(func(tx *database.DB) error {
		l2BedrockStartingHeight := big.NewInt(int64(ep.chainConfig.L2BedrockStartingHeight))
		if l2BedrockStartingHeight.Cmp(fromL2Height) > 0 { // OP Mainnet & OP Goerli Only
			legacyFromL2Height, legacyToL2Height := fromL2Height, toL2Height
			if l2BedrockStartingHeight.Cmp(toL2Height) <= 0 {
				legacyToL2Height = new(big.Int).Sub(l2BedrockStartingHeight, bigint.One)
			}
			if err := bridge.LegacyL2ProcessInitiatedBridgeEvents(tx, legacyFromL2Height, legacyToL2Height); err != nil {
				return err
			} else if legacyToL2Height.Cmp(toL2Height) == 0 {
				return nil
			}
			fromL2Height = l2BedrockStartingHeight
		}
		return bridge.L2ProcessInitiatedBridgeEvents(tx, fromL2Height, toL2Height)
	}); err != nil {
		return err
	}
	ep.LatestL2L1InitL2Header = latestL2Header
	return nil
}

func (ep *OpEventProcessor) processProvenL1Events() error {
	lastProvenL1BlockNumber := big.NewInt(int64(ep.chainConfig.L1StartingHeight))
	if ep.LatestProvenL1Header != nil {
		lastProvenL1BlockNumber = ep.LatestProvenL1Header.Number
	}
	latestProvenL1HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(common2.BlockHeader{}).Where("number > ?", lastProvenL1BlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(blocksLimit)).Select("MAX(number)"))
	}
	if latestProvenL1HeaderScope == nil {
		return nil
	}
	latestL1Header, err := ep.db.Blocks.ChainBlockHeaderWithScope(latestProvenL1HeaderScope)
	if err != nil {
		return fmt.Errorf("failed to query for latest unfinalized L1 state: %w", err)
	} else if latestL1Header == nil {
		return nil
	}
	fromL1Height, toL1Height := new(big.Int).Add(lastProvenL1BlockNumber, bigint.One), latestL1Header.Number
	if err := ep.db.Transaction(func(tx *database.DB) error {
		if err := bridge.L1ProcessProvenBridgeEvents(tx, fromL1Height, toL1Height); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	ep.LatestProvenL1Header = latestL1Header
	return nil
}

func (ep *OpEventProcessor) processFinalizedL1Events() error {
	lastFinalizedL1BlockNumber := big.NewInt(int64(ep.chainConfig.L1StartingHeight))
	if ep.LatestFinalizedL1Header != nil {
		lastFinalizedL1BlockNumber = ep.LatestFinalizedL1Header.Number
	}
	latestFinalizedL1HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(common2.BlockHeader{}).Where("number > ?", lastFinalizedL1BlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(blocksLimit)).Select("MAX(number)"))
	}
	if latestFinalizedL1HeaderScope == nil {
		return nil
	}
	latestL1Header, err := ep.db.Blocks.L1BlockHeaderWithScope(latestFinalizedL1HeaderScope)
	if err != nil {
		return fmt.Errorf("failed to query for latest unfinalized L1 state: %w", err)
	} else if latestL1Header == nil {
		return nil
	}
	fromL1Height, toL1Height := new(big.Int).Add(lastFinalizedL1BlockNumber, bigint.One), latestL1Header.Number
	if err := ep.db.Transaction(func(tx *database.DB) error {
		l1BedrockStartingHeight := big.NewInt(int64(ep.chainConfig.L1BedrockStartingHeight))
		if l1BedrockStartingHeight.Cmp(fromL1Height) > 0 {
			legacyFromL1Height, legacyToL1Height := fromL1Height, toL1Height
			if l1BedrockStartingHeight.Cmp(toL1Height) <= 0 {
				legacyToL1Height = new(big.Int).Sub(l1BedrockStartingHeight, bigint.One)
			}
			if err := bridge.LegacyL1ProcessFinalizedBridgeEvents(tx, legacyFromL1Height, legacyToL1Height); err != nil {
				return err
			} else if legacyToL1Height.Cmp(toL1Height) == 0 {
				return nil
			}
			fromL1Height = l1BedrockStartingHeight
		}
		return bridge.L1ProcessFinalizedBridgeEvents(tx, fromL1Height, toL1Height)
	}); err != nil {
		return err
	}
	ep.LatestFinalizedL1Header = latestL1Header
	return nil
}

func (ep *OpEventProcessor) processFinalizedL2Events() error {
	lastFinalizedL2BlockNumber := big.NewInt(int64(ep.chainConfig.L2StartingHeight))
	latestL1L2BlockNumber := big.NewInt(int64(ep.chainConfig.L2StartingHeight))
	if ep.LatestL1L2FinalizedL2Header != nil {
		lastFinalizedL2BlockNumber = ep.LatestL1L2FinalizedL2Header.Number
	}
	l2BlockHeader, err := ep.db.L1ToL2.L1L2LatestL2BlockHeader()
	if err != nil {
		return err
	}
	if l2BlockHeader != nil {
		ep.LatestL1L2L2Header = l2BlockHeader
		latestL1L2BlockNumber = l2BlockHeader.Number
	}
	latestL2HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(common2.L2BlockHeader{}).Where("number > ? AND number <= ?", lastFinalizedL2BlockNumber, latestL1L2BlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(blocksLimit)).Select("MAX(number)"))
	}
	latestL2Header, err := ep.db.Blocks.ChainBlockHeaderWithScope(latestL2HeaderScope)
	if err != nil {
		return fmt.Errorf("failed to query for latest unfinalized L2 state: %w", err)
	} else if latestL2Header == nil {
		ep.LatestL1L2FinalizedL2Header = latestL2Header
		return nil
	}
	fromL2Height, toL2Height := new(big.Int).Add(lastFinalizedL2BlockNumber, bigint.One), latestL2Header.Number
	if err := ep.db.Transaction(func(tx *database.DB) error {
		l2BedrockStartingHeight := big.NewInt(int64(ep.chainConfig.L2BedrockStartingHeight))
		if l2BedrockStartingHeight.Cmp(fromL2Height) > 0 {
			legacyFromL2Height, legacyToL2Height := fromL2Height, toL2Height
			if l2BedrockStartingHeight.Cmp(toL2Height) <= 0 {
				legacyToL2Height = new(big.Int).Sub(l2BedrockStartingHeight, bigint.One)
			}
			if err := bridge.LegacyL2ProcessFinalizedBridgeEvents(tx, legacyFromL2Height, legacyToL2Height); err != nil {
				return err
			} else if legacyToL2Height.Cmp(toL2Height) == 0 {
				return nil
			}
			fromL2Height = l2BedrockStartingHeight
		}
		return bridge.L2ProcessFinalizedBridgeEvents(tx, fromL2Height, toL2Height)
	}); err != nil {
		return err
	}
	ep.LatestL1L2FinalizedL2Header = latestL2Header
	return nil
}

func (ep *OpEventProcessor) processRollupStateRoot() error {
	lastStateRootL1BlockNumber := big.NewInt(int64(ep.chainConfig.L1StartingHeight))
	if ep.LatestStateRootL1Header != nil {
		lastStateRootL1BlockNumber = ep.LatestStateRootL1Header.Number
	}
	latestRollupL1HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(common2.BlockHeader{}).Where("number > ?", lastStateRootL1BlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(blocksLimit)).Select("MAX(number)"))
	}
	latestL1StateRootHeader, err := ep.db.Blocks.ChainBlockHeaderWithScope(latestRollupL1HeaderScope)
	if err != nil {
		return fmt.Errorf("failed to query new L1 state: %w", err)
	} else if latestL1StateRootHeader == nil {
		return nil
	}
	fromL1Height, toL1Height := new(big.Int).Add(lastStateRootL1BlockNumber, bigint.One), latestL1StateRootHeader.Number
	if err := ep.db.Transaction(func(tx *database.DB) error {
		if err := stateroot.L2OutputEvent(tx, fromL1Height, toL1Height); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	ep.LatestStateRootL1Header = latestL1StateRootHeader
	return nil
}
