package contracts

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/utils"

	common2 "github.com/cornerstone-labs/acorus/event/base/common"
	common3 "github.com/cornerstone-labs/acorus/event/base/common"
	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
)

type OptimismPortalTransactionDepositEvent struct {
	Event       *event.ContractEvent
	DepositTx   *common2.DepositTx
	FromAddress common.Address
	ToAddress   common.Address
	ETHAmount   *big.Int
	ERC20Amount *big.Int
	Data        utils.Bytes
	Timestamp   uint64
	GasLimit    *big.Int
}

type OptimismPortalWithdrawalProvenEvent struct {
	*bindings.OptimismPortalWithdrawalProven
	Event *event.ContractEvent
}

type OptimismPortalWithdrawalFinalizedEvent struct {
	*bindings.OptimismPortalWithdrawalFinalized
	Event *event.ContractEvent
}

type OptimismPortalProvenWithdrawal struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2OutputIndex *big.Int
}

func OptimismPortalTransactionDepositEvents(contractAddress common.Address, chainId string,
	db *database.DB, fromHeight, toHeight *big.Int) ([]OptimismPortalTransactionDepositEvent, error) {
	optimismPortalAbi, err := bindings.OptimismPortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	transactionDepositedEventAbi := optimismPortalAbi.Events["TransactionDeposited"]
	if transactionDepositedEventAbi.ID != common3.DepositEventABIHash {
		return nil, errors.New("op-node DepositEventABIHash & optimism portal TransactionDeposited ID mismatch")
	}
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: transactionDepositedEventAbi.ID}
	transactionDepositEvents, err := db.ContractEvents.ContractEventsWithFilter(chainId, contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	optimismPortalTxDeposits := make([]OptimismPortalTransactionDepositEvent, len(transactionDepositEvents))
	for i := range transactionDepositEvents {
		depositTx, err := common2.UnmarshalDepositLogEvent(transactionDepositEvents[i].RLPLog)
		if err != nil {
			return nil, err
		}

		txDeposit := bindings.OptimismPortalTransactionDeposited{Raw: *transactionDepositEvents[i].RLPLog}
		err = UnpackLog(&txDeposit, transactionDepositEvents[i].RLPLog, transactionDepositedEventAbi.Name, optimismPortalAbi)
		if err != nil {
			return nil, err
		}

		mint := depositTx.Mint
		if mint == nil {
			mint = bigint.Zero
		}
		dEth := depositTx.EthValue
		if dEth == nil {
			dEth = bigint.Zero
		}

		optimismPortalTxDeposits[i] = OptimismPortalTransactionDepositEvent{
			Event:       &transactionDepositEvents[i],
			DepositTx:   depositTx,
			GasLimit:    new(big.Int).SetUint64(depositTx.Gas),
			FromAddress: txDeposit.From,
			ToAddress:   txDeposit.To,
			ETHAmount:   dEth,
			Data:        depositTx.Data,
			Timestamp:   transactionDepositEvents[i].Timestamp,
			ERC20Amount: depositTx.Value,
		}
	}
	return optimismPortalTxDeposits, nil
}

func OptimismPortalWithdrawalProvenEvents(contractAddress common.Address, db *database.DB,
	fromHeight, toHeight *big.Int, l1ChainId, l2ChainId string) ([]OptimismPortalWithdrawalProvenEvent, error) {
	optimismPortalAbi, err := bindings.OptimismPortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	withdrawalProvenEventAbi := optimismPortalAbi.Events["WithdrawalProven"]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: withdrawalProvenEventAbi.ID}
	withdrawalProvenEvents, err := db.ContractEvents.ContractEventsWithFilter(l1ChainId, contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	provenWithdrawals := make([]OptimismPortalWithdrawalProvenEvent, len(withdrawalProvenEvents))
	for i := range withdrawalProvenEvents {
		withdrawalProven := bindings.OptimismPortalWithdrawalProven{Raw: *withdrawalProvenEvents[i].RLPLog}
		err := UnpackLog(&withdrawalProven, withdrawalProvenEvents[i].RLPLog, withdrawalProvenEventAbi.Name, optimismPortalAbi)
		if err != nil {
			return nil, err
		}

		provenWithdrawals[i] = OptimismPortalWithdrawalProvenEvent{
			OptimismPortalWithdrawalProven: &withdrawalProven,
			Event:                          &withdrawalProvenEvents[i],
		}
	}

	return provenWithdrawals, nil
}

func OptimismPortalWithdrawalFinalizedEvents(contractAddress common.Address, db *database.DB,
	fromHeight, toHeight *big.Int, l1ChainId, l2ChainId string) ([]OptimismPortalWithdrawalFinalizedEvent, error) {
	optimismPortalAbi, err := bindings.OptimismPortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	withdrawalFinalizedEventAbi := optimismPortalAbi.Events["WithdrawalFinalized"]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: withdrawalFinalizedEventAbi.ID}
	withdrawalFinalizedEvents, err := db.ContractEvents.ContractEventsWithFilter(l1ChainId, contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	finalizedWithdrawals := make([]OptimismPortalWithdrawalFinalizedEvent, len(withdrawalFinalizedEvents))
	for i := range withdrawalFinalizedEvents {
		withdrawalFinalized := bindings.OptimismPortalWithdrawalFinalized{Raw: *withdrawalFinalizedEvents[i].RLPLog}
		err := UnpackLog(&withdrawalFinalized, withdrawalFinalizedEvents[i].RLPLog, withdrawalFinalizedEventAbi.Name, optimismPortalAbi)
		if err != nil {
			return nil, err
		}
		finalizedWithdrawals[i] = OptimismPortalWithdrawalFinalizedEvent{
			OptimismPortalWithdrawalFinalized: &withdrawalFinalized,
			Event:                             &withdrawalFinalizedEvents[i],
		}
	}
	return finalizedWithdrawals, nil
}
