package common

import (
	common3 "github.com/cornerstone-labs/acorus/common"
	"github.com/cornerstone-labs/acorus/common/global_const"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	L1BedrockStartHeight = big.NewInt(0)
	L2BedrockStartHeight = big.NewInt(0)

	LegacyStateCommitmentChain      = common.HexToAddress("0x0000000000000000000000000000000000000000")
	LegacyCanonicalTransactionChain = common.HexToAddress("0x0000000000000000000000000000000000000000")

	OptimismPortalProxy         = getOptimismPortalProxy()
	L1CrossDomainMessengerProxy = getL1CrossDomainMessengerProxy()
	L1StandardBridgeProxy       = getL1StandardBridgeProxy()
	L2OutputOracleProxy         = getL2OutputOracleProxy()

	L2ToL1MessagePasser    = common.HexToAddress("0x4200000000000000000000000000000000000016")
	L2CrossDomainMessenger = common.HexToAddress("0x4200000000000000000000000000000000000007")
	L2StandardBridge       = common.HexToAddress("0x4200000000000000000000000000000000000010")

	ETHWithdrawEventSignature   = "0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5"
	ERC20WithdrawEventSignature = "0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf"
	IsMainnet                   = true
)

func InitAddress() {
	rpcs := common3.ConfigInfo.RPCs
	for i := 0; i < len(rpcs); i++ {
		chainId := rpcs[i].ChainId
		if chainId == global_const.OpChinId {
			IsMainnet = rpcs[i].IsMainnet
			break
		}
	}
}

func getOptimismPortalProxy() common.Address {
	if IsMainnet {
		return common.HexToAddress("0xbEb5Fc579115071764c7423A4f12eDde41f106Ed")
	}
	return common.HexToAddress("0x16Fc5058F25648194471939df75CF27A2fdC48BC")
}

func getL1CrossDomainMessengerProxy() common.Address {
	if IsMainnet {
		return common.HexToAddress("0x25ace71c97B33Cc4729CF772ae268934F7ab5fA1")
	}
	return common.HexToAddress("0x58Cc85b8D04EA49cC6DBd3CbFFd00B4B8D6cb3ef")
}

func getL1StandardBridgeProxy() common.Address {
	if IsMainnet {
		return common.HexToAddress("0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1")
	}
	return common.HexToAddress("0xFBb0621E0B23b5478B630BD55a5f21f67730B0F1")
}

func getL2OutputOracleProxy() common.Address {
	if IsMainnet {
		common.HexToAddress("0xdfe97868233d1aa22e815a266982f2cf17685a27")
	}
	return common.HexToAddress("0x90E9c4f8a994a250F6aEfd61CAFb4F2e895D458F")
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
