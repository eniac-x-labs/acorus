package polygon_worker

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
)

type WorkerProcessor struct {
	db      *database.DB
	chainId string
	tasks   tasks.Group
}

func NewWorkerProcessor(db *database.DB, chainId string, shutdown context.CancelCauseFunc) (*WorkerProcessor, error) {
	workerProcessor := WorkerProcessor{
		db:      db,
		chainId: chainId,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in worker processor: %w", err))
		}},
	}
	return &workerProcessor, nil
}

func (b *WorkerProcessor) WorkerStart() error {
	tickerRunMarked := time.NewTicker(time.Second * 5)
	b.tasks.Go(func() error {
		for range tickerRunMarked.C {
			err := b.marked()
			if err != nil {
				log.Println("marked ", "err", err)
				continue
			}
		}
		return nil
	})
	return nil
}

func (b *WorkerProcessor) marked() error {
	var errs error
	if err := b.markedL1ToL2Finalized(); err != nil {
		errors.Join(errs, err)
	}
	if err := b.markedL2ToL1Finalized(); err != nil {
		errors.Join(errs, err)
	}
	return errs
}

func (b *WorkerProcessor) markedL1ToL2Finalized() error {
	log.Println("start marked l1 to l2 finalized")
	finalizedList, err := b.db.RelayMessage.RelayMessageUnRelatedList(b.chainId)
	if err != nil {
		return err
	}
	var depositL2ToL1List []worker.L1ToL2
	var needMarkDepositList []event.RelayMessage
	for i := range finalizedList {
		finalized := finalizedList[i]
		l1l2Tx := worker.L1ToL2{
			L2TransactionHash: finalized.RelayTransactionHash,
			L2BlockNumber:     finalized.BlockNumber,
			MessageHash:       finalized.MessageHash,
		}
		withdrawTx, _ := b.db.L1ToL2.L1ToL2TransactionDeposit(b.chainId, finalized.MessageHash)
		if withdrawTx != nil {
			depositL2ToL1List = append(depositL2ToL1List, l1l2Tx)
			needMarkDepositList = append(needMarkDepositList, finalized)
		}
	}
	if err := b.db.Transaction(func(tx *database.DB) error {
		if len(depositL2ToL1List) > 0 {
			if err := b.db.L1ToL2.MarkL1ToL2TransactionDepositFinalized(b.chainId, depositL2ToL1List); err != nil {
				log.Println("Marked l2 to l1 transaction withdraw proven fail", "err", err)
				return err
			}
			if err := b.db.RelayMessage.MarkedRelayMessageRelated(b.chainId, needMarkDepositList); err != nil {
				log.Println("Marked withdraw proven related fail", "err", err)
				return err
			}
			log.Println("marked deposit transaction success", "deposit size", len(depositL2ToL1List), "marked size", len(needMarkDepositList))
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (b *WorkerProcessor) markedL2ToL1Finalized() error {
	log.Println("start marked l2 to l1 finalized")
	withdrawList, err := b.db.WithdrawFinalized.WithdrawFinalizedUnRelatedList(b.chainId)
	if err != nil {
		log.Println("fetch withdraw finalized un-related list fail", "err", err)
		return err
	}
	var withdrawL2ToL1List []worker.L2ToL1
	var withdrawL2ToL1ListV0 []worker.L2ToL1
	var needMarkWithdrawList []event.WithdrawFinalized
	var needMarkWithdrawListV0 []event.WithdrawFinalized
	for i := range withdrawList {
		finalizedTxn := withdrawList[i]
		l2l1Tx := worker.L2ToL1{
			L1FinalizeTxHash: finalizedTxn.FinalizedTransactionHash,
			L1BlockNumber:    finalizedTxn.BlockNumber,
		}
		withdrawTx, _ := b.db.L2ToL1.L2ToL1TransactionMsgHash(b.chainId, finalizedTxn.MessageHash)
		if withdrawTx != nil {
			if withdrawTx.Version != 0 {
				withdrawL2ToL1List = append(withdrawL2ToL1List, l2l1Tx)
				needMarkWithdrawList = append(needMarkWithdrawList, finalizedTxn)
			} else {
				withdrawL2ToL1ListV0 = append(withdrawL2ToL1ListV0, l2l1Tx)
				needMarkWithdrawListV0 = append(needMarkWithdrawListV0, finalizedTxn)
			}

		}
	}
	if err := b.db.Transaction(func(tx *database.DB) error {
		if len(withdrawL2ToL1List) > 0 {
			if err := b.db.L2ToL1.MarkL2ToL1TransactionWithdrawalFinalized(b.chainId, withdrawL2ToL1List); err != nil {
				log.Println("Marked l2 to l1 transaction withdraw finalized fail", "err", err)
				return err
			}
			if err := b.db.WithdrawFinalized.MarkedWithdrawFinalizedRelated(b.chainId, needMarkWithdrawList); err != nil {
				log.Println("Marked withdraw finalized related fail", "err", err)
				return err
			}
			log.Println("marked finalized transaction success", "withdraw size", len(withdrawList), "marked size", len(needMarkWithdrawList))
		}
		if len(withdrawL2ToL1ListV0) > 0 {
			if err := b.db.L2ToL1.MarkL2ToL1TransactionWithdrawalFinalizedV0(b.chainId, withdrawL2ToL1ListV0); err != nil {
				log.Println("Marked l2 to l1 transaction withdraw proven fail", "err", err)
				return err
			}
			if err := b.db.WithdrawFinalized.MarkedWithdrawFinalizedRelated(b.chainId, needMarkWithdrawListV0); err != nil {
				log.Println("Marked withdraw proven related fail", "err", err)
				return err
			}
			log.Println("marked proven v0 transaction success", "withdraw size", len(withdrawL2ToL1ListV0), "marked size", len(needMarkWithdrawList))
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
