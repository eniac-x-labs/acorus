package event

import (
	"context"
	"github.com/cornerstone-labs/acorus/database/common"

	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/synchronizer"
)

var blocksLimit = 10_000

type EventHandle struct {
	db                          *database.DB
	resourceCtx                 context.Context
	resourceCancel              context.CancelFunc
	tasks                       tasks.Group
	synchronizer                *synchronizer.Synchronizer
	config                      *config.Config
	LatestL1L2InitL1Header      *common.BlockHeader
	LatestL1L2L2Header          *common.BlockHeader
	LatestL1L2FinalizedL2Header *common.BlockHeader
	LatestStateRootL1Header     *common.BlockHeader
	LatestMantleDAL1Header      *common.BlockHeader
	LatestL2L1InitL2Header      *common.BlockHeader
	LatestProvenL1Header        *common.BlockHeader
	LatestFinalizedL1Header     *common.BlockHeader
}

func NewEventHandle(db *database.DB, synchronizer *synchronizer.Synchronizer, conf config.Config, shutdown context.CancelCauseFunc, chainId string) (*EventHandle, error) {
	return nil, nil
}
