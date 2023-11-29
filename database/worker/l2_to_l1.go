package worker

import (
	"errors"
	"gorm.io/gorm"
	"math/big"
	"time"

	"github.com/google/uuid"

	"github.com/acmestack/gorm-plus/gplus"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/ethereum/go-ethereum/common"
)

type L2ToL1 struct {
	GUID                      uuid.UUID      `gorm:"primaryKey"`
	L2BlockHash               common.Hash    `gorm:"column:l2_block_hash;serializer:bytes" db:"l2_block_hash" json:"l2_block_hash" form:"l2_block_hash"`
	MsgNonce                  *big.Int       `gorm:"column:msg_nonce;serializer:u256" db:"msg_nonce" json:"msg_nonce" form:"msg_nonce"`
	L2TransactionHash         common.Hash    `gorm:"column:l2_transaction_hash;serializer:bytes" db:"l2_transaction_hash" json:"l2_transaction_hash" form:"l2_transaction_hash"`
	L2WithdrawTransactionHash common.Hash    `gorm:"column:l2_withdraw_transaction_hash;serializer:bytes" db:"l2_withdraw_transaction_hash" json:"l2_withdraw_transaction_hash" form:"l2_withdraw_transaction_hash"`
	L1ProveTxHash             common.Hash    `gorm:"column:l1_prove_tx_hash;serializer:bytes" db:"l1_prove_tx_hash" json:"l1_prove_tx_hash" form:"l1_prove_tx_hash"`
	L1FinalizeTxHash          common.Hash    `gorm:"column:l1_finalize_tx_hash;serializer:bytes" db:"l1_finalize_tx_hash" json:"l1_finalize_tx_hash" form:"l1_finalize_tx_hash"`
	Status                    int64          `gorm:"column:status" db:"status" json:"status" form:"status"`
	FromAddress               common.Address `gorm:"column:from_address;serializer:bytes" db:"from_address" json:"from_address" form:"from_address"`
	ETHAmount                 *big.Int       `gorm:"serializer:u256;column:eth_amount" json:"eth_amount"`
	ERC20Amount               *big.Int       `gorm:"serializer:u256;column:erc20_amount" json:"erc20_amount"`
	GasLimit                  *big.Int       `gorm:"serializer:u256;column:gas_limit" json:"gas_limit"`
	TimeLeft                  *big.Int       `gorm:"serializer:u256;column:time_left" json:"time_left"`
	ToAddress                 common.Address `gorm:"column:to_address;serializer:bytes" db:"to_address" json:"to_address" form:"to_address"`
	L1TokenAddress            common.Address `gorm:"column:l1_token_address;serializer:bytes" db:"l1_token_address" json:"l1_token_address" form:"l1_token_address"`
	L2TokenAddress            common.Address `gorm:"column:l2_token_address;serializer:bytes" db:"l2_token_address" json:"l2_token_address" form:"l2_token_address"`
	Timestamp                 int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
	AssetType                 int64          `gorm:"column:asset_type" db:"asset_type" json:"asset_type" form:"asset_type"`
	TokenIds                  string         `gorm:"column:token_ids" db:"token_ids" json:"token_ids" form:"token_ids"`
	TokenAmounts              string         `gorm:"column:token_amounts" db:"token_amounts" json:"token_amounts" form:"token_amounts"`
	MsgHash                   common.Hash    `gorm:"column:msg_hash;serializer:bytes" db:"msg_hash" json:"msg_hash" form:"msg_hash"`
}

func (L2ToL1) TableName() string {
	return "l2_to_l1"
}

type L2ToL1DB interface {
	L2ToL1View
	StoreL2ToL1Transactions([]L2ToL1) error
	UpdateTokenPairs(l1L2List []L2ToL1) error
	UpdateReadyForProvedStatus(blockTimestamp int64) error
	UpdateReadyForClaimStatus(withdrawalHash common.Hash, provenL1EventGuid uuid.UUID) error
	UpdateL1FinalizeStatus(withdrawalHash common.Hash, finalizedL1EventGuid uuid.UUID) error
	UpdateClaimedStatus(withdrawalHash common.Hash) error
	UpdateTimeLeft() error
	UpdateL2ToL1MsgHashByL2TxHash(l2L1 L2ToL1) error
	UpdateL2ToL1L1TxHashByMsgHash(l2L1 L2ToL1) error
}

type L2ToL1View interface {
	L2ToL1List(string, int, int) (*gplus.Page[L2ToL1], error)
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
	result = result.UpdateColumn("msg_hash", l2L1Stu.MsgHash.String())
	return result.Error
}

func (l2l1 l2ToL1DB) UpdateL2ToL1L1TxHashByMsgHash(l2L1Stu L2ToL1) error {
	result := l2l1.gorm.Model(&l2L1Stu).Where(&L2ToL1{MsgHash: l2L1Stu.MsgHash})
	result = result.UpdateColumn("l1_finalize_tx_hash", l2L1Stu.L1FinalizeTxHash.String()).UpdateColumn("status", 1)
	return result.Error
}

func (l2l1 l2ToL1DB) StoreL2ToL1Transactions(l1L2List []L2ToL1) error {
	result := l2l1.gorm.CreateInBatches(&l1L2List, len(l1L2List))
	return result.Error
}

func (l2l1 l2ToL1DB) L2ToL1List(address string, page int, pageSize int) (*gplus.Page[L2ToL1], error) {
	query, l2l1s := gplus.NewQuery[L2ToL1]()
	query.Eq(&l2l1s.FromAddress, address)
	dbPage := gplus.NewPage[L2ToL1](page, pageSize)
	rst, rstDB := gplus.SelectPage(dbPage, query)
	if rstDB.Error != nil {
		return nil, rstDB.Error
	}
	return rst, nil
}

func (l2l1 l2ToL1DB) UpdateTokenPairs(l1L2List []L2ToL1) error {
	for i := 0; i < len(l1L2List); i++ {
		var l2ToL1 = L2ToL1{}
		result := l2l1.gorm.Where(&L2ToL1{L2WithdrawTransactionHash: l1L2List[i].L2WithdrawTransactionHash}).Take(&l2ToL1)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l2ToL1.L1TokenAddress = l1L2List[i].L1TokenAddress
		l2ToL1.L2TokenAddress = l1L2List[i].L2TokenAddress

		err := l2l1.gorm.Save(l2ToL1).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (l2l1 l2ToL1DB) UpdateReadyForProvedStatus(blockTimestamp int64) error {

	var l2ToL1 = L2ToL1{}

	err := l2l1.gorm.Model(&l2ToL1).Where("timestamp < ? AND status = ?", blockTimestamp, common2.Pending).Updates(map[string]interface{}{"status": common2.ReadyForProved, "time_left": uint64(time.Hour*24*7) / 1e9}).Error
	if err != nil {
		return err
	}

	return nil
}

func (l2l1 l2ToL1DB) UpdateReadyForClaimStatus(withdrawalHash common.Hash, provenL1EventGuid uuid.UUID) error {
	var l2ToL1 = L2ToL1{}
	result := l2l1.gorm.Where(&L2ToL1{L2WithdrawTransactionHash: withdrawalHash}).Take(&l2ToL1)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}

	var provenHash string
	result = l2l1.gorm.Table("l1_contract_events").Where("guid = ?", provenL1EventGuid).Select("transaction_hash").Find(&provenHash)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}

	result = l2l1.gorm.Model(&L2ToL1{}).Where("l2_block_number = ?", l2ToL1.L2BlockHash.String()).Updates(map[string]interface{}{"l1_prove_tx_hash": provenHash, "status": common2.ReadyForClaim})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}

	return nil
}

func (l2l1 l2ToL1DB) UpdateL1FinalizeStatus(withdrawalHash common.Hash, finalizedL1EventGuid uuid.UUID) error {
	var l2ToL1 = L2ToL1{}
	result := l2l1.gorm.Where(&L2ToL1{L2WithdrawTransactionHash: withdrawalHash}).Take(&l2ToL1)
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

	result = l2l1.gorm.Model(&L2ToL1{}).Where("l2_block_number = ?", l2ToL1.L2BlockHash.String()).Updates(map[string]interface{}{"l1_finalize_tx_hash": finalizedHash})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}

	return nil
}

func (l2l1 l2ToL1DB) UpdateClaimedStatus(withdrawalHash common.Hash) error {
	var l2ToL1 = L2ToL1{}
	result := l2l1.gorm.Where(&L2ToL1{L2WithdrawTransactionHash: withdrawalHash}).Take(&l2ToL1)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}

	result = l2l1.gorm.Model(&L2ToL1{}).Where("l2_block_number = ?", l2ToL1.L2BlockHash.String()).Updates(map[string]interface{}{"status": common2.Claimed})
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

	result = l2l1.gorm.Model(&L2ToL1{}).Where("time_left = ? and status = ?", 0, common2.ReadyForProved).Updates(map[string]interface{}{"status": common2.ReadyForClaim})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}

	return nil
}
