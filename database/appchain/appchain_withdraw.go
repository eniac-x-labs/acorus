package appchain

import (
	"errors"
	common2 "github.com/cornerstone-labs/acorus/common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
)

type AppChainWithdraw struct {
	GUID        uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	BlockNumber *big.Int       `json:"block_number" gorm:"serializer:u256"`
	TxHash      common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	Shares      *big.Int       `json:"shares" gorm:"serializer:u256"`
	Operator    common.Address `json:"operator" gorm:"serializer:bytes"`
	Strategy    common.Address `json:"strategy" gorm:"serializer:bytes"`
	Staker      common.Address `json:"staker" gorm:"serializer:bytes"`
	ChainId     string         `json:"chain_id"`
	Created     uint64         `json:"created"`
}

func (AppChainWithdraw) TableName() string {
	return "ac_chain_withdraw"
}

type appChainWithdrawDB struct {
	gorm *gorm.DB
}

type AppChainWithdrawDB interface {
	StoreAppchainWithdraw(chainWithdrawBatch AppChainWithdraw) error
	AppChainWithdrawDBView
}

type AppChainWithdrawDBView interface {
	ListAppChainWithdraw(page, pageSize uint32, staker, strategy string) ([]AppChainWithdraw, int64, int64)
}

func NewAppChainWithdrawDB(db *gorm.DB) AppChainWithdrawDB {
	return &appChainWithdrawDB{gorm: db}
}

func (db appChainWithdrawDB) StoreAppchainWithdraw(chainWithdrawBatch AppChainWithdraw) error {
	if chainWithdrawBatch.TxHash.String() == "" {
		return nil
	}
	var exits AppChainWithdraw
	err := db.gorm.Table(chainWithdrawBatch.TableName()).Where(AppChainWithdraw{TxHash: chainWithdrawBatch.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(chainWithdrawBatch.TableName()).Omit("guid").Create(&chainWithdrawBatch)
			return result.Error
		}
	}
	return err
}

func (db appChainWithdrawDB) ListAppChainWithdraw(page, pageSize uint32, staker, strategy string) ([]AppChainWithdraw, int64, int64) {
	var appChainWithdraw []AppChainWithdraw
	var totalRecord int64
	offset := common2.CalculateOffset(uint(page), uint(pageSize))
	table := db.gorm.Table(AppChainWithdraw{}.TableName()).Select("*")
	querySR := db.gorm.Table("(?) as temp ", table)
	if staker != "" {
		querySR = querySR.Where(AppChainWithdraw{Staker: common.HexToAddress(staker)})
	}
	if strategy != "" {
		querySR = querySR.Where(AppChainWithdraw{Strategy: common.HexToAddress(strategy)})
	}
	resultTotal := querySR.Count(&totalRecord)
	if resultTotal.Error != nil {
		log.Error("appchain withdraw record", "get withdraw records by address count fail", resultTotal.Error)
	}
	querySR = querySR.Offset(int(offset)).Limit(int(pageSize))
	resultList := querySR.Find(&appChainWithdraw)
	if resultList.Error != nil {
		log.Error("appchain withdraw record", "get withdraw records by address fail", resultList.Error)
	}
	return appChainWithdraw, totalRecord, int64(page)
}
