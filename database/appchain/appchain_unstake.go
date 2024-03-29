package appchain

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
)

type AppChainUnStake struct {
	GUID          uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	BlockNumber   *big.Int       `json:"block_number" gorm:"serializer:u256"`
	TxHash        common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	EthAmount     *big.Int       `json:"eth_amount" gorm:"serializer:u256"`
	DETHLocked    *big.Int       `json:"locked_amount" gorm:"serializer:u256"`
	ClaimTxHash   common.Hash    `json:"claim_tx_hash" gorm:"serializer:bytes"`
	L2Strategy    common.Address `json:"l2_strategy" gorm:"serializer:bytes"`
	Staker        common.Address `json:"staker" gorm:"serializer:bytes"`
	Bridge        common.Address `json:"bridge" gorm:"serializer:bytes"`
	SourceChainId string         `json:"source_chain_id"`
	DestChainId   string         `json:"dest_chain_id"`
	Status        uint8          `json:"status"`
	NotifyRelayer bool           `json:"notify_relayer"`
	Created       uint64         `json:"created"`
	Updated       uint64         `json:"updated"`
}

func (AppChainUnStake) TableName() string {
	return "ac_chain_unstake"
}

type appChainUnStakeDB struct {
	gorm *gorm.DB
}

type AppChainUnStakeDB interface {
	AppChainUnStakeDBView
	StoreAppChainUnStake(cainUnStakeBatch AppChainUnStake) error
	NotifyAppChainUnStake(txHash string) error
	ClaimAppChainUnStake(chainUnStakeBatch AppChainUnStake, noClaim uint8) error
}

type AppChainUnStakeDBView interface {
	ListAppChainUnStake() (chainUnStakeBatch []AppChainUnStake)
}

func NewAppChainUnStakeDB(db *gorm.DB) AppChainUnStakeDB {
	return &appChainUnStakeDB{gorm: db}
}

func (db appChainUnStakeDB) StoreAppChainUnStake(chainUnStakeBatch AppChainUnStake) error {
	if chainUnStakeBatch.TxHash.String() == "" {
		return nil
	}
	var exits AppChainUnStake
	err := db.gorm.Table(chainUnStakeBatch.TableName()).Where(AppChainUnStake{TxHash: chainUnStakeBatch.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(chainUnStakeBatch.TableName()).Omit("guid").Create(&chainUnStakeBatch)
			return result.Error
		}
	}
	return err
}

func (db appChainUnStakeDB) ListAppChainUnStake() (chainUnStakeBatch []AppChainUnStake) {
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

func (db appChainUnStakeDB) NotifyAppChainUnStake(txHash string) error {
	chainUnStakeBatch := new(AppChainUnStake)
	result := db.gorm.Table(chainUnStakeBatch.TableName()).Where("tx_hash = ?", txHash).Updates(map[string]interface{}{"notify_relayer": true})
	return result.Error
}

func (db appChainUnStakeDB) ClaimAppChainUnStake(chainUnStakeBatch AppChainUnStake, noClaim uint8) error {
	var exits AppChainUnStake
	err := db.gorm.Table(chainUnStakeBatch.TableName()).Where(AppChainUnStake{L2Strategy: chainUnStakeBatch.L2Strategy,
		SourceChainId: chainUnStakeBatch.SourceChainId,
		DestChainId:   chainUnStakeBatch.DestChainId,
		Staker:        chainUnStakeBatch.Staker,
		Status:        noClaim,
	}).Take(&exits)
	if err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return err.Error
	}
	exits.Status = chainUnStakeBatch.Status
	exits.Updated = chainUnStakeBatch.Updated
	exits.ClaimTxHash = chainUnStakeBatch.ClaimTxHash
	exits.Bridge = chainUnStakeBatch.Bridge

	return db.gorm.Table(chainUnStakeBatch.TableName()).Save(exits).Error
}
