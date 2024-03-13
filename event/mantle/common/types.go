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

	MNTWithdrawEventSignature   = "0x74bbfec0d26a17c2367408038090a9a4e1cd1671129dc8fdf57f146a499fe3d5"
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
		OptimismPortalProxy = common.HexToAddress("0xbEb5Fc579115071764c7423A4f12eDde41f106Ed")
	} else {
		OptimismPortalProxy = common.HexToAddress("0xB3db4bd5bc225930eD674494F9A4F6a11B8EFBc8")
	}
}

func getL1CrossDomainMessengerProxy() {
	if IsMainnet {
		L1CrossDomainMessengerProxy = common.HexToAddress("0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1")
	} else {
		L1CrossDomainMessengerProxy = common.HexToAddress("0x37dAC5312e31Adb8BB0802Fc72Ca84DA5cDfcb4c")

	}
}

func getL1StandardBridgeProxy() {
	if IsMainnet {
		L1StandardBridgeProxy = common.HexToAddress("0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1")
	} else {
		L1StandardBridgeProxy = common.HexToAddress("0x21F308067241B2028503c07bd7cB3751FFab0Fb2")

	}
}

func getL2OutputOracleProxy() {
	if IsMainnet {
		L2OutputOracleProxy = common.HexToAddress("0xdfe97868233d1aa22e815a266982f2cf17685a27")
	} else {
		L2OutputOracleProxy = common.HexToAddress("0x4121dc8e48Bc6196795eb4867772A5e259fecE07")
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
