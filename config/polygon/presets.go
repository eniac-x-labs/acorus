package polygon

import "github.com/ethereum/go-ethereum/common"

type Preset struct {
	Name        string
	PlContracts PlContracts
}

var Presets = map[int]Preset{
	137: {
		Name: "polygon",
		PlContracts: PlContracts{
			Preset: 137,
			L1Contracts: L1Contracts{
				PolygonZKEVMBridgeAddr:         common.HexToAddress("0x2a3DD3EB832aF982ec71669E178424b10Dca2EDe"),
				PolygonZkEVMGlobalExitRootAddr: common.HexToAddress("0x580bda1e7A0CFAe92Fa7F6c20A3794F169CE3CFb"),
				PolygonRollupManagerAddr:       common.HexToAddress("0x580bda1e7A0CFAe92Fa7F6c20A3794F169CE3CFb"),
			},
			L2Contracts: L2Contracts{
				PolygonZKEVMBridgeAddr:         common.HexToAddress("0x2a3DD3EB832aF982ec71669E178424b10Dca2EDe"),
				PolygonZkEVMGlobalExitRootAddr: common.HexToAddress("0x580bda1e7A0CFAe92Fa7F6c20A3794F169CE3CFb"),
				PolygonRollupManagerAddr:       common.HexToAddress("0x580bda1e7A0CFAe92Fa7F6c20A3794F169CE3CFb"),
			},
		},
	},
}
