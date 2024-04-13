package appchain

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
	"math/big"
)

type AppChainDappLinkBridge struct {
	GUID                  string         `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')"`
	TxHash                common.Hash    `json:"tx_hash"gorm:"serializer:bytes"`
	ChainId               string         `json:"chain_id"`
	SourceChainId         string         `json:"source_chain_id"`
	DestChainId           string         `json:"dest_chain_id"`
	Bridge                common.Address `json:"bridge"gorm:"serializer:bytes"`
	StakingManagerAddress common.Address `json:"staking_manager_address"gorm:"serializer:bytes"`
	TokenAddress          common.Address `json:"token_address"gorm:"serializer:bytes"`
	Amount                *big.Int       `json:"amount"gorm:"serializer:u256"`
	BatchId               *big.Int       `json:"batch_id"gorm:"serializer:u256"`
	Nonce                 *big.Int       `json:"nonce"gorm:"serializer:u256"`
	NotifyRelayer         bool           `json:"notify_relayer"`
	Created               uint64         `json:"created"`
}

func (AppChainDappLinkBridge) TableName() string {
	return "ac_chain_dapplink_bridge"
}

type AppChainDappLinkBridgeDB interface {
	StroeAppChainDappLinkBridge(appChainDappLinkBridge AppChainDappLinkBridge) error
	NotifyAppChainDappLinkBridge(batchId *big.Int) error
	AppChainDappLinkBridgeDBView
}

type AppChainDappLinkBridgeDBView interface {
	ListDataByNoNotifyRelayer(chainId string) []AppChainDappLinkBridge
}

type appChainDappLinkBridgeDB struct {
	gorm *gorm.DB
}

func NewAppChainDappLinkBridgeDB(db *gorm.DB) AppChainDappLinkBridgeDB {
	return &appChainDappLinkBridgeDB{gorm: db}
}

func (db appChainDappLinkBridgeDB) StroeAppChainDappLinkBridge(appChainDappLinkBridge AppChainDappLinkBridge) error {
	if appChainDappLinkBridge.TxHash.String() == "" {
		return nil
	}
	var exits AppChainDappLinkBridge
	err := db.gorm.Table(appChainDappLinkBridge.TableName()).Where(AppChainDappLinkBridge{TxHash: appChainDappLinkBridge.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(appChainDappLinkBridge.TableName()).Omit("guid").Create(&appChainDappLinkBridge)
			return result.Error
		}
	}
	return nil
}

func (db appChainDappLinkBridgeDB) NotifyAppChainDappLinkBridge(batchId *big.Int) error {
	return db.gorm.Table(AppChainDappLinkBridge{}.TableName()).Where(AppChainDappLinkBridge{BatchId: batchId}).Update("notify_relayer", true).Error
}

func (db appChainDappLinkBridgeDB) ListDataByNoNotifyRelayer(chainId string) []AppChainDappLinkBridge {
	var appChainDappLinkBridge []AppChainDappLinkBridge
	err := db.gorm.Table(AppChainDappLinkBridge{}.TableName()).Where(AppChainDappLinkBridge{ChainId: chainId}).Where("notify_relayer = ?", false).Find(&appChainDappLinkBridge)
	if err.Error != nil {
		log.Error("ListDataByNoNotifyRelayer", "err", err.Error)
	}
	return appChainDappLinkBridge
}
