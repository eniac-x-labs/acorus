package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Database struct {
	DbHost     string `yaml:"db_host"`
	DbPort     int    `yaml:"db_port"`
	DbName     string `yaml:"db_name"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
}

type RPC struct {
	RpcUrl             string   `yaml:"rpc_url"`
	ChainId            uint64   `yaml:"chain_id"`
	StartBlock         uint64   `yaml:"start_block"`
	L1EventUnpackBlock uint64   `yaml:"l1_event_unpack_block"`
	Contracts          []string `yaml:"contracts"`
	EventContracts     []string `yaml:"event_contracts"`
}

type Relayer struct {
	ChainId         string   `yaml:"chain_id"`
	EventStartBlock uint64   `yaml:"eventStartBlock"`
	Contracts       []string `yaml:"contracts"`
}

type Config struct {
	Server         Server     `yaml:"server"`
	RPCs           []*RPC     `yaml:"rpcs"`
	Metrics        Server     `yaml:"metrics"`
	BridgeGrpcUrl  string     `yaml:"bridge_grpc_url"`
	MasterDb       Database   `yaml:"master_db"`
	SlaveDb        Database   `yaml:"slave_db"`
	SlaveDbEnable  bool       `yaml:"slave_db_enable"`
	EnableApiCache bool       `yaml:"enable_api_cache"`
	Relayers       []*Relayer `yaml:"relayers"`
}

func New(path string) (*Config, error) {
	var config = new(Config)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
