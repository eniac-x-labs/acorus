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
				MessengerAddr:          common.HexToAddress("0xda7c91Ed60DACD28Cb97B180108958c9ACC7698a"),
				ETHGatewayAddr:         common.HexToAddress("0x567671187b5FFbcDFe0B6EcF3e56C05508a31A87"),
				WETHGatewayAddr:        common.HexToAddress("0x3b03aE2F27d62E0B2b6740CA20Fc07Af4338B791"),
				StandardERC20Gateway:   common.HexToAddress("0xb00cb1F6f7C43D2EE8C4e2163a6bEA22441A5B7c"),
				CustomERC20GatewayAddr: common.HexToAddress("0x63CCb38E9d21A72777b203267F2e4ba5C974fC62"),
				ERC721GatewayAddr:      common.HexToAddress("0xE2c36a2D8B5528719aE7A42A778b2D08b18d134a"),
				ERC1155GatewayAddr:     common.HexToAddress("0xfF14870512e42BFb85a9B7bEfDc06e9aB5A37269"),
				USDCGatewayAddr:        common.HexToAddress("0x97D5799CDC8eE2A7452913d7548c7cEE285719FA"),
				LIDOGatewayAddr:        common.HexToAddress("0xE9c5C9f67ec7B773fC76440845751F657bb953FF"),
				DAIGatewayAddr:         common.HexToAddress("0xC5034eB8F682b73F93C9246aa95A8eBbF82793aA"),
			},
		},
	},
}
