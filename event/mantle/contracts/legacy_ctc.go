package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/utils"
	legacy_bindings "github.com/cornerstone-labs/acorus/event/mantle/op-bindings/legacy-bindings"
)

type LegacyCTCDepositEvent struct {
	Event       *event.ContractEvent
	FromAddress common.Address
	ToAddress   common.Address
	ETHAmount   *big.Int
	ERC20Amount *big.Int
	Data        utils.Bytes
	Timestamp   uint64
	TxHash      common.Hash
	GasLimit    *big.Int
}

func LegacyCTCDepositEvents(contractAddress common.Address, db *database.DB,
	fromHeight, toHeight *big.Int, l1ChainId, l2ChainId string) ([]LegacyCTCDepositEvent, error) {
	ctcAbi, err := legacy_bindings.CanonicalTransactionChainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	transactionEnqueuedEventAbi := ctcAbi.Events["TransactionEnqueued"]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: transactionEnqueuedEventAbi.ID}
	events, err := db.ContractEvents.ContractEventsWithFilter(l1ChainId, contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	ctcTxDeposits := make([]LegacyCTCDepositEvent, len(events))
	for i := range events {
		txEnqueued := legacy_bindings.CanonicalTransactionChainTransactionEnqueued{Raw: *events[i].RLPLog}
		err = UnpackLog(&txEnqueued, events[i].RLPLog, transactionEnqueuedEventAbi.Name, ctcAbi)
		if err != nil {
			return nil, err
		}

		ctcTxDeposits[i] = LegacyCTCDepositEvent{
			Event:       &events[i],
			GasLimit:    txEnqueued.GasLimit,
			TxHash:      types.NewTransaction(0, txEnqueued.Target, bigint.Zero, txEnqueued.GasLimit.Uint64(), nil, txEnqueued.Data).Hash(),
			FromAddress: txEnqueued.L1TxOrigin,
			ToAddress:   txEnqueued.Target,
			ETHAmount:   bigint.Zero,
			Data:        txEnqueued.Data,
			Timestamp:   events[i].Timestamp,
			ERC20Amount: bigint.Zero,
		}
	}
	return ctcTxDeposits, nil
}
