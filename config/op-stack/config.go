package op_stack

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/event/processors/op-stack/mantle/op-bindings/predeploys"
)

type L1Contracts struct {
	AddressManager                  common.Address `toml:"address-manager"`
	SystemConfigProxy               common.Address `toml:"system-config"`
	OptimismPortalProxy             common.Address `toml:"optimism-portal"`
	L2OutputOracleProxy             common.Address `toml:"l2-output-oracle"`
	L1CrossDomainMessengerProxy     common.Address `toml:"l1-cross-domain-messenger"`
	L1StandardBridgeProxy           common.Address `toml:"l1-standard-bridge"`
	L1ERC721BridgeProxy             common.Address `toml:"l1-erc721-bridge"`
	LegacyCanonicalTransactionChain common.Address `toml:"-"`
	LegacyStateCommitmentChain      common.Address `toml:"-"`
}

func (c L1Contracts) ForEach(cb func(string, common.Address) error) error {
	contracts := reflect.ValueOf(c)
	fields := reflect.VisibleFields(reflect.TypeOf(c))
	for _, field := range fields {
		addr := (contracts.FieldByName(field.Name).Interface()).(common.Address)
		if err := cb(field.Name, addr); err != nil {
			return err
		}
	}
	return nil
}

type L2Contracts struct {
	L2ToL1MessagePasser    common.Address
	L2CrossDomainMessenger common.Address
	L2StandardBridge       common.Address
	L2ERC721Bridge         common.Address
}

func L2ContractsFromPredeploys() L2Contracts {
	return L2Contracts{
		L2ToL1MessagePasser:    predeploys.L2ToL1MessagePasserAddr,
		L2CrossDomainMessenger: predeploys.L2CrossDomainMessengerAddr,
		L2StandardBridge:       predeploys.L2StandardBridgeAddr,
		L2ERC721Bridge:         predeploys.L2ERC721BridgeAddr,
	}
}

func (c L2Contracts) ForEach(cb func(string, common.Address) error) error {
	contracts := reflect.ValueOf(c)
	fields := reflect.VisibleFields(reflect.TypeOf(c))
	for _, field := range fields {
		addr := (contracts.FieldByName(field.Name).Interface()).(common.Address)
		if err := cb(field.Name, addr); err != nil {
			return err
		}
	}
	return nil
}

type OpContracts struct {
	Preset      int
	L1Contracts L1Contracts `toml:"l1-contracts"`
	L2Contracts L2Contracts `toml:"l2-contracts"`
}
