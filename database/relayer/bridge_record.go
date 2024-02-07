package relayer

import (
	"math/big"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type BridgeRecords struct {
	GUID               uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')"`
	SourceChainId      string         `json:"source_chain_id"`
	DestChainId        string         `json:"dest_chain_id"`
	SourceTxHash       common.Hash    `json:"source_tx_hash" gorm:"serializer:bytes"`
	DestTxHash         common.Hash    `json:"dest_tx_hash" gorm:"serializer:bytes"`
	SourceBlockNumber  *big.Int       `json:"source_block_number" gorm:"serializer:u256"`
	DestBlockNumber    *big.Int       `json:"dest_block_number" gorm:"serializer:u256"`
	SourceTokenAddress common.Address `json:"source_token_address" gorm:"serializer:bytes"`
	DestTokenAddress   common.Address `json:"dest_token_address" gorm:"serializer:bytes"`
	MsgHash            common.Hash    `json:"msg_hash" gorm:"serializer:bytes"`
	From               common.Address `json:"from" gorm:"serializer:bytes"`
	To                 common.Address `json:"to" gorm:"serializer:bytes"`
	Status             int            `json:"status"`
	Amount             *big.Int       `json:"amount" gorm:"serializer:u256"`
	Nonce              *big.Int       `json:"nonce" gorm:"serializer:u256"`
	Fee                *big.Int       `json:"fee" gorm:"serializer:u256"`
	OperaType          int            `json:"opera_type"`
	AssetType          int            `json:"asset_type"`
	MsgSentTimestamp   uint64         `json:"msg_sent_timestamp"`
	ClaimTimestamp     uint64         `json:"claim_timestamp"`
}

func (BridgeRecords) TableName() string {
	return "bridge_record"
}

type bridgeRecordsDB struct {
	gorm *gorm.DB
}

type BridgeRecordDB interface {
	BridgeRecordDBView
	StoreBridgeRecord(bridgeRecord BridgeRecords) error
	StoreBridgeRecords(bridgeRecord []BridgeRecords) error
}

type BridgeRecordDBView interface {
	GetBridgeRecordList(address string, page int, pageSize int, order string) (bridgeRecords []BridgeRecords, total int64)
}

func NewBridgeRecordDB(db *gorm.DB) BridgeRecordDB {
	return &bridgeRecordsDB{gorm: db}
}

func (db bridgeRecordsDB) GetBridgeRecordList(address string, page int, pageSize int, order string) (bR []BridgeRecords, total int64) {
	var totalRecord int64
	var bridgeRecords []BridgeRecords
	queryBR := db.gorm.Table("bridge_record")
	if address != "0x00" {
		err := db.gorm.Table("bridge_record").Select("guid").Where("from = ?", address).Or("to = ?", address).Count(&totalRecord).Error
		if err != nil {
			log.Error("get bridge records by address count fail")
		}
		queryBR.Where("from = ?", address).Or(" to = ?", address).Offset((page - 1) * pageSize).Limit(pageSize)
	} else {
		err := db.gorm.Table("bridge_record").Select("guid").Count(&totalRecord).Error
		if err != nil {
			log.Error("get bridge records no address count fail ")
		}
		queryBR.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	if strings.ToLower(order) == "asc" {
		queryBR.Order("msg_sent_timestamp asc")
	} else {
		queryBR.Order("msg_sent_timestamp desc")
	}
	qErr := queryBR.Find(&bridgeRecords).Error
	if qErr != nil {
		log.Error("get bridge records fail", "err", qErr)
	}
	return bridgeRecords, totalRecord
}

func (db bridgeRecordsDB) StoreBridgeRecord(bridge BridgeRecords) error {
	bridgeRecords := new(BridgeRecords)
	result := db.gorm.Table(bridgeRecords.TableName()).Omit("guid").Create(&bridge)
	return result.Error
}

func (db bridgeRecordsDB) StoreBridgeRecords(brs []BridgeRecords) error {
	bridgeRecords := new(BridgeRecords)
	result := db.gorm.Table(bridgeRecords.TableName()).Omit("guid").Create(&brs)
	return result.Error
}
