package event

import (
	"gorm.io/gorm"
	"math/big"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/log"

	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/ethereum/go-ethereum/common"
)

type WithdrawProven struct {
	GUID                  uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','';serializer:uuid"`
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
	StoreWithdrawProven(string, []WithdrawProven) error
	MarkedWithdrawProvenRelated(chainId string, withdrawProvenList []WithdrawProven) error
	UpdateWithdrawProvenInfo(chainId string, withdrawProvenList []WithdrawProven) error
}

type WithdrawProvenView interface {
	WithdrawProvenL1BlockHeader(l1ChainId, chainId string) (*common2.BlockHeader, error)
	WithdrawProvenUnRelatedList(string) ([]WithdrawProven, error)
}

type withdrawProvenDB struct {
	gorm *gorm.DB
}

func NewWithdrawProvenDB(db *gorm.DB) WithdrawProvenDB {
	return &withdrawProvenDB{gorm: db}
}

func (w withdrawProvenDB) WithdrawProvenL1BlockHeader(l1ChainId, chainId string) (*common2.BlockHeader, error) {
	l1Query := w.gorm.Table("block_headers_"+l1ChainId).Where("number = (?)", w.gorm.Table("withdraw_proven_"+chainId).Select("MAX(block_number)"))
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

func (w withdrawProvenDB) StoreWithdrawProven(chainId string, withdrawProvenList []WithdrawProven) error {
	result := w.gorm.Omit("guid").Table("withdraw_proven_" + chainId).Create(&withdrawProvenList)
	return result.Error
}

func (w withdrawProvenDB) MarkedWithdrawProvenRelated(chainId string, withdrawProvenList []WithdrawProven) error {
	for i := 0; i < len(withdrawProvenList); i++ {
		var withdrawProvens = WithdrawProven{}
		withdrawHash := withdrawProvenList[i].WithdrawHash
		hash := common.Hash{}
		if withdrawHash == hash {
			continue
		}
		result := w.gorm.Table("withdraw_proven_" + chainId).Where(&WithdrawProven{WithdrawHash: withdrawHash}).Take(&withdrawProvens)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		withdrawProvens.Related = true
		err := w.gorm.Table("withdraw_proven_" + chainId).Save(withdrawProvens).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (w withdrawProvenDB) UpdateWithdrawProvenInfo(chainId string, withdrawProvenList []WithdrawProven) error {
	for i := 0; i < len(withdrawProvenList); i++ {
		var withdrawProvens = WithdrawProven{}
		withdrawHash := withdrawProvenList[i].WithdrawHash
		hash := common.Hash{}
		if withdrawHash == hash {
			continue
		}
		result := w.gorm.Table("withdraw_proven_" + chainId).Where(&WithdrawProven{WithdrawHash: withdrawHash}).Take(&withdrawProvens)
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
		err := w.gorm.Table("withdraw_proven_" + chainId).Save(withdrawProvens).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (w withdrawProvenDB) WithdrawProvenUnRelatedList(chainId string) ([]WithdrawProven, error) {
	var unRelatedProvenList []WithdrawProven
	err := w.gorm.Table("withdraw_proven_"+chainId).Where("related = ?", false).Find(&unRelatedProvenList).Error
	if err != nil {
		log.Error("get unrelated withdraw proven fail", "err", err)
	}
	return unRelatedProvenList, nil
}
