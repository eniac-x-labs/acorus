package common

import (
	"errors"
	"math/big"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

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

type ChainBlockHeader struct {
	BlockHeader `gorm:"embedded"`
}

type BlocksView interface {
	ChainBlockHeader(string, common.Hash) (*BlockHeader, error)
	ChainBlockHeaderWithFilter(string, BlockHeader) (*BlockHeader, error)
	ChainBlockHeaderWithScope(func(db *gorm.DB) *gorm.DB, string) (*BlockHeader, error)
	ChainLatestBlockHeader(string) (*BlockHeader, error)

	LatestObservedEpochForChain(string, *big.Int, uint64) (*BlockHeader, error)

	BlockTimeStampByNum(string, uint64) (int64, error)
}

type BlocksDB interface {
	BlocksView
	StoreBlockHeaders(string, []ChainBlockHeader) error
}

type blocksDB struct {
	gorm *gorm.DB
}

func NewBlocksDB(db *gorm.DB) BlocksDB {
	return &blocksDB{gorm: db}
}

func (db *blocksDB) StoreBlockHeaders(chainId string, headers []ChainBlockHeader) error {
	result := db.gorm.Table("block_header_"+chainId).CreateInBatches(&headers, common2.BatchInsertSize)
	return result.Error
}

func (db *blocksDB) ChainBlockHeader(chainId string, hash common.Hash) (*BlockHeader, error) {
	return db.ChainBlockHeaderWithFilter(chainId, BlockHeader{Hash: hash})
}

func (db *blocksDB) ChainBlockHeaderWithFilter(chainId string, filter BlockHeader) (*BlockHeader, error) {
	return db.ChainBlockHeaderWithScope(func(gorm *gorm.DB) *gorm.DB { return gorm.Where(&filter) }, chainId)
}

func (db *blocksDB) ChainBlockHeaderWithScope(scope func(*gorm.DB) *gorm.DB, chainId string) (*BlockHeader, error) {
	var header BlockHeader
	result := db.gorm.Table("block_header_" + chainId).Scopes(scope).Take(&header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &header, nil
}

func (db *blocksDB) ChainLatestBlockHeader(chainId string) (*BlockHeader, error) {
	var header BlockHeader
	result := db.gorm.Table("block_header_" + chainId).Order("number DESC").Take(&header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &header, nil
}

func (db *blocksDB) LatestObservedEpochForChain(chainId string, fromL1Height *big.Int, maxL1Range uint64) (*BlockHeader, error) {
	var startBlockNumber *big.Int
	if fromL1Height != nil {
		startBlockNumber = fromL1Height
	} else {
		startBlockNumber = big.NewInt(0)
	}
	query := db.gorm.Table("block_header_"+chainId).Where("number >= ? AND number < ?", startBlockNumber, maxL1Range).Order("number DESC").Limit(1)
	var epoch BlockHeader
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

func (db *blocksDB) BlockTimeStampByNum(chainId string, blockNumber uint64) (int64, error) {
	var blockTimestamp int64
	result := db.gorm.Table("block_header_"+chainId).Where("number = ?", blockNumber).Select("timestamp").Take(&blockTimestamp)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, result.Error
	}
	return blockTimestamp, nil
}
