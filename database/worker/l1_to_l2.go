package worker

import (
	"errors"
	"gorm.io/gorm"
	"math/big"
	"strings"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/google/uuid"

	common3 "github.com/cornerstone-labs/acorus/common"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/ethereum/go-ethereum/common"
)

type L1ToL2 struct {
	GUID                  uuid.UUID      `gorm:"primaryKey" json:"guid"`
	L1BlockNumber         *big.Int       `gorm:"serializer:u256;column:l1_block_number" db:"l1_block_number" json:"l1BlockNumber" form:"l1_block_number"`
	L2BlockNumber         *big.Int       `gorm:"serializer:u256;column:l2_block_number" db:"l2_block_number" json:"l2BlockNumber" form:"l2_block_number"`
	QueueIndex            *big.Int       `gorm:"serializer:u256;column:queue_index" json:"queueIndex"`
	L1TransactionHash     common.Hash    `gorm:"column:l1_transaction_hash;serializer:bytes"  db:"l1_transaction_hash" json:"l1TransactionHash" form:"l1_transaction_hash"`
	L2TransactionHash     common.Hash    `gorm:"column:l2_transaction_hash;serializer:bytes" db:"l2_transaction_hash" json:"l2TransactionHash" form:"l2_transaction_hash"`
	TransactionSourceHash common.Hash    `gorm:"column:transaction_source_hash;serializer:bytes" db:"transaction_source_hash" json:"transactionSourceHash" form:"transaction_source_hash"`
	MessageHash           common.Hash    `gorm:"serializer:bytes"`
	L1TxOrigin            common.Hash    `gorm:"column:l1_tx_origin;serializer:bytes" db:"l1_tx_origin" json:"l1TxOrigin" form:"l1_tx_origin"`
	Status                int64          `gorm:"column:status" db:"status" json:"status" form:"status"`
	FromAddress           common.Address `gorm:"column:from_address;serializer:bytes" db:"from_address" json:"fromAddress" form:"from_address"`
	ToAddress             common.Address `gorm:"column:to_address;serializer:bytes" db:"to_address" json:"toAddress" form:"to_address"`
	L1TokenAddress        common.Address `gorm:"column:l1_token_address;serializer:bytes" db:"l1_token_address" json:"l1TokenAddress" form:"l1_token_address"`
	L2TokenAddress        common.Address `gorm:"column:l2_token_address;serializer:bytes" db:"l2_token_address" json:"l2TokenAddress" form:"l2_token_address"`
	ETHAmount             *big.Int       `gorm:"serializer:u256;column:eth_amount" json:"ETHAmount"`
	ERC20Amount           *big.Int       `gorm:"serializer:u256;column:erc20_amount" json:"ERC20Amount"`
	GasLimit              *big.Int       `gorm:"serializer:u256;column:gas_limit" json:"gasLimit"`
	Timestamp             int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
	TokenIds              string         `gorm:"column:token_ids" db:"token_ids" json:"tokenIds" form:"token_ids"`
	AssetType             int64          `gorm:"column:asset_type" db:"asset_type" json:"assetType" form:"asset_type"`
	TokenAmounts          string         `gorm:"column:token_amounts" db:"token_amounts" json:"tokenAmounts" form:"token_amounts"`
}

func (L1ToL2) TableName() string {
	return "l1_to_l2"
}

type L1ToL2DB interface {
	L1ToL2View
	StoreL1ToL2Transactions([]L1ToL2) error
	UpdateTokenPairAndAddress(l1L2List []L1ToL2) error
	UpdateMessageHash(l1L2List []L1ToL2) error
	RelayedL1ToL2Transaction(l1L2List []L1ToL2) error
	FinalizedL1ToL2Transaction(l1L2List []L1ToL2) error
}

type L1ToL2View interface {
	GetBlockNumberFromHash(blockHash common.Hash) (*big.Int, error)
	L1LatestBlockHeader() (*common2.L1BlockHeader, error)
	L2LatestFinalizedBlockHeader() (*common2.L2BlockHeader, error)
	L1ToL2List(string, int, int, string) (*gplus.Page[L1ToL2], error)
	L1ToL2Transaction(common.Hash) (*L1ToL2, error)
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

func (l1l2 l1ToL2DB) StoreL1ToL2Transactions(l1L2List []L1ToL2) error {
	result := l1l2.gorm.CreateInBatches(&l1L2List, len(l1L2List))
	return result.Error
}

func (l1l2 l1ToL2DB) L1ToL2List(address string, page int, pageSize int, order string) (*gplus.Page[L1ToL2], error) {
	query, l2l1s := gplus.NewQuery[L1ToL2]()
	if address != "0x00" {
		query.Eq(&l2l1s.FromAddress, address).Or().Eq(&l2l1s.ToAddress, address)
	}
	if strings.ToLower(order) == "asc" {
		query.OrderByAsc("timestamp")
	} else {
		query.OrderByDesc("timestamp")
	}
	dbPage := gplus.NewPage[L1ToL2](page, pageSize)
	rst, rstDB := gplus.SelectPage(dbPage, query)
	if rstDB.Error != nil {
		return nil, rstDB.Error
	}
	return rst, nil
}

func (l1l2 l1ToL2DB) UpdateTokenPairAndAddress(l1L2List []L1ToL2) error {
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
		l1ToL2.FromAddress = l1L2List[i].FromAddress
		l1ToL2.ToAddress = l1L2List[i].ToAddress
		l1ToL2.ERC20Amount = l1L2List[i].ERC20Amount
		err := l1l2.gorm.Save(l1ToL2).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l1l2 l1ToL2DB) UpdateMessageHash(l1L2List []L1ToL2) error {
	for i := 0; i < len(l1L2List); i++ {
		var l1ToL2 = L1ToL2{}
		result := l1l2.gorm.Where(&L1ToL2{L1TransactionHash: l1L2List[i].L1TransactionHash}).Take(&l1ToL2)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l1ToL2.MessageHash = l1L2List[i].MessageHash
		err := l1l2.gorm.Save(l1ToL2).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l1l2 l1ToL2DB) RelayedL1ToL2Transaction(l1L2List []L1ToL2) error {
	for i := 0; i < len(l1L2List); i++ {
		var l1ToL2 = L1ToL2{}
		result := l1l2.gorm.Where(&L1ToL2{MessageHash: l1L2List[i].MessageHash}).Take(&l1ToL2)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l1ToL2.L2TransactionHash = l1L2List[i].L2TransactionHash
		err := l1l2.gorm.Save(l1ToL2).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l1l2 l1ToL2DB) FinalizedL1ToL2Transaction(l1L2List []L1ToL2) error {
	for i := 0; i < len(l1L2List); i++ {
		var l1ToL2 = L1ToL2{}
		result := l1l2.gorm.Where(&L1ToL2{MessageHash: l1L2List[i].MessageHash}).Take(&l1ToL2)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l1ToL2.L2TransactionHash = l1L2List[i].L2TransactionHash
		l1ToL2.L2BlockNumber = l1L2List[i].L2BlockNumber
		l1ToL2.Status = l1L2List[i].Status
		err := l1l2.gorm.Save(l1ToL2).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l1l2 l1ToL2DB) GetBlockNumberFromHash(blockHash common.Hash) (*big.Int, error) {
	var l1BlockNumber uint64
	result := l1l2.gorm.Table("l1_block_headers").Where("hash = ?", blockHash.String()).Select("number").Take(&l1BlockNumber)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return new(big.Int).SetUint64(l1BlockNumber), nil
}

func (l1l2 *l1ToL2DB) L1LatestBlockHeader() (*common2.L1BlockHeader, error) {
	l1Query := l1l2.gorm.Where("timestamp = (?)", l1l2.gorm.Table("l1_to_l2").Select("MAX(timestamp)"))
	var l1Header common2.L1BlockHeader
	result := l1Query.Take(&l1Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l1Header, nil
}

func (l1l2 *l1ToL2DB) L2LatestFinalizedBlockHeader() (*common2.L2BlockHeader, error) {
	l1Query := l1l2.gorm.Where("number = (?)", l1l2.gorm.Table("l1_to_l2").Where("status != (?)", common3.L1ToL2Claimed).Select("MAX(l2_block_number)"))
	var l2Header common2.L2BlockHeader
	result := l1Query.Take(&l2Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l2Header, nil
}

func (l1l2 *l1ToL2DB) L1ToL2Transaction(msgHash common.Hash) (*L1ToL2, error) {
	var l1tol2Tx L1ToL2
	filterMessageHash := L1ToL2{MessageHash: msgHash}
	result := l1l2.gorm.Where(&filterMessageHash).Take(&l1tol2Tx)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l1tol2Tx, nil
}
