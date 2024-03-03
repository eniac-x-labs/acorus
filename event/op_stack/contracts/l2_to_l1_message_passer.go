package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/utils"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
)

type L2ToL1MessagePasserMessagePassed struct {
	Event          *event.ContractEvent
	WithdrawalHash common.Hash
	GasLimit       *big.Int
	Nonce          *big.Int
	FromAddress    common.Address
	ToAddress      common.Address
	ETHAmount      *big.Int
	ERC20Amount    *big.Int
	Data           utils.Bytes
	Timestamp      uint64
}

func L2ToL1MessagePasserMessagePassedEvents(contractAddress common.Address, db *database.DB,
	fromHeight, toHeight *big.Int, l1ChainId, l2ChainId string) ([]L2ToL1MessagePasserMessagePassed, error) {
	l2ToL1MessagePasserAbi, err := bindings.L2ToL1MessagePasserMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	messagePassedAbi := l2ToL1MessagePasserAbi.Events["MessagePassed"]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: messagePassedAbi.ID}
	messagePassedEvents, err := db.ContractEvents.ContractEventsWithFilter(l2ChainId, contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	messagesPassed := make([]L2ToL1MessagePasserMessagePassed, len(messagePassedEvents))
	for i := range messagePassedEvents {
		messagePassed := bindings.L2ToL1MessagePasserMessagePassed{Raw: *messagePassedEvents[i].RLPLog}
		err := UnpackLog(&messagePassed, messagePassedEvents[i].RLPLog, messagePassedAbi.Name, l2ToL1MessagePasserAbi)
		if err != nil {
			return nil, err
		}

		messagesPassed[i] = L2ToL1MessagePasserMessagePassed{
			Event:          &messagePassedEvents[i],
			WithdrawalHash: messagePassed.WithdrawalHash,
			Nonce:          messagePassed.Nonce,
			GasLimit:       messagePassed.GasLimit,
			FromAddress:    messagePassed.Sender,
			ToAddress:      messagePassed.Target,
			ETHAmount:      messagePassed.Value,
			Data:           messagePassed.Data,
			Timestamp:      messagePassedEvents[i].Timestamp,
			ERC20Amount:    messagePassed.Value,
		}
	}

	return messagesPassed, nil
}
