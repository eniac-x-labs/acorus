package relation

import (
	"errors"
	"gorm.io/gorm"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/common"
)

type MsgHashRelation struct {
	GUID    uuid.UUID   `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
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

type msgHashRelationDB struct {
	gorm *gorm.DB
}

func NewMsgHashRelationDB(db *gorm.DB) MsgHashRelationDB {
	return &msgHashRelationDB{gorm: db}
}

func (m msgHashRelationDB) MsgHashRelationStore(msgHashRelation MsgHashRelation, chainId string) error {
	tableNameByChainId := msgHashRelation.TableNameByChainId(chainId)
	var exist MsgHashRelation
	err := m.gorm.Table(tableNameByChainId).Where("tx_hash = ?", msgHashRelation.TxHash.String()).Take(&exist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := m.gorm.Table(tableNameByChainId).Omit("guid").Create(msgHashRelation)
			return result.Error
		}
	}
	return err
}
