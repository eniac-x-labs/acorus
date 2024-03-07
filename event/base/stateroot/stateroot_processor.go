package stateroot

import (
	"github.com/cornerstone-labs/acorus/database"
	common3 "github.com/cornerstone-labs/acorus/event/base/common"
	"github.com/cornerstone-labs/acorus/event/base/contracts"
	"log"
	"math/big"
)

func LegacyL1ProcessSCCEvent(log log.Logger, db *database.DB, l1chainId, l2chainId string, fromHeight, toHeight *big.Int) error {
	sccEvents, err := contracts.LegacySCCBatchAppendedEvent(l1chainId, common3.LegacyStateCommitmentChain, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(sccEvents) > 0 {
		log.Println("detected legacy scc state batch appended event", "size", len(sccEvents))
		if err := db.StateRoots.StoreBatchStateRoots(l1chainId, sccEvents); err != nil {
			return err
		}
	}
	return nil

}

func L2OutputEvent(db *database.DB, fromHeight, toHeight *big.Int, l1ChainId, l2ChainId string) error {
	log.Println("L2OutputEvent", "fromHeight", fromHeight, "toHeight", toHeight)
	l2OutputProposedEvents, err := contracts.L2OutputProposedEvent(common3.L2OutputOracleProxy, db,
		fromHeight, toHeight, l1ChainId, l2ChainId)
	if err != nil {
		return err
	}
	if len(l2OutputProposedEvents) > 0 {
		log.Println("detected l2output proposed event", "size", len(l2OutputProposedEvents))
		if err := db.StateRoots.StoreBatchStateRoots(l2ChainId, l2OutputProposedEvents); err != nil {
			log.Println("Store batch state roots fail")
			return err
		}
	}
	return nil
}
