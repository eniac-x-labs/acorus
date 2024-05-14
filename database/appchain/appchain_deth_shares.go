package appchain

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
)

type AppChainDethShares struct {
	GUID            uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	Shares          *big.Int       `json:"shares" gorm:"serializer:u256"`
	StrategyAddress common.Address `json:"strategy_address" gorm:"serializer:bytes"`
	Staker          common.Address `json:"staker" gorm:"serializer:bytes"`
	TokenAddress    common.Address `json:"token_address" gorm:"serializer:bytes"`
	ChainId         string         `json:"chain_id"`
	Created         uint64         `json:"created"`
}

func (AppChainDethShares) TableName() string {
	return "ac_chain_deth_shares"
}

type appChainDethSharesDB struct {
	gorm *gorm.DB
}

type AppChainDethSharesDB interface {
	AppChainDethSharesDBView
	StoreAppChainDethShares(dethShares AppChainDethShares) error
	UpdateAppChainDethShares(txHash common.Hash) error
}

type AppChainDethSharesDBView interface {
	ListDethSharesByAddress(address common.Address) []AppChainDethShares
}

func NewAppChainDethSharesDB(db *gorm.DB) AppChainDethSharesDB {
	return &appChainDethSharesDB{gorm: db}
}

func (db appChainDethSharesDB) StoreAppChainDethShares(dethShares AppChainDethShares) error {
	var exits AppChainDethShares
	err := db.gorm.Table(dethShares.TableName()).Where(AppChainDethShares{Staker: dethShares.Staker,
		StrategyAddress: dethShares.StrategyAddress,
		ChainId:         dethShares.ChainId,
	}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(dethShares.TableName()).Omit("guid").Create(&dethShares)
			return result.Error
		}
	}
	return err
}

func (db appChainDethSharesDB) UpdateAppChainDethShares(txHash common.Hash) error {
	//TODO implement me
	panic("implement me")
}

func (db appChainDethSharesDB) ListDethSharesByAddress(address common.Address) []AppChainDethShares {
	var appChainDethShares AppChainDethShares
	var appChainDethSharesList []AppChainDethShares
	result := db.gorm.Table(appChainDethShares.TableName()).Where(AppChainDethShares{
		Staker: address,
	}).Order("chain_id desc").Find(&appChainDethSharesList)
	if result.Error != nil {
		log.Error("ListDethSharesByAddress", "err", result.Error)
	}
	return appChainDethSharesList
}
