package config

import (
	"time"

	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/flag"
)

const (
	defaultLoopInterval     = 5000
	defaultHeaderBufferSize = 500
)

type Config struct {
	Migrations        string
	Chain             ChainConfig
	RPCs              RPCsConfig
	DA                DAConfig
	MasterDB          DBConfig
	SlaveDB           DBConfig
	SlaveDbEnable     bool
	ApiCacheEnable    bool
	HTTPServer        ServerConfig
	MetricsServer     ServerConfig
	StartDataStoreId  uint32
	FraudProofWindows time.Duration
}

type ChainConfig struct {
	L1StartingHeight        uint
	L2StartingHeight        uint
	L1BedrockStartingHeight uint
	L2BedrockStartingHeight uint
	L1Contracts             []string
	L2Contracts             []string
	L1ConfirmationDepth     uint
	L2ConfirmationDepth     uint
	L1PollingInterval       uint
	L2PollingInterval       uint
	L1HeaderBufferSize      uint
	L2HeaderBufferSize      uint
}

type RPCsConfig struct {
	L1RPC string
	L2RPC string
}

type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type ServerConfig struct {
	Host string
	Port int
}

type DAConfig struct {
	RetrieverSocket          string
	RetrieverTimeout         time.Duration
	GraphProvider            string
	DataStorePollingDuration time.Duration
}

func LoadConfig(log log.Logger, cliCtx *cli.Context) (Config, error) {
	var cfg Config
	cfg = NewConfig(cliCtx)
	if cfg.Chain.L1PollingInterval == 0 {
		cfg.Chain.L1PollingInterval = defaultLoopInterval
	}
	if cfg.Chain.L2PollingInterval == 0 {
		cfg.Chain.L2PollingInterval = defaultLoopInterval
	}
	if cfg.Chain.L1HeaderBufferSize == 0 {
		cfg.Chain.L1HeaderBufferSize = defaultHeaderBufferSize
	}
	if cfg.Chain.L2HeaderBufferSize == 0 {
		cfg.Chain.L2HeaderBufferSize = defaultHeaderBufferSize
	}
	return cfg, nil
}

func NewConfig(ctx *cli.Context) Config {
	return Config{
		Migrations: ctx.String(flag.MigrationsFlag.Name),
		Chain: ChainConfig{
			L1StartingHeight:        ctx.Uint(flag.L1StartingHeightFlag.Name),
			L2StartingHeight:        ctx.Uint(flag.L2StartingHeightFlag.Name),
			L1BedrockStartingHeight: ctx.Uint(flag.L1BedrockStartingHeightFlag.Name),
			L2BedrockStartingHeight: ctx.Uint(flag.L2BedrockStartingHeightFlag.Name),
			L1Contracts:             ctx.StringSlice(flag.L1ContractsFlag.String()),
			L2Contracts:             ctx.StringSlice(flag.L1ContractsFlag.String()),
			L1ConfirmationDepth:     ctx.Uint(flag.L1ConfirmationDepthFlag.Name),
			L2ConfirmationDepth:     ctx.Uint(flag.L2ConfirmationDepthFlag.Name),
			L1PollingInterval:       ctx.Uint(flag.L1PollingIntervalFlag.Name),
			L2PollingInterval:       ctx.Uint(flag.L2PollingIntervalFlag.Name),
			L1HeaderBufferSize:      ctx.Uint(flag.L1HeaderBufferSizeFlag.Name),
			L2HeaderBufferSize:      ctx.Uint(flag.L2HeaderBufferSizeFlag.Name),
		},

		RPCs: RPCsConfig{
			L1RPC: ctx.String(flag.L1EthRpcFlag.Name),
			L2RPC: ctx.String(flag.L2EthRpcFlag.Name),
		},
		DA: DAConfig{
			RetrieverSocket:          ctx.String(flag.RetrieverSocketFlag.Name),
			RetrieverTimeout:         ctx.Duration(flag.RetrieverTimeoutFlag.Name),
			GraphProvider:            ctx.String(flag.GraphProviderFlag.Name),
			DataStorePollingDuration: ctx.Duration(flag.DataStorePollingDurationFlag.Name),
		},

		MasterDB: DBConfig{
			Host:     ctx.String(flag.MasterDbHostFlag.Name),
			Port:     ctx.Int(flag.MasterDbPortFlag.Name),
			Name:     ctx.String(flag.MasterDbNameFlag.Name),
			User:     ctx.String(flag.MasterDbUserFlag.Name),
			Password: ctx.String(flag.MasterDbPasswordFlag.Name),
		},
		SlaveDB: DBConfig{
			Host:     ctx.String(flag.SlaveDbHostFlag.Name),
			Port:     ctx.Int(flag.SlaveDbPortFlag.Name),
			Name:     ctx.String(flag.SlaveDbNameFlag.Name),
			User:     ctx.String(flag.SlaveDbUserFlag.Name),
			Password: ctx.String(flag.SlaveDbPasswordFlag.Name),
		},
		SlaveDbEnable:  ctx.Bool(flag.SlaveDbEnableFlag.Name),
		ApiCacheEnable: ctx.Bool(flag.EnableApiCacheFlag.Name),
		HTTPServer: ServerConfig{
			Host: ctx.String(flag.HttpHostFlag.Name),
			Port: ctx.Int(flag.HttpPortFlag.Name),
		},
		MetricsServer: ServerConfig{
			Host: ctx.String(flag.MetricsHostFlag.Name),
			Port: ctx.Int(flag.MetricsPortFlag.Name),
		},
		StartDataStoreId:  uint32(ctx.Uint64(flag.StartDataStoreIdFlag.Name)),
		FraudProofWindows: ctx.Duration(flag.FraudProofWindowsFlags.Name),
	}
}
