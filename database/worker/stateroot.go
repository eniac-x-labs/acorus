package worker

import (
	"errors"
	"gorm.io/gorm"
	"math/big"
	"strings"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	common2 "github.com/cornerstone-labs/acorus/database/common"
)

type StateRoot struct {
	GUID              uuid.UUID   `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','';serializer:uuid" json:"guid"`
	BlockHash         common.Hash `gorm:"serializer:bytes;column:block_hash;primaryKey" json:"blockHash"`
	TransactionHash   common.Hash `gorm:"serializer:bytes;column:transaction_hash" json:"transactionHash"`
	L1BlockNumber     *big.Int    `gorm:"serializer:u256" json:"l1BlockNumber"`
	L2BlockNumber     *big.Int    `gorm:"serializer:u256" json:"l2BlockNumber"`
	OutputIndex       *big.Int    `gorm:"serializer:u256" json:"outputIndex"`
	PrevTotalElements *big.Int    `gorm:"serializer:u256" json:"prevTotalElements"`
	Status            int         `json:"status"`
	OutputRoot        string      `json:"outputRoot"`
	Canonical         bool        `json:"canonical"`
	BatchSize         *big.Int    `gorm:"serializer:u256" json:"batchSize"`
	Timestamp         uint64      `json:"timestamp"`
	BlockSize         uint64      `json:"blockSize"`
}

type StateRootDB interface {
	StateRootView
	StoreBatchStateRoots(string, []StateRoot) error
}

type StateRootView interface {
	StateRootL1BlockHeader(chainId, l1ChainId string) (*common2.BlockHeader, error)
	StateRootList(string, int, int, string) ([]StateRoot, int64)
	GetLatestStateRootL2BlockNumber(chainId string) (uint64, error)
	UpdateSafeStatus(chainId string, safeBlockNumber *big.Int) error
	UpdateFinalizedStatus(chainId string, finalizedBlockNumber *big.Int) error
}

type stateRootDB struct {
	gorm *gorm.DB
}

func (s stateRootDB) StoreBatchStateRoots(chainId string, roots []StateRoot) error {
	var firstBlockSize uint64
	result := s.gorm.Table("state_root_"+chainId).Where("timestamp < ?", roots[0].Timestamp).Select("l2_block_number").Take(&firstBlockSize)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			roots[0].BlockSize = roots[0].L2BlockNumber.Uint64()
		} else {
			return result.Error
		}
	}
	roots[0].BlockSize = new(big.Int).Sub(roots[0].L2BlockNumber, new(big.Int).SetUint64(firstBlockSize)).Uint64()
	for i := 1; i < len(roots); i++ {
		roots[i].BlockSize = new(big.Int).Sub(roots[i].L2BlockNumber, roots[i-1].L2BlockNumber).Uint64()
	}
	result = s.gorm.Table("state_root_" + chainId).Omit("guid").Create(&roots)
	return result.Error
}

func (s stateRootDB) StateRootList(chainId string, page int, pageSize int, order string) ([]StateRoot, int64) {
	var totalRecord int64
	var stateRoots []StateRoot
	err := s.gorm.Table("state_root_" + chainId).Select("output_index").Count(&totalRecord).Error
	if err != nil {
		log.Error("get state root count fail")
	}
	queryStateRoot := s.gorm.Table("state_root").Offset((page - 1) * pageSize).Limit(pageSize)
	if strings.ToLower(order) == "asc" {
		queryStateRoot.Order("timestamp asc")
	} else {
		queryStateRoot.Order("timestamp desc")
	}
	qErr := queryStateRoot.Find(&stateRoots).Error
	if qErr != nil {
		log.Error("get state root list fail", "err", err)
	}
	return stateRoots, totalRecord
}

func (s stateRootDB) GetLatestStateRootL2BlockNumber(chainId string) (uint64, error) {
	var l2BlockNumber uint64
	err := s.gorm.Table("state_root_" + chainId).Select("l2_block_number").Order("timestamp DESC").Limit(1).Find(&l2BlockNumber).Error
	if err != nil {
		return 0, err
	}
	return l2BlockNumber, nil

}

func (s stateRootDB) UpdateSafeStatus(chainId string, safeBlockNumber *big.Int) error {
	err := s.gorm.Table("state_root_"+chainId).Where("l1_block_number < ? AND status = ?", safeBlockNumber.Uint64(), 0).Updates(map[string]interface{}{"status": 1}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s stateRootDB) UpdateFinalizedStatus(chainId string, finalizedBlockNumber *big.Int) error {
	err := s.gorm.Table("state_root_"+chainId).Where("l1_block_number < ?", finalizedBlockNumber.Uint64()).Updates(map[string]interface{}{"status": 2}).Error
	if err != nil {
		return err
	}
	return nil
}

func NewStateRootDB(db *gorm.DB) StateRootDB {
	return &stateRootDB{gorm: db}
}

func (s stateRootDB) StateRootL1BlockHeader(chainId, l1ChainId string) (*common2.BlockHeader, error) {
	l1Query := s.gorm.Table("block_headers_"+l1ChainId).Where("number = (?)", s.gorm.Table("state_root_"+chainId).Select("MAX(l1_block_number)"))
	var l1Header common2.BlockHeader
	result := l1Query.Take(&l1Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l1Header, nil
}
