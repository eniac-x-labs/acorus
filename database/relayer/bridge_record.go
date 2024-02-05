package relayer

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
)

type BridgeRecord struct {
	GUID             uuid.UUID      `json:"guid" gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')"`
	SourceChainId    string         `json:"source_chain_id"`
	TargetChainId    string         `json:"target_chain_id"`
	L1TxHash         common.Hash    `json:"l1_tx_hash" gorm:"serializer:bytes"`
	L2TxHash         common.Hash    `json:"l2_tx_hash" gorm:"serializer:bytes"`
	L1BlockNumber    *big.Int       `json:"l1_block_number" gorm:"serializer:u256"`
	L2BlockNumber    *big.Int       `json:"l2_block_number" gorm:"serializer:u256"`
	L1TokenAddress   common.Address `json:"l1_token_address" gorm:"serializer:bytes"`
	L2TokenAddress   common.Address `json:"l2_token_address" gorm:"serializer:bytes"`
	MsgHash          common.Hash    `json:"msg_hash" gorm:"serializer:bytes"`
	From             common.Address `json:"from" gorm:"serializer:bytes"`
	To               common.Address `json:"to" gorm:"serializer:bytes"`
	Status           int            `json:"status"`
	Amount           *big.Int       `json:"amount" gorm:"serializer:u256"`
	Nonce            *big.Int       `json:"nonce" gorm:"serializer:u256"`
	Fee              *big.Int       `json:"fee" gorm:"serializer:u256"`
	TxType           int            `json:"tx_type"`
	AssetType        int            `json:"asset_type"`
	MsgSentTimestamp uint64         `json:"msg_sent_timestamp"`
	ClaimTimestamp   uint64         `json:"claim_timestamp"`
}

func (BridgeRecord) TableName() string {
	return "bridge_recode"
}

type bridgeRecordDB struct {
	gorm *gorm.DB
}

type BridgeRecordDB interface {
	BridgeRecordDBView
}

type BridgeRecordDBView interface {
}

func NewBridgeRecordDB(db *gorm.DB) BridgeRecordDB {
	return &bridgeRecordDB{gorm: db}
}
