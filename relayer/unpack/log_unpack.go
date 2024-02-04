package unpack

import (
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

func MessageSent(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := MsgUnpack.ParseMessageSent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	fmt.Println(uEvent)
	return nil
}
func MessageClaimed(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := MsgUnpack.ParseMessageClaimed(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	fmt.Println(uEvent)
	return nil
}
func InitiateETH(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := L1Unpack.ParseInitiateETH(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	fmt.Println(uEvent)
	return nil
}
func InitiateWETH(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := L1Unpack.ParseInitiateWETH(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	fmt.Println(uEvent)
	return nil
}
func InitiateERC20(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := L1Unpack.ParseInitiateERC20(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	fmt.Println(uEvent)
	return nil
}
func FinalizeETH(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := L1Unpack.ParseFinalizeETH(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	fmt.Println(uEvent)
	return nil
}
func FinalizeWETH(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := L1Unpack.ParseFinalizeWETH(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	fmt.Println(uEvent)
	return nil
}
func FinalizeERC20(chainId string, event event.ContractEvent, db *database.DB) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := L1Unpack.ParseFinalizeERC20(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	fmt.Println(uEvent)
	return nil

}

func StarkingERC20Event(chainId string, event event.ContractEvent, db *database.DB) error {
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
		ChainId:     chainId,
		Status:      1,
		AssetType:   common2.ERC20,
		Timestamp:   event.Timestamp,
	}
	return db.StakeRecord.StoreStakingRecord(stake)
}

func StakingETHEvent(chainId string, event event.ContractEvent, db *database.DB) error {
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
		ChainId:     chainId,
		AssetType:   common2.ETH,
		Timestamp:   event.Timestamp,
	}
	return db.StakeRecord.StoreStakingRecord(stake)
}
func StakingWETHEvent(chainId string, event event.ContractEvent, db *database.DB) error {
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
		ChainId:     chainId,
		Status:      1,
		AssetType:   common2.WETH,
		Timestamp:   event.Timestamp,
	}
	return db.StakeRecord.StoreStakingRecord(stake)
}

func ClaimEvent(chainId string, event event.ContractEvent, db *database.DB) error {
	// todo what is this ?
	rlpLog := event.RLPLog
	claimEvent, unpackErr := L1Unpack.ParseClaimEvent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	fmt.Println(claimEvent)
	return nil
}
