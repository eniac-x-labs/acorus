package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/relayer"
	"github.com/cornerstone-labs/acorus/relayer/bindings"
	"github.com/cornerstone-labs/acorus/relayer/unpack"
	"github.com/cornerstone-labs/acorus/rpc/bridge"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"time"
)

type RelayerListener struct {
	db                *database.DB
	chainId           string
	l1Contracts       []string
	l2Contracts       []string
	l1EventStartBlock uint64
	l2EventStartBlock uint64
	l1StartHeight     *big.Int
	l2StartHeight     *big.Int
	layerType         int
	resourceCtx       context.Context
	resourceCancel    context.CancelFunc
	tasks             tasks.Group
	loopInterval      time.Duration
	epoch             uint64
	bridgeRpcService  bridge.BridgeRpcService
}

func NewRelayerListener(
	db *database.DB,
	bridgeRpcService bridge.BridgeRpcService,
	chainId string,
	l1Contracts []string,
	l2Contracts []string,
	l1EventStartBlock uint64,
	l2EventStartBlock uint64,
	layerType int,
	shutdown context.CancelCauseFunc,
	loopInterval time.Duration, epoch uint64) (*RelayerListener, error) {
	resCtx, resCancel := context.WithCancel(context.Background())
	return &RelayerListener{
		db:                db,
		chainId:           chainId,
		l1EventStartBlock: l1EventStartBlock,
		l2EventStartBlock: l2EventStartBlock,
		l1Contracts:       l1Contracts,
		l2Contracts:       l2Contracts,
		layerType:         layerType,
		resourceCtx:       resCtx,
		resourceCancel:    resCancel,
		bridgeRpcService:  bridgeRpcService,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
		loopInterval: loopInterval,
		epoch:        epoch,
	}, nil
}

func (rl *RelayerListener) Start() error {

	if rl.layerType == global_const.LayerTypeOne {
		relayerEventOn1 := time.NewTicker(rl.loopInterval)
		rl.tasks.Go(func() error {
			for range relayerEventOn1.C {
				err := rl.onL1Data()
				if err != nil {
					log.Println("no more l1 etl updates. shutting down l1 task")
					return err
				}
			}
			return nil
		})
	} else {
		relayerEventOn2 := time.NewTicker(rl.loopInterval)

		rl.tasks.Go(func() error {
			for range relayerEventOn2.C {
				err := rl.onL2Data()
				if err != nil {
					log.Println("no more l1 etl updates. shutting down l1 task")
					return err
				}
			}
			return nil
		})
	}
	tickerBridgeRel := time.NewTicker(rl.loopInterval)
	// start relation worker
	rl.tasks.Go(func() error {
		for range tickerBridgeRel.C {
			err := rl.relationBridge()
			if err != nil {
				log.Println("shutting down relationBridge")
				return err
			}
		}
		return nil
	})

	tickerCross := time.NewTicker(10 * time.Second)
	rl.tasks.Go(func() error {
		for range tickerCross.C {
			err := rl.CrossChainTransfer()
			if err != nil {
				log.Println("shutting down relationBridge")
				return err
			}
		}
		return nil
	})

	tickerTrans := time.NewTicker(10 * time.Second)
	rl.tasks.Go(func() error {
		for range tickerTrans.C {
			err := rl.ChangeTransferStatus()
			if err != nil {
				log.Println("shutting down relationBridge")
				return err
			}
		}
		return nil
	})
	return nil
}

func (rl *RelayerListener) onL1Data() error {

	if rl.l1StartHeight == nil {
		lastListenBlock, err := rl.db.BridgeBlockListener.GetLastBlockNumber(rl.chainId)
		if err != nil {
			log.Println("l1 failed to get last block heard", "err", err)
			return err
		}
		if lastListenBlock == nil {
			lastListenBlock = &relayer.BridgeBlockListener{
				BlockNumber: big.NewInt(0),
			}
		}
		rl.l1StartHeight = lastListenBlock.BlockNumber
		if rl.l1StartHeight.Cmp(big.NewInt(int64(rl.l1EventStartBlock))) == -1 {
			rl.l1StartHeight = big.NewInt(int64(rl.l1EventStartBlock))
		}
	} else {
		rl.l1StartHeight = new(big.Int).Add(rl.l1StartHeight, bigint.One)
	}

	fromL1Height := rl.l1StartHeight
	toL1Height := new(big.Int).Add(fromL1Height, big.NewInt(int64(rl.epoch)))
	chainLatestBlockHeader, err := rl.db.Blocks.ChainLatestBlockHeader(rl.chainId)
	if err != nil {
		return err
	}
	if chainLatestBlockHeader == nil {
		return nil
	}
	if chainLatestBlockHeader.Number.Cmp(fromL1Height) == -1 {
		return nil
	} else {
		if chainLatestBlockHeader.Number.Cmp(toL1Height) == -1 {
			toL1Height = chainLatestBlockHeader.Number
		}
	}
	if err := rl.db.Transaction(func(tx *database.DB) error {
		l1EventsFetchErr := rl.l1EventsFetch(fromL1Height, toL1Height)
		if l1EventsFetchErr != nil {
			return l1EventsFetchErr
		}
		lastBlock := relayer.BridgeBlockListener{
			ChainId:     rl.chainId,
			BlockNumber: toL1Height,
		}
		updateErr := rl.db.BridgeBlockListener.SaveOrUpdateLastBlockNumber(lastBlock)
		if updateErr != nil {
			log.Println("update last block err :", updateErr)
			return updateErr
		}
		return nil
	}); err != nil {
		return err
	}
	rl.l1StartHeight = toL1Height
	return nil
}

func (rl *RelayerListener) onL2Data() error {
	if rl.l2StartHeight == nil {
		lastListenBlock, err := rl.db.BridgeBlockListener.GetLastBlockNumber(rl.chainId)
		if err != nil {
			log.Println("l2 failed to get last block heard", "err", err)
			return err
		}
		if lastListenBlock == nil {
			if rl.l2EventStartBlock > 0 {
				lastListenBlock = &relayer.BridgeBlockListener{
					BlockNumber: big.NewInt(int64(rl.l2EventStartBlock)),
				}
			} else {
				lastListenBlock = &relayer.BridgeBlockListener{
					BlockNumber: big.NewInt(0),
				}
			}
		}
		rl.l2StartHeight = lastListenBlock.BlockNumber
	} else {
		rl.l2StartHeight = new(big.Int).Add(rl.l2StartHeight, bigint.One)
	}
	fromL2Height := rl.l2StartHeight
	toL2Height := new(big.Int).Add(fromL2Height, big.NewInt(int64(rl.epoch)))
	chainLatestBlockHeader, err := rl.db.Blocks.ChainLatestBlockHeader(rl.chainId)
	if err != nil {
		return err
	}
	if chainLatestBlockHeader == nil {
		return nil
	}
	if chainLatestBlockHeader.Number.Cmp(fromL2Height) == -1 {
		return nil
	} else {
		if chainLatestBlockHeader.Number.Cmp(toL2Height) == -1 {
			toL2Height = chainLatestBlockHeader.Number
		}
	}
	if err := rl.db.Transaction(func(tx *database.DB) error {
		l2EventsFetchErr := rl.l2EventsFetch(fromL2Height, toL2Height)
		if l2EventsFetchErr != nil {
			return l2EventsFetchErr
		}
		lastBlock := relayer.BridgeBlockListener{
			ChainId:     rl.chainId,
			BlockNumber: toL2Height,
		}
		updateErr := rl.db.BridgeBlockListener.SaveOrUpdateLastBlockNumber(lastBlock)
		if updateErr != nil {
			return updateErr
		}
		return nil
	}); err != nil {
		return err
	}
	rl.l2StartHeight = toL2Height
	return nil
}

func (rl *RelayerListener) l1EventsFetch(fromL1Height, toL1Height *big.Int) error {
	log.Println("relayer l1EventsFetch", "fromL1Height", fromL1Height, "toL1Height", toL1Height)
	l1Contracts := rl.l1Contracts
	for _, l1contract := range l1Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l1contract)}
		events, err := rl.db.ContractEvents.ContractEventsWithFilter(rl.chainId, contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			log.Println("failed to index L1ContractEventsWithFilter ", "err", err)
			return err
		}
		for _, contractEvent := range events {
			unpackErr := rl.eventUnpack(contractEvent)
			if unpackErr != nil {
				log.Println("failed to index L1 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
		}
	}
	return nil
}

func (rl *RelayerListener) l2EventsFetch(fromL1Height, toL1Height *big.Int) error {
	log.Println("relayer l2EventsFetch", "fromL1Height", fromL1Height, "toL1Height", toL1Height)

	chainIdStr := rl.chainId
	l2Contracts := rl.l2Contracts
	for _, l2contract := range l2Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l2contract)}
		events, err := rl.db.ContractEvents.ContractEventsWithFilter(chainIdStr, contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			log.Println("failed to index L2ContractEventsWithFilter ", "err", err)
			return err
		}
		for _, contractEvent := range events {
			unpackErr := rl.eventUnpack(contractEvent)
			if unpackErr != nil {
				log.Println("failed to index L2 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
		}
	}
	return nil
}

func (rl *RelayerListener) eventUnpack(event event.ContractEvent) error {
	switch event.EventSignature.String() {
	case bindings.MsgAbi.Events["MessageSent"].ID.String():
		err := unpack.MessageSent(event, rl.db)
		return err
	case bindings.MsgAbi.Events["MessageClaimed"].ID.String():
		err := unpack.MessageClaimed(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["FinalizeETH"].ID.String():
		err := unpack.FinalizeETH(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["FinalizeWETH"].ID.String():
		err := unpack.FinalizeWETH(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["FinalizeERC20"].ID.String():
		err := unpack.FinalizeERC20(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["InitiateWETH"].ID.String():
		err := unpack.InitiateWETH(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["InitiateETH"].ID.String():
		err := unpack.InitiateETH(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["InitiateERC20"].ID.String():
		err := unpack.InitiateERC20(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["StarkingERC20Event"].ID.String():
		err := unpack.StarkingERC20Event(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["StakingETHEvent"].ID.String():
		err := unpack.StakingETHEvent(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["StakingWETHEvent"].ID.String():
		err := unpack.StakingWETHEvent(event, rl.db)
		return err
	case bindings.L1PoolAbi.Events["ClaimEvent"].ID.String():
		err := unpack.ClaimEvent(event, rl.db)
		return err
	case bindings.IL1PoolAbi.Events["Withdraw"].ID.String():
		err := unpack.Withdraw(event, rl.db)
		return err
	case bindings.IL1PoolAbi.Events["ClaimReward"].ID.String():
		err := unpack.ClaimReward(event, rl.db)
		return err
	}

	return nil
}

func (rl *RelayerListener) relationBridge() error {
	if err := rl.db.Transaction(func(tx *database.DB) error {

		// step 1
		log.Println("RelationClaim")
		err := rl.db.BridgeFinalize.RelationClaim()
		if err != nil {
			log.Println("relationBridge failed", "err", err)
			return err
		}
		// step 2
		log.Println("RelationMsgHash")
		err = rl.db.BridgeMsgHash.RelationMsgHash()
		if err != nil {
			log.Println("relationBridge failed", "err", err)
			return err
		}
		// step 3
		log.Println("RelationMsgSent")
		err = rl.db.BridgeClaim.RelationMsgSent()
		if err != nil {
			log.Println("RelationMsgSent failed", "err", err)
			return err
		}

		list, err := rl.db.BridgeMsgSent.GetCanSaveDecodeList()
		if err != nil {
			return err
		}
		bridgeRecordSaveList := make([]relayer.BridgeRecords, 0)
		for _, v := range list {
			data := v.Data
			var bridgeRecord relayer.BridgeRecords
			if unMarErr := json.Unmarshal([]byte(data), &bridgeRecord); unMarErr != nil {
				return unMarErr
			}
			bridgeRecord.DestTokenAddress = v.DestToken
			bridgeRecord.DestBlockNumber = v.DestBlockNumber
			bridgeRecord.DestTxHash = v.DestHash
			bridgeRecord.ClaimTimestamp = v.DestTimestamp
			bridgeRecord.Nonce = v.MsgNonce
			bridgeRecord.Fee = v.Fee
			bridgeRecordSaveList = append(bridgeRecordSaveList, bridgeRecord)
		}
		if len(bridgeRecordSaveList) > 0 {
			log.Println("StoreBridgeRecords")
			err = rl.db.BridgeRecord.StoreBridgeRecords(bridgeRecordSaveList)
			if err != nil {
				return err
			}
		}
		log.Println("CleanMsgSent")
		err = rl.db.BridgeMsgSent.CleanMsgSent()
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (rl *RelayerListener) CrossChainTransfer() error {
	list := rl.db.BridgeMsgSent.GetCanCrossDataList()
	if len(list) > 0 {
		for _, v := range list {
			data := v.Data
			var bridgeRecord relayer.BridgeRecords
			if unMarErr := json.Unmarshal([]byte(data), &bridgeRecord); unMarErr != nil {
				continue
			}
			sourceChainId := bridgeRecord.SourceChainId
			destChainId := bridgeRecord.DestChainId
			amount := bridgeRecord.Amount.String()
			fee := bridgeRecord.Fee.String()
			nonce := bridgeRecord.Nonce.String()
			to := bridgeRecord.To.String()
			tokenAddress := bridgeRecord.SourceTokenAddress.String()
			transfer, err := rl.bridgeRpcService.CrossChainTransfer(sourceChainId, destChainId, amount, to, tokenAddress, fee, nonce)
			if err != nil {
				continue
			}
			log.Println("CrossChainTransfer", "transfer", transfer.Success)
			// todo if call rpc fail ,need add retry times
			if transfer.Success {
				rl.db.BridgeMsgSent.UpdateCrossStatus(v.TxHash.String())
			}
		}
	}
	return nil
}

func (rl *RelayerListener) ChangeTransferStatus() error {
	list := rl.db.BridgeMsgSent.GetCanChangeTransStatusList()
	if len(list) > 0 {
		for _, v := range list {
			data := v.Data
			var bridgeRecord relayer.BridgeRecords
			if unMarErr := json.Unmarshal([]byte(data), &bridgeRecord); unMarErr != nil {
				continue
			}
			sourceChainId := bridgeRecord.SourceChainId
			destChainId := bridgeRecord.DestChainId
			hash := v.DestHash.String()
			changeStatus, err := rl.bridgeRpcService.ChangeTransferStatus(sourceChainId, destChainId, hash)
			if err != nil {
				continue
			}
			log.Println("ChangeTransferStatus", "changeStatus", changeStatus.Success)
			// todo if call rpc fail ,need add retry times
			if changeStatus.Success {
				rl.db.BridgeMsgSent.UpdateChangeStatus(v.TxHash.String())
			}
		}
	}
	return nil
}
