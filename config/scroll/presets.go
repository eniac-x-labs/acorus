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
				MessengerAddr:          common.HexToAddress("0x7318152B19c3c97c886D5ee6C2525E62ce8e2abA"),
				ETHGatewayAddr:         common.HexToAddress("0xd165b42d857eae2915625819464a2a1f91E5d0A5"),
				WETHGatewayAddr:        common.HexToAddress("0xb0255e4C1a919619D1CafBA51021d638c4F71b89"),
				StandardERC20Gateway:   common.HexToAddress("0x00fEc01A9b975bA37466B4E9006dF2C71BFE0e48"),
				CustomERC20GatewayAddr: common.HexToAddress("0xD8874B0E6C3CC43C00B69D60c21Ef452d1159bDe"),
				ERC721GatewayAddr:      common.HexToAddress("0x131B46649F6882d686a766cb8b68c4cB0ACdeb24"),
				ERC1155GatewayAddr:     common.HexToAddress("0xCeE721789FAA05c7F4463efB664520656aB7C7d5"),
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
