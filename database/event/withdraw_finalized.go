package event

import (
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
	"math/big"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"

	common2 "github.com/cornerstone-labs/acorus/database/common"
)

type WithdrawFinalized struct {
	GUID                     uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','';serializer:uuid"`
	BlockNumber              *big.Int       `gorm:"serializer:u256;column:block_number"`
	WithdrawHash             common.Hash    `gorm:"serializer:bytes"`
	MessageHash              common.Hash    `gorm:"serializer:bytes"`
	FinalizedTransactionHash common.Hash    `gorm:"serializer:bytes"`
	L1TokenAddress           common.Address `gorm:"column:l1_token_address;serializer:bytes"`
	L2TokenAddress           common.Address `gorm:"column:l2_token_address;serializer:bytes"`
	ETHAmount                *big.Int       `gorm:"serializer:u256;column:eth_amount"`
	ERC20Amount              *big.Int       `gorm:"serializer:u256;column:erc20_amount"`
	Related                  bool           `json:"related"`
	Timestamp                uint64
}

func (WithdrawFinalized) TableName() string {
	return "withdraw_finalized"
}

type WithdrawFinalizedDB interface {
	WithdrawFinalizedView
	StoreWithdrawFinalized(string, []WithdrawFinalized) error
	MarkedWithdrawFinalizedRelated(chainId string, withdrawFinalizedList []WithdrawFinalized) error
	UpdateWithdrawFinalizedInfo(chainId string, withdrawFinalizedList []WithdrawFinalized) error
	UpdateMsgHashFinalizedInfo(chainId string, withdrawFinalizedList []WithdrawFinalized) error
}

type WithdrawFinalizedView interface {
	WithdrawFinalizedL1BlockHeader(l1chainId, chainId string) (*common2.BlockHeader, error)
	WithdrawFinalizedUnRelatedList(chainId string) ([]WithdrawFinalized, error)
}

type withdrawFinalizedDB struct {
	gorm *gorm.DB
}

func NewWithdrawFinalizedDB(db *gorm.DB) WithdrawFinalizedDB {
	return &withdrawFinalizedDB{gorm: db}
}

func (w withdrawFinalizedDB) WithdrawFinalizedL1BlockHeader(l1chainId, chainId string) (*common2.BlockHeader, error) {
	l1Query := w.gorm.Table("block_headers_"+l1chainId).Where("number = (?)", w.gorm.Table("withdraw_finalized_"+chainId).Select("MAX(block_number)"))
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

func (w withdrawFinalizedDB) StoreWithdrawFinalized(chainId string, withdrawFinalizedList []WithdrawFinalized) error {
	tableName := "withdraw_finalized_" + chainId
	newWithdrawFinalizedList := make([]WithdrawFinalized, 0)
	for i := 0; i < len(withdrawFinalizedList); i++ {
		var withdrawFinalizeds = WithdrawFinalized{}
		newWithdrawFinalized := withdrawFinalizedList[i]
		withdrawHash := newWithdrawFinalized.WithdrawHash
		hash := common.Hash{}
		if withdrawHash == hash {
			continue
		}
		result := w.gorm.Table(tableName).Where(&WithdrawFinalized{WithdrawHash: withdrawHash}).Take(&withdrawFinalizeds)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				newWithdrawFinalizedList = append(newWithdrawFinalizedList, newWithdrawFinalized)
				continue
			}
			return result.Error
		}
	}
	if len(newWithdrawFinalizedList) > 0 {
		result := w.gorm.Omit("guid").Table("withdraw_finalized_" + chainId).Create(&newWithdrawFinalizedList)
		return result.Error
	}
	return nil
}

func (w withdrawFinalizedDB) UpdateMsgHashFinalizedInfo(chainId string, withdrawFinalizedList []WithdrawFinalized) error {
	for i := 0; i < len(withdrawFinalizedList); i++ {
		var wthdrawFinalized = WithdrawFinalized{}

		finalizedTransactionHash := withdrawFinalizedList[i].FinalizedTransactionHash
		hash := common.Hash{}
		if finalizedTransactionHash == hash {
			continue
		}
		result := w.gorm.Table("withdraw_finalized_" + chainId).Where(&WithdrawFinalized{FinalizedTransactionHash: finalizedTransactionHash}).Take(&wthdrawFinalized)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		wthdrawFinalized.MessageHash = withdrawFinalizedList[i].MessageHash
		err := w.gorm.Table("withdraw_finalized_" + chainId).Save(wthdrawFinalized).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (w withdrawFinalizedDB) MarkedWithdrawFinalizedRelated(chainId string, withdrawFinalizedList []WithdrawFinalized) error {
	for i := 0; i < len(withdrawFinalizedList); i++ {
		var withdrawFinalizeds = WithdrawFinalized{}
		withdrawHash := withdrawFinalizedList[i].WithdrawHash
		hash := common.Hash{}
		if withdrawHash == hash {
			continue
		}
		result := w.gorm.Table("withdraw_finalized_" + chainId).Where(&WithdrawFinalized{WithdrawHash: withdrawHash}).Take(&withdrawFinalizeds)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		withdrawFinalizeds.Related = true
		err := w.gorm.Table("withdraw_finalized_" + chainId).Save(withdrawFinalizeds).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (w withdrawFinalizedDB) UpdateWithdrawFinalizedInfo(chainId string, withdrawFinalizedList []WithdrawFinalized) error {
	for i := 0; i < len(withdrawFinalizedList); i++ {
		var withdrawFinalizeds = WithdrawFinalized{}
		finalizedTransactionHash := withdrawFinalizedList[i].FinalizedTransactionHash
		hash := common.Hash{}
		if finalizedTransactionHash == hash {
			continue
		}
		result := w.gorm.Table("withdraw_finalized_" + chainId).Where(&WithdrawFinalized{FinalizedTransactionHash: finalizedTransactionHash}).Take(&withdrawFinalizeds)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		withdrawFinalizeds.L1TokenAddress = withdrawFinalizedList[i].L1TokenAddress
		withdrawFinalizeds.L2TokenAddress = withdrawFinalizedList[i].L2TokenAddress
		withdrawFinalizeds.ETHAmount = withdrawFinalizedList[i].ETHAmount
		withdrawFinalizeds.ERC20Amount = withdrawFinalizedList[i].ERC20Amount
		err := w.gorm.Table("withdraw_finalized_" + chainId).Save(withdrawFinalizeds).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (w withdrawFinalizedDB) WithdrawFinalizedUnRelatedList(chainId string) ([]WithdrawFinalized, error) {
	var unRelatedFinalizedList []WithdrawFinalized
	err := w.gorm.Table("withdraw_finalized_"+chainId).Where("related = ?", false).Find(&unRelatedFinalizedList).Error
	if err != nil {
		log.Error("get unrelated withdraw finalized fail", "err", err)
	}
	return unRelatedFinalizedList, nil
}
