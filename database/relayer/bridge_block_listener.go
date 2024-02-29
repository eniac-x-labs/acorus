package relayer

import (
	"errors"
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
	"math/big"
	"time"
)

type BridgeBlockListener struct {
	GUID        string   `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	ChainId     string   `json:"chain_id"`
	BlockNumber *big.Int `json:"block_number" gorm:"serializer:u256"`
	Created     uint64   `json:"created"`
	Updated     uint64   `json:"updated"`
}

func (BridgeBlockListener) TableName() string {
	return "bridge_block_listener"
}

type bridgeBlockListenerDB struct {
	gorm *gorm.DB
}

type BridgeBlockListenerDB interface {
	BridgeBlockListenerDBView
	SaveOrUpdateLastBlockNumber(lastBlock BridgeBlockListener) error
}

type BridgeBlockListenerDBView interface {
	GetLastBlockNumber(chainId string) (lastBlock *BridgeBlockListener, err error)
}

func NewBlockListenerDB(db *gorm.DB) BridgeBlockListenerDB {
	return &bridgeBlockListenerDB{gorm: db}
}

func (db bridgeBlockListenerDB) SaveOrUpdateLastBlockNumber(lastBlock BridgeBlockListener) error {
	bbl := new(BridgeBlockListener)
	var exitsLastBlock BridgeBlockListener

	err := db.gorm.Table(bbl.TableName()).Where("chain_id = ?", lastBlock.ChainId).Take(&exitsLastBlock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info("No last block number found, will insert it")
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

func (db bridgeBlockListenerDB) GetLastBlockNumber(chainId string) (lastBlock *BridgeBlockListener, err error) {
	bbl := new(BridgeBlockListener)
	err = db.gorm.Table(bbl.TableName()).Where("chain_id = ?", chainId).Take(&lastBlock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return lastBlock, nil
}
