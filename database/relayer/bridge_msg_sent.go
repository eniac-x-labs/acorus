package relayer

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
)

type BridgeMsgSent struct {
	GUID                uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	TxHash              common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	MsgHash             common.Hash    `json:"msg_hash" gorm:"serializer:bytes"`
	DestHash            common.Hash    `json:"dest_hash" gorm:"serializer:bytes"`
	DestBlockNumber     *big.Int       `json:"dest_block_number" gorm:"serializer:u256"`
	DestTimestamp       uint64         `json:"det_timestamp"`
	DestToken           common.Address `json:"dest_token" gorm:"serializer:bytes"`
	Fee                 *big.Int       `json:"fee" gorm:"serializer:u256"`
	MsgNonce            *big.Int       `json:"msg_nonce" gorm:"serializer:u256"`
	MsgHashRelation     bool           `json:"msg_hash_relation"`
	BridgeRelation      bool           `json:"bridge_relation"`
	ToChangeTransStatus bool           `json:"to_change_trans_status"`
	ToCrossStatus       bool           `json:"to_cross_status"`
	ToBridgeRecord      bool           `json:"to_bridge_record"`
	Data                string         `json:"data"`
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
	CleanMsgSent() error
	UpdateChangeStatus(txHash string) error
	UpdateCrossStatus(txHash string) error
}

type BridgeMsgSentDBView interface {
	GetCanSaveDecodeList() (mList []BridgeMsgSent, err error)
	GetCanCrossDataList() []BridgeMsgSent
	GetCanChangeTransStatusList() []BridgeMsgSent
}

func NewBridgeMsgSentDB(db *gorm.DB) BridgeMsgSentDB {
	return &bridgeMsgSentDB{gorm: db}
}

func (db bridgeMsgSentDB) StoreBridgeMsgSent(msgSent BridgeMsgSent) error {
	msgSentRecord := new(BridgeMsgSent)
	var exist BridgeMsgSent
	err := db.gorm.Table(msgSentRecord.TableName()).Where("tx_hash = ?", msgSent.TxHash.String()).Take(&exist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(msgSentRecord.TableName()).Omit("guid").Create(&msgSent)
			return result.Error
		}
	}
	return err
}

func (db bridgeMsgSentDB) StoreBridgeMsgSents(msgSentList []BridgeMsgSent) error {
	msgSentRecord := new(BridgeMsgSent)
	result := db.gorm.Table(msgSentRecord.TableName()).Omit("guid").Create(&msgSentList)
	return result.Error
}

func (db bridgeMsgSentDB) GetCanSaveDecodeList() (mList []BridgeMsgSent, err error) {
	var msgSentList []BridgeMsgSent
	selectSql := `
		SELECT DISTINCT ON ("a"."tx_hash") * FROM bridge_msg_sent "a" WHERE "a"."msg_hash_relation" = true
			AND "a"."bridge_relation" = false AND "a"."to_bridge_record" = false  LIMIT 1000
	`
	result := db.gorm.Raw(selectSql).Find(&msgSentList)
	if result.Error != nil {
		return nil, result.Error
	}
	return msgSentList, nil
}

func (db bridgeMsgSentDB) CleanMsgSent() error {
	cleanSql := `
		update "bridge_msg_sent" a  set to_bridge_record=true
				from  bridge_record b 
				where  a.msg_hash=b.msg_hash  and a.to_bridge_record=false 
	`
	err := db.gorm.Exec(cleanSql).Error
	return err
}

func (db bridgeMsgSentDB) GetCanCrossDataList() []BridgeMsgSent {
	getSql := `
		select DISTINCT ON ("a"."tx_hash") * from bridge_msg_sent a where a.bridge_relation=false
			and a.to_bridge_record = true and a.to_cross_status = false and a.msg_hash_relation = true  LIMIT 20
	`
	var bridgeMsgSent []BridgeMsgSent
	result := db.gorm.Raw(getSql).Find(&bridgeMsgSent)
	err := result.Error
	if err != nil {
		return nil
	}
	return bridgeMsgSent
}

func (db bridgeMsgSentDB) GetCanChangeTransStatusList() []BridgeMsgSent {
	getSql := `
		select DISTINCT ON ("a"."tx_hash") * from bridge_msg_sent a where a.bridge_relation=true
			and a.to_change_trans_status = false and a.msg_hash_relation = true  LIMIT 20
	`
	var bridgeMsgSent []BridgeMsgSent
	result := db.gorm.Raw(getSql).Find(&bridgeMsgSent)
	err := result.Error
	if err != nil {
		return nil
	}
	return bridgeMsgSent
}

func (db bridgeMsgSentDB) UpdateChangeStatus(txHash string) error {
	msgSentRecord := new(BridgeMsgSent)
	result := db.gorm.Table(msgSentRecord.TableName()).Where("tx_hash = ?", txHash).Updates(map[string]interface{}{"to_change_trans_status": true})
	return result.Error
}
func (db bridgeMsgSentDB) UpdateCrossStatus(txHash string) error {
	msgSentRecord := new(BridgeMsgSent)
	result := db.gorm.Table(msgSentRecord.TableName()).Where("tx_hash = ?", txHash).Updates(map[string]interface{}{"to_cross_status": true})
	return result.Error
}
