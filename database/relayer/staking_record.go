package relayer

import (
	"errors"
	"math/big"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type StakingRecord struct {
	GUID        uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')"`
	TxHash      common.Hash    `json:"txHash" gorm:"serializer:bytes"`
	BlockNumber *big.Int       `json:"blockNumber" gorm:"serializer:u256"`
	UserAddress common.Address `json:"userAddress" gorm:"serializer:bytes"`
	Token       common.Address `json:"token" gorm:"serializer:bytes"`
	Amount      *big.Int       `json:"amount" gorm:"serializer:u256"`
	Reward      *big.Int       `json:"reward" gorm:"serializer:u256"`
	StartPoolId *big.Int       `json:"start_pool_id" gorm:"serializer:u256"`
	EndPoolId   *big.Int       `json:"end_pool_id" gorm:"serializer:u256"`
	TxType      int            `json:"txType"`
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
	GetStakingRecords(address string, page int, pageSize int, order string) (stakingRecords []StakingRecord, total int64)
}

func NewStakingRecordDB(db *gorm.DB) StakingRecordDB {
	return &stakingRecordDB{gorm: db}
}

func (db stakingRecordDB) StoreStakingRecord(stake StakingRecord) error {
	stakingRecord := new(StakingRecord)
	var exitsStake StakingRecord
	err := db.gorm.Table(stakingRecord.TableName()).Where("tx_hash = ?", stake.TxHash).Take(&exitsStake).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(stakingRecord.TableName()).Omit("guid").Create(&stake)
			return result.Error
		}
	}
	return err
}

func (db stakingRecordDB) StoreStakingRecords(stakes []StakingRecord) error {
	stakingRecord := new(StakingRecord)
	result := db.gorm.Table(stakingRecord.TableName()).Omit("guid").Create(&stakes)
	return result.Error
}

func (db stakingRecordDB) GetStakingRecords(address string, page int, pageSize int, order string) (sR []StakingRecord, total int64) {
	var totalRecord int64
	var stakingRecords []StakingRecord
	querySR := db.gorm.Table("staking_record")
	if address != "0x00" {
		err := db.gorm.Table("staking_record").Select("DISTINCT ON (tx_hash) guid").Where("user_address = ?", address).Count(&totalRecord).Error
		if err != nil {
			log.Error("get staking records by address count fail")
		}
		querySR.Where("user_address = ?", address).Offset((page - 1) * pageSize).Limit(pageSize)
	} else {
		err := db.gorm.Table("staking_record").Select("DISTINCT ON (tx_hash) guid").Count(&totalRecord).Error
		if err != nil {
			log.Error("get staking records no address count fail ")
		}
		querySR.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	if strings.ToLower(order) == "asc" {
		querySR.Order("timestamp asc")
	} else {
		querySR.Order("timestamp desc")
	}
	qErr := querySR.Select("DISTINCT ON (tx_hash) *").Find(&stakingRecords).Error
	if qErr != nil {
		log.Error("get staking records fail", "err", qErr)
	}
	return stakingRecords, totalRecord
}
