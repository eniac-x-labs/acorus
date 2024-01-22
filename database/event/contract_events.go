package event

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ContractEvent struct {
	GUID            uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')"`
	BlockHash       common.Hash    `gorm:"serializer:bytes"`
	ContractAddress common.Address `gorm:"serializer:bytes"`
	TransactionHash common.Hash    `gorm:"serializer:bytes"`
	LogIndex        uint64
	BlockNumber     *big.Int    `gorm:"serializer:u256"`
	EventSignature  common.Hash `gorm:"serializer:bytes"`
	Timestamp       uint64
	RLPLog          *types.Log `gorm:"serializer:rlp;column:rlp_bytes"`
}

func ContractEventFromLog(log *types.Log, timestamp uint64, blockNumber *big.Int) ContractEvent {
	eventSig := common.Hash{}
	if len(log.Topics) > 0 {
		eventSig = log.Topics[0]
	}
	return ContractEvent{
		BlockHash:       log.BlockHash,
		TransactionHash: log.TxHash,
		ContractAddress: log.Address,
		EventSignature:  eventSig,
		BlockNumber:     blockNumber,
		LogIndex:        uint64(log.Index),
		Timestamp:       timestamp,
		RLPLog:          log,
	}
}

func (c *ContractEvent) AfterFind(tx *gorm.DB) error {
	// Fill in some of the derived fields that are not
	// populated when decoding the RLPLog from RLP
	c.RLPLog.BlockHash = c.BlockHash
	c.RLPLog.TxHash = c.TransactionHash
	c.RLPLog.Index = uint(c.LogIndex)
	c.RLPLog.BlockNumber = c.BlockNumber.Uint64()
	return nil
}

type ChainContractEvent struct {
	ContractEvent `gorm:"embedded"`
}

type ContractEventsView interface {
	ChainContractEvent(string, uuid.UUID) (*ContractEvent, error)
	ChainContractEventWithFilter(string, ContractEvent) (*ContractEvent, error)
	ChainContractEventsWithFilter(string, ContractEvent, *big.Int, *big.Int) ([]ContractEvent, error)
	ChainLatestContractEventWithFilter(string, ContractEvent) (*ContractEvent, error)

	ContractEventsWithFilter(string, ContractEvent, *big.Int, *big.Int) ([]ContractEvent, error)
}

type ContractEventsDB interface {
	ContractEventsView

	StoreChainContractEvents(string, []ChainContractEvent) error
}

type contractEventsDB struct {
	gorm *gorm.DB
}

func NewContractEventsDB(db *gorm.DB) ContractEventsDB {
	return &contractEventsDB{gorm: db}
}

func (db *contractEventsDB) StoreChainContractEvents(chainId string, events []ChainContractEvent) error {
	result := db.gorm.Omit("guid").Table("contract_events_" + chainId).Create(&events)
	return result.Error
}

func (db *contractEventsDB) ChainContractEvent(chainId string, uuid uuid.UUID) (*ContractEvent, error) {
	return db.ChainContractEventWithFilter(chainId, ContractEvent{GUID: uuid})
}

func (db *contractEventsDB) ChainContractEventWithFilter(chainId string, filter ContractEvent) (*ContractEvent, error) {
	var contractEvent ContractEvent
	result := db.gorm.Table("contract_events_" + chainId).Where(&filter).Take(&contractEvent)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &contractEvent, nil
}

func (db *contractEventsDB) ChainContractEventsWithFilter(chainId string, filter ContractEvent, fromHeight, toHeight *big.Int) ([]ContractEvent, error) {
	if fromHeight == nil {
		fromHeight = big.NewInt(0)
	}
	if toHeight == nil {
		return nil, errors.New("end height unspecified")
	}
	if fromHeight.Cmp(toHeight) > 0 {
		return nil, fmt.Errorf("fromHeight %d is greater than toHeight %d", fromHeight, toHeight)
	}
	contractEventsTable := " contract_events_" + chainId + " as contract_events "
	blockHeadersTable := " block_headers_" + chainId + " as block_headers "
	query := db.gorm.Table(contractEventsTable).Where(&filter)
	query = query.Joins("INNER JOIN " + blockHeadersTable + " ON contract_events.block_hash = block_headers.hash")
	query = query.Where("block_headers.number >= ? AND block_headers.number <= ?", fromHeight, toHeight)
	query = query.Order("block_headers.number ASC").Select("contract_events.*")

	var events []ContractEvent
	result := query.Find(&events)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return events, nil
}

func (db *contractEventsDB) ChainLatestContractEventWithFilter(chainId string, filter ContractEvent) (*ContractEvent, error) {
	var contractEvent ContractEvent
	result := db.gorm.Table("contract_events_" + chainId).Where(&filter).Order("timestamp DESC").Take(&contractEvent)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &contractEvent, nil
}

func (db *contractEventsDB) ContractEventsWithFilter(chainId string, filter ContractEvent, fromHeight, toHeight *big.Int) ([]ContractEvent, error) {
	eventList, err := db.ChainContractEventsWithFilter(chainId, filter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}
	return eventList, nil
}
