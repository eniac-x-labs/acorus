package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"time"
)

var (
	DepositEventSignatureHash  = crypto.Keccak256Hash([]byte("BridgeEvent(uint8,uint32,address,uint32,address,uint256,bytes,uint32)")) // Used in oldBridge as well
	WithdrawEventSignatureHash = crypto.Keccak256Hash([]byte("BridgeEvent(uint8,uint32,address,uint32,address,uint256,bytes,uint32)")) // Used in oldBridge as well
	ClaimEventSignatureHash    = crypto.Keccak256Hash([]byte("ClaimEvent(uint256,uint32,address,address,uint256)"))
	L1_ETH                     = "0x0000000000000000000000000000000000000000"
)

// Block struct
type Block struct {
	ID              uint64
	BlockNumber     uint64
	BlockHash       common.Hash
	ParentHash      common.Hash
	NetworkID       uint
	GlobalExitRoots []GlobalExitRoot
	Deposits        []Deposit
	Claims          []Claim
	Tokens          []TokenWrapped
	VerifiedBatches []VerifiedBatch
	ReceivedAt      time.Time
}

// GlobalExitRoot struct
type GlobalExitRoot struct {
	BlockID        uint64
	BlockNumber    uint64
	ExitRoots      []common.Hash
	GlobalExitRoot common.Hash
}

// Deposit struct
type Deposit struct {
	LeafType           uint8
	OriginalNetwork    uint
	OriginalAddress    common.Address
	Amount             *big.Int
	DestinationNetwork uint
	DestinationAddress common.Address
	DepositCount       uint
	BlockID            uint64
	BlockNumber        uint64
	NetworkID          uint
	TxHash             common.Hash
	Metadata           []byte
	// it is only used for the bridge service
	ReadyForClaim bool
}

type Withdraw struct {
	LeafType           uint8
	OriginalNetwork    uint
	OriginalAddress    common.Address
	Amount             *big.Int
	DestinationNetwork uint
	DestinationAddress common.Address
	DepositCount       uint
	BlockID            uint64
	BlockNumber        uint64
	NetworkID          uint
	TxHash             common.Hash
	Metadata           []byte
	// it is only used for the bridge service
	ReadyForClaim bool
}

// Claim struct
type Claim struct {
	MainnetFlag        bool
	RollupIndex        uint64
	Index              uint
	OriginalNetwork    uint
	OriginalAddress    common.Address
	Amount             *big.Int
	DestinationAddress common.Address
	BlockID            uint64
	BlockNumber        uint64
	NetworkID          uint
	TxHash             common.Hash
}

// TokenWrapped struct
type TokenWrapped struct {
	TokenMetadata
	OriginalNetwork      uint
	OriginalTokenAddress common.Address
	WrappedTokenAddress  common.Address
	BlockID              uint64
	BlockNumber          uint64
	NetworkID            uint
}

// TokenMetadata is a metadata of ERC20 token.
type TokenMetadata struct {
	Name     string
	Symbol   string
	Decimals uint8
}

type VerifiedBatch struct {
	BlockNumber   uint64
	BatchNumber   uint64
	RollupID      uint
	LocalExitRoot common.Hash
	TxHash        common.Hash
	StateRoot     common.Hash
	Aggregator    common.Address
}

// RollupExitLeaf struct
type RollupExitLeaf struct {
	ID       uint64
	BlockID  uint64
	Leaf     common.Hash
	RollupId uint
	Root     common.Hash
}
