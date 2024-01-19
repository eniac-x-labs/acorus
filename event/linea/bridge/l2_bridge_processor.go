package bridge

import (
	"github.com/cornerstone-labs/acorus/database"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func L2SentMessageEvent(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	return nil
}

func L2ClaimedMessageEvent(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	return nil
}
