package mantle

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
	"math/big"
	"time"

	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/event/mantle/bridge"
	common3 "github.com/cornerstone-labs/acorus/event/mantle/common"
	"github.com/cornerstone-labs/acorus/event/mantle/stateroot"
)

type MantleEventProcessor struct {
	db                          *database.DB
	resourceCtx                 context.Context
	resourceCancel              context.CancelFunc
	tasks                       tasks.Group
	L1StartHeight               *big.Int
	L2StartHeight               *big.Int
	L1BedrockStartHeight        *big.Int
	L2BedrockStartHeight        *big.Int
	LatestL1L2InitL1Header      *common2.BlockHeader
	LatestL1L2L2Header          *common2.BlockHeader
	LatestL1L2FinalizedL2Header *common2.BlockHeader
	LatestStateRootL1Header     *common2.BlockHeader
	LatestL2L1InitL2Header      *common2.BlockHeader
	LatestProvenL1Header        *common2.BlockHeader
	LatestFinalizedL1Header     *common2.BlockHeader
	loopInterval                time.Duration
	epoch                       uint64
	l1ChainId                   string
	l2ChainId                   string
}

func NewBridgeProcessor(db *database.DB, l1StartHeight *big.Int, l2StartHeight *big.Int,
	shutdown context.CancelCauseFunc, loopInterval time.Duration,
	epoch uint64, l1ChainId, l2ChainId string, isMainnet bool) (*MantleEventProcessor, error) {

	common3.InitAddress(isMainnet)

	latestL1L2InitL1Header, err := db.L1ToL2.L1LatestBlockHeader(l2ChainId, l1ChainId)
	if err != nil {
		return nil, err
	}
	latestL1L2L2Header, err := db.L1ToL2.L2LatestBlockHeader(l2ChainId)
	if err != nil {
		return nil, err
	}
	latestL1L2FinalizedL2Header, err := db.L1ToL2.L2LatestFinalizedBlockHeader(l2ChainId)
	if err != nil {
		return nil, err
	}
	latestStateRootL1Header, err := db.StateRoots.StateRootL1BlockHeader(l2ChainId, l1ChainId)
	if err != nil {
		return nil, err
	}

	latestL2L1InitL2Header, err := db.L2ToL1.L2LatestBlockHeader(l2ChainId)
	if err != nil {
		return nil, err
	}
	latestProvenL1Header, err := db.WithdrawProven.WithdrawProvenL1BlockHeader(l1ChainId, l2ChainId)
	if err != nil {
		return nil, err
	}
	latestFinalizedL1Header, err := db.WithdrawFinalized.WithdrawFinalizedL1BlockHeader(l1ChainId, l2ChainId)
	if err != nil {
		return nil, err
	}
	resCtx, resCancel := context.WithCancel(context.Background())

	return &MantleEventProcessor{
		db:                   db,
		resourceCtx:          resCtx,
		resourceCancel:       resCancel,
		L1StartHeight:        l1StartHeight,
		L2StartHeight:        l2StartHeight,
		L1BedrockStartHeight: common3.L1BedrockStartHeight,
		L2BedrockStartHeight: common3.L2BedrockStartHeight,
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
		loopInterval:                loopInterval,
		epoch:                       epoch,
		l1ChainId:                   l1ChainId,
		l2ChainId:                   l2ChainId,
	}, nil
}

func (ep *MantleEventProcessor) StartUnpack() error {
	tickerEventOn1 := time.NewTicker(ep.loopInterval)
	tickerEventOn2 := time.NewTicker(ep.loopInterval)
	ep.tasks.Go(func() error {
		for range tickerEventOn1.C {
			err := ep.onL1Data()
			if err != nil {
				log.Error("failed to process L1 data", "err", err)
				continue
			}
		}
		return nil
	})

	ep.tasks.Go(func() error {
		for range tickerEventOn2.C {
			err := ep.onL2Data()
			if err != nil {
				log.Error("failed to process L2 data", "err", err)
				continue
			}
		}
		return nil
	})
	return nil
}

func (ep *MantleEventProcessor) ClosetUnpack() error {
	ep.resourceCancel()
	return ep.tasks.Wait()
}

func (ep *MantleEventProcessor) onL1Data() error {
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

func (ep *MantleEventProcessor) onL2Data() error {
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

func (ep *MantleEventProcessor) processInitiatedL1Events() error {
	lastL1BlockNumber := ep.L1StartHeight
	if ep.LatestL1L2InitL1Header != nil {
		lastL1BlockNumber = ep.LatestL1L2InitL1Header.Number
	} else {
		lastL1BlockNumber = new(big.Int).Sub(lastL1BlockNumber, bigint.One)
	}
	latestL1HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Table("block_headers_"+ep.l1ChainId).Where("number > ?", lastL1BlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(int(ep.epoch))).Select("MAX(number)"))
	}
	latestL1Header, err := ep.db.Blocks.ChainBlockHeaderWithScope(latestL1HeaderScope, ep.l1ChainId)
	if err != nil {
		return fmt.Errorf("failed to query new L1 state: %w", err)
	} else if latestL1Header == nil {
		return nil
	}
	fromL1Height, toL1Height := new(big.Int).Add(lastL1BlockNumber, bigint.One), latestL1Header.Number
	if err := ep.db.Transaction(func(tx *database.DB) error {
		l1BedrockStartingHeight := ep.L1BedrockStartHeight
		if l1BedrockStartingHeight.Cmp(fromL1Height) > 0 {
			legacyFromL1Height, legacyToL1Height := fromL1Height, toL1Height
			if l1BedrockStartingHeight.Cmp(toL1Height) <= 0 {
				legacyToL1Height = new(big.Int).Sub(l1BedrockStartingHeight, bigint.One)
			}
			if err := bridge.LegacyL1ProcessInitiatedBridgeEvents(tx, legacyFromL1Height, legacyToL1Height,
				ep.l1ChainId, ep.l2ChainId); err != nil {
				return err
			} else if legacyToL1Height.Cmp(toL1Height) == 0 {
				return nil
			}
			fromL1Height = l1BedrockStartingHeight
		}
		return bridge.L1ProcessInitiatedBridgeEvents(tx, fromL1Height, toL1Height, ep.l1ChainId, ep.l2ChainId)
	}); err != nil {
		return err
	}
	ep.LatestL1L2InitL1Header = latestL1Header
	return nil
}

func (ep *MantleEventProcessor) processInitiatedL2Events() error {
	lastL2BlockNumber := ep.L2StartHeight
	if ep.LatestL2L1InitL2Header != nil {
		lastL2BlockNumber = ep.LatestL2L1InitL2Header.Number
	} else {
		lastL2BlockNumber = new(big.Int).Sub(lastL2BlockNumber, bigint.One)
	}
	latestL2HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Table("block_headers_"+ep.l2ChainId).Where("number > ?", lastL2BlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(int(ep.epoch))).Select("MAX(number)"))
	}
	latestL2Header, err := ep.db.Blocks.ChainBlockHeaderWithScope(latestL2HeaderScope, ep.l2ChainId)
	if err != nil {
		return fmt.Errorf("failed to query new L2 state: %w", err)
	} else if latestL2Header == nil {
		return nil
	}
	fromL2Height, toL2Height := new(big.Int).Add(lastL2BlockNumber, bigint.One), latestL2Header.Number
	if err := ep.db.Transaction(func(tx *database.DB) error {
		l2BedrockStartingHeight := ep.L2BedrockStartHeight
		if l2BedrockStartingHeight.Cmp(fromL2Height) > 0 { // OP Mainnet & OP Goerli Only
			legacyFromL2Height, legacyToL2Height := fromL2Height, toL2Height
			if l2BedrockStartingHeight.Cmp(toL2Height) <= 0 {
				legacyToL2Height = new(big.Int).Sub(l2BedrockStartingHeight, bigint.One)
			}
			if err := bridge.LegacyL2ProcessInitiatedBridgeEvents(tx, legacyFromL2Height,
				legacyToL2Height, ep.l1ChainId, ep.l2ChainId); err != nil {
				return err
			} else if legacyToL2Height.Cmp(toL2Height) == 0 {
				return nil
			}
			fromL2Height = l2BedrockStartingHeight
		}
		return bridge.L2ProcessInitiatedBridgeEvents(tx, fromL2Height, toL2Height, ep.l1ChainId, ep.l2ChainId)
	}); err != nil {
		return err
	}
	ep.LatestL2L1InitL2Header = latestL2Header
	return nil
}

func (ep *MantleEventProcessor) processProvenL1Events() error {
	lastProvenL1BlockNumber := ep.L1StartHeight
	if ep.LatestProvenL1Header != nil {
		lastProvenL1BlockNumber = ep.LatestProvenL1Header.Number
	} else {
		lastProvenL1BlockNumber = new(big.Int).Sub(lastProvenL1BlockNumber, bigint.One)
	}
	latestProvenL1HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Table("block_headers_"+ep.l1ChainId).Where("number > ?", lastProvenL1BlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers",
			headers.Order("number ASC").Limit(int(ep.epoch))).Select("MAX(number)"))
	}
	if latestProvenL1HeaderScope == nil {
		return nil
	}
	latestL1Header, err := ep.db.Blocks.ChainBlockHeaderWithScope(latestProvenL1HeaderScope, ep.l1ChainId)
	if err != nil {
		return fmt.Errorf("failed to query for latest unfinalized L1 state: %w", err)
	} else if latestL1Header == nil {
		return nil
	}
	fromL1Height, toL1Height := new(big.Int).Add(lastProvenL1BlockNumber, bigint.One), latestL1Header.Number
	if err := ep.db.Transaction(func(tx *database.DB) error {
		if err := bridge.L1ProcessProvenBridgeEvents(tx, fromL1Height, toL1Height,
			ep.l1ChainId, ep.l2ChainId); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	ep.LatestProvenL1Header = latestL1Header
	return nil
}

func (ep *MantleEventProcessor) processFinalizedL1Events() error {
	lastFinalizedL1BlockNumber := ep.L1StartHeight
	if ep.LatestFinalizedL1Header != nil {
		lastFinalizedL1BlockNumber = ep.LatestFinalizedL1Header.Number
	} else {
		lastFinalizedL1BlockNumber = new(big.Int).Sub(lastFinalizedL1BlockNumber, bigint.One)
	}
	latestFinalizedL1HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Table("block_headers_"+ep.l1ChainId).Where("number > ?", lastFinalizedL1BlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers",
			headers.Order("number ASC").Limit(int(ep.epoch))).Select("MAX(number)"))
	}
	if latestFinalizedL1HeaderScope == nil {
		return nil
	}
	// latestL1Header, err := ep.db.Blocks.BlockHeaderWithScope(latestFinalizedL1HeaderScope)
	latestL1Header, err := ep.db.Blocks.ChainBlockHeaderWithScope(latestFinalizedL1HeaderScope, ep.l1ChainId)
	if err != nil {
		return fmt.Errorf("failed to query for latest unfinalized L1 state: %w", err)
	} else if latestL1Header == nil {
		return nil
	}
	fromL1Height, toL1Height := new(big.Int).Add(lastFinalizedL1BlockNumber, bigint.One), latestL1Header.Number
	if err := ep.db.Transaction(func(tx *database.DB) error {
		l1BedrockStartingHeight := ep.L1BedrockStartHeight
		if l1BedrockStartingHeight.Cmp(fromL1Height) > 0 {
			legacyFromL1Height, legacyToL1Height := fromL1Height, toL1Height
			if l1BedrockStartingHeight.Cmp(toL1Height) <= 0 {
				legacyToL1Height = new(big.Int).Sub(l1BedrockStartingHeight, bigint.One)
			}
			if err := bridge.LegacyL1ProcessFinalizedBridgeEvents(tx, legacyFromL1Height, legacyToL1Height,
				ep.l1ChainId, ep.l2ChainId); err != nil {
				return err
			} else if legacyToL1Height.Cmp(toL1Height) == 0 {
				return nil
			}
			fromL1Height = l1BedrockStartingHeight
		}
		return bridge.L1ProcessFinalizedBridgeEvents(tx, fromL1Height, toL1Height, ep.l1ChainId, ep.l2ChainId)
	}); err != nil {
		return err
	}
	ep.LatestFinalizedL1Header = latestL1Header
	return nil
}

func (ep *MantleEventProcessor) processFinalizedL2Events() error {
	lastFinalizedL2BlockNumber := ep.L2StartHeight
	latestL1L2BlockNumber := ep.L2StartHeight
	if ep.LatestL1L2FinalizedL2Header != nil {
		lastFinalizedL2BlockNumber = ep.LatestL1L2FinalizedL2Header.Number
	} else {
		lastFinalizedL2BlockNumber = new(big.Int).Sub(lastFinalizedL2BlockNumber, bigint.One)
	}
	l2BlockHeader, err := ep.db.L1ToL2.L2LatestBlockHeader(ep.l2ChainId)
	if err != nil {
		log.Error("get latest l2 block header fail", "err", err)
		return err
	}
	if l2BlockHeader != nil {
		ep.LatestL1L2L2Header = l2BlockHeader
		latestL1L2BlockNumber = l2BlockHeader.Number
	}
	latestL2HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Table("block_headers_"+ep.l2ChainId).Where("number > ? AND number <= ?", lastFinalizedL2BlockNumber, latestL1L2BlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(int(ep.epoch))).Select("MAX(number)"))
	}
	latestL2Header, err := ep.db.Blocks.ChainBlockHeaderWithScope(latestL2HeaderScope, ep.l2ChainId)
	if err != nil {
		return fmt.Errorf("failed to query for latest unfinalized L2 state: %w", err)
	} else if latestL2Header == nil {
		ep.LatestL1L2FinalizedL2Header = latestL2Header
		return nil
	}
	fromL2Height, toL2Height := new(big.Int).Add(lastFinalizedL2BlockNumber, bigint.One), latestL2Header.Number
	if err := ep.db.Transaction(func(tx *database.DB) error {
		l2BedrockStartingHeight := ep.L2BedrockStartHeight
		if l2BedrockStartingHeight.Cmp(fromL2Height) > 0 {
			legacyFromL2Height, legacyToL2Height := fromL2Height, toL2Height
			if l2BedrockStartingHeight.Cmp(toL2Height) <= 0 {
				legacyToL2Height = new(big.Int).Sub(l2BedrockStartingHeight, bigint.One)
			}
			if err := bridge.LegacyL2ProcessFinalizedBridgeEvents(tx, legacyFromL2Height, legacyToL2Height,
				ep.l1ChainId, ep.l2ChainId); err != nil {
				return err
			} else if legacyToL2Height.Cmp(toL2Height) == 0 {
				return nil
			}
			fromL2Height = l2BedrockStartingHeight
		}
		return bridge.L2ProcessFinalizedBridgeEvents(tx, fromL2Height, toL2Height, ep.l1ChainId, ep.l2ChainId)
	}); err != nil {
		return err
	}
	ep.LatestL1L2FinalizedL2Header = latestL2Header
	return nil
}

func (ep *MantleEventProcessor) processRollupStateRoot() error {
	lastStateRootL1BlockNumber := ep.L1StartHeight
	if ep.LatestStateRootL1Header != nil {
		lastStateRootL1BlockNumber = ep.LatestStateRootL1Header.Number
	} else {
		lastStateRootL1BlockNumber = new(big.Int).Sub(lastStateRootL1BlockNumber, bigint.One)
	}
	latestRollupL1HeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Table("block_headers_"+ep.l1ChainId).Where("number > ?", lastStateRootL1BlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(int(ep.epoch))).Select("MAX(number)"))
	}
	latestL1StateRootHeader, err := ep.db.Blocks.ChainBlockHeaderWithScope(latestRollupL1HeaderScope, ep.l1ChainId)
	if err != nil {
		return fmt.Errorf("failed to query new L1 state: %w", err)
	} else if latestL1StateRootHeader == nil {
		return nil
	}
	fromL1Height, toL1Height := new(big.Int).Add(lastStateRootL1BlockNumber, bigint.One), latestL1StateRootHeader.Number
	if err := ep.db.Transaction(func(tx *database.DB) error {
		if err := stateroot.L2OutputEvent(tx, fromL1Height, toL1Height, ep.l1ChainId, ep.l2ChainId); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	ep.LatestStateRootL1Header = latestL1StateRootHeader
	return nil
}
