package common

import (
	"errors"
	"github.com/ethereum/go-ethereum/log"
	"math/big"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	common2 "github.com/cornerstone-labs/acorus/database/utils"
)

type BlockHeader struct {
	Hash       common.Hash `gorm:"primaryKey;serializer:bytes"`
	ParentHash common.Hash `gorm:"serializer:bytes"`
	Number     *big.Int    `gorm:"serializer:u256"`
	Timestamp  uint64
	RLPHeader  *common2.RLPHeader `gorm:"serializer:rlp;column:rlp_bytes"`
}

func BlockHeaderFromHeader(header *types.Header) BlockHeader {
	return BlockHeader{
		Hash:       header.Hash(),
		ParentHash: header.ParentHash,
		Number:     header.Number,
		Timestamp:  header.Time,
		RLPHeader:  (*common2.RLPHeader)(header),
	}
}

type L1BlockHeader struct {
	BlockHeader `gorm:"embedded"`
}

type L2BlockHeader struct {
	BlockHeader `gorm:"embedded"`
}

type BlocksView interface {
	L1BlockHeader(common.Hash) (*L1BlockHeader, error)
	L1BlockHeaderWithFilter(BlockHeader) (*L1BlockHeader, error)
	L1BlockHeaderWithScope(func(db *gorm.DB) *gorm.DB) (*L1BlockHeader, error)
	L1LatestBlockHeader() (*L1BlockHeader, error)

	L2BlockHeader(common.Hash) (*L2BlockHeader, error)
	L2BlockHeaderWithFilter(BlockHeader) (*L2BlockHeader, error)
	L2BlockHeaderWithScope(func(db *gorm.DB) *gorm.DB) (*L2BlockHeader, error)
	L2LatestBlockHeader() (*L2BlockHeader, error)

	LatestObservedEpochForL1(*big.Int, uint64) (*EpochL1, error)

	LatestObservedEpochForL2(*big.Int, uint64) (*EpochL2, error)

	L2BlockTimeStampByNum(blockNumber uint64) (int64, error)
}

type BlocksDB interface {
	BlocksView

	StoreL1BlockHeaders([]L1BlockHeader) error
	StoreL2BlockHeaders([]L2BlockHeader) error
}

/**
 * Implementation
 */

type blocksDB struct {
	gorm *gorm.DB
}

func NewBlocksDB(db *gorm.DB) BlocksDB {
	return &blocksDB{gorm: db}
}

// L1 module

func (db *blocksDB) StoreL1BlockHeaders(headers []L1BlockHeader) error {
	result := db.gorm.CreateInBatches(&headers, common2.BatchInsertSize)
	return result.Error
}

func (db *blocksDB) L1BlockHeader(hash common.Hash) (*L1BlockHeader, error) {
	return db.L1BlockHeaderWithFilter(BlockHeader{Hash: hash})
}

func (db *blocksDB) L1BlockHeaderWithFilter(filter BlockHeader) (*L1BlockHeader, error) {
	return db.L1BlockHeaderWithScope(func(gorm *gorm.DB) *gorm.DB { return gorm.Where(&filter) })
}

func (db *blocksDB) L1BlockHeaderWithScope(scope func(*gorm.DB) *gorm.DB) (*L1BlockHeader, error) {
	var l1Header L1BlockHeader
	result := db.gorm.Scopes(scope).Take(&l1Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l1Header, nil
}

func (db *blocksDB) L1LatestBlockHeader() (*L1BlockHeader, error) {
	var l1Header L1BlockHeader
	result := db.gorm.Order("number DESC").Take(&l1Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &l1Header, nil
}

// L2 module

func (db *blocksDB) StoreL2BlockHeaders(headers []L2BlockHeader) error {
	result := db.gorm.CreateInBatches(&headers, common2.BatchInsertSize)
	return result.Error
}

func (db *blocksDB) L2BlockHeader(hash common.Hash) (*L2BlockHeader, error) {
	return db.L2BlockHeaderWithFilter(BlockHeader{Hash: hash})
}

func (db *blocksDB) L2BlockHeaderWithFilter(filter BlockHeader) (*L2BlockHeader, error) {
	return db.L2BlockHeaderWithScope(func(gorm *gorm.DB) *gorm.DB { return gorm.Where(&filter) })
}

func (db *blocksDB) L2BlockHeaderWithScope(scope func(*gorm.DB) *gorm.DB) (*L2BlockHeader, error) {
	var l2Header L2BlockHeader
	result := db.gorm.Scopes(scope).Take(&l2Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &l2Header, nil
}

func (db *blocksDB) L2LatestBlockHeader() (*L2BlockHeader, error) {
	var l2Header L2BlockHeader
	result := db.gorm.Order("number DESC").Take(&l2Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &l2Header, nil
}

// Auxiliary Methods on both L1 & L2

type EpochL1 struct {
	L1BlockHeader L1BlockHeader `gorm:"embedded"`
}

type EpochL2 struct {
	L2BlockHeader L2BlockHeader `gorm:"embedded"`
}

func (db *blocksDB) LatestObservedEpochForL1(fromL1Height *big.Int, maxL1Range uint64) (*EpochL1, error) {
	var startBlockNumber *big.Int
	if fromL1Height != nil {
		startBlockNumber = fromL1Height
	} else {
		startBlockNumber = big.NewInt(0)
	}
	query := db.gorm.Table("l1_block_headers").Where("number >= ? AND number < ?", startBlockNumber, maxL1Range).Order("number DESC").Limit(1)
	var epoch EpochL1
	result := query.Take(&epoch)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Info("BlockHeader record not found according to epoch")
			return nil, nil
		}
		return nil, result.Error
	}
	return &epoch, nil
}

func (db *blocksDB) LatestObservedEpochForL2(fromL2Height *big.Int, maxL1Range uint64) (*EpochL2, error) {
	var startBlockNumber *big.Int
	if fromL2Height != nil {
		startBlockNumber = fromL2Height
	} else {
		startBlockNumber = big.NewInt(0)
	}
	query := db.gorm.Table("l2_block_headers").Where("number >= ? AND number < ?", startBlockNumber, maxL1Range).Order("number DESC").Limit(1)
	var epoch EpochL2
	result := query.Take(&epoch)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Info("BlockHeader record not found according to epoch")
			return nil, nil
		}
		return nil, result.Error
	}
	return &epoch, nil
}

func (db *blocksDB) L2BlockTimeStampByNum(blockNumber uint64) (int64, error) {
	var l2BlockTimestamp int64
	result := db.gorm.Table("l2_block_headers").Where("number = ?", blockNumber).Select("timestamp").Take(&l2BlockTimestamp)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, result.Error
	}
	return l2BlockTimestamp, nil
}
