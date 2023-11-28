package scroll

import "github.com/ethereum/go-ethereum/common"

type Preset struct {
	Name         string
	SclContracts SclContracts
}

var Presets = map[int]Preset{
	1: {
		Name: "eth",
		SclContracts: SclContracts{
			Preset: 1,
			L1Contracts: L1Contracts{
				MessengerAddr:          common.HexToAddress("0x6774Bcbd5ceCeF1336b5300fb5186a12DDD8b367"),
				ETHGatewayAddr:         common.HexToAddress("0x7F2b8C31F88B6006c382775eea88297Ec1e3E905"),
				WETHGatewayAddr:        common.HexToAddress("0x7AC440cAe8EB6328de4fA621163a792c1EA9D4fE"),
				StandardERC20Gateway:   common.HexToAddress("0xD8A791fE2bE73eb6E6cF1eb0cb3F36adC9B3F8f9"),
				CustomERC20GatewayAddr: common.HexToAddress("0xb2b10a289A229415a124EFDeF310C10cb004B6ff"),
				ERC721GatewayAddr:      common.HexToAddress("0x6260aF48e8948617b8FA17F4e5CEa2d21D21554B"),
				ERC1155GatewayAddr:     common.HexToAddress("0xb94f7F6ABcb811c5Ac709dE14E37590fcCd975B6"),
				USDCGatewayAddr:        common.HexToAddress("0x37ba659D6CC380D12Fb96567CC52FC8e1DF4E334"),
				LIDOGatewayAddr:        common.HexToAddress("0x892dDB2899325aBBA1fD00FDA8249B40Cbbc33F9"),
				DAIGatewayAddr:         common.HexToAddress("0xD8dD7787f89c7E6243AD32E0d0cCf460243C8130"),
			},
			L2Contracts: L2Contracts{
				MessengerAddr:          common.HexToAddress("0x781e90f1c8Fc4611c9b7497C3B47F99Ef6969CbC"),
				ETHGatewayAddr:         common.HexToAddress("0x6EA73e05AdC79974B931123675ea8F78FfdacDF0"),
				WETHGatewayAddr:        common.HexToAddress("0x7003E7B7186f0E6601203b99F7B8DECBfA391cf9"),
				StandardERC20Gateway:   common.HexToAddress("0xE2b4795039517653c5Ae8C2A9BFdd783b48f447A"),
				CustomERC20GatewayAddr: common.HexToAddress("0x64CCBE37c9A82D85A1F2E74649b7A42923067988"),
				ERC721GatewayAddr:      common.HexToAddress("0x7bC08E1c04fb41d75F1410363F0c5746Eae80582"),
				ERC1155GatewayAddr:     common.HexToAddress("0x62597Cc19703aF10B58feF87B0d5D29eFE263bcc"),
				USDCGatewayAddr:        common.HexToAddress("0x987e300fDfb06093859358522a79098848C33852"),
				LIDOGatewayAddr:        common.HexToAddress("0xE9c5C9f67ec7B773fC76440845751F657bb953FF"),
				DAIGatewayAddr:         common.HexToAddress("0xC5034eB8F682b73F93C9246aa95A8eBbF82793aA"),
				L2GasOracle:            common.HexToAddress("0x987e300fDfb06093859358522a79098848C33852"),
			},
		},
	},
}
