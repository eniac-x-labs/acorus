package worker

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"

	"github.com/acmestack/gorm-plus/gplus"

	"github.com/ethereum/go-ethereum/common"
)

type L1ToL2 struct {
	GUID              uuid.UUID      `gorm:"primaryKey"`
	L1BlockHash       common.Hash    `gorm:"column:l1_block_hash;serializer:bytes" db:"l1_block_hash" json:"l1_block_hash" form:"l1_block_hash"`
	QueueIndex        *big.Int       `gorm:"serializer:u256;column:queue_index" json:"queue_index"`
	L1TransactionHash common.Hash    `gorm:"column:l1_transaction_hash;serializer:bytes"  db:"l1_transaction_hash" json:"l1_transaction_hash" form:"l1_transaction_hash"`
	L1DepositHash     common.Hash    `gorm:"column:l1_deposit_hash;serializer:bytes"  db:"l1_deposit_hash" json:"l1_deposit_hash" form:"l1_deposit_hash"`
	L2TransactionHash common.Hash    `gorm:"column:l2_transaction_hash;serializer:bytes" db:"l2_transaction_hash" json:"l2_transaction_hash" form:"l2_transaction_hash"`
	L1TxOrigin        common.Hash    `gorm:"column:l1_tx_origin;serializer:bytes" db:"l1_tx_origin" json:"l1_tx_origin" form:"l1_tx_origin"`
	Status            int64          `gorm:"column:status" db:"status" json:"status" form:"status"`
	FromAddress       common.Address `gorm:"column:from_address;serializer:bytes" db:"from_address" json:"from_address" form:"from_address"`
	ToAddress         common.Address `gorm:"column:to_address;serializer:bytes" db:"to_address" json:"to_address" form:"to_address"`
	L1TokenAddress    common.Address `gorm:"column:l1_token_address;serializer:bytes" db:"l1_token_address" json:"l1_token_address" form:"l1_token_address"`
	L2TokenAddress    common.Address `gorm:"column:l2_token_address;serializer:bytes" db:"l2_token_address" json:"l2_token_address" form:"l2_token_address"`
	ETHAmount         *big.Int       `gorm:"serializer:u256;column:eth_amount" json:"eth_amount"`
	ERC20Amount       *big.Int       `gorm:"serializer:u256;column:erc20_amount" json:"erc20_amount"`
	GasLimit          *big.Int       `gorm:"serializer:u256;column:gas_limit" json:"gas_limit"`
	Timestamp         int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
	AssetType         int64          `gorm:"column:asset_type" db:"asset_type" json:"asset_type" form:"asset_type"`
	TokenIds          string         `gorm:"column:token_ids" db:"token_ids" json:"token_ids" form:"token_ids"`
	TokenAmounts      string         `gorm:"column:token_amounts" db:"token_amounts" json:"token_amounts" form:"token_amounts"`
	MsgHash           common.Hash    `gorm:"column:msg_hash;serializer:bytes" db:"msg_hash" json:"msg_hash" form:"msg_hash"`
	ChainId           int64          `gorm:"column:chain_id" db:"chain_id" json:"chain_id" form:"chain_id"`
}

func (L1ToL2) TableName() string {
	return "l1_to_l2"
}

type L1ToL2DB interface {
	L1ToL2View
	StoreL1ToL2Transactions([]*L1ToL2) error
	UpdateTokenPairs(l1L2List []L1ToL2) error
	UpdateL1ToL2MsgHashByL1TxHash(l1L2 L1ToL2) error
	UpdateL1ToL2L2TxHashByMsgHash(l1L2 L1ToL2) error
}

type L1ToL2View interface {
	L1ToL2List(string, int, int) (*gplus.Page[L1ToL2], error)
}

/**
 * Implementation
 */

type l1ToL2DB struct {
	gorm *gorm.DB
}

func NewL1ToL2DB(db *gorm.DB) L1ToL2DB {
	gplus.Init(db)
	return &l1ToL2DB{gorm: db}
}

func (l1l2 l1ToL2DB) StoreL1ToL2Transactions(l1L2List []*L1ToL2) error {
	result := l1l2.gorm.CreateInBatches(l1L2List, 3000)
	return result.Error
}

func (l1l2 l1ToL2DB) UpdateL1ToL2MsgHashByL1TxHash(l1L2Stu L1ToL2) error {
	result := l1l2.gorm.Model(&l1L2Stu).Where(&L1ToL2{L1TransactionHash: l1L2Stu.L1TransactionHash})
	result = result.UpdateColumn("msg_hash", l1L2Stu.MsgHash.String())
	return result.Error
}
func (l1l2 l1ToL2DB) UpdateL1ToL2L2TxHashByMsgHash(l1L2Stu L1ToL2) error {
	result := l1l2.gorm.Model(&l1L2Stu).Where(&L1ToL2{MsgHash: l1L2Stu.MsgHash})
	result = result.UpdateColumn("l2_transaction_hash", l1L2Stu.L2TransactionHash.String()).UpdateColumn("status", 1)
	return result.Error
}

func (l1l2 l1ToL2DB) L1ToL2List(address string, page int, pageSize int) (*gplus.Page[L1ToL2], error) {
	query, l2l1s := gplus.NewQuery[L1ToL2]()
	query.Eq(&l2l1s.FromAddress, address)
	dbPage := gplus.NewPage[L1ToL2](page, pageSize)
	rst, rstDB := gplus.SelectPage(dbPage, query)
	if rstDB.Error != nil {
		return nil, rstDB.Error
	}
	return rst, nil
}

func (l1l2 l1ToL2DB) UpdateTokenPairs(l1L2List []L1ToL2) error {
	for i := 0; i < len(l1L2List); i++ {
		var l1ToL2 = L1ToL2{}
		result := l1l2.gorm.Where(&L1ToL2{L1TransactionHash: l1L2List[i].L1TransactionHash}).Take(&l1ToL2)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l1ToL2.L1TokenAddress = l1L2List[i].L1TokenAddress
		l1ToL2.L2TokenAddress = l1L2List[i].L2TokenAddress

		err := l1l2.gorm.Save(l1ToL2).Error
		if err != nil {
			return err
		}
	}

	return nil
}
