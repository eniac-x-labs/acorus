package appchain

import (
	"errors"
	"github.com/ethereum/go-ethereum/log"

	"gorm.io/gorm"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type AppChainDEthTransfer struct {
	GUID          uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	TxHash        common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	BlockNumber   *big.Int       `json:"block_number" gorm:"serializer:u256"`
	From          common.Address `json:"from_address" column:"from_address" gorm:"serializer:bytes"`
	To            common.Address `json:"to_address" column:"to_address" gorm:"serializer:bytes"`
	Claimed       int8           `json:"claimed"`
	Shares        *big.Int       `json:"shares" gorm:"serializer:u256"`
	MessageNonce  *big.Int       `json:"message_nonce" gorm:"serializer:u256"`
	NotifyRelayer bool           `json:"notify_relayer"`
	Created       uint64         `json:"created"`
}

func (AppChainDEthTransfer) TableName() string {
	return "ac_chain_deth_transfer"
}

type appChainDEthTransferDB struct {
	gorm *gorm.DB
}

type AppChainDEthTransferDB interface {
	StoreAppChainDEthTransfer(dethTransfer AppChainDEthTransfer) error
	ClaimDethTransfer(messageNonce *big.Int) error
	NotifyTransferL2Success(TxHash common.Hash) error
	AppChainDEthTransferDBView
}

type AppChainDEthTransferDBView interface {
	ListDEthTransferNoRelayer() []AppChainDEthTransfer
}

func NewAppChainDEthTransferDB(db *gorm.DB) AppChainDEthTransferDB {
	return &appChainDEthTransferDB{gorm: db}
}

func (db appChainDEthTransferDB) StoreAppChainDEthTransfer(dethTransfer AppChainDEthTransfer) error {
	if dethTransfer.TxHash.String() == "" {
		return nil
	}
	var exits AppChainDEthTransfer
	err := db.gorm.Table(dethTransfer.TableName()).Where(AppChainDEthTransfer{
		TxHash: dethTransfer.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(dethTransfer.TableName()).Omit("guid").Create(&dethTransfer)
			return result.Error
		}
	}
	return err
}

func (db appChainDEthTransferDB) ListDEthTransferNoRelayer() []AppChainDEthTransfer {
	var appChainDEthTransfers []AppChainDEthTransfer
	err := db.gorm.Table(AppChainDEthTransfer{}.TableName()).Where("notify_relayer = ?", false).
		Limit(20).Find(&appChainDEthTransfers)
	if err.Error != nil {
		log.Error("AppChainDEthTransfer ListDEthTransferNoRelayer", "err", err.Error)
	}
	return appChainDEthTransfers
}

func (db appChainDEthTransferDB) NotifyTransferL2Success(TxHash common.Hash) error {
	return db.gorm.Table(AppChainMigrateShares{}.TableName()).Where(
		AppChainMigrateShares{TxHash: TxHash}).Update("notify_relayer", true).Error
}

func (db appChainDEthTransferDB) ClaimDethTransfer(messageNonce *big.Int) error {
	return db.gorm.Table(AppChainDEthTransfer{}.TableName()).Where(
		AppChainDEthTransfer{MessageNonce: messageNonce}).Update("claimed", 1).Error
}
