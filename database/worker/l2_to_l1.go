package worker

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math/big"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/acmestack/gorm-plus/gplus"
	common3 "github.com/cornerstone-labs/acorus/common"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/ethereum/go-ethereum/common"
)

type L2ToL1 struct {
	GUID                    uuid.UUID      `gorm:"primaryKey" json:"guid"`
	L1BlockNumber           *big.Int       `gorm:"serializer:u256;column:l1_block_number" db:"l1_block_number" json:"l1BlockNumber" form:"l1_block_number"`
	L2BlockNumber           *big.Int       `gorm:"serializer:u256;column:l2_block_number;primaryKey" db:"l2_block_number" json:"l2BlockNumber" form:"l2_block_number"`
	MsgNonce                *big.Int       `gorm:"column:msg_nonce;serializer:u256" db:"msg_nonce" json:"msgNonce" form:"msg_nonce"`
	L2TransactionHash       common.Hash    `gorm:"column:l2_transaction_hash;serializer:bytes" db:"l2_transaction_hash" json:"l2TransactionHash" form:"l2_transaction_hash"`
	WithdrawTransactionHash common.Hash    `gorm:"column:withdraw_transaction_hash;serializer:bytes" db:"withdraw_transaction_hash" json:"l2WithdrawTransactionHash" form:"withdraw_transaction_hash"`
	L1ProveTxHash           common.Hash    `gorm:"column:l1_prove_tx_hash;serializer:bytes" db:"l1_prove_tx_hash" json:"l1ProveTxHash" form:"l1_prove_tx_hash"`
	L1FinalizeTxHash        common.Hash    `gorm:"column:l1_finalize_tx_hash;serializer:bytes" db:"l1_finalize_tx_hash" json:"l1FinalizeTxHash" form:"l1_finalize_tx_hash"`
	MessageHash             common.Hash    `gorm:"column:message_hash;serializer:bytes" db:"message_hash" json:"messageHash" form:"message_hash"`
	Status                  int64          `gorm:"column:status" db:"status" json:"status" form:"status"`
	FromAddress             common.Address `gorm:"column:from_address;serializer:bytes" db:"from_address" json:"fromAddress" form:"from_address"`
	ETHAmount               *big.Int       `gorm:"serializer:u256;column:eth_amount" json:"ETHAmount"`
	ERC20Amount             *big.Int       `gorm:"serializer:u256;column:erc20_amount" json:"ERC20Amount"`
	GasLimit                *big.Int       `gorm:"serializer:u256;column:gas_limit" json:"gasLimit"`
	TimeLeft                *big.Int       `gorm:"serializer:u256;column:time_left" json:"timeLeft"`
	ToAddress               common.Address `gorm:"column:to_address;serializer:bytes" db:"to_address" json:"toAddress" form:"to_address"`
	L1TokenAddress          common.Address `gorm:"column:l1_token_address;serializer:bytes" db:"l1_token_address" json:"l1TokenAddress" form:"l1_token_address"`
	L2TokenAddress          common.Address `gorm:"column:l2_token_address;serializer:bytes" db:"l2_token_address" json:"l2TokenAddress" form:"l2_token_address"`
	Timestamp               int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
	TokenIds                string         `gorm:"column:token_ids" db:"token_ids" json:"tokenIds" form:"token_ids"`
	AssetType               int64          `gorm:"column:asset_type" db:"asset_type" json:"assetType" form:"asset_type"`
	TokenAmounts            string         `gorm:"column:token_amounts" db:"token_amounts" json:"tokenAmounts" form:"token_amounts"`
	ClaimedIndex            int64          `gorm:"column:claimed_index" db:"claimed_index" json:"claimed_index"`
}

func (L2ToL1) TableName() string {
	return "l2_to_l1"
}

type L2ToL1DB interface {
	L2ToL1View
	StoreL2ToL1Transactions([]L2ToL1) error
	UpdateTokenPairsAndAddress(l2L1List []L2ToL1) error
	UpdateReadyForProvedStatus(blockTimestamp int64, fraudProofWindows time.Duration) error
	UpdateL1FinalizeStatus(withdrawalHash common.Hash, finalizedL1EventGuid uuid.UUID) error
	UpdateTimeLeft() error
	MarkL2ToL1TransactionWithdrawalProven(l2L1List []L2ToL1) error
	MarkL2ToL1TransactionWithdrawalFinalized(l2L1List []L2ToL1) error
	UpdateL2ToL1MsgHashByL2TxHash(l2L1 L2ToL1) error
	UpdateL2ToL1L1TxHashByMsgHash(l2L1 L2ToL1) error
	UpdateL2ToL1ClaimedStatus(l1L2 L2ToL1) error
}

type L2ToL1View interface {
	L2ToL1List(string, int, int, string) (*gplus.Page[L2ToL1], error)
	L2LatestBlockHeader(chainId uint) (*common2.L2BlockHeader, error)
	L2ToL1TransactionWithdrawal(common.Hash) (*L2ToL1, error)
	GetBlockNumberFromHash(blockHash common.Hash) (*big.Int, error)
	L1LatestFinalizedBlockHeader(chainId uint) (*common2.L1BlockHeader, error)
}

/**
 * Implementation
 */

type l2ToL1DB struct {
	gorm *gorm.DB
}

func NewL21ToL1DB(db *gorm.DB) L2ToL1DB {
	gplus.Init(db)
	return &l2ToL1DB{gorm: db}
}

func (l2l1 l2ToL1DB) UpdateL2ToL1MsgHashByL2TxHash(l2L1Stu L2ToL1) error {
	result := l2l1.gorm.Model(&l2L1Stu).Where(&L2ToL1{L2TransactionHash: l2L1Stu.L2TransactionHash})
	result = result.UpdateColumn("message_hash", l2L1Stu.MessageHash.String())
	return result.Error
}

func (l2l1 l2ToL1DB) UpdateL2ToL1L1TxHashByMsgHash(l2L1Stu L2ToL1) error {
	result := l2l1.gorm.Model(&l2L1Stu).Where(&L2ToL1{MessageHash: l2L1Stu.MessageHash})
	result = result.UpdateColumn("l1_finalize_tx_hash", l2L1Stu.L1FinalizeTxHash.String()).UpdateColumn("status", 1)
	return result.Error
}

func (l2l1 l2ToL1DB) StoreL2ToL1Transactions(l1L2List []L2ToL1) error {
	result := l2l1.gorm.CreateInBatches(&l1L2List, len(l1L2List))
	return result.Error
}

func (l2l1 l2ToL1DB) L2ToL1List(address string, page int, pageSize int, order string) (*gplus.Page[L2ToL1], error) {
	query, l2l1s := gplus.NewQuery[L2ToL1]()
	if address != "0x00" {
		query.Eq(&l2l1s.FromAddress, address).Or().Eq(&l2l1s.ToAddress, address)
	}
	if strings.ToLower(order) == "asc" {
		query.OrderByAsc("timestamp")
	} else {
		query.OrderByDesc("timestamp")
	}
	dbPage := gplus.NewPage[L2ToL1](page, pageSize)
	rst, rstDB := gplus.SelectPage(dbPage, query)
	if rstDB.Error != nil {
		return nil, rstDB.Error
	}
	return rst, nil
}

func (l2l1 l2ToL1DB) UpdateTokenPairsAndAddress(l2L1List []L2ToL1) error {
	for i := 0; i < len(l2L1List); i++ {
		var l2ToL1 = L2ToL1{}
		result := l2l1.gorm.Where(&L2ToL1{WithdrawTransactionHash: l2L1List[i].WithdrawTransactionHash}).Take(&l2ToL1)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l2ToL1.L1BlockNumber = l2L1List[i].L1BlockNumber
		l2ToL1.FromAddress = l2L1List[i].FromAddress
		l2ToL1.ToAddress = l2L1List[i].ToAddress
		l2ToL1.L1TokenAddress = l2L1List[i].L1TokenAddress
		l2ToL1.L2TokenAddress = l2L1List[i].L2TokenAddress

		err := l2l1.gorm.Save(l2ToL1).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (l2l1 l2ToL1DB) MarkL2ToL1TransactionWithdrawalProven(l2L1List []L2ToL1) error {
	for i := 0; i < len(l2L1List); i++ {
		var l2ToL1 = L2ToL1{}
		result := l2l1.gorm.Where(&L2ToL1{WithdrawTransactionHash: l2L1List[i].WithdrawTransactionHash}).Take(&l2ToL1)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l2ToL1.L1BlockNumber = l2L1List[i].L1BlockNumber
		l2ToL1.L1ProveTxHash = l2L1List[i].L1ProveTxHash
		l2ToL1.Status = l2L1List[i].Status // in challenge period
		err := l2l1.gorm.Save(l2ToL1).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l2l1 l2ToL1DB) MarkL2ToL1TransactionWithdrawalFinalized(l2L1List []L2ToL1) error {
	for i := 0; i < len(l2L1List); i++ {
		var l2ToL1 = L2ToL1{}
		result := l2l1.gorm.Where(&L2ToL1{WithdrawTransactionHash: l2L1List[i].WithdrawTransactionHash}).Take(&l2ToL1)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l2ToL1.L1BlockNumber = l2L1List[i].L1BlockNumber
		l2ToL1.L1FinalizeTxHash = l2L1List[i].L1FinalizeTxHash
		l2ToL1.Status = l2L1List[i].Status // relayed
		err := l2l1.gorm.Save(l2ToL1).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l2l1 l2ToL1DB) UpdateReadyForProvedStatus(blockTimestamp int64, fraudProofWindows time.Duration) error {
	var l2ToL1 = L2ToL1{}
	err := l2l1.gorm.Model(&l2ToL1).Where("timestamp < ? AND status = ?", blockTimestamp, 0).Updates(map[string]interface{}{"status": common3.L2ToL1ReadyForProved, "time_left": uint64(fraudProofWindows)}).Error
	if err != nil {
		return err
	}
	return nil
}

func (l2l1 l2ToL1DB) UpdateL1FinalizeStatus(withdrawalHash common.Hash, finalizedL1EventGuid uuid.UUID) error {
	var l2ToL1 = L2ToL1{}
	result := l2l1.gorm.Where(&L2ToL1{WithdrawTransactionHash: withdrawalHash}).Take(&l2ToL1)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	var finalizedHash string
	result = l2l1.gorm.Table("l1_contract_events").Where("guid = ?", finalizedL1EventGuid).Select("transaction_hash").Find(&finalizedHash)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}

	result = l2l1.gorm.Model(&L2ToL1{}).Where("l2_block_number = ?", l2ToL1.L2BlockNumber.String()).Updates(map[string]interface{}{"l1_finalize_tx_hash": finalizedHash})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}

	return nil
}

func (l2l1 l2ToL1DB) UpdateTimeLeft() error {
	result := l2l1.gorm.Model(&L2ToL1{}).Where("time_left > ?", 0).Updates(map[string]interface{}{"time_left": gorm.Expr("GREATEST(time_left - 1, 0)")})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}

	result = l2l1.gorm.Model(&L2ToL1{}).Where("time_left = ? and status = ?", 0, common3.L2ToL1InChallengePeriod).Updates(map[string]interface{}{"status": common3.L2ToL1ReadyForClaim})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}

	return nil
}

func (l2l1 l2ToL1DB) GetBlockNumberFromHash(blockHash common.Hash) (*big.Int, error) {
	var l2BlockNumber uint64
	result := l2l1.gorm.Table("l2_block_headers").Where("hash = ?", blockHash.String()).Select("number").Take(&l2BlockNumber)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return new(big.Int).SetUint64(l2BlockNumber), nil
}

func (l2l1 l2ToL1DB) L2ToL1TransactionWithdrawal(withdrawalHash common.Hash) (*L2ToL1, error) {
	var l2ToL1Withdrawal L2ToL1
	result := l2l1.gorm.Where(&L2ToL1{WithdrawTransactionHash: withdrawalHash}).Take(&l2ToL1Withdrawal)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l2ToL1Withdrawal, nil
}

func (l2l1 l2ToL1DB) L2LatestBlockHeader(chainId uint) (*common2.L2BlockHeader, error) {
	tableName := fmt.Sprintf("l2_to_l1_%d", chainId)
	l2Query := l2l1.gorm.Where("timestamp = (?)", l2l1.gorm.Table(tableName).Select("MAX(timestamp)"))
	var l2Header common2.L2BlockHeader
	result := l2Query.Take(&l2Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l2Header, nil
}

func (l2l1 l2ToL1DB) L1LatestFinalizedBlockHeader(chainId uint) (*common2.L1BlockHeader, error) {
	tableName := fmt.Sprintf("l2_to_l1_%d", chainId)
	l1Query := l2l1.gorm.Where("number = (?)", l2l1.gorm.Table(tableName).Where("status != (?)", 4).Select("MAX(l1_block_number)"))
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

func (l2l1 l2ToL1DB) UpdateL2ToL1ClaimedStatus(l2L1 L2ToL1) error {

	result := l2l1.gorm.Model(&L2ToL1{}).Where("claimed_index = ?", l2L1.ClaimedIndex).Updates(map[string]interface{}{"status": l2L1.Status, "l1_finalize_tx_hash": l2L1.L1FinalizeTxHash})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	return nil
}
