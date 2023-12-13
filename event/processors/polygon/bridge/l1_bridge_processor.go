package bridge

import (
	"github.com/cornerstone-labs/acorus/event/processors/polygon/abi"
	"github.com/cornerstone-labs/acorus/event/processors/polygon/utils"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"

	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	l1_l2 "github.com/cornerstone-labs/acorus/database/event/l1-l2"
	"github.com/cornerstone-labs/acorus/database/worker"
)

func L1Deposit(polygonBridge *abi.Polygonzkevmbridge, contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1DepositETHSig := utils.DepositEventSignatureHash
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1DepositETHSig}
	depositEvents, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l1ToL2s := make([]worker.L1ToL2, len(depositEvents))
	transactionDeposits := make([]l1_l2.L1TransactionDeposit, len(depositEvents))
	if len(depositEvents) > 0 {
		log.Info("detected transaction deposits", "size", len(depositEvents))
	}
	for i := range depositEvents {
		DepositTx := depositEvents[i]
		rlpLog := *DepositTx.RLPLog
		deposit := utils.Deposit{}

		d, unpackErr := polygonBridge.ParseBridgeEvent(rlpLog)
		if unpackErr != nil {
			return unpackErr
		}

		deposit.Amount = d.Amount
		deposit.BlockNumber = rlpLog.BlockNumber
		deposit.OriginalNetwork = uint(d.OriginNetwork)
		deposit.DestinationAddress = d.DestinationAddress
		deposit.DestinationNetwork = uint(d.DestinationNetwork)
		deposit.OriginalAddress = d.OriginAddress
		deposit.DepositCount = uint(d.DepositCount)
		deposit.TxHash = rlpLog.TxHash
		deposit.Metadata = d.Metadata
		deposit.LeafType = d.LeafType

		transactionDeposits[i] = l1_l2.L1TransactionDeposit{
			SourceHash:           DepositTx.TransactionHash,
			InitiatedL1EventGUID: DepositTx.GUID,
			Tx: l1_l2.Transaction{
				FromAddress: deposit.DestinationAddress,
				ToAddress:   deposit.DestinationAddress,
				ETHAmount:   big.NewInt(0),
				ERC20Amount: big.NewInt(0),
				Data:        deposit.Metadata,
				Timestamp:   DepositTx.Timestamp,
			},
			GasLimit: big.NewInt(0),
		}
		l1ToL2s[i] = worker.L1ToL2{
			GUID:              uuid.New(),
			L1BlockHash:       rlpLog.BlockHash,
			QueueIndex:        nil,
			L1TransactionHash: rlpLog.TxHash,
			L2TransactionHash: common.Hash{},
			L1TxOrigin:        common.Hash{},
			Status:            0,
			FromAddress:       deposit.DestinationAddress,
			ToAddress:         deposit.DestinationAddress,
			L1TokenAddress:    deposit.OriginalAddress,
			L2TokenAddress:    common.Address{},
			ETHAmount:         big.NewInt(0),
			ERC20Amount:       big.NewInt(0),
			GasLimit:          big.NewInt(0),
			Timestamp:         int64(DepositTx.Timestamp),
			AssetType:         int64(common2.ETH),
			MsgHash:           common.Hash{},
		}
		if deposit.OriginalAddress.String() == utils.L1_ETH {
			transactionDeposits[i].Tx.ETHAmount = deposit.Amount
			l1ToL2s[i].ETHAmount = deposit.Amount
			l1ToL2s[i].AssetType = int64(common2.ETH)
		} else {
			transactionDeposits[i].Tx.ERC20Amount = deposit.Amount
			l1ToL2s[i].ERC20Amount = deposit.Amount
			l1ToL2s[i].AssetType = int64(common2.ERC20)
		}
	}
	if len(l1ToL2s) > 0 {
		if err := db.BridgeTransactions.StoreL1TransactionDeposits(transactionDeposits); err != nil {
			return err
		}
		//marshal, _ := json.Marshal(l1ToL2s)
		if err := db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s); err != nil {
			return err
		}
	}
	return nil
}
