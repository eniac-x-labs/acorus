package worker

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
	"math/big"
	"strings"

	"github.com/google/uuid"

	common3 "github.com/cornerstone-labs/acorus/common"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/ethereum/go-ethereum/common"
)

type L1ToL2 struct {
	GUID                  uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid" json:"guid"`
	L1BlockNumber         *big.Int       `gorm:"serializer:u256;column:l1_block_number" db:"l1_block_number" json:"l1BlockNumber" form:"l1_block_number"`
	L2BlockNumber         *big.Int       `gorm:"serializer:u256;column:l2_block_number" db:"l2_block_number" json:"l2BlockNumber" form:"l2_block_number"`
	QueueIndex            *big.Int       `gorm:"serializer:u256;column:queue_index" json:"queueIndex"`
	L1TransactionHash     common.Hash    `gorm:"column:l1_transaction_hash;serializer:bytes"  db:"l1_transaction_hash" json:"l1TransactionHash" form:"l1_transaction_hash"`
	L2TransactionHash     common.Hash    `gorm:"column:l2_transaction_hash;serializer:bytes" db:"l2_transaction_hash" json:"l2TransactionHash" form:"l2_transaction_hash"`
	TransactionSourceHash common.Hash    `gorm:"column:transaction_source_hash;serializer:bytes" db:"transaction_source_hash" json:"transactionSourceHash" form:"transaction_source_hash"`
	MessageHash           common.Hash    `gorm:"serializer:bytes"`
	L1TxOrigin            common.Address `gorm:"column:l1_tx_origin;serializer:bytes" db:"l1_tx_origin" json:"l1TxOrigin" form:"l1_tx_origin"`
	Status                int64          `gorm:"column:status" db:"status" json:"status" form:"status"`
	FromAddress           common.Address `gorm:"column:from_address;serializer:bytes" db:"from_address" json:"fromAddress" form:"from_address"`
	ToAddress             common.Address `gorm:"column:to_address;serializer:bytes" db:"to_address" json:"toAddress" form:"to_address"`
	L1TokenAddress        common.Address `gorm:"column:l1_token_address;serializer:bytes" db:"l1_token_address" json:"l1TokenAddress" form:"l1_token_address"`
	L2TokenAddress        common.Address `gorm:"column:l2_token_address;serializer:bytes" db:"l2_token_address" json:"l2TokenAddress" form:"l2_token_address"`
	ETHAmount             *big.Int       `gorm:"serializer:u256;column:eth_amount" json:"ETHAmount"`
	GasLimit              *big.Int       `gorm:"serializer:u256;column:gas_limit" json:"gasLimit"`
	Timestamp             int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
	TokenIds              string         `gorm:"column:token_ids" db:"token_ids" json:"tokenIds" form:"token_ids"`
	AssetType             int64          `gorm:"column:asset_type" db:"asset_type" json:"assetType" form:"asset_type"`
	TokenAmounts          string         `gorm:"column:token_amounts" db:"token_amounts" json:"tokenAmounts" form:"token_amounts"`
}

type L1ToL2DB interface {
	L1ToL2View
	StoreL1ToL2Transactions(string, []L1ToL2) error
	UpdateTokenPairAndAddress(chainId string, l1L2List []L1ToL2) error
	UpdateMessageHash(chainId string, l1L2List []L1ToL2) error
	MarkL1ToL2TransactionDepositFinalized(chainId string, l1l2List []L1ToL2) error
	RelayedL1ToL2Transaction(chainId string, l1L2List []L1ToL2) error
	FinalizedL1ToL2Transaction(chainId string, l1L2List []L1ToL2) error
	UpdateL1ToL2MsgHashByL1TxHash(chainId string, l1L2 L1ToL2) error
	UpdateL1ToL2L2TxHashByMsgHash(chainId string, l1L2 L1ToL2) error
}

type L1ToL2View interface {
	GetBlockNumberFromHash(chainId string, blockHash common.Hash) (*big.Int, error)
	L1LatestBlockHeader(l2chainId, destChainId string) (*common2.BlockHeader, error)
	L2LatestBlockHeader(chainId string) (*common2.BlockHeader, error)
	L1LatestFinalizedBlockHeader(chainId, destChainId string) (*common2.BlockHeader, error)
	L2LatestFinalizedBlockHeader(chainId string) (*common2.BlockHeader, error)
	L1ToL2List(string, string, int, int, string) ([]L1ToL2, int64)
	L1ToL2TransactionDeposit(string, common.Hash) (*L1ToL2, error)
	L1ToL2Transaction(string, common.Hash) (*L1ToL2, error)
}

type l1ToL2DB struct {
	gorm *gorm.DB
}

func NewL1ToL2DB(db *gorm.DB) L1ToL2DB {
	return &l1ToL2DB{gorm: db}
}

func (l1l2 l1ToL2DB) StoreL1ToL2Transactions(chainId string, l1L2List []L1ToL2) error {
	result := l1l2.gorm.Omit("guid").Table("l1_to_l2_" + chainId).Create(&l1L2List)
	return result.Error
}

func (l1l2 l1ToL2DB) UpdateL1ToL2MsgHashByL1TxHash(chainId string, l1L2Stu L1ToL2) error {
	result := l1l2.gorm.Table("l1_to_l2_" + chainId).Where(&L1ToL2{L1TransactionHash: l1L2Stu.L1TransactionHash})
	result = result.UpdateColumn("message_hash", l1L2Stu.MessageHash.String())
	return result.Error
}
func (l1l2 l1ToL2DB) UpdateL1ToL2L2TxHashByMsgHash(chainId string, l1L2Stu L1ToL2) error {
	result := l1l2.gorm.Table("l1_to_l2_" + chainId).Where(&L1ToL2{MessageHash: l1L2Stu.MessageHash})
	result = result.UpdateColumn("l2_transaction_hash", l1L2Stu.L2TransactionHash.String()).UpdateColumn("status", 1)
	result = result.UpdateColumn("l2_block_number", l1L2Stu.L2BlockNumber.String())
	return result.Error
}

func (l1l2 l1ToL2DB) L1ToL2TransactionDeposit(chainId string, messageHash common.Hash) (*L1ToL2, error) {
	var l1ToL2Withdrawal L1ToL2
	result := l1l2.gorm.Table("l1_to_l2_" + chainId).Where(&L2ToL1{MessageHash: messageHash}).Take(&l1ToL2Withdrawal)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l1ToL2Withdrawal, nil
}

func (l1l2 l1ToL2DB) L1ToL2List(chainId string, address string, page int, pageSize int, order string) (l1l2List []L1ToL2, total int64) {
	var totalRecord int64
	var l1ToL2List []L1ToL2
	queryStateRoot := l1l2.gorm.Table("l1_to_l2_" + chainId)
	if address != "0x00" {
		err := l1l2.gorm.Table("l1_to_l2_"+chainId).Select("l2_block_number").Where("from_address = ?", address).Or(" to_address = ?", address).Count(&totalRecord).Error
		if err != nil {
			log.Error("get l1 to l2 by address count fail")
		}
		queryStateRoot.Where("from_address = ?", address).Or(" to_address = ?", address).Offset((page - 1) * pageSize).Limit(pageSize)
	} else {
		err := l1l2.gorm.Table("l1_to_l2_" + chainId).Select("l2_block_number").Count(&totalRecord).Error
		if err != nil {
			log.Error("get l1 to l2 no address count fail ")
		}
		queryStateRoot.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	if strings.ToLower(order) == "asc" {
		queryStateRoot.Order("timestamp asc")
	} else {
		queryStateRoot.Order("timestamp desc")
	}
	qErr := queryStateRoot.Find(&l1ToL2List).Error
	if qErr != nil {
		log.Error("get l1 to l2 list fail", "err", qErr)
	}
	return l1ToL2List, totalRecord
}

func (l1l2 l1ToL2DB) UpdateTokenPairAndAddress(chainId string, l1L2List []L1ToL2) error {
	for i := 0; i < len(l1L2List); i++ {
		var l1ToL2 = L1ToL2{}
		result := l1l2.gorm.Table("l1_to_l2_" + chainId).Where(&L1ToL2{L1TransactionHash: l1L2List[i].L1TransactionHash}).Take(&l1ToL2)
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
		l1ToL2.TokenAmounts = l1L2List[i].TokenAmounts
		l1ToL2.ETHAmount = l1L2List[i].ETHAmount
		l1ToL2.TokenAmounts = l1L2List[i].TokenAmounts
		l1ToL2.AssetType = l1L2List[i].AssetType
		err := l1l2.gorm.Table("l1_to_l2_" + chainId).Save(l1ToL2).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l1l2 l1ToL2DB) UpdateMessageHash(chainId string, l1L2List []L1ToL2) error {
	for i := 0; i < len(l1L2List); i++ {
		var l1ToL2 = L1ToL2{}
		result := l1l2.gorm.Table("l1_to_l2_" + chainId).Where(&L1ToL2{L1TransactionHash: l1L2List[i].L1TransactionHash}).Take(&l1ToL2)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l1ToL2.MessageHash = l1L2List[i].MessageHash
		err := l1l2.gorm.Table("l1_to_l2_" + chainId).Save(l1ToL2).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l1l2 l1ToL2DB) MarkL1ToL2TransactionDepositFinalized(chainId string, L1l2List []L1ToL2) error {
	for i := 0; i < len(L1l2List); i++ {
		var l1ToL2 = L1ToL2{}
		if L1l2List[i].L2BlockNumber.Uint64() <= 0 {
			continue
		}
		result := l1l2.gorm.Table("l1_to_l2_" + chainId).Where(&L2ToL1{MessageHash: L1l2List[i].MessageHash}).Take(&l1ToL2)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		log.Info("mark transaction finalized", "L2BlockNumber", L1l2List[i].L2BlockNumber, "L2TransactionHash", L1l2List[i].L2TransactionHash)
		l1ToL2.L2BlockNumber = L1l2List[i].L2BlockNumber
		l1ToL2.L2TransactionHash = L1l2List[i].L2TransactionHash
		l1ToL2.Status = common3.L1ToL2Claimed // relayed
		err := l1l2.gorm.Table("l1_to_l2_" + chainId).Save(&l1ToL2).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l1l2 l1ToL2DB) RelayedL1ToL2Transaction(chainId string, l1L2List []L1ToL2) error {
	for i := 0; i < len(l1L2List); i++ {
		var l1ToL2 = L1ToL2{}
		result := l1l2.gorm.Table("l1_to_l2_" + chainId).Where(&L1ToL2{MessageHash: l1L2List[i].MessageHash}).Take(&l1ToL2)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l1ToL2.L2TransactionHash = l1L2List[i].L2TransactionHash
		err := l1l2.gorm.Table("l1_to_l2_" + chainId).Save(l1ToL2).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l1l2 l1ToL2DB) FinalizedL1ToL2Transaction(chainId string, l1L2List []L1ToL2) error {
	for i := 0; i < len(l1L2List); i++ {
		var l1ToL2 = L1ToL2{}
		result := l1l2.gorm.Table("l1_to_l2_" + chainId).Where(&L1ToL2{MessageHash: l1L2List[i].MessageHash}).Take(&l1ToL2)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l1ToL2.L2TransactionHash = l1L2List[i].L2TransactionHash
		l1ToL2.L2BlockNumber = l1L2List[i].L2BlockNumber
		l1ToL2.Status = l1L2List[i].Status
		err := l1l2.gorm.Table("l1_to_l2_" + chainId).Save(l1ToL2).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l1l2 l1ToL2DB) GetBlockNumberFromHash(chainId string, blockHash common.Hash) (*big.Int, error) {
	var l1BlockNumber uint64
	result := l1l2.gorm.Table("block_headers_"+chainId).Where("hash = ?", blockHash.String()).Select("number").Take(&l1BlockNumber)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return new(big.Int).SetUint64(l1BlockNumber), nil
}

func (l1l2 l1ToL2DB) L1LatestBlockHeader(l2chainId, destChainId string) (*common2.BlockHeader, error) {
	tableName := fmt.Sprintf("l1_to_l2_%s", l2chainId)
	blockHeaderSTableName := fmt.Sprintf("block_headers_%s", destChainId)
	l1Query := l1l2.gorm.Table(blockHeaderSTableName).Where("number = (?)", l1l2.gorm.Table(tableName).Select("MAX(l1_block_number)"))
	var l1Header common2.BlockHeader
	result := l1Query.Take(&l1Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l1Header, nil
}

func (l1l2 l1ToL2DB) L2LatestBlockHeader(chainId string) (*common2.BlockHeader, error) {
	var l2Header common2.BlockHeader
	blockHeaderSTableName := fmt.Sprintf("block_headers_%s", chainId)
	result := l1l2.gorm.Table(blockHeaderSTableName).Order("number desc").Limit(1).Take(&l2Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l2Header, nil
}

func (l1l2 l1ToL2DB) latestBlockHeaderWithChainId(chainId, destChainId string) (*common2.BlockHeader, error) {
	tableName := fmt.Sprintf("l1_to_l2_%s", chainId)
	l1Query := l1l2.gorm.Where("timestamp = (?)", l1l2.gorm.Table(tableName).Select("MAX(timestamp)"))
	var l1Header common2.BlockHeader
	blockHeaderSTableName := fmt.Sprintf("block_headers_%s", destChainId)
	result := l1Query.Table(blockHeaderSTableName).Take(&l1Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l1Header, nil
}

func (l1l2 l1ToL2DB) L1LatestFinalizedBlockHeader(chainId, destChainId string) (*common2.BlockHeader, error) {
	return l1l2.latestFinalizedBlockHeaderWithChainId(chainId, destChainId)
}

func (l1l2 l1ToL2DB) L2LatestFinalizedBlockHeader(chainId string) (*common2.BlockHeader, error) {
	return l1l2.latestFinalizedBlockHeaderWithChainId(chainId, chainId)
}

func (l1l2 l1ToL2DB) latestFinalizedBlockHeaderWithChainId(l2ChainId, destChainId string) (*common2.BlockHeader, error) {
	tableName := fmt.Sprintf("l1_to_l2_%s", l2ChainId)
	l1Query := l1l2.gorm.Where("number = (?)", l1l2.gorm.Table(tableName).Where("status != (?)", common3.L1ToL2Claimed).Select("MAX(l2_block_number)"))
	blockHeaderSTableName := fmt.Sprintf("block_headers_%s", destChainId)
	var l2Header common2.BlockHeader
	result := l1Query.Table(blockHeaderSTableName).Take(&l2Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l2Header, nil
}

func (l1l2 l1ToL2DB) L1ToL2Transaction(chainId string, msgHash common.Hash) (*L1ToL2, error) {
	var l1tol2Tx L1ToL2
	filterMessageHash := L1ToL2{MessageHash: msgHash}
	result := l1l2.gorm.Table("l1_to_l2_" + chainId).Where(&filterMessageHash).Take(&l1tol2Tx)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l1tol2Tx, nil
}
