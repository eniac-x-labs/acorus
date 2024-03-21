package bridge

import (
	"math/big"

	"github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/oldpolygonzkevmbridge"
	erc20 "github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/pol"
	"github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/polygonzkevmbridge"

	"github.com/ethereum/go-ethereum/common"

	common3 "github.com/cornerstone-labs/acorus/common"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/okx/utils"
)

func L1Deposit(l1ChainId, l2chainId string, polygonBridge *polygonzkevmbridge.Polygonzkevmbridge,
	eventInfo event.ContractEvent, db *database.DB) error {
	rlpLog := eventInfo.RLPLog
	d, unpackErr := polygonBridge.ParseBridgeEvent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	l1ToL2 := worker.L1ToL2{
		QueueIndex:        nil,
		L1TransactionHash: rlpLog.TxHash,
		L2TransactionHash: common.Hash{},
		L1TxOrigin:        common.Address{},
		L1BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
		Status:            common3.L1ToL2Pending,
		FromAddress:       common.Address{},
		ToAddress:         d.DestinationAddress,
		L1TokenAddress:    d.OriginAddress,
		L2TokenAddress:    common.Address{},
		GasLimit:          big.NewInt(0),
		Timestamp:         int64(eventInfo.Timestamp),
		MessageHash:       common.BigToHash(big.NewInt(int64(d.DepositCount))),
		ETHAmount:         big.NewInt(0),
		TokenAmounts:      "0",
	}
	// is erc20
	contractEventFilter := event.ContractEvent{TransactionHash: rlpLog.TxHash, EventSignature: utils.TokenTransferSignatureHash}
	transferEvent, dataErr := db.ContractEvents.ChainContractEventWithFilter(l1ChainId, contractEventFilter)
	if dataErr != nil {
		return dataErr
	}
	if transferEvent != nil {
		pol, newErc20Err := erc20.NewPol(transferEvent.ContractAddress, nil)
		if newErc20Err != nil {
			return newErc20Err
		}
		transfer, erc20UnpackErr := pol.ParseTransfer(*transferEvent.RLPLog)
		if erc20UnpackErr != nil {
			return erc20UnpackErr
		}
		l1ToL2.FromAddress = transfer.From
		l1ToL2.TokenAmounts = transfer.Value.String()
		l1ToL2.AssetType = int64(common2.ERC20)
	} else {
		l1ToL2.ETHAmount = d.Amount
		l1ToL2.AssetType = int64(common2.ETH)
	}

	l1ToL2s := make([]worker.L1ToL2, 0)
	l1ToL2s = append(l1ToL2s, l1ToL2)
	err := db.L1ToL2.StoreL1ToL2Transactions(l2chainId, l1ToL2s)
	return err
}

func L1ClaimedOld(l1chainId, l2chainId string, polygonBridge *oldpolygonzkevmbridge.Oldpolygonzkevmbridge,
	eventInfo event.ContractEvent, db *database.DB) error {
	rlpLog := eventInfo.RLPLog
	c, unpackErr := polygonBridge.ParseClaimEvent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	index := c.Index
	indexBig := new(big.Int).SetUint64(uint64(index))

	withDrawMessage := event.WithdrawFinalized{
		BlockNumber:              big.NewInt(int64(rlpLog.BlockNumber)),
		FinalizedTransactionHash: rlpLog.TxHash,
		WithdrawHash:             rlpLog.TxHash,
		MessageHash:              common.BigToHash(indexBig),
		Related:                  false,
		Timestamp:                eventInfo.Timestamp,
		ERC20Amount:              big.NewInt(0),
		ETHAmount:                big.NewInt(0),
	}
	// is erc20
	contractEventFilter := event.ContractEvent{TransactionHash: rlpLog.TxHash, EventSignature: utils.TokenTransferSignatureHash}
	transferEvent, err := db.ContractEvents.ChainContractEventWithFilter(l1chainId, contractEventFilter)
	if err != nil {
		return err
	}
	if transferEvent != nil {
		withDrawMessage.L1TokenAddress = transferEvent.ContractAddress
		withDrawMessage.ERC20Amount = c.Amount
	} else {
		withDrawMessage.ETHAmount = c.Amount
	}
	withdrawMessageList := make([]event.WithdrawFinalized, 0)
	withdrawMessageList = append(withdrawMessageList, withDrawMessage)
	err = db.WithdrawFinalized.StoreWithdrawFinalized(l2chainId, withdrawMessageList)
	return err
}

func L1Claimed(l1chainId, l2chainId string, polygonBridge *polygonzkevmbridge.Polygonzkevmbridge,
	eventInfo event.ContractEvent, db *database.DB) error {
	rlpLog := eventInfo.RLPLog
	c, unpackErr := polygonBridge.ParseClaimEvent(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}

	index := c.GlobalIndex
	_, _, localRootIndex, decodeErr := utils.DecodeGlobalIndex(index)
	if decodeErr != nil {
		return decodeErr
	}

	withDrawMessage := event.WithdrawFinalized{
		BlockNumber:              big.NewInt(int64(rlpLog.BlockNumber)),
		FinalizedTransactionHash: rlpLog.TxHash,
		WithdrawHash:             rlpLog.TxHash,
		MessageHash:              common.BigToHash(localRootIndex),
		Related:                  false,
		Timestamp:                eventInfo.Timestamp,
		ERC20Amount:              big.NewInt(0),
		ETHAmount:                big.NewInt(0),
	}
	// is erc20
	contractEventFilter := event.ContractEvent{TransactionHash: rlpLog.TxHash, EventSignature: utils.TokenTransferSignatureHash}
	transferEvent, err := db.ContractEvents.ChainContractEventWithFilter(l1chainId, contractEventFilter)
	if err != nil {
		return err
	}
	if transferEvent != nil {
		withDrawMessage.L1TokenAddress = transferEvent.ContractAddress
		withDrawMessage.ERC20Amount = c.Amount
	} else {
		withDrawMessage.ETHAmount = c.Amount
	}
	withdrawMessageList := make([]event.WithdrawFinalized, 0)
	withdrawMessageList = append(withdrawMessageList, withDrawMessage)
	err = db.WithdrawFinalized.StoreWithdrawFinalized(l2chainId, withdrawMessageList)
	return err
}
