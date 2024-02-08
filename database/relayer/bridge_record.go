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
	GUID                 uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')"`
	SourceChainId        string         `json:"source_chain_id"`
	DestChainId          string         `json:"dest_chain_id"`
	SourceTxHash         common.Hash    `json:"source_tx_hash" gorm:"serializer:bytes"`
	DestTxHash           common.Hash    `json:"dest_tx_hash" gorm:"serializer:bytes"`
	SourceBlockNumber    *big.Int       `json:"source_block_number" gorm:"serializer:u256"`
	DestBlockNumber      *big.Int       `json:"dest_block_number" gorm:"serializer:u256"`
	SourceTokenAddress   common.Address `json:"source_token_address" gorm:"serializer:bytes"`
	DestTokenAddress     common.Address `json:"dest_token_address" gorm:"serializer:bytes"`
	MsgHash              common.Hash    `json:"msg_hash" gorm:"serializer:bytes"`
	FromAddress          common.Address `json:"fromAddress" gorm:"serializer:bytes"`
	ToAddress            common.Address `json:"toAddress" gorm:"serializer:bytes"`
	Status               int            `json:"status"`
	Amount               *big.Int       `json:"amount" gorm:"serializer:u256"`
	Nonce                *big.Int       `json:"nonce" gorm:"serializer:u256"`
	Fee                  *big.Int       `json:"fee" gorm:"serializer:u256"`
	BridgeRecordRelation bool           `json:"bridge_record_relation""`
	AssetType            int            `json:"asset_type"`
	MsgSentTimestamp     uint64         `json:"msg_sent_timestamp"`
	ClaimTimestamp       uint64         `json:"claim_timestamp"`
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
	RelationClaimData() error
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
	table := db.gorm.Table("bridge_record").Select("DISTINCT ON (source_tx_hash) *")
	queryBR := db.gorm.Table("(?) as temp ", table)
	if address != "0x00" {
		err := db.gorm.Table("bridge_record").Select("DISTINCT ON (source_tx_hash) guid").Where("from_address = ?", address).Or("to_address = ?", address).Count(&totalRecord).Error
		if err != nil {
			log.Error("get bridge records by address count fail")
		}
		queryBR.Where("from_address = ?", address).Or(" to_address = ?", address).Offset((page - 1) * pageSize).Limit(pageSize)
	} else {
		err := db.gorm.Table("bridge_record").Select("DISTINCT ON (source_tx_hash) guid").Count(&totalRecord).Error
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

func (db bridgeRecordsDB) RelationClaimData() error {
	relationSql := `
		update  bridge_record a  set dest_tx_hash=b.dest_hash,dest_token_address=b.dest_token,
		                               dest_block_number=b.dest_block_number,claim_timestamp=b.dest_timestamp,
		                               status=1
		                               bridge_record_relation=true 
				from bridge_msg_sent  b 
				where  a.msg_hash=b.msg_hash  and b.bridge_relation=true and a.bridge_record_relation=false and a.status=0
	`
	err := db.gorm.Exec(relationSql).Error
	return err
}
