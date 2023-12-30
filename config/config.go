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
	Migrations        string        `toml:"migrations"`
	Chain             ChainConfig   `toml:"chain"`
	RPCs              RPCsConfig    `toml:"rpcs"`
	DA                DAConfig      `toml:"da"`
	MasterDB          DBConfig      `toml:"master_db"`
	SlaveDB           DBConfig      `toml:"slave_db"`
	SlaveDbEnable     bool          `toml:"slave_db_enable"`
	ApiCacheEnable    bool          `toml:"api_cache_enable"`
	HTTPServer        ServerConfig  `toml:"http_server"`
	MetricsServer     ServerConfig  `toml:"metrics_server"`
	StartDataStoreId  uint32        `toml:"start_data_store_id"`
	FraudProofWindows time.Duration `toml:"fraud_proof_windows"`
}

type ChainConfig struct {
	ChainId                 uint     `toml:"chain_id"`
	L1StartingHeight        uint     `toml:"l1_starting_height"`
	L2StartingHeight        uint     `toml:"l2_starting_height"`
	L1BedrockStartingHeight uint     `toml:"l1_bedrock_starting_height"`
	L2BedrockStartingHeight uint     `toml:"l2_bedrock_starting_height"`
	L1Contracts             []string `toml:"l1_contracts"`
	L2Contracts             []string `toml:"l2_contracts"`
	L1ConfirmationDepth     uint     `toml:"l1_confirmation_depth"`
	L2ConfirmationDepth     uint     `toml:"l2_confirmation_depth"`
	L1PollingInterval       uint     `toml:"l1_polling_interval"`
	L2PollingInterval       uint     `toml:"l2_polling_interval"`
	L1HeaderBufferSize      uint     `toml:"l1_header_buffer_size"`
	L2HeaderBufferSize      uint     `toml:"l2_header_buffer_size"`
}

type RPCsConfig struct {
	L1RPC string `toml:"l1_rpc"`
	L2RPC string `toml:"l2_rpc"`
}

type DBConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Name     string `toml:"name"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

type ServerConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type DAConfig struct {
	RetrieverSocket          string        `toml:"retriever_socket"`
	RetrieverTimeout         time.Duration `toml:"retriever_timeout"`
	GraphProvider            string        `toml:"graph_provider"`
	DataStorePollingDuration time.Duration `toml:"data_store_polling_duration"`
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
			ChainId:                 ctx.Uint(flag.ChainIdFlag.Name),
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
