package scroll

import (
	"github.com/ethereum/go-ethereum/common"
	"reflect"
)

type L1Contracts struct {
	MessengerAddr common.Address `toml:"MessengerAddr"`
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
	MessengerAddr common.Address `toml:"MessengerAddr"`
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

type LineaContracts struct {
	Preset      int
	L1Contracts L1Contracts `toml:"l1-contracts"`
	L2Contracts L2Contracts `toml:"l2-contracts"`
}
