package appchain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
)

type AppChainIncreaseBatch struct {
	GUID            uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	StrategyAddress common.Address `json:"strategy_address"gorm:"serializer:bytes"`
	Staker          common.Address `json:"staker"gorm:"serializer:bytes"`
	Operator        common.Address `json:"operator"gorm:"serializer:bytes"`
	Shares          *big.Int       `json:"shares" gorm:"serializer:u256"`
	ChainId         string         `json:"chain_id"`
	BatchId         string         `json:"batch_id"`
	Created         uint64         `json:"created"`
	NotifyRelayer   *bool          `json:"notify_relayer"`
}

func (AppChainIncreaseBatch) TableName() string {
	return "ac_chain_increase_batch"
}

type appChainIncreaseBatchDB struct {
	gorm *gorm.DB
}

type AppChainIncreaseBatchDB interface {
	StoreAppChainIncreasedBatch(appChainIncreaseBatch []AppChainIncreaseBatch) error
	NotifyBatchMintSuccess(batchId string) error
	AppChainIncreaseBatchDBView
}

type AppChainIncreaseBatchDBView interface {
	ListBatchDataByBatchId(batchId string) []AppChainIncreaseBatch
}

func NewAppChainIncreaseBatchDB(db *gorm.DB) AppChainIncreaseBatchDB {
	return &appChainIncreaseBatchDB{gorm: db}
}

func (db appChainIncreaseBatchDB) StoreAppChainIncreasedBatch(appChainIncreaseBatch []AppChainIncreaseBatch) error {
	return db.gorm.Omit("guid").Create(&appChainIncreaseBatch).Error
}

func (db appChainIncreaseBatchDB) ListBatchDataByBatchId(batchId string) []AppChainIncreaseBatch {
	var appChainIncreaseBatch []AppChainIncreaseBatch
	var noNotify = false
	err := db.gorm.Table(AppChainIncreaseBatch{}.TableName()).Where(AppChainIncreaseBatch{BatchId: batchId, NotifyRelayer: &noNotify}).Find(&appChainIncreaseBatch)
	if err != nil {
		log.Error("ListBatchDataByBatchId", "err", err)
	}
	return appChainIncreaseBatch
}

func (db appChainIncreaseBatchDB) NotifyBatchMintSuccess(batchId string) error {
	return db.gorm.Table(AppChainIncreaseBatch{}.TableName()).Where(AppChainIncreaseBatch{BatchId: batchId}).Update("notify_relayer", true).Error
}
