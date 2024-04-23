package appchain

import (
	"errors"
	"gorm.io/gorm"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"

	common2 "github.com/cornerstone-labs/acorus/common"
)

type AppChainStake struct {
	GUID            uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	BlockNumber     *big.Int       `json:"block_number" gorm:"serializer:u256"`
	TxHash          common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	Shares          *big.Int       `json:"shares" gorm:"serializer:u256"`
	StrategyAddress common.Address `json:"strategy_address" gorm:"serializer:bytes"`
	Staker          common.Address `json:"staker" gorm:"serializer:bytes"`
	TokenAddress    common.Address `json:"token_address" gorm:"serializer:bytes"`
	ChainId         string         `json:"chain_id"`
	Created         uint64         `json:"created"`
}

func (AppChainStake) TableName() string {
	return "ac_chain_stake"
}

type appChainStakeDB struct {
	gorm *gorm.DB
}

type AppChainStakeDB interface {
	AppChainStakeDBView
	StoreAppChainStake(cainStakeBatch AppChainStake) error
}

type AppChainStakeDBView interface {
	ListAppChainStake(page, pageSize uint32, staker, strategy string) ([]AppChainStake, int64, int64)
}

func NewAppChainStakeDB(db *gorm.DB) AppChainStakeDB {
	return &appChainStakeDB{gorm: db}
}

func (db appChainStakeDB) StoreAppChainStake(chainUnStakeBatch AppChainStake) error {
	if chainUnStakeBatch.TxHash.String() == "" {
		return nil
	}
	var exits AppChainStake
	err := db.gorm.Table(chainUnStakeBatch.TableName()).Where(AppChainStake{TxHash: chainUnStakeBatch.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(chainUnStakeBatch.TableName()).Omit("guid").Create(&chainUnStakeBatch)
			return result.Error
		}
	}
	return err
}

func (db appChainStakeDB) ListAppChainStake(page, pageSize uint32, staker, strategy string) ([]AppChainStake, int64, int64) {
	var chainStakeBatch []AppChainStake
	var totalRecord int64
	offset := common2.CalculateOffset(uint(page), uint(pageSize))
	table := db.gorm.Table(AppChainStake{}.TableName()).Select("*")
	querySR := db.gorm.Table("(?) as temp ", table)
	if staker != "" {
		querySR = querySR.Where(AppChainStake{Staker: common.HexToAddress(staker)})
	}
	if strategy != "" {
		querySR = querySR.Where(AppChainStake{StrategyAddress: common.HexToAddress(strategy)})
	}
	resultTotal := querySR.Count(&totalRecord)
	if resultTotal.Error != nil {
		log.Error("appchain stake record", "get staking records by address count fail", resultTotal.Error)
	}
	querySR = querySR.Order("created desc").Offset(int(offset)).Limit(int(pageSize))
	resultList := querySR.Find(&chainStakeBatch)
	if resultList.Error != nil {
		log.Error("appchain stake record", "get staking records by address fail", resultList.Error)
	}
	return chainStakeBatch, totalRecord, int64(page)
}
