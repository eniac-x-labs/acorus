package relayer

import (
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
	"math/big"
)

type BridgeMsgHash struct {
	GUID     string      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')"`
	TxHash   common.Hash `json:"tx_hash" gorm:"serializer:bytes"`
	Fee      *big.Int    `json:"fee" gorm:"serializer:u256"`
	MsgNonce *big.Int    `json:"msg_nonce" gorm:"serializer:u256"`
	MsgHash  common.Hash `json:"msg_hash" gorm:"serializer:bytes"`
}

func (BridgeMsgHash) TableName() string {
	return "bridge_msg_hash"
}

type bridgeMsgHashDB struct {
	gorm *gorm.DB
}

type BridgeMsgHashDB interface {
	BridgeMsgHashView
	StoreBridgeMsgHash(msgHash BridgeMsgHash) error
	StoreBridgeMsgHashs(msgHashList []BridgeMsgHash) error
	RelationMsgHash() error
}

type BridgeMsgHashView interface {
}

func NewBridgeMsgHashDB(db *gorm.DB) BridgeMsgHashDB {
	return &bridgeMsgHashDB{gorm: db}
}

func (db bridgeMsgHashDB) StoreBridgeMsgHash(msgHash BridgeMsgHash) error {
	msgHashRecord := new(BridgeMsgHash)
	result := db.gorm.Table(msgHashRecord.TableName()).Omit("guid").Create(&msgHash)
	return result.Error
}

func (db bridgeMsgHashDB) StoreBridgeMsgHashs(msgHashList []BridgeMsgHash) error {
	msgHashRecord := new(BridgeMsgHash)
	result := db.gorm.Table(msgHashRecord.TableName()).Omit("guid").Create(&msgHashList)
	return result.Error
}

func (db bridgeMsgHashDB) RelationMsgHash() error {
	relationSql := `
		update  bridge_msg_sent a  set msg_hash=b.msg_hash,fee=b.fee,
		                               msg_nonce=b.msg_nonce, msg_hash_relation=true 
				from bridge_msg_hash  b 
				where  a.tx_hash=b.tx_hash  and a.msg_hash_relation=false 
	`
	err := db.gorm.Exec(relationSql).Error
	return err
}
