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
		OptimismPortalProxy = common.HexToAddress("0x49048044D57e1C92A77f79988d21Fa8fAF74E97e")
	} else {
		OptimismPortalProxy = common.HexToAddress("0x49f53e41452C74589E85cA1677426Ba426459e85")
	}
}

func getL1CrossDomainMessengerProxy() {
	if IsMainnet {
		L1CrossDomainMessengerProxy = common.HexToAddress("0x866E82a600A1414e583f7F13623F1aC5d58b0Afa")
	} else {
		L1CrossDomainMessengerProxy = common.HexToAddress("0xC34855F4De64F1840e5686e64278da901e261f20")

	}
}

func getL1StandardBridgeProxy() {
	if IsMainnet {
		L1StandardBridgeProxy = common.HexToAddress("0x3154Cf16ccdb4C6d922629664174b904d80F2C35")
	} else {
		L1StandardBridgeProxy = common.HexToAddress("0xfd0Bf71F60660E2f608ed56e1659C450eB113120")

	}
}

func getL2OutputOracleProxy() {
	if IsMainnet {
		L2OutputOracleProxy = common.HexToAddress("0x56315b90c40730925ec5485cf004d835058518A0")
	} else {
		L2OutputOracleProxy = common.HexToAddress("0x84457ca9D0163FbC4bbfe4Dfbb20ba46e48DF254")
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
