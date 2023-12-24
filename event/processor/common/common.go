package common

import (
	"context"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/event/processor/polygon/abi"
	"github.com/cornerstone-labs/acorus/synchronizer"
)

const (
	OpChainId      = 1
	ScrollChainId  = 2
	PolygonChainId = 3
)

type IProcessor struct {
	Log                   log.Logger
	Db                    *database.DB
	ResourceCtx           context.Context
	ResourceCancel        context.CancelFunc
	Tasks                 tasks.Group
	L1Sync                *synchronizer.L1Sync
	L2Sync                *synchronizer.L2Sync
	ChainConfig           config.ChainConfig
	LastL1Header          *common2.L1BlockHeader
	LastL2Header          *common2.L2BlockHeader
	LastRollupL1Header    *common2.L1BlockHeader
	LastFinalizedL1Header *common2.L1BlockHeader
	LastFinalizedL2Header *common2.L2BlockHeader
	L1PolygonBridge       *abi.Polygonzkevmbridge
	L2PolygonBridge       *abi.Polygonzkevmbridge
}

func (i *IProcessor) Start() error {
	return nil
}

func (i *IProcessor) Close() error {
	return nil
}

func GetLastBlockHeardByChain(log log.Logger, db *database.DB,
	chainConfig config.ChainConfig) (
	*common2.L1BlockHeader,
	*common2.L2BlockHeader,
	*common2.L1BlockHeader,
	*common2.L2BlockHeader, error) {

	chainId := chainConfig.ChainId
	log = log.New("processor", "bridge", "chainId=", chainId)

	latestL1Header, err := db.L1ToL2.L1LatestBlockHeader(chainId)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	latestL2Header, err := db.L2ToL1.L2LatestBlockHeader(chainId)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	latestFinalizedL1Header, err := db.L2ToL1.L1LatestFinalizedBlockHeader(chainId)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	latestFinalizedL2Header, err := db.L1ToL2.L2LatestFinalizedBlockHeader(chainId)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	log.Info("detected indexed bridge state",
		"l1_block", latestL1Header, "l2_block", latestL2Header,
		"finalized_l1_block", latestFinalizedL1Header, "finalized_l2_block", latestFinalizedL2Header)
	return latestL1Header, latestL2Header, latestFinalizedL1Header, latestFinalizedL2Header, nil
}
