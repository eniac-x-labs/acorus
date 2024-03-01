package relayer

import (
	"errors"
	"fmt"
	common3 "github.com/cornerstone-labs/acorus/common"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BridgeFundingPoolUpdate struct {
	GUID           uuid.UUID `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	TxHash         string    `json:"tx_hash"`
	SourceChainId  string    `json:"source_chain_id"`
	DestChainId    string    `json:"dest_chain_id"`
	ReceiveAddress string    `json:"timestamp"`
	Amount         string    `json:"amount" `
	TokenAddress   string    `json:"token_address"`
	LayerType      int
	OnSend         bool
}

func (BridgeFundingPoolUpdate) TableName() string {
	return "bridge_funding_pool_update"
}

type BridgeFundingPoolUpdateDB interface {
	BridgeFundingPoolUpdateView
	StoreBridgeFundingPools(bridgeFundingPoolUpdate []BridgeFundingPoolUpdate) error
	StoreBridgeFundingPool(bridgeFundingPoolUpdate BridgeFundingPoolUpdate) error
	L1GetCanStoreTransactions(chainId, receiveAddress string) ([]worker.L1ToL2, error)
	L2GetCanStoreTransactions(chainId, receiveAddress string) ([]worker.L2ToL1, error)
	UpdateCrossStatus(txHash string) error
}
type BridgeFundingPoolUpdateView interface {
	BridgeFundingPoolNoSendList() []BridgeFundingPoolUpdate
}

type bridgeFundingPoolUpdateDB struct {
	gorm *gorm.DB
}

func NewBridgeFundingPoolUpdateDB(db *gorm.DB) BridgeFundingPoolUpdateDB {
	return &bridgeFundingPoolUpdateDB{gorm: db}
}

func (db bridgeFundingPoolUpdateDB) BridgeFundingPoolNoSendList() []BridgeFundingPoolUpdate {
	var bridgeFundingPoolUpdate []BridgeFundingPoolUpdate
	result := db.gorm.Where("on_send = ?", false).Find(&bridgeFundingPoolUpdate)
	if result.Error != nil {
		return nil
	}
	return bridgeFundingPoolUpdate
}

func (db bridgeFundingPoolUpdateDB) StoreBridgeFundingPools(bridgeFundingPoolUpdate []BridgeFundingPoolUpdate) error {
	var needSave []BridgeFundingPoolUpdate
	for i := 0; i < len(bridgeFundingPoolUpdate); i++ {
		bridgeFundingPoolInfo := bridgeFundingPoolUpdate[i]
		var existsBridge BridgeFundingPoolUpdate
		result := db.gorm.Table("bridge_funding_pool_update").Where(&BridgeFundingPoolUpdate{TxHash: bridgeFundingPoolInfo.TxHash}).Take(&existsBridge).Error
		if result != nil {
			if errors.Is(result, gorm.ErrRecordNotFound) {
				needSave = append(needSave, bridgeFundingPoolInfo)
				continue
			}
			return result
		}
	}
	if len(needSave) > 0 {
		result := db.gorm.Omit("guid").Table("bridge_funding_pool_update").Create(&needSave)
		return result.Error
	}
	return nil
}

func (db bridgeFundingPoolUpdateDB) StoreBridgeFundingPool(bridgeFundingPoolUpdate BridgeFundingPoolUpdate) error {
	result := db.gorm.Omit("guid").Table("bridge_funding_pool_update").Create(&bridgeFundingPoolUpdate)
	return result.Error
}

func (db bridgeFundingPoolUpdateDB) L1GetCanStoreTransactions(chainId, receiveAddress string) ([]worker.L1ToL2, error) {
	l1TableName := "l1_to_l2_" + chainId
	l1_get_can_store_sql := `
			SELECT * from ? a where not EXISTS (
			SELECT 1 from bridge_funding_pool_update b where a.l1_transaction_hash=b.tx_hash and b.receive_address = ? and a.layer_type=1
		) and a.status=?
	`
	l1_get_can_store_sql = fmt.Sprintf(l1_get_can_store_sql, l1TableName, receiveAddress, common3.L1ToL2Claimed, receiveAddress)
	var l1ToL2s []worker.L1ToL2
	result := db.gorm.Raw(l1_get_can_store_sql).Find(&l1ToL2s)
	return l1ToL2s, result.Error
}

func (db bridgeFundingPoolUpdateDB) L2GetCanStoreTransactions(chainId, receiveAddress string) ([]worker.L2ToL1, error) {
	l1TableName := "l2_to_l1_" + chainId
	l2_get_can_store_sql := `
			SELECT * from ? a where not EXISTS (
			SELECT 1 from bridge_funding_pool_update b where a.l2_transaction_hash=b.tx_hash and b.receive_address = ? and a.layer_type=2
		) and a.status=? and a.to_address=?
	`
	l2_get_can_store_sql = fmt.Sprintf(l2_get_can_store_sql, l1TableName, receiveAddress, common3.L2ToL1Claimed, receiveAddress)
	var l2ToL1s []worker.L2ToL1
	result := db.gorm.Raw(l2_get_can_store_sql).Find(&l2ToL1s)
	return l2ToL1s, result.Error
}
func (db bridgeFundingPoolUpdateDB) UpdateCrossStatus(txHash string) error {
	updateSql := `
		UPDATE bridge_funding_pool_update SET on_send=true WHERE tx_hash = ?
	`
	err := db.gorm.Exec(updateSql, txHash).Error
	return err
}
