package ovm1

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/event/base/bridge/ovm1/crossdomain"
)

func CalcTransaction(legacyWithdrawal *crossdomain.LegacyWithdrawal, l1CrossdomainMessengerAddress *common.Address, l2ChainID *big.Int) (common.Hash, error) {
	withdrawal, err := crossdomain.CalcWithdrawalHash(legacyWithdrawal, l1CrossdomainMessengerAddress, l2ChainID)
	if err != nil {
		return common.Hash{}, err
	}
	hash, err := withdrawal.Hash() // WithdrawalHash
	if err != nil {
		return common.Hash{}, err
	}
	log.Println("withdraw hash", hash)
	return hash, nil
}
