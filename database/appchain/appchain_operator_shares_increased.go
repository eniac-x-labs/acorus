package appchain

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
)

type AppChainOperatorSharesIncreased struct {
	GUID            uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	BlockNumber     *big.Int       `json:"block_number" gorm:"serializer:u256"`
	TxHash          common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	Shares          *big.Int       `json:"shares" gorm:"serializer:u256"`
	UseShares       *big.Int       `json:"use_shares" gorm:"serializer:u256"`
	StrategyAddress common.Address `json:"strategy_address" gorm:"serializer:bytes"`
	Operator        common.Address `json:"staker" gorm:"serializer:bytes"`
	Staker          common.Address `json:"token_address" gorm:"serializer:bytes"`
	Status          int            `json:"status"`
	ChainId         string         `json:"chain_id"`
	Created         uint64         `json:"created"`
}

func (AppChainOperatorSharesIncreased) TableName() string {
	return "ac_operator_shares_increased"
}

type appChainOperatorSharesIncreasedDB struct {
	gorm *gorm.DB
}

type AppChainOperatorSharesIncreasedDB interface {
	StoreAppChainOperatorSharesIncreased(acOperatorSharesIncreased AppChainOperatorSharesIncreased) error
	UpdateOperatorUseShares(acOperatorSharesIncreased AppChainOperatorSharesIncreased) error
	AppChainOperatorSharesIncreasedDBView
}
type AppChainOperatorSharesIncreasedDBView interface {
	GetNeedStakeShares(chainId string) []*AppChainOperatorSharesIncreased
}

func NewAppChainOperatorSharesIncreasedDB(db *gorm.DB) AppChainOperatorSharesIncreasedDB {
	return &appChainOperatorSharesIncreasedDB{gorm: db}
}

func (db appChainOperatorSharesIncreasedDB) StoreAppChainOperatorSharesIncreased(acOperatorSharesIncreased AppChainOperatorSharesIncreased) error {
	if acOperatorSharesIncreased.TxHash.String() == "" {
		return nil
	}
	var exits AppChainOperatorSharesIncreased
	err := db.gorm.Table(acOperatorSharesIncreased.TableName()).Where(AppChainOperatorSharesIncreased{TxHash: acOperatorSharesIncreased.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(acOperatorSharesIncreased.TableName()).Omit("guid").Create(&acOperatorSharesIncreased)
			return result.Error
		}
	}
	return err
}

func (db appChainOperatorSharesIncreasedDB) GetNeedStakeShares(chainId string) []*AppChainOperatorSharesIncreased {
	var resultList []*AppChainOperatorSharesIncreased
	err := db.gorm.Table(AppChainOperatorSharesIncreased{}.TableName()).
		Where("chain_id = ? and status = ?", chainId, 0).Order("created asc").
		Scan(&resultList).Error
	if err != nil {
		log.Error("GetNeedStakeShares", "err", err)
		return resultList
	}
	return resultList
}

func (db appChainOperatorSharesIncreasedDB) UpdateOperatorUseShares(acOperatorSharesIncreased AppChainOperatorSharesIncreased) error {
	var exits AppChainOperatorSharesIncreased
	err := db.gorm.Table(AppChainOperatorSharesIncreased{}.TableName()).Where(AppChainOperatorSharesIncreased{TxHash: acOperatorSharesIncreased.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("UpdateOperatorUseShares", "err", err)
			return err
		}
	}
	result := db.gorm.Table(AppChainOperatorSharesIncreased{}.TableName()).Where(AppChainOperatorSharesIncreased{GUID: acOperatorSharesIncreased.GUID}).Updates(acOperatorSharesIncreased)
	if result.Error != nil {
		log.Error("UpdateOperatorUseShares", "err", result.Error)
		return result.Error
	}
	return nil
}
