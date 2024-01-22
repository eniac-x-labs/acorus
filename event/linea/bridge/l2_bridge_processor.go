package bridge

import (
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
)

func L2SentMessageEvent(event event.ContractEvent) (*worker.L1ToL2, error) {
	return nil, nil
}

func L2ClaimedMessageEvent(event event.ContractEvent) (*worker.L1ToL2, error) {
	return nil, nil
}
