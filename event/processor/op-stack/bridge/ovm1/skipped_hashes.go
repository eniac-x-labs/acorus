package ovm1

import "github.com/ethereum/go-ethereum/common"

var PortalWithdrawalTransactions = map[common.Hash]bool{
	common.HexToHash("0x877e8b0417ad5ac9a8f60b16588ef83ab17ca98054dff017b4fc89088df7c194"): true,
}

var L1RelayedMessages = map[common.Hash]bool{
	common.HexToHash("0xba6348e51701c2637723ea761561794442861e78c471282b3a982a59cadccfc5"): true,
}
