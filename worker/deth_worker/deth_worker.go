package deth_worker

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/rpc/bridge"
	"github.com/cornerstone-labs/acorus/rpc/bridge/protobuf/pb"
)

type WorkerProcessor struct {
	db               *database.DB
	bridgeRpcService bridge.BridgeRpcService
	tasks            tasks.Group
}

func NewWorkerProcessor(db *database.DB, bridgeRpcService bridge.BridgeRpcService, shutdown context.CancelCauseFunc) (*WorkerProcessor, error) {
	workerProcessor := WorkerProcessor{
		db:               db,
		bridgeRpcService: bridgeRpcService,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in worker processor: %w", err))
		}},
	}
	return &workerProcessor, nil
}

func (b *WorkerProcessor) WorkerStart() error {
	ticker := time.NewTicker(time.Second)
	b.tasks.Go(func() error {
		for range ticker.C {
			if err := b.TranserL2Shares(); err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})
	return nil
}

func (b *WorkerProcessor) TranserL2Shares() error {
	listDEthTransferNoRelayer := b.db.AppChainDEthTransfer.ListDEthTransferNoRelayer()
	log.Info("TranserL2Shares", "listDEthTransferNoRelayer", len(listDEthTransferNoRelayer))
	if len(listDEthTransferNoRelayer) == 0 {
		return nil
	}

	for i := 0; i < len(listDEthTransferNoRelayer); i++ {
		dethTransfer := listDEthTransferNoRelayer[i]
		err := b.TransferL2Shares(dethTransfer.From, dethTransfer.To, dethTransfer.Shares,
			dethTransfer.MessageNonce.Uint64(), dethTransfer.TxHash)
		if err != nil {
			log.Error("TranserL2Shares", "TransferL2SharesErr", err)
		}
	}

	return nil
}

func (b *WorkerProcessor) TransferL2Shares(fromAddress, toAddress common.Address,
	shares *big.Int, stakeMessageNonce uint64, txHash common.Hash) error {
	fromAddressDethList := b.db.AppChainDethShares.ListDethSharesByAddress(fromAddress)
	log.Info("TransferL2Shares", "fromAddressDethList", len(fromAddressDethList))
	if len(fromAddressDethList) == 0 {
		return nil
	}
	sharesMap := make(map[uint64]*pb.ShareMap)
	for i := 0; i < len(fromAddressDethList); i++ {
		dethSharesInfo := fromAddressDethList[i]
		theStrategyShares := dethSharesInfo.Shares
		chainIdStr := dethSharesInfo.ChainId
		chainId, _ := big.NewInt(0).SetString(chainIdStr, 10)
		chainIdUint := chainId.Uint64()
		totalShares := GetTotalShares(sharesMap)
		needShares := big.NewInt(0).Sub(shares, totalShares)
		if needShares.Cmp(shares) == 0 {
			break
		}
		strategyAddress := dethSharesInfo.StrategyAddress
		if needShares.Cmp(theStrategyShares) == -1 {
			AddSharesToMap(sharesMap, chainIdUint, strategyAddress, needShares)
			break
		}
		if needShares.Cmp(theStrategyShares) == 0 {
			AddSharesToMap(sharesMap, chainIdUint, strategyAddress, needShares)
			break
		}
		if needShares.Cmp(theStrategyShares) == 1 {
			AddSharesToMap(sharesMap, chainIdUint, strategyAddress, theStrategyShares)
		}
	}
	l2ShareReq := &pb.TransferL2ShareRequest{
		ShareRequest:      sharesMap,
		From:              fromAddress.String(),
		To:                toAddress.String(),
		StakeMessageNonce: stakeMessageNonce,
	}
	transferL2ShareResp, err := b.bridgeRpcService.TransferL2Share(l2ShareReq)
	log.Info("TransferL2Shares", "TransferL2ShareResp", transferL2ShareResp)

	if err != nil {
		log.Error("TransferL2Shares", "TransferL2ShareErr", err)
		return err
	}
	if transferL2ShareResp.Success {
		log.Info("TransferL2Shares", "call rpc success", transferL2ShareResp.Success)
		return b.db.AppChainDEthTransfer.NotifyTransferL2Success(txHash)
	} else {
		return fmt.Errorf("TransferL2Shares call Rpc Err %s", transferL2ShareResp.Message)
	}
}

func GetTotalShares(sharesMap map[uint64]*pb.ShareMap) *big.Int {
	totalShares := bigint.Zero
	for _, shareMapInfo := range sharesMap {
		shareMap := shareMapInfo.GetShareMap()
		for _, shares := range shareMap {
			sharesBig, _ := big.NewInt(0).SetString(shares, 10)
			totalShares = big.NewInt(0).Add(totalShares, sharesBig)
		}
	}
	return totalShares
}

func AddSharesToMap(sharesMap map[uint64]*pb.ShareMap, chainId uint64, strategyAddress common.Address, shares *big.Int) {
	shareMap, exists := sharesMap[chainId]
	if exists {
		strategyShares, existsStrategy := shareMap.ShareMap[strategyAddress.String()]
		if existsStrategy {
			existsSharesBig, _ := big.NewInt(0).SetString(strategyShares, 10)
			existsSharesBig = big.NewInt(0).Add(existsSharesBig, shares)
			shareMap.ShareMap[strategyAddress.String()] = existsSharesBig.String()
		} else {
			shareMap.ShareMap[strategyAddress.String()] = shares.String()
		}
	} else {
		strategyAmountMap := make(map[string]string)
		strategyAmountMap[strategyAddress.String()] = shares.String()
		sharesMap[chainId] = &pb.ShareMap{
			ShareMap: strategyAmountMap,
		}
	}
}
