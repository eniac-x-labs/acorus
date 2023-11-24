package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	common2 "github.com/cornerstone-labs/acorus/common"
	op_stack "github.com/cornerstone-labs/acorus/config/op-stack"
	"github.com/ethereum/go-ethereum/log"
)

const (
	defaultLoopInterval     = 5000
	defaultHeaderBufferSize = 500
)

type ChainConfig struct {
	L1RPC string `toml:"l1-rpc"`
	L2RPC string `toml:"l2-rpc"`

	L1StartHeight uint `toml:"l1-start-height"`
	L2StartHeight uint `toml:"l2-start-height"`

	L1ConfirmationDepth uint `toml:"l1-confirmation-depth"`
	L2ConfirmationDepth uint `toml:"l2-confirmation-depth"`

	L1PollingInterval uint `toml:"l1-polling-interval"`
	L2PollingInterval uint `toml:"l2-polling-interval"`

	L1HeaderBufferSize uint `toml:"l1-header-buffer-size"`
	L2HeaderBufferSize uint `toml:"l2-header-buffer-size"`
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
	Chain         ChainConfig          `toml:"chain"`
	OpContracts   op_stack.OpContracts `toml:"op_contracts"`
	Redis         RedisConfig          `toml:"redis"`
	MasterDB      DBConfig             `toml:"master_db"`
	SlaveDB       DBConfig             `toml:"slave_db"`
	SlaveDbEnable bool                 `toml:"slave_db_enable"`
	HTTPServer    ServerConfig         `toml:"http"`
	MetricsServer ServerConfig         `toml:"metrics"`
}

func LoadConfig(log log.Logger, path string, chainBridge string) (Config, error) {
	log.Debug("loading config", "path", path)
	var cfg Config
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	data = []byte(os.ExpandEnv(string(data)))
	log.Debug("parsed config file", "data", string(data))

	if _, err := toml.Decode(string(data), &cfg); err != nil {
		log.Error("failed to decode config file", "err", err)
		return cfg, err
	}

	if chainBridge == common2.Op {
		if cfg.OpContracts.Preset == op_stack.DevnetPresetId {
			preset, err := op_stack.DevnetPreset()
			if err != nil {
				return cfg, err
			}
			log.Info("detected preset", "preset", op_stack.DevnetPresetId, "name", preset.Name)
			cfg.OpContracts = preset.OpContracts
		} else if cfg.OpContracts.Preset != 0 {
			preset, ok := op_stack.Presets[cfg.OpContracts.Preset]
			if !ok {
				return cfg, fmt.Errorf("unknown preset: %d", cfg.OpContracts.Preset)
			}
			log.Info("detected preset", "preset", cfg.OpContracts.Preset, "name", preset.Name)
			cfg.OpContracts = preset.OpContracts
		}
		cfg.OpContracts.L2Contracts = op_stack.L2ContractsFromPredeploys()

		if cfg.OpContracts.Preset > 0 {
			if _, err := toml.Decode(string(data), &cfg); err != nil {
				log.Error("failed to decode config file", "err", err)
				return cfg, err
			}
		}
		log.Info("loaded chain config", "config", cfg.OpContracts)
	} else if chainBridge == common2.Polygon {
		// todo: handle polygon config here
	} else if chainBridge == common2.Zksync {
		// todo: handle zksync config here
	} else if chainBridge == common2.Mantle {
		// todo: handle mantle config here
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
