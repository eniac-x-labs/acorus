package worker

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math/big"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/log"

	common3 "github.com/cornerstone-labs/acorus/common"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/ethereum/go-ethereum/common"
)

type L2ToL1 struct {
	GUID                    uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','';serializer:uuid" json:"guid"`
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
	GasLimit                *big.Int       `gorm:"serializer:u256;column:gas_limit" json:"gasLimit"`
	TimeLeft                *big.Int       `gorm:"serializer:u256;column:time_left" json:"timeLeft"`
	ToAddress               common.Address `gorm:"column:to_address;serializer:bytes" db:"to_address" json:"toAddress" form:"to_address"`
	L1TokenAddress          common.Address `gorm:"column:l1_token_address;serializer:bytes" db:"l1_token_address" json:"l1TokenAddress" form:"l1_token_address"`
	L2TokenAddress          common.Address `gorm:"column:l2_token_address;serializer:bytes" db:"l2_token_address" json:"l2TokenAddress" form:"l2_token_address"`
	Version                 int64          `gorm:"column:version" json:"version"`
	Timestamp               int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
	TokenIds                string         `gorm:"column:token_ids" db:"token_ids" json:"tokenIds" form:"token_ids"`
	AssetType               int64          `gorm:"column:asset_type" db:"asset_type" json:"assetType" form:"asset_type"`
	TokenAmounts            string         `gorm:"column:token_amounts" db:"token_amounts" json:"tokenAmounts" form:"token_amounts"`
	ClaimedIndex            int64          `gorm:"column:claimed_index" db:"claimed_index" json:"claimed_index"`
}

type L2ToL1DB interface {
	L2ToL1View
	StoreL2ToL1Transactions(string, []L2ToL1) error
	UpdateTokenPairsAndAddress(chainId string, l2L1List []L2ToL1) error
	UpdateMessageHash(chainId string, l1L2List []L2ToL1) error
	UpdateReadyForProvedStatus(chainId string, blockTimestamp int64, fraudProofWindows time.Duration) error
	UpdateL1FinalizeStatus(chainId string, withdrawalHash common.Hash, finalizedL1EventGuid uuid.UUID) error
	UpdateTimeLeft(chainId string) error
	MarkL2ToL1TransactionWithdrawalProven(chainId string, l2L1List []L2ToL1) error
	MarkL2ToL1TransactionWithdrawalProvenV0(chainId string, l2L1List []L2ToL1) error
	MarkL2ToL1TransactionWithdrawalFinalized(chainId string, l2L1List []L2ToL1) error
	MarkL2ToL1TransactionWithdrawalFinalizedV0(chainId string, l2L1List []L2ToL1) error
	MarkL2ToL1TransactionMsgHashFinalized(chainId string, l2L1List []L2ToL1) error
	MarkL2ToL1TransactionMsgHashFinalizedV0(chainId string, l2L1List []L2ToL1) error
	UpdateL2ToL1ClaimedStatus(chainId string, l1L2 L2ToL1) error
}

type L2ToL1View interface {
	L2ToL1List(string, string, int, int, string) ([]L2ToL1, int64)
	L1LatestBlockHeader(l2chainId, destChainId string) (*common2.BlockHeader, error)
	L2LatestBlockHeader(chainId string) (*common2.BlockHeader, error)
	L2ToL1TransactionWithdrawal(string, common.Hash) (*L2ToL1, error)
	L2ToL1TransactionMsgHash(string, common.Hash) (*L2ToL1, error)
	GetBlockNumberFromHash(chainId string, blockHash common.Hash) (*big.Int, error)
	L1LatestFinalizedBlockHeader(chainId string) (*common2.BlockHeader, error)
	L2LatestFinalizedBlockHeader(chainId string) (*common2.BlockHeader, error)
}

type l2ToL1DB struct {
	gorm *gorm.DB
}

func NewL21ToL1DB(db *gorm.DB) L2ToL1DB {
	return &l2ToL1DB{gorm: db}
}

func (l2l1 l2ToL1DB) StoreL2ToL1Transactions(chainId string, l1L2List []L2ToL1) error {
	result := l2l1.gorm.Omit("guid").Table("l2_to_l1_" + chainId).Create(&l1L2List)
	return result.Error
}

func (l2l1 l2ToL1DB) L2ToL1List(chainId string, address string, page int, pageSize int, order string) (l2L1List []L2ToL1, total int64) {
	var totalRecord int64
	var l2ToL1List []L2ToL1
	queryStateRoot := l2l1.gorm.Table("l2_to_l1_" + chainId)
	if address != "0x00" {
		err := l2l1.gorm.Table("l2_to_l1_"+chainId).Select("l2_block_number").Where("from_address = ?", address).Or(" to_address = ?", address).Count(&totalRecord).Error
		if err != nil {
			log.Error("get l2 to l1 by address count fail")
		}
		queryStateRoot.Where("from_address = ?", address).Or(" to_address = ?", address).Offset((page - 1) * pageSize).Limit(pageSize)
	} else {
		err := l2l1.gorm.Table("l2_to_l1").Select("l2_block_number").Count(&totalRecord).Error
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
	qErr := queryStateRoot.Find(&l2ToL1List).Error
	if qErr != nil {
		log.Error("get l2 to l1 list fail", "err", qErr)
	}
	return l2ToL1List, totalRecord
}

func (l2l1 l2ToL1DB) UpdateTokenPairsAndAddress(chainId string, l2L1List []L2ToL1) error {
	for i := 0; i < len(l2L1List); i++ {
		var l2ToL1 = L2ToL1{}
		withdrawTransactionHash := l2L1List[i].WithdrawTransactionHash
		nilHash := common.Hash{}
		if withdrawTransactionHash == nilHash {
			continue
		}
		result := l2l1.gorm.Table("l2_to_l1_" + chainId).Where(&L2ToL1{WithdrawTransactionHash: withdrawTransactionHash}).Take(&l2ToL1)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l2ToL1.L2TransactionHash = l2L1List[i].L2TransactionHash
		l2ToL1.FromAddress = l2L1List[i].FromAddress
		l2ToL1.ToAddress = l2L1List[i].ToAddress
		l2ToL1.L1TokenAddress = l2L1List[i].L1TokenAddress
		l2ToL1.L2TokenAddress = l2L1List[i].L2TokenAddress
		l2ToL1.TokenAmounts = l2L1List[i].TokenAmounts
		l2ToL1.ETHAmount = l2L1List[i].ETHAmount
		l2ToL1.AssetType = l2L1List[i].AssetType
		err := l2l1.gorm.Table("l2_to_l1_" + chainId).Save(l2ToL1).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l2l1 l2ToL1DB) UpdateMessageHash(chainId string, l2L1List []L2ToL1) error {
	for i := 0; i < len(l2L1List); i++ {
		var l2ToL1 = L2ToL1{}
		nilHash := common.Hash{}
		l2TransactionHash := l2L1List[i].L2TransactionHash
		if l2TransactionHash == nilHash {
			continue
		}
		result := l2l1.gorm.Table("l2_to_l1_" + chainId).Where(&L2ToL1{L2TransactionHash: l2TransactionHash}).Take(&l2ToL1)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l2ToL1.MessageHash = l2L1List[i].MessageHash
		err := l2l1.gorm.Table("l2_to_l1_" + chainId).Save(l2ToL1).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l2l1 l2ToL1DB) MarkL2ToL1TransactionWithdrawalProven(chainId string, l2L1List []L2ToL1) error {
	for i := 0; i < len(l2L1List); i++ {
		var l2ToL1 = L2ToL1{}
		withdrawTransactionHash := l2L1List[i].WithdrawTransactionHash
		nilHash := common.Hash{}
		if withdrawTransactionHash == nilHash {
			continue
		}
		result := l2l1.gorm.Table("l2_to_l1_" + chainId).Where(&L2ToL1{WithdrawTransactionHash: withdrawTransactionHash}).Take(&l2ToL1)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l2ToL1.L1BlockNumber = l2L1List[i].L1BlockNumber
		l2ToL1.L1ProveTxHash = l2L1List[i].L1ProveTxHash
		l2ToL1.Status = common3.L2ToL1InChallengePeriod // in challenge period
		err := l2l1.gorm.Table("l2_to_l1_" + chainId).Save(l2ToL1).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l2l1 l2ToL1DB) MarkL2ToL1TransactionWithdrawalProvenV0(chainId string, l2L1List []L2ToL1) error {
	for i := 0; i < len(l2L1List); i++ {
		var l2ToL1 = L2ToL1{}
		if l2L1List[i].L1BlockNumber.Uint64() <= 0 {
			continue
		}
		messageHash := l2L1List[i].MessageHash
		nilHash := common.Hash{}
		if messageHash == nilHash {
			continue
		}
		result := l2l1.gorm.Table("l2_to_l1_" + chainId).Where(&L2ToL1{MessageHash: messageHash}).Take(&l2ToL1)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		log.Info("mark transaction v0 prove", "L1BlockNumber",
			l2L1List[i].L1BlockNumber, "L1ProveTxHash", l2L1List[i].L1ProveTxHash,
			"WithdrawTransactionHash", l2L1List[i].WithdrawTransactionHash)
		l2ToL1.L1BlockNumber = l2L1List[i].L1BlockNumber
		l2ToL1.L1ProveTxHash = l2L1List[i].L1ProveTxHash
		if l2ToL1.TimeLeft.Uint64() > 0 {
			l2ToL1.Status = common3.L2ToL1InChallengePeriod // in challenge period
		} else {
			l2ToL1.Status = common3.L2ToL1ReadyForClaim // ready for claim
		}
		err := l2l1.gorm.Table("l2_to_l1_" + chainId).Save(&l2ToL1).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l2l1 l2ToL1DB) MarkL2ToL1TransactionWithdrawalFinalized(chainId string, l2L1List []L2ToL1) error {
	for i := 0; i < len(l2L1List); i++ {
		var l2ToL1 = L2ToL1{}
		withdrawTransactionHash := l2L1List[i].WithdrawTransactionHash
		nilHash := common.Hash{}
		if withdrawTransactionHash == nilHash {
			continue
		}
		result := l2l1.gorm.Table("l2_to_l1_" + chainId).Where(&L2ToL1{WithdrawTransactionHash: withdrawTransactionHash}).Take(&l2ToL1)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l2ToL1.L1BlockNumber = l2L1List[i].L1BlockNumber
		l2ToL1.L1FinalizeTxHash = l2L1List[i].L1FinalizeTxHash
		l2ToL1.Status = common3.L2ToL1Claimed // relayed
		err := l2l1.gorm.Table("l2_to_l1_" + chainId).Save(l2ToL1).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l2l1 l2ToL1DB) MarkL2ToL1TransactionWithdrawalFinalizedV0(chainId string, l2L1List []L2ToL1) error {
	for i := 0; i < len(l2L1List); i++ {
		var l2ToL1 = L2ToL1{}
		if l2L1List[i].L1BlockNumber.Uint64() <= 0 {
			continue
		}
		withdrawTransactionHash := l2L1List[i].WithdrawTransactionHash
		nilHash := common.Hash{}
		if withdrawTransactionHash == nilHash {
			continue
		}
		result := l2l1.gorm.Table("l2_to_l1_" + chainId).Where(&L2ToL1{WithdrawTransactionHash: withdrawTransactionHash}).Take(&l2ToL1)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		log.Info("mark transaction v0 finalized",
			"L1BlockNumber", l2L1List[i].L1BlockNumber, "L1FinalizeTxHash", l2L1List[i].L1FinalizeTxHash,
			"WithdrawTransactionHash", l2L1List[i].WithdrawTransactionHash)
		l2ToL1.L1BlockNumber = l2L1List[i].L1BlockNumber
		l2ToL1.L1FinalizeTxHash = l2L1List[i].L1FinalizeTxHash
		l2ToL1.Status = common3.L2ToL1Claimed // relayed
		err := l2l1.gorm.Table("l2_to_l1_" + chainId).Save(&l2ToL1).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l2l1 l2ToL1DB) MarkL2ToL1TransactionMsgHashFinalized(chainId string, l2L1List []L2ToL1) error {
	for i := 0; i < len(l2L1List); i++ {
		var l2ToL1 = L2ToL1{}
		messageHash := l2L1List[i].MessageHash
		nilHash := common.Hash{}
		if messageHash == nilHash {
			continue
		}
		result := l2l1.gorm.Table("l2_to_l1_" + chainId).Where(&L2ToL1{MessageHash: messageHash}).Take(&l2ToL1)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		l2ToL1.L1BlockNumber = l2L1List[i].L1BlockNumber
		l2ToL1.L1FinalizeTxHash = l2L1List[i].L1FinalizeTxHash
		l2ToL1.Status = common3.L2ToL1Claimed // relayed
		err := l2l1.gorm.Table("l2_to_l1_" + chainId).Save(l2ToL1).Error
		if err != nil {
			return err
		}
	}
	return nil
}
func (l2l1 l2ToL1DB) MarkL2ToL1TransactionMsgHashFinalizedV0(chainId string, l2L1List []L2ToL1) error {
	for i := 0; i < len(l2L1List); i++ {
		var l2ToL1 = L2ToL1{}
		if l2L1List[i].L1BlockNumber.Uint64() <= 0 {
			continue
		}
		messageHash := l2L1List[i].MessageHash
		nilHash := common.Hash{}
		if messageHash == nilHash {
			continue
		}
		result := l2l1.gorm.Table("l2_to_l1_" + chainId).Where(&L2ToL1{MessageHash: messageHash}).Take(&l2ToL1)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}
		log.Info("mark transaction v0 finalized",
			"L1BlockNumber", l2L1List[i].L1BlockNumber, "L1FinalizeTxHash", l2L1List[i].L1FinalizeTxHash,
			"MessageHash", l2L1List[i].WithdrawTransactionHash)
		l2ToL1.L1BlockNumber = l2L1List[i].L1BlockNumber
		l2ToL1.L1FinalizeTxHash = l2L1List[i].L1FinalizeTxHash
		l2ToL1.Status = common3.L2ToL1Claimed // relayed
		err := l2l1.gorm.Table("l2_to_l1_" + chainId).Save(&l2ToL1).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (l2l1 l2ToL1DB) UpdateReadyForProvedStatus(chainId string, blockTimestamp int64, fraudProofWindows time.Duration) error {
	err := l2l1.gorm.Table("l2_to_l1_"+chainId).Where("timestamp < ? AND status = ?", blockTimestamp, 0).Updates(map[string]interface{}{"status": common3.L2ToL1ReadyForProved, "time_left": uint64(fraudProofWindows)}).Error
	if err != nil {
		return err
	}
	return nil
}

func (l2l1 l2ToL1DB) UpdateL1FinalizeStatus(chainId string, withdrawalHash common.Hash, finalizedL1EventGuid uuid.UUID) error {
	var l2ToL1 = L2ToL1{}
	nilHash := common.Hash{}
	if withdrawalHash == nilHash {
		return nil
	}
	result := l2l1.gorm.Table("l2_to_l1_" + chainId).Where(&L2ToL1{WithdrawTransactionHash: withdrawalHash}).Take(&l2ToL1)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	var finalizedHash string
	result = l2l1.gorm.Table("contract_events_"+chainId).Where("guid = ?", finalizedL1EventGuid).Select("transaction_hash").Find(&finalizedHash)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	result = l2l1.gorm.Table("l2_to_l1_"+chainId).Where("l2_block_number = ?", l2ToL1.L2BlockNumber.String()).Updates(map[string]interface{}{"l1_finalize_tx_hash": finalizedHash})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	return nil
}

func (l2l1 l2ToL1DB) UpdateTimeLeft(chainId string) error {
	result := l2l1.gorm.Table("l2_to_l1_"+chainId).Where("time_left > ?", 0).Updates(map[string]interface{}{"time_left": gorm.Expr("GREATEST(time_left - 1, 0)")})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}

	result = l2l1.gorm.Table("l2_to_l1_"+chainId).Where("time_left = ? and status = ?", 0, common3.L2ToL1InChallengePeriod).Updates(map[string]interface{}{"status": common3.L2ToL1ReadyForClaim})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	return nil
}

func (l2l1 l2ToL1DB) GetBlockNumberFromHash(chainId string, blockHash common.Hash) (*big.Int, error) {
	var l2BlockNumber uint64
	result := l2l1.gorm.Table("block_headers_"+chainId).Where("hash = ?", blockHash.String()).Select("number").Take(&l2BlockNumber)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return new(big.Int).SetUint64(l2BlockNumber), nil
}

func (l2l1 l2ToL1DB) L2ToL1TransactionWithdrawal(chainId string, withdrawalHash common.Hash) (*L2ToL1, error) {
	var l2ToL1Withdrawal L2ToL1
	nilHash := common.Hash{}
	if withdrawalHash.String() == nilHash.String() {
		return nil, nil
	}
	result := l2l1.gorm.Table("l2_to_l1_" + chainId).Where(&L2ToL1{WithdrawTransactionHash: withdrawalHash}).Take(&l2ToL1Withdrawal)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l2ToL1Withdrawal, nil
}

func (l2l1 l2ToL1DB) L2ToL1TransactionMsgHash(chainId string, messageHash common.Hash) (*L2ToL1, error) {
	var l2ToL1Withdrawal L2ToL1
	nilHash := common.Hash{}
	if messageHash.String() == nilHash.String() {
		return nil, nil
	}
	result := l2l1.gorm.Table("l2_to_l1_" + chainId).Where(&L2ToL1{MessageHash: messageHash}).Take(&l2ToL1Withdrawal)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l2ToL1Withdrawal, nil
}

func (l2l1 l2ToL1DB) L1LatestBlockHeader(l2chainId, destChainId string) (*common2.BlockHeader, error) {
	tableName := fmt.Sprintf("l2_to_l1_%s", l2chainId)
	blockHeaderSTableName := fmt.Sprintf("block_headers_%s", destChainId)
	l2Query := l2l1.gorm.Table(blockHeaderSTableName).Where("number = (?)", l2l1.gorm.Table(tableName).Select("MAX(l2_block_number)"))
	var l2Header common2.BlockHeader
	result := l2Query.Take(&l2Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l2Header, nil
	return l2l1.latestBlockHeaderWithChainId(l2chainId, destChainId)

}

func (l2l1 l2ToL1DB) L2LatestBlockHeader(chainId string) (*common2.BlockHeader, error) {
	return l2l1.latestBlockHeaderWithChainId(chainId, chainId)
}

func (l2l1 l2ToL1DB) latestBlockHeaderWithChainId(chainId, destChainId string) (*common2.BlockHeader, error) {
	tableName := fmt.Sprintf("l2_to_l1_%s", chainId)
	l2Query := l2l1.gorm.Where("timestamp = (?)", l2l1.gorm.Table(tableName).Select("MAX(timestamp)"))
	var l2Header common2.BlockHeader
	blockHeaderSTableName := fmt.Sprintf("block_headers_%s", destChainId)
	result := l2Query.Table(blockHeaderSTableName).Take(&l2Header)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &l2Header, nil
}

func (l2l1 l2ToL1DB) L1LatestFinalizedBlockHeader(chainId string) (*common2.BlockHeader, error) {
	return l2l1.latestFinalizedBlockHeaderWithChainId(chainId, "1")
}

func (l2l1 l2ToL1DB) L2LatestFinalizedBlockHeader(chainId string) (*common2.BlockHeader, error) {
	return l2l1.latestFinalizedBlockHeaderWithChainId(chainId, chainId)
}

func (l2l1 l2ToL1DB) latestFinalizedBlockHeaderWithChainId(chainId, destChainId string) (*common2.BlockHeader, error) {
	tableName := fmt.Sprintf("l2_to_l1_%s", chainId)
	l1Query := l2l1.gorm.Where("number = (?)", l2l1.gorm.Table(tableName).Where("status != (?)", 4).Select("MAX(l1_block_number)"))
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

func (l2l1 l2ToL1DB) UpdateL2ToL1ClaimedStatus(chainId string, l2L1 L2ToL1) error {

	result := l2l1.gorm.Table("l2_to_l1_"+chainId).Where("claimed_index = ?", l2L1.ClaimedIndex).Updates(map[string]interface{}{"status": l2L1.Status, "l1_finalize_tx_hash": l2L1.L1FinalizeTxHash})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	return nil
}
