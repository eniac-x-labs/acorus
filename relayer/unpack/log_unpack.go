package unpack

import (
	"encoding/json"
	"fmt"
	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/relayer"
	"github.com/cornerstone-labs/acorus/relayer/bindings"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var (
	L1Unpack, _  = bindings.NewL1PoolManagerFilterer(common.Address{}, nil)
	MsgUnpack, _ = bindings.NewIMessageManagerFilterer(common.Address{}, nil)
)

func MessageSent(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := MsgUnpack.ParseMessageSent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	msgHash := relayer.BridgeMsgHash{
		TxHash:   rlpLog.TxHash,
		MsgHash:  uEvent.MessageHash,
		Fee:      uEvent.Fee,
		MsgNonce: uEvent.Nonce,
	}
	return db.BridgeMsgHash.StoreBridgeMsgHash(msgHash)
}
func MessageClaimed(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := MsgUnpack.ParseMessageClaimed(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	msgClaim := relayer.BridgeClaimed{
		TxHash:        rlpLog.TxHash,
		MsgHash:       uEvent.MessageHash,
		BlockNumber:   big.NewInt(int64(rlpLog.BlockNumber)),
		Timestamp:     event.Timestamp,
		TokenRelation: false,
	}
	return db.BridgeClaim.StoreBridgeClaim(msgClaim)
}
func InitiateETH(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := L1Unpack.ParseInitiateETH(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	bridgeRecord := relayer.BridgeRecords{
		SourceChainId:      uEvent.SourceChainId.String(),
		DestChainId:        uEvent.DestChainId.String(),
		From:               uEvent.From,
		To:                 uEvent.To,
		Amount:             uEvent.Value,
		AssetType:          common2.ETH,
		SourceTokenAddress: common.Address{},
		DestTokenAddress:   common.Address{},
		MsgHash:            common.Hash{},
		MsgSentTimestamp:   event.Timestamp,
		Status:             0,
		SourceTxHash:       rlpLog.TxHash,
		SourceBlockNumber:  big.NewInt(int64(rlpLog.BlockNumber)),
	}

	// todo 调用rpc，发送代币给to地址
	msgData, marsharlErr := json.Marshal(bridgeRecord)
	if marsharlErr != nil {
		return marsharlErr
	}

	bridgeMsgSent := relayer.BridgeMsgSent{
		TxHash: rlpLog.TxHash,
		Data:   string(msgData),
	}
	return db.BridgeMsgSent.StoreBridgeMsgSent(bridgeMsgSent)
}
func InitiateWETH(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := L1Unpack.ParseInitiateWETH(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	bridgeRecord := relayer.BridgeRecords{
		SourceChainId:      uEvent.SourceChainId.String(),
		DestChainId:        uEvent.DestChainId.String(),
		From:               uEvent.From,
		To:                 uEvent.To,
		Amount:             uEvent.Value,
		AssetType:          common2.WETH,
		SourceTokenAddress: common.Address{},
		DestTokenAddress:   common.Address{},
		MsgHash:            common.Hash{},
		MsgSentTimestamp:   event.Timestamp,
		Status:             0,
		SourceTxHash:       rlpLog.TxHash,
		SourceBlockNumber:  big.NewInt(int64(rlpLog.BlockNumber)),
	}

	// todo 调用rpc，发送代币给to地址
	msgData, marsharlErr := json.Marshal(bridgeRecord)
	if marsharlErr != nil {
		return marsharlErr
	}
	bridgeMsgSent := relayer.BridgeMsgSent{
		TxHash: rlpLog.TxHash,
		Data:   string(msgData),
	}
	return db.BridgeMsgSent.StoreBridgeMsgSent(bridgeMsgSent)
}
func InitiateERC20(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := L1Unpack.ParseInitiateERC20(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	bridgeRecord := relayer.BridgeRecords{
		SourceChainId:      uEvent.SourceChainId.String(),
		DestChainId:        uEvent.DestChainId.String(),
		From:               uEvent.From,
		To:                 uEvent.To,
		Amount:             uEvent.Value,
		AssetType:          common2.ERC20,
		SourceTokenAddress: uEvent.ERC20Address,
		DestTokenAddress:   common.Address{},
		MsgHash:            common.Hash{},
		MsgSentTimestamp:   event.Timestamp,
		Status:             0,
		SourceTxHash:       rlpLog.TxHash,
		SourceBlockNumber:  big.NewInt(int64(rlpLog.BlockNumber)),
	}

	// todo 调用rpc，发送代币给to地址
	msgData, marsharlErr := json.Marshal(bridgeRecord)
	if marsharlErr != nil {
		return marsharlErr
	}
	bridgeMsgSent := relayer.BridgeMsgSent{
		TxHash: rlpLog.TxHash,
		Data:   string(msgData),
	}
	return db.BridgeMsgSent.StoreBridgeMsgSent(bridgeMsgSent)
}
func FinalizeETH(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	//uEvent, unpackErr := L1Unpack.ParseFinalizeETH(*rlpLog)
	//if unpackErr != nil {
	//	return unpackErr
	//}
	bridgeFinalize := relayer.BridgeFinalize{
		TxHash:    rlpLog.TxHash,
		DestToken: common.Address{},
	}
	return db.BridgeFinalize.StoreBridgeFinalize(bridgeFinalize)
}
func FinalizeWETH(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	//uEvent, unpackErr := L1Unpack.ParseFinalizeWETH(*rlpLog)
	//if unpackErr != nil {
	//	return unpackErr
	//}
	bridgeFinalize := relayer.BridgeFinalize{
		TxHash:    rlpLog.TxHash,
		DestToken: common.Address{},
	}
	return db.BridgeFinalize.StoreBridgeFinalize(bridgeFinalize)
}
func FinalizeERC20(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := L1Unpack.ParseFinalizeERC20(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	bridgeFinalize := relayer.BridgeFinalize{
		TxHash:    rlpLog.TxHash,
		DestToken: uEvent.ERC20Address,
	}
	return db.BridgeFinalize.StoreBridgeFinalize(bridgeFinalize)
}

func StarkingERC20Event(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	erc20Event, unpackErr := L1Unpack.ParseStarkingERC20Event(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	stake := relayer.StakingRecord{
		TxHash:      rlpLog.TxHash,
		BlockNumber: big.NewInt(int64(rlpLog.BlockNumber)),
		UserAddress: erc20Event.User,
		Token:       erc20Event.Token,
		Amount:      erc20Event.Amount,
		Status:      1,
		TxType:      1,
		AssetType:   common2.ERC20,
		Timestamp:   event.Timestamp,
	}
	return db.StakeRecord.StoreStakingRecord(stake)
}

func StakingETHEvent(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	ethEvent, unpackErr := L1Unpack.ParseStakingETHEvent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	stake := relayer.StakingRecord{
		TxHash:      rlpLog.TxHash,
		BlockNumber: big.NewInt(int64(rlpLog.BlockNumber)),
		UserAddress: ethEvent.User,
		Token:       common.Address{},
		Amount:      ethEvent.Amount,
		Status:      1,
		TxType:      1,
		AssetType:   common2.ETH,
		Timestamp:   event.Timestamp,
	}
	return db.StakeRecord.StoreStakingRecord(stake)
}
func StakingWETHEvent(event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	wethEvent, unpackErr := L1Unpack.ParseStakingWETHEvent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	stake := relayer.StakingRecord{
		TxHash:      rlpLog.TxHash,
		BlockNumber: big.NewInt(int64(rlpLog.BlockNumber)),
		UserAddress: wethEvent.User,
		Token:       common.HexToAddress(global_const.WEthAddress),
		Amount:      wethEvent.Amount,
		Status:      1,
		TxType:      1,
		AssetType:   common2.WETH,
		Timestamp:   event.Timestamp,
	}
	return db.StakeRecord.StoreStakingRecord(stake)
}

func ClaimEvent(event event.ContractEvent, db *database.DB) error {
	// todo what is this ?
	rlpLog := event.RLPLog
	claimEvent, unpackErr := L1Unpack.ParseClaimEvent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	fmt.Println(claimEvent)
	return nil
}
