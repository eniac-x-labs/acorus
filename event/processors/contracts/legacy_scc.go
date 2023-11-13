package contracts

import (
	"encoding/hex"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/business"
	"github.com/cornerstone-labs/acorus/database/event"
	legacy_bindings "github.com/cornerstone-labs/acorus/event/op-bindings/legacy-bindings"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func LegacySCCBatchAppendedEvent(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) ([]business.StateRoot, error) {
	sccAbi, err := legacy_bindings.StateCommitmentChainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	stateBatchAppendedEventAbi := sccAbi.Events["StateBatchAppended"]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: stateBatchAppendedEventAbi.ID}
	stateBatchAppendedEvents, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}
	stateRoots := make([]business.StateRoot, len(stateBatchAppendedEvents))
	for i := range stateBatchAppendedEvents {
		stateBatchAppended := legacy_bindings.StateCommitmentChainStateBatchAppended{Raw: *stateBatchAppendedEvents[i].RLPLog}
		err = UnpackLog(&stateBatchAppended, stateBatchAppendedEvents[i].RLPLog, stateBatchAppendedEventAbi.Name, sccAbi)
		if err != nil {
			return nil, err
		}

		stateRoots[i] = business.StateRoot{
			TransactionHash:   stateBatchAppendedEvents[i].TransactionHash,
			OutputRoot:        hex.EncodeToString(stateBatchAppended.BatchRoot[:]),
			BatchSize:         stateBatchAppended.BatchSize,
			OutputIndex:       stateBatchAppended.BatchIndex,
			PrevTotalElements: stateBatchAppended.PrevTotalElements,
			Timestamp:         stateBatchAppendedEvents[i].Timestamp,
		}
	}
	return stateRoots, nil
}
