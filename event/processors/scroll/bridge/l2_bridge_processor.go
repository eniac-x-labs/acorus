package bridge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/database"
)

func L2WithdrawETH(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	//l2WithdrawETHSig := abi.L2WithdrawETHSig
	//contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l2WithdrawETHSig}
	//withdrawEvents, err := db.ContractEvents.L2ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	//if err != nil {
	//	return err
	//}
	//l2ToL1s := make([]worker.L2ToL1, len(withdrawEvents))
	//transactionWithdrawals := make([]l1_l2.L2TransactionWithdrawal, len(withdrawEvents))
	//if len(withdrawEvents) > 0 {
	//	log.Info("detected eth transaction deposits", "size", len(transactionWithdrawals))
	//}
	//for i := range withdrawEvents {
	//	ethDepositTx := withdrawEvents[i]
	//	rlpLog := *ethDepositTx.RLPLog
	//	depositEvent := abi.DepositETH{}
	//
	//	unpackErr := utils.UnpackLog(abi.L1ETHGatewayABI, &depositEvent, "WithdrawETH", rlpLog)
	//	if unpackErr != nil {
	//		return unpackErr
	//	}
	//
	//	transactionWithdrawals[i] = l1_l2.L2TransactionWithdrawal{
	//
	//		Tx: l1_l2.Transaction{
	//			FromAddress: depositEvent.From,
	//			ToAddress:   depositEvent.To,
	//			ETHAmount:   depositEvent.Amount,
	//			ERC20Amount: big.NewInt(0),
	//			Data:        depositEvent.Data,
	//			Timestamp:   ethDepositTx.Timestamp,
	//		},
	//		GasLimit: big.NewInt(0),
	//	}
	//	l2ToL1s[i] = worker.L2ToL1{
	//		GUID:              uuid.New(),
	//		L2BlockHash:       rlpLog.BlockHash,
	//		L1TransactionHash: rlpLog.TxHash,
	//		L2TransactionHash: rlpLog.TxHash,
	//		L1TxOrigin:        common.Hash{},
	//		Status:            0,
	//		FromAddress:       depositEvent.From,
	//		ToAddress:         depositEvent.To,
	//		L1TokenAddress:    common.Address{},
	//		L2TokenAddress:    common.Address{},
	//		ETHAmount:         depositEvent.Amount,
	//		ERC20Amount:       big.NewInt(0),
	//		GasLimit:          big.NewInt(0),
	//		Timestamp:         int64(ethDepositTx.Timestamp),
	//		AssetType:         int64(common2.ETH),
	//		MsgHash:           common.Hash{},
	//	}
	//}
	//if len(l1ToL2s) > 0 {
	//	if err := db.BridgeTransactions.StoreL1TransactionDeposits(transactionDeposits); err != nil {
	//		return err
	//	}
	//	//marshal, _ := json.Marshal(l1ToL2s)
	//	if err := db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s); err != nil {
	//		return err
	//	}
	//}
	return nil
}

func L2WithdrawERC20() {

}

func L2WithdrawERC721() {

}

func L2WithdrawERC1155() {

}

func L2BatchWithdrawERC721() {

}

func L2BatchWithdrawERC1155() {

}

func L2SentMessageEvent() {

}

func L2RelayedMessageEvent() {

	// todo msgHash
	// todo

}
