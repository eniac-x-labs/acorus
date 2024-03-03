package event

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"math/big"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	common2 "github.com/cornerstone-labs/acorus/database/common"
)

type RelayMessage struct {
	GUID                 uuid.UUID      `gorm:"primaryKey;serializer:uuid"`
	BlockNumber          *big.Int       `gorm:"serializer:u256;column:block_number" db:"block_number"`
	RelayTransactionHash common.Hash    `gorm:"serializer:bytes"`
	MessageHash          common.Hash    `gorm:"serializer:bytes"`
	L1TokenAddress       common.Address `gorm:"column:l1_token_address;serializer:bytes"`
	L2TokenAddress       common.Address `gorm:"column:l2_token_address;serializer:bytes"`
	ETHAmount            *big.Int       `gorm:"serializer:u256;column:eth_amount"`
	ERC20Amount          *big.Int       `gorm:"serializer:u256;column:erc20_amount"`
	Related              bool           `json:"related"`
	Timestamp            uint64
}

func (RelayMessage) TableName() string {
	return "relay_message"
}

type RelayMessageDB interface {
	RelayMessageView
	StoreRelayMessage(string, []RelayMessage) error
	MarkedRelayMessageRelated(chainId string, relayMessageList []RelayMessage) error
	UpdateRelayMessageInfo(chainId string, relayMessageList []RelayMessage) error
}

type RelayMessageView interface {
	RelayMessageL1BlockHeader(l1ChainId, chainId string) (*common2.BlockHeader, error)
	RelayMessageUnRelatedList(string) ([]RelayMessage, error)
}

type relayMessageDB struct {
	gorm *gorm.DB
}

func NewRelayMessageDB(db *gorm.DB) RelayMessageDB {
	return &relayMessageDB{gorm: db}
}

func (rm relayMessageDB) RelayMessageL1BlockHeader(l1ChainId, chainId string) (*common2.BlockHeader, error) {
	l1Query := rm.gorm.Table("block_headers_"+l1ChainId).Where("number = (?)", rm.gorm.Table("relay_message_"+chainId).Select("MAX(block_number)"))
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

func (rm relayMessageDB) StoreRelayMessage(chainId string, relayMessageList []RelayMessage) error {
	tableName := "relay_message_" + chainId
	newRelayMessageList := make([]RelayMessage, 0)

	for i := 0; i < len(relayMessageList); i++ {
		var relayMessages = RelayMessage{}
		messageNew := relayMessageList[i]
		result := rm.gorm.Table(tableName).Where(&RelayMessage{MessageHash: messageNew.MessageHash}).Take(&relayMessages)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				newRelayMessageList = append(newRelayMessageList, messageNew)
				continue
			}
			return result.Error
		}
	}
	if len(newRelayMessageList) > 0 {
		return rm.gorm.Table(tableName).Omit("guid").Create(newRelayMessageList).Error
	}
	return nil
}

func (rm relayMessageDB) MarkedRelayMessageRelated(chainId string, relayMessageList []RelayMessage) error {
	for i := 0; i < len(relayMessageList); i++ {
		var relayMessages = RelayMessage{}
		messageHash := relayMessageList[i].MessageHash
		hash := common.Hash{}
		if messageHash == hash {
			continue
		}
		result := rm.gorm.Table("relay_message_" + chainId).Where(&RelayMessage{MessageHash: messageHash}).Take(&relayMessages)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		relayMessages.Related = true
		err := rm.gorm.Table("relay_message_" + chainId).Save(relayMessages).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (rm relayMessageDB) UpdateRelayMessageInfo(chainId string, relayMessageList []RelayMessage) error {
	for i := 0; i < len(relayMessageList); i++ {
		var relayMessages = RelayMessage{}
		messageHash := relayMessageList[i].MessageHash
		hash := common.Hash{}
		if messageHash == hash {
			continue
		}
		result := rm.gorm.Table("relay_message_" + chainId).Where(&RelayMessage{MessageHash: messageHash}).Take(&relayMessages)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		relayMessages.L1TokenAddress = relayMessageList[i].L1TokenAddress
		relayMessages.L2TokenAddress = relayMessageList[i].L2TokenAddress
		relayMessages.ETHAmount = relayMessageList[i].ETHAmount
		relayMessages.ERC20Amount = relayMessageList[i].ERC20Amount
		relayMessages.MessageHash = messageHash
		err := rm.gorm.Table("relay_message_" + chainId).Save(relayMessages).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (rm relayMessageDB) RelayMessageUnRelatedList(chainId string) ([]RelayMessage, error) {
	var unRelatedRelayList []RelayMessage
	err := rm.gorm.Table("relay_message_"+chainId).Where("related = ?", false).Find(&unRelatedRelayList).Error
	if err != nil {
		log.Error("get unrelated deposit finalized fail", "err", err)
	}
	return unRelatedRelayList, nil
}
