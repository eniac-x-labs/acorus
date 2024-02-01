package relation

import (
	"gorm.io/gorm"
	"math/big"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/common"
)

type RelayRelation struct {
	GUID        uuid.UUID   `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')"`
	TxHash      common.Hash `gorm:"serializer:bytes"`
	MsgHash     common.Hash `gorm:"serializer:bytes"`
	BlockNumber *big.Int    `gorm:"serializer:u256"`
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

type relayRelationViewDB struct {
	gorm *gorm.DB
}

func NewEvmRelayRelationViewDB(db *gorm.DB) RelayRelationDB {
	return &relayRelationViewDB{gorm: db}
}

func (m relayRelationViewDB) RelayRelationStore(relayRelation RelayRelation, chainId string) error {
	tableNameByChainId := relayRelation.TableNameByChainId(chainId)
	err := m.gorm.Table(tableNameByChainId).Omit("guid").Create(relayRelation).Error
	return err
}
