package appchain

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
)

type AppChainUnStakeSingle struct {
	GUID          uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	BlockNumber   *big.Int       `json:"block_number" gorm:"serializer:u256"`
	TxHash        common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	LockedAmount  *big.Int       `json:"locked_amount" gorm:"serializer:u256"`
	Staker        common.Address `json:"staker" gorm:"serializer:bytes"`
	L2Strategy    common.Address `json:"l2_strategy" gorm:"serializer:bytes"`
	NotifyRelayer bool           `json:"notify_relayer"`
	Created       uint64         `json:"created"`
	Updated       uint64         `json:"updated"`
}

func (AppChainUnStakeSingle) TableName() string {
	return "ac_chain_unstake_single"
}

type appChainUnStakeSingleDB struct {
	gorm *gorm.DB
}

type AppChainUnStakeSingleDB interface {
	AppChainUnStakeSingleDBView
	StoreAppChainUnStakeSingle(chainUnStakeSingle AppChainUnStakeSingle) error
	NotifyAppChainUnStakeSingle(txHash string) error
}

type AppChainUnStakeSingleDBView interface {
	ListAppChainUnStakeSingle() (chainUnStakeSingle []AppChainUnStakeSingle)
}

func NewAppChainUnStakeSingleDB(db *gorm.DB) AppChainUnStakeSingleDB {
	return &appChainUnStakeSingleDB{gorm: db}
}

func (db appChainUnStakeSingleDB) StoreAppChainUnStakeSingle(chainUnStakeSingle AppChainUnStakeSingle) error {
	if chainUnStakeSingle.TxHash.String() == "" {
		return nil
	}
	var exits AppChainUnStakeSingle
	err := db.gorm.Table(chainUnStakeSingle.TableName()).Where(AppChainUnStakeSingle{TxHash: chainUnStakeSingle.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(chainUnStakeSingle.TableName()).Omit("guid").Create(&chainUnStakeSingle)
			return result.Error
		}
	}
	return err
}

func (db appChainUnStakeSingleDB) ListAppChainUnStakeSingle() (chainUnStakeSingle []AppChainUnStakeSingle) {
	getSql := `
		select * from  ac_chain_unstake_single where notify_relayer = false  LIMIT 20
	`
	result := db.gorm.Raw(getSql).Find(&chainUnStakeSingle)
	err := result.Error
	if err != nil {
		return nil
	}
	return chainUnStakeSingle
}

func (db appChainUnStakeSingleDB) NotifyAppChainUnStakeSingle(txHash string) error {
	chainUnStakeSingle := new(AppChainUnStakeSingle)
	result := db.gorm.Table(chainUnStakeSingle.TableName()).Where("tx_hash = ?", txHash).Updates(map[string]interface{}{"notify_relayer": true})
	return result.Error
}
