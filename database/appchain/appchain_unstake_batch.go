package appchain

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
)

type AppChainUnStakeBatch struct {
	GUID            uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	BlockNumber     *big.Int       `json:"block_number" gorm:"serializer:u256"`
	TxHash          common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	BatchId         *big.Int       `json:"batch_id" gorm:"serializer:u256"`
	BatchEthAmount  *big.Int       `json:"batch_amount" gorm:"serializer:u256"`
	BatchDETHLocked *big.Int       `json:"batch_locked_amount" gorm:"serializer:u256"`
	ClaimTxHash     common.Hash    `json:"claim_tx_hash" gorm:"serializer:bytes"`
	Staker          common.Address `json:"staker" gorm:"serializer:bytes"`
	Bridge          common.Address `json:"bridge" gorm:"serializer:bytes"`
	SourceChainId   string         `json:"source_chain_id"`
	DestChainId     string         `json:"dest_chain_id"`
	Status          uint8          `json:"status"`
	NotifyRelayer   bool           `json:"notify_relayer"`
	Created         uint64         `json:"created"`
	Updated         uint64         `json:"updated"`
}

func (AppChainUnStakeBatch) TableName() string {
	return "ac_chain_unstake_batch"
}

type appChainUnStakeBatchDB struct {
	gorm *gorm.DB
}

type AppChainUnStakeBatchDB interface {
	AppChainUnStakeBatchDBView
	StoreAppChainUnStakeBatch(cainUnStakeBatch AppChainUnStakeBatch) error
	NotifyAppChainUnStakeBatch(txHash string) error
	ClaimAppChainUnStakeBatch(chainUnStakeBatch AppChainUnStakeBatch) error
}

type AppChainUnStakeBatchDBView interface {
	ListAppChainUnStakeBatch() (chainUnStakeBatch []AppChainUnStakeBatch)
}

func NewAppChainUnStakeBatchDB(db *gorm.DB) AppChainUnStakeBatchDB {
	return &appChainUnStakeBatchDB{gorm: db}
}

func (db appChainUnStakeBatchDB) StoreAppChainUnStakeBatch(chainUnStakeBatch AppChainUnStakeBatch) error {
	if chainUnStakeBatch.TxHash.String() == "" {
		return nil
	}
	var exits AppChainUnStakeBatch
	err := db.gorm.Table(chainUnStakeBatch.TableName()).Where(AppChainUnStakeBatch{TxHash: chainUnStakeBatch.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(chainUnStakeBatch.TableName()).Omit("guid").Create(&chainUnStakeBatch)
			return result.Error
		}
	}
	return err
}

func (db appChainUnStakeBatchDB) ListAppChainUnStakeBatch() (chainUnStakeBatch []AppChainUnStakeBatch) {
	getSql := `
		select * from  ac_chain_unstake_batch where notify_relayer = false  LIMIT 20
	`
	result := db.gorm.Raw(getSql).Find(&chainUnStakeBatch)
	err := result.Error
	if err != nil {
		return nil
	}
	return chainUnStakeBatch
}

func (db appChainUnStakeBatchDB) NotifyAppChainUnStakeBatch(txHash string) error {
	chainUnStakeBatch := new(AppChainUnStakeBatch)
	result := db.gorm.Table(chainUnStakeBatch.TableName()).Where("tx_hash = ?", txHash).Updates(map[string]interface{}{"notify_relayer": true})
	return result.Error
}

func (db appChainUnStakeBatchDB) ClaimAppChainUnStakeBatch(chainUnStakeBatch AppChainUnStakeBatch) error {
	var exits AppChainUnStakeBatch
	err := db.gorm.Table(chainUnStakeBatch.TableName()).Where(AppChainUnStakeBatch{BatchId: chainUnStakeBatch.BatchId}).Take(&exits)
	if err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return err.Error
	}
	exits.Status = chainUnStakeBatch.Status
	exits.SourceChainId = chainUnStakeBatch.SourceChainId
	exits.DestChainId = chainUnStakeBatch.DestChainId
	exits.Updated = chainUnStakeBatch.Updated
	exits.ClaimTxHash = chainUnStakeBatch.ClaimTxHash
	exits.Staker = chainUnStakeBatch.Staker
	exits.Bridge = chainUnStakeBatch.Bridge

	return db.gorm.Table(chainUnStakeBatch.TableName()).Save(exits).Error
}
