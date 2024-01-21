package stateroot

import (
	"github.com/cornerstone-labs/acorus/common/global_const"
	common3 "github.com/cornerstone-labs/acorus/event/op_stack/common"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/event/op_stack/contracts"
)

func LegacyL1ProcessSCCEvent(log log.Logger, db *database.DB, fromHeight, toHeight *big.Int) error {
	sccEvents, err := contracts.LegacySCCBatchAppendedEvent(common3.LegacyStateCommitmentChain, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(sccEvents) > 0 {
		log.Info("detected legacy scc state batch appended event", "size", len(sccEvents))
		if err := db.StateRoots.StoreBatchStateRoots(strconv.FormatUint(global_const.OpChinId, 10), sccEvents); err != nil {
			return err
		}
	}
	return nil

}

func L2OutputEvent(db *database.DB, fromHeight, toHeight *big.Int) error {
	l2OutputProposedEvents, err := contracts.L2OutputProposedEvent(common3.L2OutputOracleProxy, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(l2OutputProposedEvents) > 0 {
		log.Info("detected l2output proposed event", "size", len(l2OutputProposedEvents))
		if err := db.StateRoots.StoreBatchStateRoots(strconv.FormatUint(global_const.OpChinId, 10), l2OutputProposedEvents); err != nil {
			log.Error("Store batch state roots fail")
			return err
		}
	}
	return nil
}
