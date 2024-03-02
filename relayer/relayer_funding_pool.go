package relayer

import (
	"context"
	"fmt"
	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/relayer"
	"github.com/cornerstone-labs/acorus/rpc/bridge"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strconv"
	"strings"
	"time"
)

type RelayerFundingPool struct {
	db               *database.DB
	bridgeRpcService bridge.BridgeRpcService
	resourceCtx      context.Context
	resourceCancel   context.CancelFunc
	tasks            tasks.Group
	rpcS             []*config.RPC
	isMainnet        bool
}

func NewRelayerFundingPool(
	db *database.DB,
	bridgeRpcService bridge.BridgeRpcService,
	rpcS []*config.RPC,
	shutdown context.CancelCauseFunc) (*RelayerFundingPool, error) {
	resCtx, resCancel := context.WithCancel(context.Background())
	isMainnet := true
	for i := 0; i < len(rpcS); i++ {
		if rpcS[i].ChainId == global_const.EthereumChainId {
			isMainnet = rpcS[i].IsMainnet
			break
		}
	}
	return &RelayerFundingPool{
		db:               db,
		bridgeRpcService: bridgeRpcService,
		resourceCancel:   resCancel,
		resourceCtx:      resCtx,
		rpcS:             rpcS,
		isMainnet:        isMainnet,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
	}, nil
}

func (rfp *RelayerFundingPool) Start() error {
	tickerScan := time.NewTicker(5 * time.Second)
	// start rundingPool worker
	rfp.tasks.Go(func() error {
		for range tickerScan.C {
			err := rfp.ScanL1NeedFundBalance()
			if err != nil {
				log.Println(" shutting down ScanL1NeedFundBalance ", "err", err)
				continue
			}
		}
		return nil
	})

	tickerCross := time.NewTicker(5 * time.Second)
	// start FundingPoolCross worker
	rfp.tasks.Go(func() error {
		for range tickerCross.C {
			rfp.FundingPoolCross()
		}
		return nil
	})
	return nil
}

func (rfp *RelayerFundingPool) ScanL1NeedFundBalance() error {
	for i := 0; i < len(rfp.rpcS); i++ {
		rpc := rfp.rpcS[i]
		chainId := rpc.ChainId
		if chainId == 1 {
			continue
		}
		poolContract := rpc.PoolContract
		chainIdStr := strconv.FormatUint(chainId, 10)
		err := rfp.scanL1NeedFundBalanceByChainId(chainIdStr, poolContract)
		if err != nil {
			log.Println(" shutting down scanL1NeedFundBalanceByChainId ", "err", err)
			return err
		}
	}
	return nil
}

func (rfp *RelayerFundingPool) scanL1NeedFundBalanceByChainId(chainId, toAddress string) error {
	nilAddress := common.Address{}
	if err := rfp.db.Transaction(func(tx *database.DB) error {
		// get all l1 need fund balance
		needFundBalances, err := rfp.db.BridgeFundingPoolDB.L1GetCanStoreTransactions(chainId, strings.ToLower(toAddress))
		if err != nil {
			return err
		}

		if len(needFundBalances) > 0 {
			var needStoreList []relayer.BridgeFundingPoolUpdate
			for i := 0; i < len(needFundBalances); i++ {
				l1ToL2 := needFundBalances[i]
				bridgeFundingPool := relayer.BridgeFundingPoolUpdate{
					LayerType:      global_const.LayerTypeOne,
					TxHash:         l1ToL2.L1TransactionHash.String(),
					DestChainId:    chainId,
					ReceiveAddress: l1ToL2.ToAddress.String(),
				}
				if rfp.isMainnet {
					bridgeFundingPool.SourceChainId = "1"
				} else {
					bridgeFundingPool.SourceChainId = "11155111"
				}
				if l1ToL2.L2TokenAddress == nilAddress {
					bridgeFundingPool.TokenAddress = l1ToL2.L1TokenAddress.String()
					bridgeFundingPool.Amount = l1ToL2.TokenAmounts
				} else {
					bridgeFundingPool.Amount = l1ToL2.ETHAmount.String()
				}
				needStoreList = append(needStoreList, bridgeFundingPool)
			}
			err = rfp.db.BridgeFundingPoolDB.StoreBridgeFundingPools(needStoreList)
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (rfp *RelayerFundingPool) FundingPoolCross() {
	log.Println("FundingPoolCross, start")
	list := rfp.db.BridgeFundingPoolDB.BridgeFundingPoolNoSendList()
	if len(list) > 0 {
		for _, v := range list {
			destChainId := v.DestChainId
			sourceChainId := v.SourceChainId
			amount := v.Amount
			receiveAddress := v.ReceiveAddress
			tokenAddress := v.TokenAddress
			txHash := v.TxHash
			updateFundingPool, err := rfp.bridgeRpcService.UpdateFundingPoolBalance(sourceChainId, destChainId, amount,
				receiveAddress, tokenAddress, txHash)
			if err != nil {
				log.Println("FundingPoolCross", "error", err)
				continue
			}
			log.Println("FundingPoolCross", "Update", updateFundingPool.Success)
			if updateFundingPool.Success {
				rfp.db.BridgeFundingPoolDB.UpdateCrossStatus(txHash)
			}
		}
	}
}
