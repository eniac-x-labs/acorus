package scroll

import "github.com/ethereum/go-ethereum/common"

type Preset struct {
	Name           string
	LineaContracts LineaContracts
}

var Presets = map[int]Preset{
	59144: {
		Name: "linea",
		LineaContracts: LineaContracts{
			Preset: 1,
			L1Contracts: L1Contracts{
				MessengerAddr: common.HexToAddress("0xd19d4B5d358258f05D7B411E21A1460D11B0876F"),
			},
			L2Contracts: L2Contracts{
				MessengerAddr: common.HexToAddress("0x508Ca82Df566dCD1B0DE8296e70a96332cD644ec"),
			},
		},
	},
}
