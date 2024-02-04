package relayer

import (
	"gorm.io/gorm"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type StakingRecord struct {
	GUID        uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')"`
	TxHash      common.Hash    `json:"txHash" gorm:"serializer:bytes"`
	BlockNumber *big.Int       `json:"blockNumber" gorm:"serializer:u256"`
	UserAddress common.Address `json:"userAddress" gorm:"serializer:bytes"`
	Token       common.Address `json:"token" gorm:"serializer:bytes"`
	Amount      *big.Int       `json:"amount"`
	ChainId     string         `json:"chainIid"`
	Status      int            `json:"status"`
	AssetType   int            `json:"assetType"`
	Timestamp   uint64         `json:"timestamp"`
}

func (StakingRecord) TableName() string {
	return "staking_record"
}

type stakingRecordDB struct {
	gorm *gorm.DB
}

type StakingRecordDB interface {
	StakingRecordView
	StoreStakingRecords(stakes []StakingRecord) error
	StoreStakingRecord(stake StakingRecord) error
}

type StakingRecordView interface {
}

func NewStakingRecordDB(db *gorm.DB) StakingRecordDB {
	return &stakingRecordDB{gorm: db}
}

func (db stakingRecordDB) StoreStakingRecord(stake StakingRecord) error {
	stakingRecord := new(StakingRecord)
	result := db.gorm.Table(stakingRecord.TableName()).Omit("guid").Create(&stake)
	return result.Error
}

func (db stakingRecordDB) StoreStakingRecords(stakes []StakingRecord) error {
	stakingRecord := new(StakingRecord)
	result := db.gorm.Table(stakingRecord.TableName()).Omit("guid").Create(&stakes)
	return result.Error
}
