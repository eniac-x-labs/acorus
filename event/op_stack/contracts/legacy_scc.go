package contracts

import (
	"encoding/hex"
	"math/big"

	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/ethereum/go-ethereum/common"

	legacy_bindings "github.com/ethereum-optimism/optimism/op-bindings/legacy-bindings"
)

func LegacySCCBatchAppendedEvent(chainId string, contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) ([]worker.StateRoot, error) {
	sccAbi, err := legacy_bindings.StateCommitmentChainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	stateBatchAppendedEventAbi := sccAbi.Events["StateBatchAppended"]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: stateBatchAppendedEventAbi.ID}
	stateBatchAppendedEvents, err := db.ContractEvents.ContractEventsWithFilter(chainId, contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}
	stateRoots := make([]worker.StateRoot, len(stateBatchAppendedEvents))
	for i := range stateBatchAppendedEvents {
		stateBatchAppended := legacy_bindings.StateCommitmentChainStateBatchAppended{Raw: *stateBatchAppendedEvents[i].RLPLog}
		err = UnpackLog(&stateBatchAppended, stateBatchAppendedEvents[i].RLPLog, stateBatchAppendedEventAbi.Name, sccAbi)
		if err != nil {
			return nil, err
		}

		stateRoots[i] = worker.StateRoot{
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
