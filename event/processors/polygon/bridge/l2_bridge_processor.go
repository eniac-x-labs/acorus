package bridge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"

	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	l1_l2 "github.com/cornerstone-labs/acorus/database/event/l1-l2"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/processors/polygon/abi"
	"github.com/cornerstone-labs/acorus/event/processors/polygon/utils"
)

func L2Withdraw(polygonBridge *abi.Polygonzkevmbridge, contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l2WithdrawETHSig := utils.WithdrawEventSignatureHash
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l2WithdrawETHSig}
	withdrawEvents, err := db.ContractEvents.L2ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l2ToL1s := make([]worker.L2ToL1, len(withdrawEvents))
	transactionWithdrawals := make([]l1_l2.L2TransactionWithdrawal, len(withdrawEvents))
	if len(withdrawEvents) > 0 {
		log.Info("detected transaction withdraw", "size", len(transactionWithdrawals))
	}
	for i := range withdrawEvents {
		WithdrawTx := withdrawEvents[i]
		rlpLog := *WithdrawTx.RLPLog
		withdraw := utils.Withdraw{}

		w, unpackErr := polygonBridge.ParseBridgeEvent(rlpLog)
		if unpackErr != nil {
			return unpackErr
		}

		withdraw.Amount = w.Amount
		withdraw.BlockNumber = rlpLog.BlockNumber
		withdraw.OriginalNetwork = uint(w.OriginNetwork)
		withdraw.DestinationAddress = w.DestinationAddress
		withdraw.DestinationNetwork = uint(w.DestinationNetwork)
		withdraw.OriginalAddress = w.OriginAddress
		withdraw.DepositCount = uint(w.DepositCount)
		withdraw.TxHash = rlpLog.TxHash
		withdraw.Metadata = w.Metadata
		withdraw.LeafType = w.LeafType

		transactionWithdrawals[i] = l1_l2.L2TransactionWithdrawal{
			WithdrawalHash:       WithdrawTx.TransactionHash,
			InitiatedL2EventGUID: WithdrawTx.GUID,
			Tx: l1_l2.Transaction{
				FromAddress: withdraw.DestinationAddress,
				ToAddress:   withdraw.DestinationAddress,
				ETHAmount:   big.NewInt(0),
				ERC20Amount: big.NewInt(0),
				Data:        withdraw.Metadata,
				Timestamp:   WithdrawTx.Timestamp,
			},
			GasLimit: big.NewInt(0),
		}
		l2ToL1s[i] = worker.L2ToL1{
			GUID:                      uuid.New(),
			L2BlockHash:               rlpLog.BlockHash,
			L1FinalizeTxHash:          common.Hash{},
			L2TransactionHash:         rlpLog.TxHash,
			L2WithdrawTransactionHash: common.Hash{},
			L1ProveTxHash:             common.Hash{},
			Status:                    0,
			TimeLeft:                  big.NewInt(0),
			FromAddress:               withdraw.DestinationAddress,
			ToAddress:                 withdraw.DestinationAddress,
			L1TokenAddress:            withdraw.OriginalAddress,
			L2TokenAddress:            common.Address{},
			ETHAmount:                 big.NewInt(0),
			ERC20Amount:               big.NewInt(0),
			GasLimit:                  big.NewInt(0),
			Timestamp:                 int64(WithdrawTx.Timestamp),
			AssetType:                 int64(common2.ETH),
			MsgNonce:                  big.NewInt(0),
			MsgHash:                   common.Hash{},
		}
		if withdraw.OriginalAddress.String() == utils.L1_ETH {
			transactionWithdrawals[i].Tx.ETHAmount = withdraw.Amount
			l2ToL1s[i].ETHAmount = withdraw.Amount
			l2ToL1s[i].AssetType = int64(common2.ETH)
		} else {
			transactionWithdrawals[i].Tx.ERC20Amount = withdraw.Amount
			l2ToL1s[i].ERC20Amount = withdraw.Amount
			l2ToL1s[i].AssetType = int64(common2.ERC20)
		}
	}
	if len(l2ToL1s) > 0 {
		if err := db.BridgeTransactions.StoreL2TransactionWithdrawals(transactionWithdrawals); err != nil {
			return err
		}
		if err := db.L2ToL1.StoreL2ToL1Transactions(l2ToL1s); err != nil {
			return err
		}
	}
	return nil
}
