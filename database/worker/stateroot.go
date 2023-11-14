package worker

import (
	"gorm.io/gorm"
	"math/big"

	"github.com/acmestack/gorm-plus/gplus"

	"github.com/ethereum/go-ethereum/common"
)

type StateRoot struct {
	BlockHash         common.Hash `gorm:"serializer:bytes;column:block_hash;primaryKey" json:"block_hash"`
	TransactionHash   common.Hash `gorm:"serializer:bytes;column:transaction_hash" json:"transaction_hash"`
	L1BlockNumber     *big.Int    `gorm:"serializer:u256" json:"l1_block_number"`
	L2BlockNumber     *big.Int    `gorm:"serializer:u256" json:"l2_block_number"`
	OutputIndex       *big.Int    `gorm:"serializer:u256" json:"output_index"`
	PrevTotalElements *big.Int    `gorm:"serializer:u256" json:"prev_total_elements"`
	Status            int         `json:"status"`
	OutputRoot        string      `json:"output_root"`
	Canonical         bool        `json:"canonical"`
	BatchSize         *big.Int    `gorm:"serializer:u256" json:"batch_size"`
	Timestamp         uint64      `json:"timestamp"`
}

func (StateRoot) TableName() string {
	return "state_root"
}

type StateRootDB interface {
	StateRootView
	StoreBatchStateRoots([]StateRoot) error
}

type StateRootView interface {
	StateRootList(int, int) (*gplus.Page[StateRoot], error)
	StateRootByIndex(index *big.Int) (*StateRoot, error)
	GetLatestStateRootL2BlockNumber() (uint64, error)
}

/**
 * Implementation
 */
type stateRootDB struct {
	gorm *gorm.DB
}

func (s stateRootDB) StoreBatchStateRoots(roots []StateRoot) error {
	result := s.gorm.CreateInBatches(&roots, len(roots))
	return result.Error
}

func (s stateRootDB) StateRootList(page int, pageSize int) (*gplus.Page[StateRoot], error) {
	query, _ := gplus.NewQuery[StateRoot]()
	dbPage := gplus.NewPage[StateRoot](page, pageSize)
	rst, rstDB := gplus.SelectPage(dbPage, query)
	if rstDB.Error != nil {
		return nil, rstDB.Error
	}
	return rst, nil
}

func (s stateRootDB) GetLatestStateRootL2BlockNumber() (uint64, error) {
	var l2BlockNumber uint64
	err := s.gorm.Table("state_root").Select("l2_block_number").Order("timestamp DESC").Limit(1).Find(&l2BlockNumber).Error
	if err != nil {
		return 0, err
	}
	return l2BlockNumber, nil

}

func (s stateRootDB) StateRootByIndex(index *big.Int) (*StateRoot, error) {
	query, sr := gplus.NewQuery[StateRoot]()
	query.Eq(&sr.OutputIndex, index)
	stateRoot, resultDb := gplus.SelectOne(query)
	if resultDb.Error != nil {
		return nil, resultDb.Error
	}
	return stateRoot, nil

}

func NewStateRootDB(db *gorm.DB) StateRootDB {
	gplus.Init(db)
	return &stateRootDB{gorm: db}
}
