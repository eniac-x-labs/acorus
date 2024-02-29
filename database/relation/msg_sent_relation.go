package relation

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MsgSentRelationStruct struct {
	GUID            uuid.UUID   `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	TxHash          common.Hash `gorm:"serializer:bytes"`
	MsgHash         common.Hash `gorm:"serializer:bytes"`
	MsgHashRelation bool
	ToL1L2Table     bool
	LayerType       int
	Data            string
}

func (MsgSentRelationStruct) TableName() string {
	return "msg_sent"
}

func (MsgSentRelationStruct) TableNameByChainId(chainId string) string {
	return "msg_sent_" + chainId
}

type MsgSentRelationDB interface {
	MsgSentRelationView
	RelationClear(chainId string) error
	MsgHashToRelation(chainId string) error
	L1RelayToRelation(chainId string) error
	L2RelayToRelation(chainId string) error
	MsgSentRelationStore(msgSentRelation MsgSentRelationStruct, chainId string) error
	L1MsgSentRelationClear(chainId string) error
	L2MsgSentRelationClear(chainId string) error
}

type MsgSentRelationView interface {
	GetCanSaveDataList(chainId string) ([]MsgSentRelationStruct, error)
}

type msgSentRelationStructDB struct {
	gorm *gorm.DB
}

func NewMsgSentRelationStructDB(db *gorm.DB) MsgSentRelationDB {
	return &msgSentRelationStructDB{gorm: db}
}

func (m msgSentRelationStructDB) MsgHashToRelation(chainId string) error {

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

func (m msgSentRelationStructDB) L1RelayToRelation(chainId string) error {
	l1toL2Table := fmt.Sprintf("l1_to_l2_%s", chainId)
	relayTable := fmt.Sprintf("relay_%s", chainId)

	updateSql := `
				update %s a  set l2_transaction_hash=b.tx_hash, 
				                 l2_block_number=b.block_number,
				                 status=1
				from %s  b 
				where  a.message_hash=b.msg_hash and b.l1_l2_relation=false
				`
	updateSql = fmt.Sprintf(updateSql, l1toL2Table, relayTable)

	err := m.gorm.Exec(updateSql).Error
	return err
}

func (m msgSentRelationStructDB) L2RelayToRelation(chainId string) error {
	l1toL2Table := fmt.Sprintf("l2_to_l1_%s", chainId)
	relayTable := fmt.Sprintf("relay_%s", chainId)

	updateSql := `
				update %s a  set l1_finalize_tx_hash=b.tx_hash, 
				                 l1_block_number=b.block_number,
				                 status=1
				from %s  b 
				where  a.message_hash=b.msg_hash and b.l1_l2_relation=false
				`
	updateSql = fmt.Sprintf(updateSql, l1toL2Table, relayTable)

	err := m.gorm.Exec(updateSql).Error
	return err
}

func (m msgSentRelationStructDB) GetCanSaveDataList(chainId string) ([]MsgSentRelationStruct, error) {
	tableNameByChainId := new(MsgSentRelationStruct).TableNameByChainId(chainId)
	var msgSentRels []MsgSentRelationStruct
	selectSql := `
		SELECT DISTINCT ON ("a"."tx_hash") * FROM %s "a" WHERE "a"."msg_hash_relation" = true
			AND "a"."to_l1_l2_table" = false LIMIT 1000
	`
	selectSql = fmt.Sprintf(selectSql, tableNameByChainId)
	result := m.gorm.Raw(selectSql).Find(&msgSentRels)
	if result.Error != nil {
		return nil, result.Error
	}
	return msgSentRels, nil
}

func (m msgSentRelationStructDB) MsgSentRelationStore(msgSentRelation MsgSentRelationStruct, chainId string) error {
	tableNameByChainId := msgSentRelation.TableNameByChainId(chainId)
	var exist RelayRelation
	err := m.gorm.Table(tableNameByChainId).Where("tx_hash = ?", msgSentRelation.TxHash.String()).Take(&exist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := m.gorm.Table(tableNameByChainId).Omit("guid").Create(msgSentRelation)
			return result.Error
		}
	}
	return err
}

func (m msgSentRelationStructDB) RelationClear(chainId string) error {
	if err := m.L1MsgSentRelationClear(chainId); err != nil {
		return err
	}
	if err := m.L2MsgSentRelationClear(chainId); err != nil {
		return err
	}
	if err := m.L1RelayRelationClear(chainId); err != nil {
		return err
	}
	if err := m.L2RelayRelationClear(chainId); err != nil {
		return err
	}
	return nil
}

func (m msgSentRelationStructDB) L1RelayRelationClear(chainId string) error {
	msgSentTable := fmt.Sprintf("relay_%s", chainId)
	l1toL2Table := fmt.Sprintf("l1_to_l2_%s", chainId)

	updateSql := `
				update %s a  set l1_l2_relation=true
				from %s  b 
				where  a.tx_hash=b.l1_transaction_hash  and a.l1_l2_relation=false
				`
	updateSql = fmt.Sprintf(updateSql, msgSentTable, l1toL2Table)

	err := m.gorm.Exec(updateSql).Error
	return err
}

func (m msgSentRelationStructDB) L2RelayRelationClear(chainId string) error {
	msgSentTable := fmt.Sprintf("relay_%s", chainId)
	l1toL2Table := fmt.Sprintf("l2_to_l1_%s", chainId)

	updateSql := `
				update %s a  set l1_l2_relation=true
				from %s  b 
				where  a.tx_hash=b.l1_finalize_tx_hash  and a.l1_l2_relation=false
				`
	updateSql = fmt.Sprintf(updateSql, msgSentTable, l1toL2Table)

	err := m.gorm.Exec(updateSql).Error
	return err
}

func (m msgSentRelationStructDB) L1MsgSentRelationClear(chainId string) error {
	msgSentTable := fmt.Sprintf("msg_sent_%s", chainId)
	l1toL2Table := fmt.Sprintf("l1_to_l2_%s", chainId)

	updateSql := `
				update %s a  set to_l1_l2_table=true
				from %s  b 
				where  a.tx_hash=b.l1_transaction_hash  and a.to_l1_l2_table=false 
				`
	updateSql = fmt.Sprintf(updateSql, msgSentTable, l1toL2Table)

	err := m.gorm.Exec(updateSql).Error
	return err
}

func (m msgSentRelationStructDB) L2MsgSentRelationClear(chainId string) error {
	msgSentTable := fmt.Sprintf("msg_sent_%s", chainId)
	l2toL1Table := fmt.Sprintf("l2_to_l1_%s", chainId)

	updateSql := `
				update %s a  set to_l1_l2_table=true
				from %s  b 
				where  a.tx_hash=b.l2_transaction_hash  and a.to_l1_l2_table=false
				`
	updateSql = fmt.Sprintf(updateSql, msgSentTable, l2toL1Table)

	err := m.gorm.Exec(updateSql).Error
	return err
}
