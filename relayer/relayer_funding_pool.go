package relayer

import (
	"context"
	"fmt"
	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	common4 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/relayer"
	"github.com/cornerstone-labs/acorus/rpc/bridge"
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
}

func NewRelayerFundingPool(
	db *database.DB,
	bridgeRpcService bridge.BridgeRpcService,
	rpcS []*config.RPC,
	shutdown context.CancelCauseFunc) (*RelayerFundingPool, error) {
	resCtx, resCancel := context.WithCancel(context.Background())
	return &RelayerFundingPool{
		db:               db,
		bridgeRpcService: bridgeRpcService,
		resourceCancel:   resCancel,
		resourceCtx:      resCtx,
		rpcS:             rpcS,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
	}, nil
}

func (rfp *RelayerFundingPool) Start() error {
	l1TickerScan := time.NewTicker(5 * time.Second)
	// start rundingPoolL1 worker
	rfp.tasks.Go(func() error {
		for range l1TickerScan.C {
			err := rfp.ScanL1NeedFundBalance()
			if err != nil {
				log.Println(" shutting down ScanL1NeedFundBalance ", "err", err)
				continue
			}
		}
		return nil
	})

	l2TickerScan := time.NewTicker(5 * time.Second)
	// start rundingPoolL2 worker
	rfp.tasks.Go(func() error {
		for range l2TickerScan.C {
			err := rfp.ScanL2NeedFundBalance()
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
		l1ChainId := rpc.L1ChainId
		if chainId == 1 {
			continue
		}
		poolContract := rpc.L2PoolContract
		if poolContract == "" {
			continue
		}
		l2ChainIdStr := strconv.FormatUint(chainId, 10)
		l1ChainIdStr := strconv.FormatUint(l1ChainId, 10)
		err := rfp.scanL1NeedFundBalanceByChainId(l1ChainIdStr, l2ChainIdStr, poolContract)
		if err != nil {
			log.Println(" shutting down scanL1NeedFundBalanceByChainId ", "err", err)
			return err
		}
	}
	return nil
}

func (rfp *RelayerFundingPool) scanL1NeedFundBalanceByChainId(l1ChainIdStr, l2ChainIdStr, toAddress string) error {
	if err := rfp.db.Transaction(func(tx *database.DB) error {
		// get all l1 need fund balance
		needFundBalances, err := rfp.db.BridgeFundingPoolDB.L1GetCanStoreTransactions(l2ChainIdStr, strings.ToLower(toAddress))
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
					DestChainId:    l2ChainIdStr,
					ReceiveAddress: l1ToL2.ToAddress.String(),
				}
				bridgeFundingPool.SourceChainId = l1ChainIdStr

				if l1ToL2.AssetType == common4.ERC20 {
					bridgeFundingPool.TokenAddress = l1ToL2.L1TokenAddress.String()
					bridgeFundingPool.Amount = l1ToL2.TokenAmounts
				} else {
					bridgeFundingPool.TokenAddress = global_const.EthAddress
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

func (rfp *RelayerFundingPool) ScanL2NeedFundBalance() error {
	for i := 0; i < len(rfp.rpcS); i++ {
		rpc := rfp.rpcS[i]
		chainId := rpc.ChainId
		l1ChainId := rpc.L1ChainId
		if chainId == 1 {
			continue
		}
		poolContract := rpc.L1PoolContract
		if poolContract == "" {
			continue
		}
		l2ChainIdStr := strconv.FormatUint(chainId, 10)
		l1ChainIdStr := strconv.FormatUint(l1ChainId, 10)
		err := rfp.scanL2NeedFundBalanceByChainId(l1ChainIdStr, l2ChainIdStr, poolContract)
		if err != nil {
			log.Println(" shutting down scanL1NeedFundBalanceByChainId ", "err", err)
			return err
		}
	}
	return nil
}

func (rfp *RelayerFundingPool) scanL2NeedFundBalanceByChainId(l1ChainIdStr, l2ChainIdStr, toAddress string) error {
	if err := rfp.db.Transaction(func(tx *database.DB) error {
		// get all l2 need fund balance
		needFundBalances, err := rfp.db.BridgeFundingPoolDB.L2GetCanStoreTransactions(l2ChainIdStr, strings.ToLower(toAddress))
		if err != nil {
			return err
		}

		if len(needFundBalances) > 0 {
			var needStoreList []relayer.BridgeFundingPoolUpdate
			for i := 0; i < len(needFundBalances); i++ {
				l2ToL1 := needFundBalances[i]
				bridgeFundingPool := relayer.BridgeFundingPoolUpdate{
					LayerType:      global_const.LayerTypeTwo,
					TxHash:         l2ToL1.L2TransactionHash.String(),
					DestChainId:    l1ChainIdStr,
					ReceiveAddress: l2ToL1.ToAddress.String(),
				}
				bridgeFundingPool.SourceChainId = l2ChainIdStr

				if l2ToL1.AssetType == common4.ERC20 {
					bridgeFundingPool.TokenAddress = l2ToL1.L1TokenAddress.String()
					bridgeFundingPool.Amount = l2ToL1.TokenAmounts
				} else {
					bridgeFundingPool.TokenAddress = global_const.EthAddress
					bridgeFundingPool.Amount = l2ToL1.ETHAmount.String()
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
