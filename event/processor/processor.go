package processor

import (
	"context"

	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
)

type IProcessor struct {
	Db             *database.DB
	CfgRpc         *config.RPC
	ResourceCtx    context.Context
	ResourceCancel context.CancelFunc
	Tasks          tasks.Group
}

func (ip *IProcessor) Start() {

}

//func GetLastBlockHeardByChain(db *database.DB,
//	chainRpc config.RPC) (
//	*common2.BlockHeader,
//	*common2.BlockHeader, error) {
//
//	chainId := chainRpc.ChainId
//
//	chainIdStr := strconv.Itoa(int(chainId))
//
//	latestL1Header, err := db.L1ToL2.LatestBlockHeader(chainIdStr)
//	if err != nil {
//		return nil, nil, err
//	}
//	latestL2Header, err := db.L2ToL1.LatestBlockHeader(chainIdStr)
//	if err != nil {
//		return nil, nil, err
//	}
//	log.Println("detected indexed bridge state",
//		"l1_block", latestL1Header, "l2_block", latestL2Header)
//	return latestL1Header, latestL2Header, nil
//}
