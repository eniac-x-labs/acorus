package appchain

import (
	"github.com/ethereum/go-ethereum/common"
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
}

func (AppChainIncreaseBatch) TableName() string {
	return "ac_chain_increase_batch"
}

type appChainIncreaseBatchDB struct {
	gorm *gorm.DB
}

type AppChainIncreaseBatchDB interface {
	StoreAppChainIncreasedBatch(appChainIncreaseBatch []AppChainIncreaseBatch) error
	AppChainIncreaseBatchDBView
}

type AppChainIncreaseBatchDBView interface {
}

func NewAppChainIncreaseBatchDB(db *gorm.DB) AppChainIncreaseBatchDB {
	return &appChainIncreaseBatchDB{gorm: db}
}

func (db appChainIncreaseBatchDB) StoreAppChainIncreasedBatch(appChainIncreaseBatch []AppChainIncreaseBatch) error {
	return db.gorm.Omit("guid").Create(&appChainIncreaseBatch).Error
}
