package stateroot

import (
	op_stack "github.com/cornerstone-labs/acorus/event/processor/op-stack"
	"math/big"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/event/processor/op-stack/contracts"
	"github.com/ethereum/go-ethereum/common"
)

func LegacyL1ProcessSCCEvent(log log.Logger, db *database.DB, fromHeight, toHeight *big.Int) error {
	sccEvents, err := contracts.LegacySCCBatchAppendedEvent(common.HexToAddress(op_stack.LegacyStateCommitmentChain), db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(sccEvents) > 0 {
		log.Info("detected legacy scc state batch appended event", "size", len(sccEvents))
		if err := db.StateRoots.StoreBatchStateRoots(sccEvents); err != nil {
			return err
		}
	}
	return nil

}

func L2OutputEvent(log log.Logger, db *database.DB, fromHeight, toHeight *big.Int) error {
	l2OutputProposedEvents, err := contracts.L2OutputProposedEvent(common.HexToAddress(op_stack.L2OutputOracleProxy), db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(l2OutputProposedEvents) > 0 {
		log.Info("detected l2output proposed event", "size", len(l2OutputProposedEvents))
		if err := db.StateRoots.StoreBatchStateRoots(l2OutputProposedEvents); err != nil {
			log.Error("Store batch state roots fail")
			return err
		}
	}
	return nil
}
