package predeploys

import "github.com/ethereum/go-ethereum/common"

const (
	BVM_ETH        = "0xdEAddEaDdeadDEadDEADDEAddEADDEAddead1111"
	LegacyERC20MNT = "0xDeadDeAddeAddEAddeadDEaDDEAdDeaDDeAD0000"
)

var (
	BVM_ETHAddr        = common.HexToAddress(BVM_ETH)
	LegacyERC20MNTAddr = common.HexToAddress(LegacyERC20MNT)
)
