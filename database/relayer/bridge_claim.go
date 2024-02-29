package relayer

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
)

type BridgeClaimed struct {
	GUID          uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	TxHash        common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	MsgHash       common.Hash    `json:"msg_hash" gorm:"serializer:bytes"`
	Timestamp     uint64         `json:"timestamp"`
	DestToken     common.Address `json:"dest_token" gorm:"serializer:bytes"`
	BlockNumber   *big.Int       `json:"block_number" gorm:"serializer:u256"`
	TokenRelation bool
}

func (BridgeClaimed) TableName() string {
	return "bridge_claim"
}

type bridgeClaimDB struct {
	gorm *gorm.DB
}

type BridgeClaimDB interface {
	BridgeClaimView
	StoreBridgeClaim(claim BridgeClaimed) error
	StoreBridgeClaims(claims []BridgeClaimed) error
	RelationMsgSent() error
}

type BridgeClaimView interface {
}

func NewBridgeClaimDB(db *gorm.DB) BridgeClaimDB {
	return &bridgeClaimDB{gorm: db}
}

func (db bridgeClaimDB) StoreBridgeClaim(claim BridgeClaimed) error {
	bridgeClaimed := new(BridgeClaimed)
	var exist BridgeClaimed
	err := db.gorm.Table(bridgeClaimed.TableName()).Where("tx_hash = ?", claim.TxHash.String()).Take(&exist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(bridgeClaimed.TableName()).Omit("guid").Create(&claim)
			return result.Error
		}
	}
	return err
}

func (db bridgeClaimDB) StoreBridgeClaims(claims []BridgeClaimed) error {
	bridgeClaimed := new(BridgeClaimed)
	result := db.gorm.Table(bridgeClaimed.TableName()).Omit("guid").Create(&claims)
	return result.Error
}

func (db bridgeClaimDB) RelationMsgSent() error {
	relationSql := `
		update  bridge_msg_sent a  set dest_hash=b.tx_hash,dest_token=b.dest_token,
		                               dest_block_number=b.block_number,dest_timestamp=b.timestamp,
		                               bridge_relation=true 
				from bridge_claim  b 
				where  a.msg_hash=b.msg_hash  and a.bridge_relation=false and b.token_relation=true
	`
	err := db.gorm.Exec(relationSql).Error
	return err
}
