package flag

import (
	"time"

	"github.com/urfave/cli/v2"
)

const envVarPrefix = "ACORUS"

func prefixEnvVars(name string) []string {
	return []string{envVarPrefix + "_" + name}
}

var (
	MigrationsFlag = &cli.StringFlag{
		Name:    "migrations-dir",
		Value:   "./migrations",
		Usage:   "path to migrations folder",
		EnvVars: prefixEnvVars("MIGRATIONS_DIR"),
	}
	L1EthRpcFlag = &cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "HTTP provider URL for L1",
		EnvVars:  prefixEnvVars("L1_RPC"),
		Required: true,
	}
	L2EthRpcFlag = &cli.StringFlag{
		Name:     "l2-eth-rpc",
		Usage:    "HTTP provider URL for L2",
		EnvVars:  prefixEnvVars("L2_PRC"),
		Required: true,
	}
	HttpHostFlag = &cli.StringFlag{
		Name:     "http-host",
		Usage:    "The host of the api",
		EnvVars:  prefixEnvVars("HTTP_HOST"),
		Required: true,
	}
	HttpPortFlag = &cli.IntFlag{
		Name:     "http-port",
		Usage:    "The port of the api",
		EnvVars:  prefixEnvVars("HTTP_PORT"),
		Value:    8987,
		Required: true,
	}
	MetricsHostFlag = &cli.StringFlag{
		Name:     "metrics-host",
		Usage:    "The host of the metrics",
		EnvVars:  prefixEnvVars("METRICS_HOST"),
		Required: true,
	}
	MetricsPortFlag = &cli.IntFlag{
		Name:     "metrics-port",
		Usage:    "The port of the metrics",
		EnvVars:  prefixEnvVars("METRICS_PORT"),
		Value:    7214,
		Required: true,
	}
	SlaveDbEnableFlag = &cli.BoolFlag{
		Name:     "slave-db-enable",
		Usage:    "Whether to use slave db",
		EnvVars:  prefixEnvVars("SLAVE_DB_ENABLE"),
		Required: true,
	}
	MasterDbHostFlag = &cli.StringFlag{
		Name:     "master-db-host",
		Usage:    "The host of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_HOST"),
		Required: true,
	}
	MasterDbPortFlag = &cli.IntFlag{
		Name:     "master-db-port",
		Usage:    "The port of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_PORT"),
		Required: true,
	}
	MasterDbUserFlag = &cli.StringFlag{
		Name:     "master-db-user",
		Usage:    "The user of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_USER"),
		Required: true,
	}
	MasterDbPasswordFlag = &cli.StringFlag{
		Name:     "master-db-password",
		Usage:    "The host of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_PASSWORD"),
		Required: true,
	}
	MasterDbNameFlag = &cli.StringFlag{
		Name:     "master-db-name",
		Usage:    "The db name of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_NAME"),
		Required: true,
	}
	StartDataStoreIdFlag = &cli.Uint64Flag{
		Name:    "start-data-store-id",
		Usage:   "Start data store id",
		EnvVars: prefixEnvVars("START_DATA_STORE_ID"),
		Value:   1,
	}
	// Optional flags
	L1PollingIntervalFlag = &cli.DurationFlag{
		Name:    "l1-polling-interval",
		Usage:   "The interval of l1 synchronization",
		EnvVars: prefixEnvVars("L1_POLLING_INTERVAL"),
		Value:   0,
	}
	L1HeaderBufferSizeFlag = &cli.IntFlag{
		Name:    "l1-header-buffer-size",
		Usage:   "The buffer size of l1 header",
		EnvVars: prefixEnvVars("L1_HEADER_BUFFER_SIZE"),
		Value:   0,
	}
	L1ConfirmationDepthFlag = &cli.IntFlag{
		Name:    "l1-confirmation-depth",
		Usage:   "The confirmation depth of l1",
		EnvVars: prefixEnvVars("L1_CONFIRMATION_DEPTH"),
		Value:   0,
	}
	L1StartingHeightFlag = &cli.IntFlag{
		Name:    "l1-starting-height",
		Usage:   "The starting height of l1",
		EnvVars: prefixEnvVars("L1_STARTING_HEIGHT"),
		Value:   0,
	}
	L2StartingHeightFlag = &cli.IntFlag{
		Name:    "l2-starting-height",
		Usage:   "The starting height of l2",
		EnvVars: prefixEnvVars("L2_STARTING_HEIGHT"),
		Value:   0,
	}

	L1ContractsFlag = &cli.MultiStringFlag{
		Target: &cli.StringSliceFlag{
			Name:     "l1-contracts",
			Usage:    "The l1 contracts",
			EnvVars:  prefixEnvVars("L1_CONTRACTS"),
			Required: true,
		},
	}

	L2ContractsFlag = &cli.MultiStringFlag{
		Target: &cli.StringSliceFlag{
			Name:     "l2-contracts",
			Usage:    "The l2 contracts",
			EnvVars:  prefixEnvVars("L2_CONTRACTS"),
			Required: true,
		},
	}

	L1BedrockStartingHeightFlag = &cli.IntFlag{
		Name:    "l1-bedrock-starting-height",
		Usage:   "The starting height of l1 upgrade to bedrock",
		EnvVars: prefixEnvVars("L1_BEDROCK_STARTING_HEIGHT"),
		Value:   0,
	}
	L2BedrockStartingHeightFlag = &cli.IntFlag{
		Name:    "l2-bedrock-starting-height",
		Usage:   "The starting height of l2 upgrade to bedrock",
		EnvVars: prefixEnvVars("L2_BEDROCK_STARTING_HEIGHT"),
		Value:   0,
	}
	L2PollingIntervalFlag = &cli.DurationFlag{
		Name:    "l2-polling-interval",
		Usage:   "The interval of l2 synchronization",
		EnvVars: prefixEnvVars("L2_POLLING_INTERVAL"),
		Value:   0,
	}
	L2HeaderBufferSizeFlag = &cli.IntFlag{
		Name:    "l2-header-buffer-size",
		Usage:   "The buffer size of l2 header",
		EnvVars: prefixEnvVars("L2_HEADER_BUFFER_SIZE"),
		Value:   500,
	}
	L2ConfirmationDepthFlag = &cli.IntFlag{
		Name:    "l2-confirmation-depth",
		Usage:   "The confirmation depth of l2",
		EnvVars: prefixEnvVars("L2_CONFIRMATION_DEPTH"),
		Value:   0,
	}
	RetrieverSocketFlag = &cli.StringFlag{
		Name:    "retriever-socket",
		Usage:   "Websocket for MantleDA disperser",
		EnvVars: prefixEnvVars("RETRIEVER_SOCKET"),
	}
	RetrieverTimeoutFlag = &cli.DurationFlag{
		Name:    "retriever-timeout",
		Usage:   "Retriever timeout",
		EnvVars: prefixEnvVars("RETRIEVER_TIMEOUT"),
	}
	FraudProofWindowsFlags = &cli.DurationFlag{
		Name:    "fraud-proof-windows",
		Usage:   "The fraud proof windows",
		Value:   time.Second * 604800,
		EnvVars: prefixEnvVars("FRAUD_PROOF_WINDOWS"),
	}
	DataStorePollingDurationFlag = &cli.DurationFlag{
		Name:    "data-store-polling-duration",
		Usage:   "Duration to store blob",
		EnvVars: prefixEnvVars("DATA_STORE_POLLING_DURATION"),
	}
	GraphProviderFlag = &cli.StringFlag{
		Name:    "graph-provider",
		Usage:   "Graph node url of MantleDA graph node",
		EnvVars: prefixEnvVars("GRAPH_PROVIDER"),
	}
	SlaveDbHostFlag = &cli.StringFlag{
		Name:    "slave-db-host",
		Usage:   "The host of the slave database",
		EnvVars: prefixEnvVars("SLAVE_DB_HOST"),
	}
	SlaveDbPortFlag = &cli.IntFlag{
		Name:    "slave-db-port",
		Usage:   "The port of the slave database",
		EnvVars: prefixEnvVars("SLAVE_DB_PORT"),
	}
	SlaveDbUserFlag = &cli.StringFlag{
		Name:    "slave-db-user",
		Usage:   "The user of the slave database",
		EnvVars: prefixEnvVars("SLAVE_DB_USER"),
	}
	SlaveDbPasswordFlag = &cli.StringFlag{
		Name:    "slave-db-password",
		Usage:   "The host of the slave database",
		EnvVars: prefixEnvVars("SLAVE_DB_PASSWORD"),
	}
	SlaveDbNameFlag = &cli.StringFlag{
		Name:    "slave-db-name",
		Usage:   "The db name of the slave database",
		EnvVars: prefixEnvVars("SLAVE_DB_NAME"),
	}
	EnableApiCacheFlag = &cli.BoolFlag{
		Name:    "apo-cache-enable",
		Usage:   "Enable api cache.",
		Value:   false,
		EnvVars: prefixEnvVars("API_CACHE_ENABLE"),
	}
)

var requiredFlags = []cli.Flag{
	MigrationsFlag,
	L1EthRpcFlag,
	L2EthRpcFlag,
	HttpPortFlag,
	HttpHostFlag,
	MetricsPortFlag,
	MetricsHostFlag,
	SlaveDbEnableFlag,
	MasterDbHostFlag,
	MasterDbPortFlag,
	MasterDbUserFlag,
	MasterDbPasswordFlag,
	MasterDbNameFlag,
	StartDataStoreIdFlag,
	L1ContractsFlag,
	L2ContractsFlag,
}

var optionalFlags = []cli.Flag{
	L1PollingIntervalFlag,
	L1HeaderBufferSizeFlag,
	L1ConfirmationDepthFlag,
	L1StartingHeightFlag,
	L2StartingHeightFlag,
	L1BedrockStartingHeightFlag,
	L2BedrockStartingHeightFlag,
	FraudProofWindowsFlags,
	L2PollingIntervalFlag,
	L2ConfirmationDepthFlag,
	L2HeaderBufferSizeFlag,
	RetrieverTimeoutFlag,
	RetrieverSocketFlag,
	GraphProviderFlag,
	DataStorePollingDurationFlag,
	SlaveDbHostFlag,
	SlaveDbPortFlag,
	SlaveDbUserFlag,
	SlaveDbPasswordFlag,
	SlaveDbNameFlag,
	EnableApiCacheFlag,
}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

var Flags []cli.Flag
