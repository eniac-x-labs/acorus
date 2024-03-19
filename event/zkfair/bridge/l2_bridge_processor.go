package bridge

//
//import (
//	erc20 "github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/pol"
//	"math/big"
//
//	"github.com/ethereum/go-ethereum/common"
//
//	common3 "github.com/cornerstone-labs/acorus/common"
//	"github.com/cornerstone-labs/acorus/database"
//	common2 "github.com/cornerstone-labs/acorus/database/common"
//	"github.com/cornerstone-labs/acorus/database/event"
//	"github.com/cornerstone-labs/acorus/database/worker"
//	"github.com/cornerstone-labs/acorus/event/polygon/utils"
//	"github.com/cornerstone-labs/acorus/event/zkfair/bindings"
//)
//
//func L2Withdraw(l1ChainId, l2chainId string,
//	polygonBridge *bindings.Polygonzkevmbridge,
//	eventInfo event.ContractEvent, db *database.DB) error {
//	rlpLog := eventInfo.RLPLog
//	w, unpackErr := polygonBridge.ParseBridgeEvent(*rlpLog)
//	if unpackErr != nil {
//		return unpackErr
//	}
//	l2ToL1 := worker.L2ToL1{
//		L1FinalizeTxHash:  common.Hash{},
//		L2TransactionHash: rlpLog.TxHash,
//		L1ProveTxHash:     common.Hash{},
//		Status:            common3.L2ToL1Pending,
//		TimeLeft:          big.NewInt(0),
//		FromAddress:       common.Address{},
//		ToAddress:         w.DestinationAddress,
//		L2BlockNumber:     big.NewInt(int64(rlpLog.BlockNumber)),
//		L1TokenAddress:    w.OriginAddress,
//		L2TokenAddress:    common.Address{},
//		Timestamp:         int64(eventInfo.Timestamp),
//		MessageHash:       common.BigToHash(big.NewInt(int64(w.DepositCount))),
//		ClaimedIndex:      int64(w.DepositCount),
//		MsgNonce:          big.NewInt(0),
//		GasLimit:          big.NewInt(0),
//		ETHAmount:         big.NewInt(0),
//		TokenAmounts:      "0",
//	}
//
//	contractEventFilter := event.ContractEvent{TransactionHash: rlpLog.TxHash, EventSignature: utils.TokenTransferSignatureHash}
//	transferEvent, dataErr := db.ContractEvents.ChainContractEventWithFilter(l2chainId, contractEventFilter)
//	if dataErr != nil {
//		return dataErr
//	}
//	if transferEvent != nil {
//		pol, newErc20Err := erc20.NewPol(transferEvent.ContractAddress, nil)
//		if newErc20Err != nil {
//			return newErc20Err
//		}
//		transfer, erc20UnpackErr := pol.ParseTransfer(*transferEvent.RLPLog)
//		if erc20UnpackErr != nil {
//			return erc20UnpackErr
//		}
//		l2ToL1.L2TokenAddress = transferEvent.ContractAddress
//		l2ToL1.TokenAmounts = transfer.Value.String()
//		l2ToL1.FromAddress = transfer.From
//		l2ToL1.AssetType = int64(common2.ERC20)
//	} else {
//		l2ToL1.ETHAmount = w.Amount
//		l2ToL1.AssetType = int64(common2.ETH)
//	}
//
//	l2ToL1s := make([]worker.L2ToL1, 0)
//	l2ToL1s = append(l2ToL1s, l2ToL1)
//	err := db.L2ToL1.StoreL2ToL1Transactions(l2chainId, l2ToL1s)
//	return err
//}
//
//func L2Claimed(l1ChainId, l2chainId string, polygonBridge *bindings.Polygonzkevmbridge,
//	eventInfo event.ContractEvent, db *database.DB) error {
//	rlpLog := eventInfo.RLPLog
//	c, unpackErr := polygonBridge.ParseClaimEvent(*rlpLog)
//	if unpackErr != nil {
//		return unpackErr
//	}
//
//	index := c.GlobalIndex
//	_, _, localRootIndex, decodeErr := utils.DecodeGlobalIndex(index)
//	if decodeErr != nil {
//		return decodeErr
//	}
//
//	relayMessage := event.RelayMessage{
//		BlockNumber:          big.NewInt(int64(rlpLog.BlockNumber)),
//		RelayTransactionHash: rlpLog.TxHash,
//		MessageHash:          common.BigToHash(localRootIndex),
//		Related:              false,
//		Timestamp:            eventInfo.Timestamp,
//		ERC20Amount:          big.NewInt(0),
//		ETHAmount:            big.NewInt(0),
//	}
//	// is erc20
//	contractEventFilter := event.ContractEvent{TransactionHash: rlpLog.TxHash, EventSignature: utils.TokenTransferSignatureHash}
//	transferEvent, err := db.ContractEvents.ChainContractEventWithFilter(l2chainId, contractEventFilter)
//	if err != nil {
//		return err
//	}
//	if transferEvent != nil {
//		relayMessage.L2TokenAddress = transferEvent.ContractAddress
//		relayMessage.ERC20Amount = c.Amount
//	} else {
//		relayMessage.ETHAmount = c.Amount
//	}
//	relayMessageList := make([]event.RelayMessage, 0)
//	relayMessageList = append(relayMessageList, relayMessage)
//	err = db.RelayMessage.StoreRelayMessage(l2chainId, relayMessageList)
//	return err
//}
