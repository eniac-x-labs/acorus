package relayer

import (
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
	"math/big"
)

type BridgeMsgSent struct {
	GUID            string         `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')"`
	TxHash          common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	MsgHash         common.Hash    `json:"msg_hash" gorm:"serializer:bytes"`
	DestHash        common.Hash    `json:"dest_hash" gorm:"serializer:bytes"`
	DestBlockNumber *big.Int       `json:"dest_block_number" gorm:"serializer:u256"`
	DestTimestamp   uint64         `json:"det_timestamp"`
	DestToken       common.Address `json:"dest_token" gorm:"serializer:bytes"`
	Fee             *big.Int       `json:"fee" gorm:"serializer:u256"`
	MsgNonce        *big.Int       `json:"msg_nonce" gorm:"serializer:u256"`
	MsgHashRelation bool           `json:"msg_hash_relation"`
	BridgeRelation  bool           `json:"bridge_relation"`
	ToBridgeRecord  bool           `json:"to_bridge_record"`
	Data            string         `json:"data"`
}

func (BridgeMsgSent) TableName() string {
	return "bridge_msg_sent"
}

type bridgeMsgSentDB struct {
	gorm *gorm.DB
}

type BridgeMsgSentDB interface {
	BridgeMsgSentDBView
	StoreBridgeMsgSent(msgSent BridgeMsgSent) error
	StoreBridgeMsgSents(msgSentList []BridgeMsgSent) error
}

type BridgeMsgSentDBView interface {
}

func NewBridgeMsgSentDB(db *gorm.DB) BridgeMsgSentDB {
	return &bridgeMsgSentDB{gorm: db}
}

func (db bridgeMsgSentDB) StoreBridgeMsgSent(msgSent BridgeMsgSent) error {
	msgSentRecord := new(BridgeMsgSent)
	result := db.gorm.Table(msgSentRecord.TableName()).Omit("guid").Create(&msgSent)
	return result.Error
}

func (db bridgeMsgSentDB) StoreBridgeMsgSents(msgSentList []BridgeMsgSent) error {
	msgSentRecord := new(BridgeMsgSent)
	result := db.gorm.Table(msgSentRecord.TableName()).Omit("guid").Create(&msgSentList)
	return result.Error
}
