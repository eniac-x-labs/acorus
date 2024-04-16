package appchain

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
)

type AppChainMigrateShares struct {
	GUID          uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	ChainId       string         `json:"chain_id"`
	TxHash        common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	Shares        *big.Int       `json:"shares" gorm:"serializer:u256"`
	Staker        common.Address `json:"staker" gorm:"serializer:bytes"`
	Strategy      common.Address `json:"strategy" gorm:"serializer:bytes"`
	BlockNumber   *big.Int       `json:"block_number" gorm:"serializer:u256"`
	NotifyRelayer bool           `json:"notify_relayer"`
	UnstakeNonce  *big.Int       `json:"unstake_nonce" gorm:"serializer:u256"`
	Created       uint64         `json:"created"`
}

func (AppChainMigrateShares) TableName() string {
	return "ac_migrate_shares"
}

type appChainMigrateSharesDB struct {
	gorm *gorm.DB
}

type AppChainMigrateSharesDB interface {
	StoreMigrateShares(migrateShares AppChainMigrateShares) error
	NotifyMigrateSharesSuccess(TxHash common.Hash) error
	AppChainMigrateSharesDBView
}
type AppChainMigrateSharesDBView interface {
	ListDataByNoNotifyRelayer() []AppChainMigrateShares
}

func NewAppChainMigrateSharesDB(db *gorm.DB) AppChainMigrateSharesDB {
	return &appChainMigrateSharesDB{gorm: db}
}

func (db appChainMigrateSharesDB) StoreMigrateShares(migrateShares AppChainMigrateShares) error {
	if migrateShares.TxHash.String() == "" {
		return nil
	}
	var exits AppChainMigrateShares
	err := db.gorm.Table(migrateShares.TableName()).Where(AppChainMigrateShares{TxHash: migrateShares.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(migrateShares.TableName()).Omit("guid").Create(&migrateShares)
			return result.Error
		}
	}
	return nil
}

func (db appChainMigrateSharesDB) NotifyMigrateSharesSuccess(TxHash common.Hash) error {
	return db.gorm.Table(AppChainMigrateShares{}.TableName()).Where(AppChainMigrateShares{TxHash: TxHash}).Update("notify_relayer", true).Error
}

func (db appChainMigrateSharesDB) ListDataByNoNotifyRelayer() []AppChainMigrateShares {
	var appChainMigrateShares []AppChainMigrateShares
	err := db.gorm.Table(AppChainMigrateShares{}.TableName()).Where("notify_relayer = ?", false).
		Limit(20).Find(&appChainMigrateShares)
	if err.Error != nil {
		log.Error("AppChainMigrateShares ListDataByNoNotifyRelayer", "err", err.Error)
	}
	return appChainMigrateShares
}
