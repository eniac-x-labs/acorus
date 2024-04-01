package appchain

import (
	"errors"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
	"time"
)

type AppChainLastBlock struct {
	GUID        uuid.UUID `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	ChainId     string    `json:"chain_id"`
	BlockNumber *big.Int  `json:"block_number" gorm:"serializer:u256"`
	Created     uint64    `json:"created"`
	Updated     uint64    `json:"updated"`
}

func (AppChainLastBlock) TableName() string {
	return "ac_last_block"
}

type appChainLastBlockDB struct {
	gorm *gorm.DB
}

type AppChainLastBlockDB interface {
	AppChainLastBlockDBView
	SaveOrUpdateLastBlockNumber(lastBlock AppChainLastBlock) error
}

type AppChainLastBlockDBView interface {
	GetLastBlockNumber(chainId string) (lastBlock *AppChainLastBlock, err error)
}

func NewAppChainLastBlockDB(db *gorm.DB) AppChainLastBlockDB {
	return &appChainLastBlockDB{gorm: db}
}

func (db appChainLastBlockDB) GetLastBlockNumber(chainId string) (lastBlock *AppChainLastBlock, err error) {
	acl := new(AppChainLastBlock)
	err = db.gorm.Table(acl.TableName()).Where("chain_id = ?", chainId).Take(&lastBlock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return lastBlock, nil
}

func (db appChainLastBlockDB) SaveOrUpdateLastBlockNumber(lastBlock AppChainLastBlock) error {
	bbl := new(AppChainLastBlock)
	var exitsLastBlock AppChainLastBlock

	err := db.gorm.Table(bbl.TableName()).Where("chain_id = ?", lastBlock.ChainId).Take(&exitsLastBlock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info("appchain listener No last block number found, will insert it")
			lastBlock.Created = uint64(time.Now().Unix())
			lastBlock.Updated = uint64(time.Now().Unix())
			result := db.gorm.Table(bbl.TableName()).Omit("guid").Create(&lastBlock)
			if result.Error != nil {
				return result.Error
			}
			return nil
		}
		return err
	} else {
		lastBlock.Updated = uint64(time.Now().Unix())
		updateResult := db.gorm.Table(bbl.TableName()).Omit("guid", "created").Where("chain_id = ?", lastBlock.ChainId).Updates(&lastBlock)
		if updateResult.Error != nil {
			return updateResult.Error
		}
	}
	return nil
}
