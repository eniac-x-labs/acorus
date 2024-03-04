package common

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	L1BedrockStartHeight = big.NewInt(0)
	L2BedrockStartHeight = big.NewInt(0)

	LegacyStateCommitmentChain      = common.HexToAddress("0x0000000000000000000000000000000000000000")
	LegacyCanonicalTransactionChain = common.HexToAddress("0x0000000000000000000000000000000000000000")

	OptimismPortalProxy         = common.Address{}
	L1CrossDomainMessengerProxy = common.Address{}
	L1StandardBridgeProxy       = common.Address{}
	L2OutputOracleProxy         = common.Address{}

	L2ToL1MessagePasser    = common.HexToAddress("0x4200000000000000000000000000000000000016")
	L2CrossDomainMessenger = common.HexToAddress("0x4200000000000000000000000000000000000007")
	L2StandardBridge       = common.HexToAddress("0x4200000000000000000000000000000000000010")

	ETHWithdrawEventSignature   = "0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5"
	ERC20WithdrawEventSignature = "0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf"
	IsMainnet                   = true
)

func InitAddress(isMainNet bool) {
	IsMainnet = isMainNet
	getOptimismPortalProxy()
	getL1CrossDomainMessengerProxy()
	getL1StandardBridgeProxy()
	getL2OutputOracleProxy()
}

func getOptimismPortalProxy() {
	if IsMainnet {
		OptimismPortalProxy = common.HexToAddress("0x9168765EE952de7C6f8fC6FaD5Ec209B960b7622")
	} else {
		OptimismPortalProxy = common.HexToAddress("0x7FD7eEA37c53ABf356cc80e71144D62CD8aF27d3")
	}
}

func getL1CrossDomainMessengerProxy() {
	if IsMainnet {
		L1CrossDomainMessengerProxy = common.HexToAddress("0x635ba609680c55C3bDd0B3627b4c5dB21b13c310")
	} else {
		L1CrossDomainMessengerProxy = common.HexToAddress("0x7Ad11bB9216BC9Dc4CBd488D7618CbFD433d1E75")

	}
}

func getL1StandardBridgeProxy() {
	if IsMainnet {
		L1StandardBridgeProxy = common.HexToAddress("0x3B95bC951EE0f553ba487327278cAc44f29715E5")
	} else {
		L1StandardBridgeProxy = common.HexToAddress("0x4638aC6b5727a8b9586D3eba5B44Be4b74ED41Fc")

	}
}

func getL2OutputOracleProxy() {
	if IsMainnet {
		L2OutputOracleProxy = common.HexToAddress("0x30c789674ad3B458886BBC9abf42EEe19EA05C1D")
	} else {
		L2OutputOracleProxy = common.HexToAddress("0x8553D4d201ef97F2b76A28F5E543701b25e55B1b")
	}
}

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
