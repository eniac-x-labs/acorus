package relation

import (
	"fmt"
	"gorm.io/gorm"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type MsgSentRelation struct {
	GUID             uuid.UUID   `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')"`
	TxHash           common.Hash `gorm:"serializer:bytes"`
	MsgHash          common.Hash `gorm:"serializer:bytes"`
	LayerHash        common.Hash `gorm:"serializer:bytes"`
	LayerBlockNumber *big.Int    `gorm:"serializer:u256"`
	MsgHashRelation  bool
	L1L2Relation     bool
	ToL1L2Table      bool
	LayerType        int
	Data             string
}

func (MsgSentRelation) TableName() string {
	return "msg_sent"
}

func (MsgSentRelation) TableNameByChainId(chainId string) string {
	return "msg_sent_" + chainId
}

type MsgSentRelationDB interface {
	MsgSentRelationView
	MsgHashRelation(chainId string) error
	RelayRelation(chainId string) error
	MsgSentRelationStore(msgSentRelation MsgSentRelation, chainId string) error
}

type MsgSentRelationView interface {
	GetCanSaveDataList(chainId string) ([]MsgSentRelation, error)
}

type msgSentRelationViewDB struct {
	gorm *gorm.DB
}

func NewMsgSentRelationViewDB(db *gorm.DB) MsgSentRelationDB {
	return &msgSentRelationViewDB{gorm: db}
}

func (m msgSentRelationViewDB) MsgHashRelation(chainId string) error {

	msgSentTable := fmt.Sprintf("msg_sent_%s", chainId)
	msgHashTable := fmt.Sprintf("msg_hash_%s", chainId)

	updateSql := `
				update %s a  set msg_hash=b.msg_hash, msg_hash_relation=true 
				from %s  b 
				where  a.tx_hash=b.tx_hash  and a.msg_hash_relation=false
				`
	updateSql = fmt.Sprintf(updateSql, msgSentTable, msgHashTable)

	err := m.gorm.Exec(updateSql).Error
	return err
}

func (m msgSentRelationViewDB) RelayRelation(chainId string) error {
	msgSentTable := fmt.Sprintf("msg_sent_%s", chainId)
	relayTable := fmt.Sprintf("relay_%s", chainId)

	updateSql := `
				update %s a  set layer_hash=b.tx_hash, 
				                 layer_block_number=b.block_number,
				                 l1_l2_relation=true
				from %s  b 
				where  a.msg_hash=b.msg_hash  and a.l1_l2_relation=false
				`
	updateSql = fmt.Sprintf(updateSql, msgSentTable, relayTable)

	err := m.gorm.Exec(updateSql).Error
	return err
}

func (m msgSentRelationViewDB) GetCanSaveDataList(chainId string) ([]MsgSentRelation, error) {
	tableNameByChainId := new(MsgSentRelation).TableNameByChainId(chainId)
	var msgSentRels []MsgSentRelation
	result := m.gorm.Table(tableNameByChainId).Where(&MsgSentRelation{MsgHashRelation: true,
		L1L2Relation: true, ToL1L2Table: false}).Limit(1000).Find(&msgSentRels)
	if result.Error != nil {
		return nil, result.Error
	}
	return msgSentRels, nil
}

func (m msgSentRelationViewDB) MsgSentRelationStore(msgSentRelation MsgSentRelation, chainId string) error {
	tableNameByChainId := msgSentRelation.TableNameByChainId(chainId)
	err := m.gorm.Omit("guid").Table(tableNameByChainId).Create(msgSentRelation).Error
	return err
}
