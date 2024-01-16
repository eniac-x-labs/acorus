package bridge

import (
	"math/big"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/common"

	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/polygon/abi"
	"github.com/cornerstone-labs/acorus/event/polygon/utils"
)

func L2Withdraw(polygonBridge *abi.Polygonzkevmbridge, event event.ContractEvent) (*worker.L2ToL1, error) {
	rlpLog := event.RLPLog
	w, unpackErr := polygonBridge.ParseBridgeEvent(*rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	l2ToL1 := worker.L2ToL1{
		GUID:              uuid.New(),
		L1FinalizeTxHash:  common.Hash{},
		L2TransactionHash: rlpLog.TxHash,
		L1ProveTxHash:     common.Hash{},
		Status:            0,
		TimeLeft:          big.NewInt(0),
		FromAddress:       w.DestinationAddress,
		ToAddress:         w.DestinationAddress,
		L1TokenAddress:    w.OriginAddress,
		L2TokenAddress:    common.Address{},
		ETHAmount:         big.NewInt(0),
		GasLimit:          big.NewInt(0),
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ETH),
		MsgNonce:          big.NewInt(0),
		MessageHash:       common.Hash{},
		ClaimedIndex:      int64(w.DepositCount),
		TokenAmounts:      "0",
	}

	if w.OriginAddress.String() == utils.L1_ETH {
		l2ToL1.ETHAmount = w.Amount
		l2ToL1.AssetType = int64(common2.ETH)
	} else {
		l2ToL1.TokenAmounts = w.Amount.String()
		l2ToL1.AssetType = int64(common2.ERC20)
	}
	return &l2ToL1, nil
}
