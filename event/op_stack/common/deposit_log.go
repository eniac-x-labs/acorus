package common

import (
	"fmt"
	"math/big"

	"github.com/holiman/uint256"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	DepositEventABI      = "TransactionDeposited(address,address,uint256,bytes)"
	DepositEventABIHash  = crypto.Keccak256Hash([]byte(DepositEventABI))
	DepositEventVersion0 = common.Hash{}
)

func UnmarshalDepositLogEvent(ev *types.Log) (*DepositTx, error) {
	if len(ev.Topics) != 4 {
		return nil, fmt.Errorf("expected 4 event topics (event identity, indexed from, indexed to, indexed version), got %d", len(ev.Topics))
	}
	if ev.Topics[0] != DepositEventABIHash {
		return nil, fmt.Errorf("invalid deposit event selector: %s, expected %s", ev.Topics[0], DepositEventABIHash)
	}
	if len(ev.Data) < 64 {
		return nil, fmt.Errorf("incomplate opaqueData slice header (%d bytes): %x", len(ev.Data), ev.Data)
	}
	if len(ev.Data)%32 != 0 {
		return nil, fmt.Errorf("expected log data to be multiple of 32 bytes: got %d bytes", len(ev.Data))
	}

	// indexed 0
	from := common.BytesToAddress(ev.Topics[1][12:])
	// indexed 1
	to := common.BytesToAddress(ev.Topics[2][12:])
	// indexed 2
	version := ev.Topics[3]
	// unindexed data
	// Solidity serializes the event's Data field as follows:
	// abi.encode(abi.encodPacked(uint256 mint, uint256 value, uint64 gasLimit, uint8 isCreation, bytes data))
	// Thus the first 32 bytes of the Data will give us the offset of the opaqueData,
	// which should always be 0x20.
	var opaqueContentOffset uint256.Int
	opaqueContentOffset.SetBytes(ev.Data[0:32])
	if !opaqueContentOffset.IsUint64() || opaqueContentOffset.Uint64() != 32 {
		return nil, fmt.Errorf("invalid opaqueData slice header offset: %d", opaqueContentOffset.Uint64())
	}
	// The next 32 bytes indicate the length of the opaqueData content.
	var opaqueContentLength uint256.Int
	opaqueContentLength.SetBytes(ev.Data[32:64])
	// Make sure the length is an uint64, it's not larger than the remaining data, and the log is using minimal padding (i.e. can't add 32 bytes without exceeding data)
	if !opaqueContentLength.IsUint64() || opaqueContentLength.Uint64() > uint64(len(ev.Data)-64) || opaqueContentLength.Uint64()+32 <= uint64(len(ev.Data)-64) {
		return nil, fmt.Errorf("invalid opaqueData slice header length: %d", opaqueContentLength.Uint64())
	}
	// The remaining data is the opaqueData which is tightly packed
	// and then padded to 32 bytes by the EVM.
	opaqueData := ev.Data[64 : 64+opaqueContentLength.Uint64()]

	var dep DepositTx

	source := UserDepositSource{
		L1BlockHash: ev.BlockHash,
		LogIndex:    uint64(ev.Index),
	}
	dep.SourceHash = source.SourceHash()
	dep.From = from
	dep.IsSystemTransaction = false

	var err error
	switch version {
	case DepositEventVersion0:
		err = unmarshalDepositVersion0(&dep, to, opaqueData)
	default:
		return nil, fmt.Errorf("invalid deposit version, got %s", version)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to decode deposit (version %s): %w", version, err)
	}
	return &dep, nil
}

func unmarshalDepositVersion0(dep *DepositTx, to common.Address, opaqueData []byte) error {
	if len(opaqueData) < 32+32+8+1 {
		return fmt.Errorf("unexpected opaqueData length: %d", len(opaqueData))
	}
	offset := uint64(0)

	// uint256 mint
	dep.Mint = new(big.Int).SetBytes(opaqueData[offset : offset+32])
	// 0 mint is represented as nil to skip minting code
	if dep.Mint.Cmp(new(big.Int)) == 0 {
		dep.Mint = nil
	}
	offset += 32

	// uint256 value
	dep.Value = new(big.Int).SetBytes(opaqueData[offset : offset+32])
	offset += 32

	// uint64 gas
	gas := new(big.Int).SetBytes(opaqueData[offset : offset+8])
	if !gas.IsUint64() {
		return fmt.Errorf("bad gas value: %x", opaqueData[offset:offset+8])
	}
	dep.Gas = gas.Uint64()
	offset += 8

	// uint8 isCreation
	// isCreation: If the boolean byte is 1 then dep.To will stay nil,
	// and it will create a contract using L2 account nonce to determine the created address.
	if opaqueData[offset] == 0 {
		dep.To = &to
	}
	offset += 1

	// The remainder of the opaqueData is the transaction data (without length prefix).
	// The data may be padded to a multiple of 32 bytes
	txDataLen := uint64(len(opaqueData)) - offset

	// remaining bytes fill the data
	dep.Data = opaqueData[offset : offset+txDataLen]

	return nil
}
