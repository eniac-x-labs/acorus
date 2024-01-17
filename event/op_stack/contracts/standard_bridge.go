package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/utils"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
)

type StandardBridgeInitiatedEvent struct {
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

type StandardBridgeFinalizedEvent struct {
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

// StandardBridgeInitiatedEvents extracts all initiated bridge events from the contracts that follow the StandardBridge ABI. The
// correlated CrossDomainMessenger nonce is also parsed from the associated messenger events.
func StandardBridgeInitiatedEvents(chainSelector string, contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) ([]StandardBridgeInitiatedEvent, error) {
	ethBridgeInitiatedEvents, err := _standardBridgeInitiatedEvents[bindings.StandardBridgeETHBridgeInitiated](contractAddress, chainSelector, db, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	erc20BridgeInitiatedEvents, err := _standardBridgeInitiatedEvents[bindings.StandardBridgeERC20BridgeInitiated](contractAddress, chainSelector, db, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	return append(ethBridgeInitiatedEvents, erc20BridgeInitiatedEvents...), nil
}

// StandardBridgeFinalizedEvents extracts all finalization bridge events from the contracts that follow the StandardBridge ABI. The
// correlated CrossDomainMessenger nonce is also parsed by looking at the parameters of the corresponding relayMessage transaction data.
func StandardBridgeFinalizedEvents(chainSelector string, contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) ([]StandardBridgeFinalizedEvent, error) {
	ethBridgeFinalizedEvents, err := _standardBridgeFinalizedEvents[bindings.StandardBridgeETHBridgeFinalized](contractAddress, chainSelector, db, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	erc20BridgeFinalizedEvents, err := _standardBridgeFinalizedEvents[bindings.StandardBridgeERC20BridgeFinalized](contractAddress, chainSelector, db, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	return append(ethBridgeFinalizedEvents, erc20BridgeFinalizedEvents...), nil
}

func _standardBridgeInitiatedEvents[BridgeEventType bindings.StandardBridgeETHBridgeInitiated | bindings.StandardBridgeERC20BridgeInitiated](
	contractAddress common.Address, chainSelector string, db *database.DB, fromHeight, toHeight *big.Int,
) ([]StandardBridgeInitiatedEvent, error) {
	standardBridgeAbi, err := bindings.StandardBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	var eventType BridgeEventType
	var eventName string
	switch any(eventType).(type) {
	case bindings.StandardBridgeETHBridgeInitiated:
		eventName = "ETHBridgeInitiated"
	case bindings.StandardBridgeERC20BridgeInitiated:
		eventName = "ERC20BridgeInitiated"
	default:
		panic("should not be here")
	}

	initiatedBridgeEventAbi := standardBridgeAbi.Events[eventName]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: initiatedBridgeEventAbi.ID}
	initiatedBridgeEvents, err := db.ContractEvents.ContractEventsWithFilter("10", contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	standardBridgeInitiatedEvents := make([]StandardBridgeInitiatedEvent, len(initiatedBridgeEvents))
	for i := range initiatedBridgeEvents {
		switch any(eventType).(type) {
		case bindings.StandardBridgeETHBridgeInitiated:
			ethBridge := bindings.StandardBridgeETHBridgeInitiated{Raw: *initiatedBridgeEvents[i].RLPLog}
			err := UnpackLog(&ethBridge, initiatedBridgeEvents[i].RLPLog, eventName, standardBridgeAbi)
			if err != nil {
				return nil, err
			}
			standardBridgeInitiatedEvents[i] = StandardBridgeInitiatedEvent{
				Event:              &initiatedBridgeEvents[i],
				LocalTokenAddress:  common.Address{},
				RemoteTokenAddress: common.Address{},
				FromAddress:        ethBridge.From,
				ToAddress:          ethBridge.To,
				ETHAmount:          ethBridge.Amount,
				ERC20Amount:        bigint.Zero,
				Data:               ethBridge.ExtraData,
				Timestamp:          initiatedBridgeEvents[i].Timestamp,
			}
		case bindings.StandardBridgeERC20BridgeInitiated:
			erc20Bridge := bindings.StandardBridgeERC20BridgeInitiated{Raw: *initiatedBridgeEvents[i].RLPLog}
			err := UnpackLog(&erc20Bridge, initiatedBridgeEvents[i].RLPLog, eventName, standardBridgeAbi)
			if err != nil {
				return nil, err
			}
			standardBridgeInitiatedEvents[i] = StandardBridgeInitiatedEvent{
				Event:              &initiatedBridgeEvents[i],
				LocalTokenAddress:  erc20Bridge.LocalToken,
				RemoteTokenAddress: erc20Bridge.RemoteToken,
				FromAddress:        erc20Bridge.From,
				ToAddress:          erc20Bridge.To,
				ETHAmount:          bigint.Zero,
				ERC20Amount:        erc20Bridge.Amount,
				Data:               erc20Bridge.ExtraData,
				Timestamp:          initiatedBridgeEvents[i].Timestamp,
			}
		}
	}
	return standardBridgeInitiatedEvents, nil
}

func _standardBridgeFinalizedEvents[BridgeEventType bindings.StandardBridgeETHBridgeFinalized | bindings.StandardBridgeERC20BridgeFinalized](
	contractAddress common.Address, chainSelector string, db *database.DB, fromHeight, toHeight *big.Int,
) ([]StandardBridgeFinalizedEvent, error) {
	standardBridgeAbi, err := bindings.StandardBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	var eventType BridgeEventType
	var eventName string
	switch any(eventType).(type) {
	case bindings.StandardBridgeETHBridgeFinalized:
		eventName = "ETHBridgeFinalized"
	case bindings.StandardBridgeERC20BridgeFinalized:
		eventName = "ERC20BridgeFinalized"
	default:
		panic("should not be here")
	}

	bridgeFinalizedEventAbi := standardBridgeAbi.Events[eventName]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: bridgeFinalizedEventAbi.ID}
	bridgeFinalizedEvents, err := db.ContractEvents.ContractEventsWithFilter("10", contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	standardBridgeFinalizedEvents := make([]StandardBridgeFinalizedEvent, len(bridgeFinalizedEvents))
	for i := range bridgeFinalizedEvents {

		switch any(eventType).(type) {
		case bindings.StandardBridgeETHBridgeFinalized:
			ethBridge := bindings.StandardBridgeETHBridgeFinalized{Raw: *bridgeFinalizedEvents[i].RLPLog}
			err := UnpackLog(&ethBridge, bridgeFinalizedEvents[i].RLPLog, eventName, standardBridgeAbi)
			if err != nil {
				return nil, err
			}

			standardBridgeFinalizedEvents[i] = StandardBridgeFinalizedEvent{
				Event:              &bridgeFinalizedEvents[i],
				LocalTokenAddress:  common.Address{},
				RemoteTokenAddress: common.Address{},
				FromAddress:        ethBridge.From,
				ToAddress:          ethBridge.To,
				ETHAmount:          ethBridge.Amount,
				ERC20Amount:        bigint.Zero,
				Data:               ethBridge.ExtraData,
				Timestamp:          bridgeFinalizedEvents[i].Timestamp,
			}

		case bindings.StandardBridgeERC20BridgeFinalized:
			erc20Bridge := bindings.StandardBridgeERC20BridgeFinalized{Raw: *bridgeFinalizedEvents[i].RLPLog}
			err := UnpackLog(&erc20Bridge, bridgeFinalizedEvents[i].RLPLog, eventName, standardBridgeAbi)
			if err != nil {
				return nil, err
			}

			standardBridgeFinalizedEvents[i] = StandardBridgeFinalizedEvent{
				Event:              &bridgeFinalizedEvents[i],
				LocalTokenAddress:  erc20Bridge.LocalToken,
				RemoteTokenAddress: erc20Bridge.RemoteToken,
				FromAddress:        erc20Bridge.From,
				ToAddress:          erc20Bridge.To,
				ETHAmount:          bigint.Zero,
				ERC20Amount:        erc20Bridge.Amount,
				Data:               erc20Bridge.ExtraData,
				Timestamp:          bridgeFinalizedEvents[i].Timestamp,
			}
		default:
			panic("should not be here")
		}
	}
	return standardBridgeFinalizedEvents, nil
}
