package global_const

const (
	ChainIdRedisKey     = "chainId:%s"
	ChainId             = "chainId"
	ChainName           = "chainName"
	Polygon             = "Polygon"
	GormInfoFmt         = "%s\n[%.3fms] [rows:%v] %s"
	ZeroAddress         = "0x0000000000000000000000000000000000000000"
	WEthAddress         = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	EthAddress          = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
	SepoliaWETH         = ""
	LogTimeFormat       = "2006-01-02"
	LayerTypeOne        = 1
	LayerTypeTwo        = 2
	StakingTypeStake    = 1
	StakingTypeReward   = 2
	StakingTypeWithdraw = 3

	BridgeOperaInitType     = 1
	BridgeOperaFinalizeType = 2

	ScrollChainId          uint64 = 534352
	PolygonChainId         uint64 = 1101
	PolygonSepoliaChainId  uint64 = 1442
	EthereumChainId        uint64 = 1
	EthereumSepoliaChainId uint64 = 11155111
	BaseChainId            uint64 = 8453
	BaseSepoliaChainId     uint64 = 84532
	MantaChainId           uint64 = 169
	MantaSepoliaChainId    uint64 = 3441006
	MantleSepoliaChainId   uint64 = 5003
	MantleChainId          uint64 = 5000
	ZkFairSepoliaChainId   uint64 = 43851
	ZkFairChainId          uint64 = 42766
	OkxSepoliaChainId      uint64 = 195
	OkxChainId             uint64 = 66
	OpChinId               uint64 = 10
	OpTestChinId           uint64 = 11155420
	LineaChainId           uint64 = 59144
	BlocksLimit                   = 10000
)
