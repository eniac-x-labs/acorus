package contracts

import (
	"encoding/hex"
	"math/big"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
)

func L2OutputProposedEvent(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) ([]worker.StateRoot, error) {
	l2OutputAbi, err := bindings.L2OutputOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	outputProposedEventAbi := l2OutputAbi.Events["OutputProposed"]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: outputProposedEventAbi.ID}
	outputProposedEvents, err := db.ContractEvents.ContractEventsWithFilter("10", contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}
	stateRoots := make([]worker.StateRoot, len(outputProposedEvents))
	for i := range outputProposedEvents {
		outputProposed := bindings.L2OutputOracleOutputProposed{Raw: *outputProposedEvents[i].RLPLog}
		err = UnpackLog(&outputProposed, outputProposedEvents[i].RLPLog, outputProposedEventAbi.Name, l2OutputAbi)
		if err != nil {
			return nil, err
		}
		l1BlockNumber, err := db.L1ToL2.GetBlockNumberFromHash("10", outputProposedEvents[i].BlockHash)
		if err != nil {
			return nil, err
		}
		stateRoots[i] = worker.StateRoot{
			GUID:            uuid.New(),
			TransactionHash: outputProposedEvents[i].TransactionHash,
			BlockHash:       outputProposedEvents[i].BlockHash,
			OutputRoot:      hex.EncodeToString(outputProposed.OutputRoot[:]),
			OutputIndex:     outputProposed.L2OutputIndex,
			BatchSize:       outputProposed.L2OutputIndex,
			L1BlockNumber:   l1BlockNumber,
			L2BlockNumber:   outputProposed.L2BlockNumber,
			Timestamp:       outputProposed.L1Timestamp.Uint64(),
		}
	}
	return stateRoots, nil
}
