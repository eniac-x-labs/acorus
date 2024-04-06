package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Server struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	GrpcPort int    `yaml:"grpc_port"`
}

type Database struct {
	DbHost     string `yaml:"db_host"`
	DbPort     int    `yaml:"db_port"`
	DbName     string `yaml:"db_name"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
}

type RPC struct {
	RpcUrl             string       `yaml:"rpc_url"`
	IsMainnet          bool         `yaml:"is_mainnet"`
	ChainId            uint64       `yaml:"chain_id"`
	L1ChainId          uint64       `yaml:"l1_chain_id"`
	StartBlock         uint64       `yaml:"start_block"`
	HeaderBufferSize   uint64       `yaml:"header_buffer_size"`
	L1EventUnpackBlock uint64       `yaml:"l1_event_unpack_block"`
	L1PoolContract     string       `yaml:"l1_pool_contract"`
	L2PoolContract     string       `yaml:"l2_pool_contract"`
	AppChainAddr       AppChainAddr `yaml:"app_chain_addr"`
	Contracts          []string     `yaml:"contracts"`
	EventContracts     []string     `yaml:"event_contracts"`
}

type AppChainAddr struct {
	L1 AppChainL1 `yaml:"l1"`
	L2 AppChainL2 `yaml:"l2"`
}

type AppChainL1 struct {
}

type AppChainL2 struct {
	L1RewardManager string `yaml:"l1_reward_manager"`
	L2RewardManager string `yaml:"l2_reward_manager"`
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
	AirDropGRpcUrl string     `yaml:"airdrop_grpc_url"`
	MasterDb       Database   `yaml:"master_db"`
	SlaveDb        Database   `yaml:"slave_db"`
	SlaveDbEnable  bool       `yaml:"slave_db_enable"`
	EnableApiCache bool       `yaml:"enable_api_cache"`
	Relayers       []*Relayer `yaml:"relayers"`
	AppChain       AppChain   `yaml:"appchain"`
}

type AppChain struct {
	L1 L1AppChain `yaml:"l1"`
	L2 L2AppChain `yaml:"l2"`
}

type L1AppChain struct {
	ChainId     string   `yaml:"chain_id"`
	Contracts   []string `yaml:"contracts"`
	StartHeight uint64   `yaml:"start_height"`
}

type L2AppChain struct {
	ChainId   string   `yaml:"chain_id"`
	Contracts []string `yaml:"contracts"`
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
