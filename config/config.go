package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

const (
	defaultLoopInterval     = 5000
	defaultHeaderBufferSize = 500
)

type ChainConfig struct {
	ChainId uint64 `toml:"chainId"`
	L1RPC   string `toml:"l1-rpc"`
	L2RPC   string `toml:"l2-rpc"`

	L1StartHeight uint `toml:"l1-start-height"`
	L2StartHeight uint `toml:"l2-start-height"`

	L1ConfirmationDepth uint `toml:"l1-confirmation-depth"`
	L2ConfirmationDepth uint `toml:"l2-confirmation-depth"`

	L1PollingInterval uint `toml:"l1-polling-interval"`
	L2PollingInterval uint `toml:"l2-polling-interval"`

	L1HeaderBufferSize uint `toml:"l1-header-buffer-size"`
	L2HeaderBufferSize uint `toml:"l2-header-buffer-size"`

	L1Contracts []common.Address `toml:"l1-contracts"`
	L2Contracts []common.Address `toml:"l2-contracts"`
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

type RedisConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

type Config struct {
	Chain         ChainConfig  `toml:"chain"`
	Redis         RedisConfig  `toml:"redis"`
	MasterDB      DBConfig     `toml:"master_db"`
	SlaveDB       DBConfig     `toml:"slave_db"`
	SlaveDbEnable bool         `toml:"slave_db_enable"`
	HTTPServer    ServerConfig `toml:"http"`
	MetricsServer ServerConfig `toml:"metrics"`
}

func LoadConfig(log log.Logger, path string, chainBridge string) (Config, error) {
	log.Info("loading config", "path", path)
	var cfg Config
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	data = []byte(os.ExpandEnv(string(data)))
	log.Info("parsed config file", "data", string(data))

	if _, err := toml.Decode(string(data), &cfg); err != nil {
		log.Error("failed to decode config file", "err", err)
		return cfg, err
	}

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
