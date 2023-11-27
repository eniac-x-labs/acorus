package scroll

import (
	"github.com/ethereum/go-ethereum/common"
	"reflect"
)

type L1Contracts struct {
	MessengerAddr          common.Address `toml:"MessengerAddr"`
	ETHGatewayAddr         common.Address `toml:"ETHGatewayAddr"`
	WETHGatewayAddr        common.Address `toml:"WETHGatewayAddr"`
	StandardERC20Gateway   common.Address `toml:"StandardERC20Gateway"`
	CustomERC20GatewayAddr common.Address `toml:"CustomERC20GatewayAddr"`
	ERC721GatewayAddr      common.Address `toml:"ERC721GatewayAddr"`
	ERC1155GatewayAddr     common.Address `toml:"ERC1155GatewayAddr"`
	USDCGatewayAddr        common.Address `toml:"USDCGatewayAddr"`
	LIDOGatewayAddr        common.Address `toml:"LIDOGatewayAddr"`
	DAIGatewayAddr         common.Address `toml:"DAIGatewayAddr"`
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
	MessengerAddr          common.Address `toml:"MessengerAddr"`
	ETHGatewayAddr         common.Address `toml:"ETHGatewayAddr"`
	WETHGatewayAddr        common.Address `toml:"WETHGatewayAddr"`
	StandardERC20Gateway   common.Address `toml:"StandardERC20Gateway"`
	CustomERC20GatewayAddr common.Address `toml:"CustomERC20GatewayAddr"`
	ERC721GatewayAddr      common.Address `toml:"ERC721GatewayAddr"`
	ERC1155GatewayAddr     common.Address `toml:"ERC1155GatewayAddr"`
	USDCGatewayAddr        common.Address `toml:"USDCGatewayAddr"`
	LIDOGatewayAddr        common.Address `toml:"LIDOGatewayAddr"`
	DAIGatewayAddr         common.Address `toml:"DAIGatewayAddr"`
	L2GasOracle            common.Address `toml:"L2GasOracle"`
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

type SclContracts struct {
	Preset      int
	L1Contracts L1Contracts `toml:"l1-contracts"`
	L2Contracts L2Contracts `toml:"l2-contracts"`
}
