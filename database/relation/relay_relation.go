package relation

import (
	"errors"
	"gorm.io/gorm"
	"math/big"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/common"
)

type RelayRelation struct {
	GUID         uuid.UUID   `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid"`
	TxHash       common.Hash `gorm:"serializer:bytes"`
	MsgHash      common.Hash `gorm:"serializer:bytes"`
	BlockNumber  *big.Int    `gorm:"serializer:u256"`
	L1L2Relation bool
	LayerType    int
}

func (RelayRelation) TableName() string {
	return "relay"
}

func (RelayRelation) TableNameByChainId(chainId string) string {
	return "relay_" + chainId
}

type RelayRelationDB interface {
	RelayRelationView
	RelayRelationStore(relayRelation RelayRelation, chainId string) error
}

type RelayRelationView interface {
}

type relayRelationDB struct {
	gorm *gorm.DB
}

func NewEvmRelayRelationDB(db *gorm.DB) RelayRelationDB {
	return &relayRelationDB{gorm: db}
}

func (m relayRelationDB) RelayRelationStore(relayRelation RelayRelation, chainId string) error {
	tableNameByChainId := relayRelation.TableNameByChainId(chainId)
	var exist RelayRelation
	err := m.gorm.Table(tableNameByChainId).Where("tx_hash = ?", relayRelation.TxHash.String()).Take(&exist).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := m.gorm.Table(tableNameByChainId).Omit("guid").Create(relayRelation)
			return result.Error
		}
	}
	return err
}
