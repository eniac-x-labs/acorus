package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cornerstone-labs/acorus/database/relayer"
	"log"
	"time"

	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/rpc/bridge"
)

type RelayerBridgeRelation struct {
	db               *database.DB
	bridgeRpcService bridge.BridgeRpcService
	resourceCtx      context.Context
	resourceCancel   context.CancelFunc
	tasks            tasks.Group
}

func NewRelayerBridgeRelation(
	db *database.DB,
	bridgeRpcService bridge.BridgeRpcService,
	shutdown context.CancelCauseFunc) (*RelayerBridgeRelation, error) {
	resCtx, resCancel := context.WithCancel(context.Background())
	return &RelayerBridgeRelation{
		db:               db,
		bridgeRpcService: bridgeRpcService,
		resourceCancel:   resCancel,
		resourceCtx:      resCtx,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
	}, nil
}

func (rbr *RelayerBridgeRelation) Start() error {
	tickerBridgeRel := time.NewTicker(10 * time.Second)
	// start relation worker
	rbr.tasks.Go(func() error {
		for range tickerBridgeRel.C {
			err := rbr.relationBridge()
			if err != nil {
				log.Println(" shutting down relationBridge")
				continue
			}
		}
		return nil
	})

	tickerCross := time.NewTicker(10 * time.Second)
	rbr.tasks.Go(func() error {
		for range tickerCross.C {
			err := rbr.CrossChainTransfer()
			if err != nil {
				log.Println("shutting down relationBridge")
				continue
			}
		}
		return nil
	})

	tickerTrans := time.NewTicker(10 * time.Second)
	rbr.tasks.Go(func() error {
		for range tickerTrans.C {
			err := rbr.ChangeTransferStatus()
			if err != nil {
				log.Println("shutting down relationBridge")
				continue
			}
		}
		return nil
	})
	return nil
}

func (rbr *RelayerBridgeRelation) relationBridge() error {
	if errRel := rbr.db.Transaction(func(tx *database.DB) error {
		// step 1
		log.Println("RelationClaim")
		err := rbr.db.BridgeFinalize.RelationClaim()
		if err != nil {
			log.Println("relationBridge failed", "err", err)
			return err
		}
		// step 2
		log.Println("RelationMsgHash")
		err = rbr.db.BridgeMsgHash.RelationMsgHash()
		if err != nil {
			log.Println("relationBridge failed", "err", err)
			return err
		}

		log.Println("CleanMsgSent")
		err = rbr.db.BridgeMsgSent.CleanMsgSent()
		if err != nil {
			return err
		}
		// step 3
		log.Println("RelationMsgSent")
		err = rbr.db.BridgeClaim.RelationMsgSent()
		if err != nil {
			log.Println("RelationMsgSent failed", "err", err)
			return err
		}

		// setp 4
		log.Println("RelationClaimData")
		err = rbr.db.BridgeRecord.RelationClaimData()
		if err != nil {
			log.Println("RelationClaimData failed", "err", err)
			return err
		}
		return nil
	}); errRel != nil {
		return errRel
	}
	if err := rbr.db.Transaction(func(tx *database.DB) error {

		list, err := rbr.db.BridgeMsgSent.GetCanSaveDecodeList()
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
			bridgeRecord.MsgHash = v.MsgHash
			bridgeRecord.Nonce = v.MsgNonce
			bridgeRecord.Fee = v.Fee
			bridgeRecordSaveList = append(bridgeRecordSaveList, bridgeRecord)
		}
		if len(bridgeRecordSaveList) > 0 {
			log.Println("StoreBridgeRecords")
			err = rbr.db.BridgeRecord.StoreBridgeRecords(bridgeRecordSaveList)
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

func (rbr *RelayerBridgeRelation) CrossChainTransfer() error {
	list := rbr.db.BridgeMsgSent.GetCanCrossDataList()
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
			fee := v.Fee.String()
			nonce := v.MsgNonce.String()
			to := bridgeRecord.ToAddress.String()
			tokenAddress := bridgeRecord.SourceTokenAddress.String()
			transfer, err := rbr.bridgeRpcService.CrossChainTransfer(sourceChainId, destChainId, amount, to, tokenAddress, fee, nonce)
			if err != nil {
				log.Println("CrossChainTransfer", "error", err)
				continue
			}
			log.Println("CrossChainTransfer", "transfer", transfer.Success)
			// todo if call rpc fail ,need add retry times
			if transfer.Success {
				rbr.db.BridgeMsgSent.UpdateCrossStatus(v.TxHash.String())
			}
		}
	}
	return nil
}

func (rbr *RelayerBridgeRelation) ChangeTransferStatus() error {
	list := rbr.db.BridgeMsgSent.GetCanChangeTransStatusList()
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
			changeStatus, err := rbr.bridgeRpcService.ChangeTransferStatus(sourceChainId, destChainId, hash)
			if err != nil {
				log.Println("ChangeTransferStatus", "error", err)
				continue
			}
			log.Println("ChangeTransferStatus", "changeStatus", changeStatus.Success)
			// todo if call rpc fail ,need add retry times
			if changeStatus.Success {
				rbr.db.BridgeMsgSent.UpdateChangeStatus(v.TxHash.String())
			}
		}
	}
	return nil
}
