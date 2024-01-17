package event

import (
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
	"math/big"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/ethereum/go-ethereum/common"
)

type WithdrawProven struct {
	GUID                  uuid.UUID      `gorm:"primaryKey"`
	BlockNumber           *big.Int       `gorm:"serializer:u256;column:block_number" db:"block_number"`
	WithdrawHash          common.Hash    `gorm:"serializer:bytes"`
	MessageHash           common.Hash    `gorm:"serializer:bytes"`
	ProvenTransactionHash common.Hash    `gorm:"serializer:bytes"`
	L1TokenAddress        common.Address `gorm:"column:l1_token_address;serializer:bytes"`
	L2TokenAddress        common.Address `gorm:"column:l2_token_address;serializer:bytes"`
	ETHAmount             *big.Int       `gorm:"serializer:u256;column:eth_amount"`
	ERC20Amount           *big.Int       `gorm:"serializer:u256;column:erc20_amount"`
	Related               bool           `json:"related"`
	Timestamp             uint64
}

func (WithdrawProven) TableName() string {
	return "withdraw_proven"
}

type WithdrawProvenDB interface {
	WithdrawProvenView
	StoreWithdrawProven([]WithdrawProven) error
	MarkedWithdrawProvenRelated(withdrawProvenList []WithdrawProven) error
	UpdateWithdrawProvenInfo(withdrawProvenList []WithdrawProven) error
}

type WithdrawProvenView interface {
	WithdrawProvenL1BlockHeader() (*common2.BlockHeader, error)
	WithdrawProvenUnRelatedList() ([]WithdrawProven, error)
}

type withdrawProvenDB struct {
	gorm *gorm.DB
}

func NewWithdrawProvenDB(db *gorm.DB) WithdrawProvenDB {
	return &withdrawProvenDB{gorm: db}
}

func (w withdrawProvenDB) WithdrawProvenL1BlockHeader() (*common2.BlockHeader, error) {
	l1Query := w.gorm.Where("number = (?)", w.gorm.Table("withdraw_proven").Select("MAX(block_number)"))
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

func (w withdrawProvenDB) StoreWithdrawProven(withdrawProvenList []WithdrawProven) error {
	result := w.gorm.CreateInBatches(&withdrawProvenList, len(withdrawProvenList))
	return result.Error
}

func (w withdrawProvenDB) MarkedWithdrawProvenRelated(withdrawProvenList []WithdrawProven) error {
	for i := 0; i < len(withdrawProvenList); i++ {
		var withdrawProvens = WithdrawProven{}
		result := w.gorm.Where(&WithdrawProven{WithdrawHash: withdrawProvenList[i].WithdrawHash}).Take(&withdrawProvens)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		withdrawProvens.Related = true
		err := w.gorm.Save(withdrawProvens).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (w withdrawProvenDB) UpdateWithdrawProvenInfo(withdrawProvenList []WithdrawProven) error {
	for i := 0; i < len(withdrawProvenList); i++ {
		var withdrawProvens = WithdrawProven{}
		result := w.gorm.Where(&WithdrawProven{WithdrawHash: withdrawProvenList[i].WithdrawHash}).Take(&withdrawProvens)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		withdrawProvens.L1TokenAddress = withdrawProvenList[i].L1TokenAddress
		withdrawProvens.L2TokenAddress = withdrawProvenList[i].L2TokenAddress
		withdrawProvens.ETHAmount = withdrawProvenList[i].ETHAmount
		withdrawProvens.ERC20Amount = withdrawProvenList[i].ERC20Amount
		withdrawProvens.MessageHash = withdrawProvenList[i].MessageHash
		err := w.gorm.Save(withdrawProvens).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (w withdrawProvenDB) WithdrawProvenUnRelatedList() ([]WithdrawProven, error) {
	var unRelatedProvenList []WithdrawProven
	err := w.gorm.Table("withdraw_proven").Where("related = ?", false).Find(&unRelatedProvenList).Error
	if err != nil {
		log.Error("get unrelated withdraw proven fail", "err", err)
	}
	return unRelatedProvenList, nil
}
