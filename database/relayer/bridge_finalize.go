package relayer

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BridgeFinalize struct {
	GUID      uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	TxHash    common.Hash    `json:"tx_hash" gorm:"serializer:bytes"`
	DestToken common.Address `json:"dest_token" gorm:"serializer:bytes"`
}

func (BridgeFinalize) TableName() string {
	return "bridge_finalize"
}

type bridgeFinalizeDB struct {
	gorm *gorm.DB
}

type BridgeFinalizeDB interface {
	BridgeFinalizeView
	StoreBridgeFinalize(finalize BridgeFinalize) error
	StoreBridgeFinalizes(finalizes []BridgeFinalize) error
	RelationClaim() error
}

type BridgeFinalizeView interface {
}

func NewBridgeFinalizeDB(db *gorm.DB) BridgeFinalizeDB {
	return &bridgeFinalizeDB{gorm: db}
}

func (db bridgeFinalizeDB) StoreBridgeFinalize(finalize BridgeFinalize) error {
	bridgeFinalize := new(BridgeFinalize)
	var exist BridgeFinalize
	err := db.gorm.Table(bridgeFinalize.TableName()).Where("tx_hash = ?", finalize.TxHash.String()).Take(&exist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(bridgeFinalize.TableName()).Omit("guid").Create(&finalize)
			return result.Error
		}
	}
	return err
}

func (db bridgeFinalizeDB) StoreBridgeFinalizes(finalizes []BridgeFinalize) error {
	bridgeFinalize := new(BridgeFinalize)
	result := db.gorm.Table(bridgeFinalize.TableName()).Omit("guid").Create(&finalizes)
	return result.Error
}

func (db bridgeFinalizeDB) RelationClaim() error {
	relationSql := `
		update  bridge_claim a  set dest_token=b.dest_token, token_relation=true 
				from bridge_finalize  b 
				where  a.tx_hash=b.tx_hash  and a.token_relation=false 
	`
	err := db.gorm.Exec(relationSql).Error
	return err
}
