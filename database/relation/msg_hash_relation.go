package relation

import (
	"gorm.io/gorm"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/common"
)

type MsgHashRelation struct {
	GUID    uuid.UUID   `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')"`
	TxHash  common.Hash `gorm:"serializer:bytes"`
	MsgHash common.Hash `gorm:"serializer:bytes"`
}

func (MsgHashRelation) TableName() string {
	return "msg_hash"
}

func (MsgHashRelation) TableNameByChainId(chainId string) string {
	return "msg_hash_" + chainId
}

type MsgHashRelationDB interface {
	MsgHashRelationView
	MsgHashRelationStore(msgHashRelation MsgHashRelation, chainId string) error
}

type MsgHashRelationView interface {
}

type msgHashRelationViewDB struct {
	gorm *gorm.DB
}

func NewMsgHashRelationViewDB(db *gorm.DB) MsgHashRelationDB {
	return &msgHashRelationViewDB{gorm: db}
}

func (m msgHashRelationViewDB) MsgHashRelationStore(msgHashRelation MsgHashRelation, chainId string) error {
	tableNameByChainId := msgHashRelation.TableNameByChainId(chainId)
	err := m.gorm.Table(tableNameByChainId).Omit("guid").Create(msgHashRelation).Error
	return err
}
