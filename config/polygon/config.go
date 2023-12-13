package polygon

import (
	"github.com/ethereum/go-ethereum/common"
	"reflect"
)

type L1Contracts struct {
	PolygonZkEVMGlobalExitRootAddr common.Address `toml:"polygonZkEVMGlobalExitRootAddress"`
	PolygonRollupManagerAddr       common.Address `toml:"polygonRollupManagerAddress"`
	PolygonZKEVMBridgeAddr         common.Address `toml:"PolygonZKEVMBridgeAddress"`
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
	PolygonZkEVMGlobalExitRootAddr common.Address `toml:"polygonZkEVMGlobalExitRootAddress"`
	PolygonRollupManagerAddr       common.Address `toml:"polygonRollupManagerAddress"`
	PolygonZKEVMBridgeAddr         common.Address `toml:"PolygonZKEVMBridgeAddr"`
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

type PlContracts struct {
	Preset      int
	L1Contracts L1Contracts `toml:"l1-contracts"`
	L2Contracts L2Contracts `toml:"l2-contracts"`
}
