package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/utils"

	bindings2 "github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
)

type LegacyBridgeEvent struct {
	Event                  *event.ContractEvent
	CrossDomainMessageHash *common.Hash
	FromAddress            common.Address
	ToAddress              common.Address
	ETHAmount              *big.Int
	ERC20Amount            *big.Int
	Data                   utils.Bytes
	LocalTokenAddress      common.Address
	RemoteTokenAddress     common.Address
	Timestamp              uint64
}

func L1StandardBridgeLegacyDepositInitiatedEvents(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) ([]LegacyBridgeEvent, error) {
	// The L1StandardBridge ABI contains the legacy events
	l1StandardBridgeAbi, err := bindings2.L1StandardBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	ethDepositEventAbi := l1StandardBridgeAbi.Events["ETHDepositInitiated"]
	erc20DepositEventAbi := l1StandardBridgeAbi.Events["ERC20DepositInitiated"]
	mntDepositEventAbi := l1StandardBridgeAbi.Events["MNTDepositInitiated"]

	// Grab both ETH & ERC20 Events
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: ethDepositEventAbi.ID}
	ethDepositEvents, err := db.ContractEvents.ContractEventsWithFilter("10", contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}
	contractEventFilter.EventSignature = erc20DepositEventAbi.ID
	erc20DepositEvents, err := db.ContractEvents.ContractEventsWithFilter("10", contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}
	contractEventFilter.EventSignature = mntDepositEventAbi.ID
	mntDepositEvents, err := db.ContractEvents.ContractEventsWithFilter("10", contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	deposits := make([]LegacyBridgeEvent, len(ethDepositEvents)+len(erc20DepositEvents)+len(mntDepositEvents))
	for i := range ethDepositEvents {
		bridgeEvent := bindings2.L1StandardBridgeETHDepositInitiated{Raw: *ethDepositEvents[i].RLPLog}
		err := UnpackLog(&bridgeEvent, &bridgeEvent.Raw, ethDepositEventAbi.Name, l1StandardBridgeAbi)
		if err != nil {
			return nil, err
		}
		deposits[i] = LegacyBridgeEvent{
			Event:              &ethDepositEvents[i],
			LocalTokenAddress:  predeploys.LegacyERC20ETHAddr,
			RemoteTokenAddress: predeploys.LegacyERC20ETHAddr,
			FromAddress:        bridgeEvent.From,
			ToAddress:          bridgeEvent.To,
			ETHAmount:          bridgeEvent.Amount,
			Data:               bridgeEvent.ExtraData,
			Timestamp:          ethDepositEvents[i].Timestamp,
		}
	}
	for i := range erc20DepositEvents {
		bridgeEvent := bindings2.L1StandardBridgeERC20DepositInitiated{Raw: *erc20DepositEvents[i].RLPLog}
		err := UnpackLog(&bridgeEvent, &bridgeEvent.Raw, erc20DepositEventAbi.Name, l1StandardBridgeAbi)
		if err != nil {
			return nil, err
		}
		deposits[len(ethDepositEvents)+i] = LegacyBridgeEvent{
			Event:              &erc20DepositEvents[i],
			LocalTokenAddress:  bridgeEvent.L1Token,
			RemoteTokenAddress: bridgeEvent.L2Token,
			FromAddress:        bridgeEvent.From,
			ToAddress:          bridgeEvent.To,
			ETHAmount:          bridgeEvent.Amount,
			Data:               bridgeEvent.ExtraData,
			Timestamp:          erc20DepositEvents[i].Timestamp,
		}
	}
	return deposits, nil
}

func L2StandardBridgeLegacyWithdrawalInitiatedEvents(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) ([]LegacyBridgeEvent, error) {
	l2StandardBridgeAbi, err := bindings2.L2StandardBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	withdrawalInitiatedEventAbi := l2StandardBridgeAbi.Events["WithdrawalInitiated"]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: withdrawalInitiatedEventAbi.ID}
	withdrawalEvents, err := db.ContractEvents.ContractEventsWithFilter("10", contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	withdrawals := make([]LegacyBridgeEvent, len(withdrawalEvents))
	for i := range withdrawalEvents {
		bridgeEvent := bindings2.L2StandardBridgeWithdrawalInitiated{Raw: *withdrawalEvents[i].RLPLog}
		err := UnpackLog(&bridgeEvent, &bridgeEvent.Raw, withdrawalInitiatedEventAbi.Name, l2StandardBridgeAbi)
		if err != nil {
			return nil, err
		}

		withdrawals[i] = LegacyBridgeEvent{
			Event:              &withdrawalEvents[i],
			LocalTokenAddress:  predeploys.LegacyERC20ETHAddr,
			RemoteTokenAddress: predeploys.LegacyERC20ETHAddr,
			FromAddress:        bridgeEvent.From,
			ToAddress:          bridgeEvent.To,
			ETHAmount:          bridgeEvent.Amount,
			Data:               bridgeEvent.ExtraData,
			Timestamp:          withdrawalEvents[i].Timestamp,
		}
	}

	return withdrawals, nil
}
