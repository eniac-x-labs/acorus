package l1l2

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/database/business"
	"github.com/cornerstone-labs/acorus/database/event/l1-l2"
	bridgeDb "github.com/cornerstone-labs/acorus/database/event/l1-l2"
)

type L1L2Processor struct {
	logger log.Logger
	view   bridgeDb.BridgeTransfersView
}

func NewL1L2Processor(logger log.Logger, bv bridgeDb.BridgeTransfersView) L1L2Processor {
	return L1L2Processor{
		logger: logger,
		view:   bv,
	}
}

func ConstructL1ToL2(deposits []l1_l2.L1BridgeDepositWithTransactionHashes) []business.L1ToL2 {
	l1ToL2s := []business.L1ToL2{}
	for i := 0; i < len(deposits); i++ {
		l1ToL2 := business.L1ToL2{
			L1BlockNumber:     deposits[i].L1BlockHash,
			QueueIndex:        new(big.Int).SetUint64(1),
			L1TransactionHash: deposits[i].L1TransactionHash,
			L2TransactionHash: deposits[i].L2TransactionHash,
			L1TxOrigin:        common.HexToHash("10000"),
			Status:            1,
			FromAddress:       deposits[i].L1BridgeDeposit.Tx.FromAddress,
			ToAddress:         deposits[i].L1BridgeDeposit.Tx.ToAddress,
			L1TokenAddress:    deposits[i].L1BridgeDeposit.TokenPair.LocalTokenAddress,
			L2TokenAddress:    deposits[i].L1BridgeDeposit.TokenPair.RemoteTokenAddress,
			ETHAmount:         deposits[i].L1BridgeDeposit.Tx.ETHAmount,
			ERC20Amount:       deposits[i].L1BridgeDeposit.Tx.ERC20Amount,
			GasLimit:          new(big.Int).SetUint64(21000),
			Timestamp:         int64(deposits[i].L1BridgeDeposit.Tx.Timestamp),
		}
		l1ToL2s = append(l1ToL2s, l1ToL2)
	}
	return l1ToL2s
}

func ConstructL2ToL1(withdraws []l1_l2.L2BridgeWithdrawalWithTransactionHashes) []business.L2ToL1 {
	l2ToL1s := []business.L2ToL1{}
	for i := 0; i < len(withdraws); i++ {
		l2ToL1 := business.L2ToL1{
			L2BlockNumber:             withdraws[i].L2BlockHash,
			MsgNonce:                  new(big.Int).SetUint64(1),
			L2TransactionHash:         withdraws[i].L2TransactionHash,
			L2WithdrawTransactionHash: withdraws[i].L2BridgeWithdrawal.TransactionWithdrawalHash,
			L1ProveTxHash:             withdraws[i].ProvenL1TransactionHash,
			L1FinalizeTxHash:          withdraws[i].FinalizedL1TransactionHash,
			Status:                    1,
			FromAddress:               withdraws[i].L2BridgeWithdrawal.Tx.FromAddress,
			ETHAmount:                 withdraws[i].L2BridgeWithdrawal.Tx.ETHAmount,
			ERC20Amount:               withdraws[i].L2BridgeWithdrawal.Tx.ERC20Amount,
			GasLimit:                  new(big.Int).SetUint64(1),
			TimeLeft:                  new(big.Int).SetUint64(uint64(time.Hour * 24 * 7)),
			ToAddress:                 withdraws[i].L2BridgeWithdrawal.Tx.ToAddress,
			L1TokenAddress:            withdraws[i].L2BridgeWithdrawal.TokenPair.LocalTokenAddress,
			L2TokenAddress:            withdraws[i].L2BridgeWithdrawal.TokenPair.RemoteTokenAddress,
			Timestamp:                 int64(withdraws[i].L2BridgeWithdrawal.Tx.Timestamp),
		}
		l2ToL1s = append(l2ToL1s, l2ToL1)
	}

	return l2ToL1s
}
