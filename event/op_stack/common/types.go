package common

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	LegacyStateCommitmentChain      = common.HexToAddress("0x2a3DD3EB832aF982ec71669E178424b10Dca2EDe")
	L2OutputOracleProxy             = common.HexToAddress("0x2a3DD3EB832aF982ec71669E178424b10Dca2EDe")
	OptimismPortalProxy             = common.HexToAddress("0x2a3DD3EB832aF982ec71669E178424b10Dca2EDe")
	L1CrossDomainMessengerProxy     = common.HexToAddress("0x2a3DD3EB832aF982ec71669E178424b10Dca2EDe")
	L1StandardBridgeProxy           = common.HexToAddress("0x2a3DD3EB832aF982ec71669E178424b10Dca2EDe")
	L2ToL1MessagePasser             = common.HexToAddress("0x2a3DD3EB832aF982ec71669E178424b10Dca2EDe")
	L2CrossDomainMessenger          = common.HexToAddress("0x2a3DD3EB832aF982ec71669E178424b10Dca2EDe")
	L2StandardBridge                = common.HexToAddress("0x2a3DD3EB832aF982ec71669E178424b10Dca2EDe")
	LegacyCanonicalTransactionChain = common.HexToAddress("0x2a3DD3EB832aF982ec71669E178424b10Dca2EDe")
	ETHWithdrawEventSignature       = "0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5"
	ERC20WithdrawEventSignature     = "0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf"
)

type DepositTx struct {
	SourceHash          common.Hash
	From                common.Address
	To                  *common.Address `rlp:"nil"`
	Mint                *big.Int        `rlp:"nil"`
	Value               *big.Int
	Gas                 uint64
	IsSystemTransaction bool
	EthValue            *big.Int `rlp:"nil"`
	Data                []byte
}
