package common

// AssetType can be ETH/ERC20/ERC1155/ERC721
type AssetType int

// MsgType can be layer1/layer2 msg
type MsgType int

func (a AssetType) String() string {
	switch a {
	case ETH:
		return "ETH"
	case ERC20:
		return "ERC20"
	case ERC1155:
		return "ERC1155"
	case ERC721:
		return "ERC721"
	}
	return "Unknown Asset Type"
}

const (
	// ETH = 0
	ETH  = 0
	WETH = 1
	// ERC20 = 1
	ERC20 = 20
	// ERC721 = 2
	ERC721 = 721
	// ERC1155 = 3
	ERC1155 = 1155
)

const (
	// UnknownMsg = 0
	UnknownMsg MsgType = iota
	// Layer1Msg = 1
	Layer1Msg
	// Layer2Msg = 2
	Layer2Msg
)
