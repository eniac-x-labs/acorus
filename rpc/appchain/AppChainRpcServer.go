package appchain

import (
	"context"
	"fmt"

	"math/big"
	"net"
	"strings"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	l1_reward_manager_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/l1_reward_manager/bindings"
	l2_reward_manager_bindings "github.com/cornerstone-labs/acorus/appchain/bindings/l2_reward_manager/bindings"
	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/rpc/appchain/protobuf/pb"
)

const MaxRecvMessageSize = 1024 * 1024 * 300

type RpcServerConfig struct {
	GrpcHostname string
	GrpcPort     string
}

type RpcServer struct {
	*RpcServerConfig
	l1RewardManagers map[uint64]*l1_reward_manager_bindings.L1RewardManager
	l2RewardManagers map[uint64]*l2_reward_manager_bindings.L2RewardManager
	tasks            tasks.Group
	stopped          atomic.Bool
	db               *database.DB
}

func NewRpcServer(grpcCfg *RpcServerConfig, chainRpcCfg []*config.RPC,
	db *database.DB,
	shutdown context.CancelCauseFunc) (*RpcServer, error) {
	l1RewardManagers := make(map[uint64]*l1_reward_manager_bindings.L1RewardManager)
	l2RewardManagers := make(map[uint64]*l2_reward_manager_bindings.L2RewardManager)
	for i := 0; i < len(chainRpcCfg); i++ {
		if chainRpcCfg[i].ChainId == global_const.EthereumSepoliaChainId ||
			chainRpcCfg[i].ChainId == global_const.EthereumChainId {
			continue
		}
		client, _ := ethclient.Dial(chainRpcCfg[i].RpcUrl)

		l1RewardManagerAddr := chainRpcCfg[i].AppChainAddr.L2.L1RewardManager
		l2RewardManagerAddr := chainRpcCfg[i].AppChainAddr.L2.L2RewardManager
		if l1RewardManagerAddr == "" || l2RewardManagerAddr == "" {
			continue
		}
		L1RewardManager, _ := l1_reward_manager_bindings.NewL1RewardManager(common.HexToAddress(l1RewardManagerAddr), client)
		L2RewardManager, _ := l2_reward_manager_bindings.NewL2RewardManager(common.HexToAddress(l2RewardManagerAddr), client)
		l1RewardManagers[chainRpcCfg[i].ChainId] = L1RewardManager
		l2RewardManagers[chainRpcCfg[i].ChainId] = L2RewardManager
	}

	return &RpcServer{
		RpcServerConfig:  grpcCfg,
		l1RewardManagers: l1RewardManagers,
		l2RewardManagers: l2RewardManagers,
		db:               db,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in acorus processor: %w", err))
		}},
	}, nil
}

func (s *RpcServer) Start(ctx context.Context) error {
	go func(s *RpcServer) error {
		addr := fmt.Sprintf("%s:%s", s.GrpcHostname, s.GrpcPort)
		log.Info("start rpc server", "addr", addr)

		listener, err := net.Listen("tcp", addr)
		if err != nil {
			log.Error("Could not start tcp listener", "err", err)
			return err
		}

		opt := grpc.MaxRecvMsgSize(MaxRecvMessageSize)

		gs := grpc.NewServer(
			opt,
			grpc.ChainUnaryInterceptor(),
		)
		reflection.Register(gs)
		pb.RegisterAppChainServiceServer(gs, s)

		log.Info("grpc info", "port", s.GrpcPort, "address", listener.Addr().String())
		if err := gs.Serve(listener); err != nil {
			log.Error("Could not GRPC server", "err", err)
			return err
		}
		return nil
	}(s)
	return nil
}

func (s *RpcServer) Stop(ctx context.Context) error {
	s.stopped.Store(true)
	return nil
}

func (s *RpcServer) Stopped() bool {
	return s.stopped.Load()
}

func (s *RpcServer) L1StakerRewardsAmount(ctx context.Context, request *pb.L1StakerRewardsAmountRequest) (*pb.L1StakerRewardsAmountResponse, error) {
	chainId := request.ChainId
	chainIdBig, _ := big.NewInt(0).SetString(chainId, 10)
	chainIdUint := chainIdBig.Uint64()
	// get call instance
	l1RewardManager := s.l1RewardManagers[chainIdUint]
	if l1RewardManager == nil {
		return &pb.L1StakerRewardsAmountResponse{
			Success: false,
			Message: "this chainId is unsupported",
		}, nil
	}
	// create req params
	stakerAddress := request.StakerAddress
	strategies := request.Strategies
	strategiesSplit := strings.Split(strategies, ",")
	strategiesParam := make([]common.Address, 0)
	for _, strategy := range strategiesSplit {
		strategiesParam = append(strategiesParam, common.HexToAddress(strategy))
	}
	l1RewardManagerCaller := l1RewardManager.L1RewardManagerCaller
	rawCaller := &l1_reward_manager_bindings.L1RewardManagerCallerRaw{Contract: &l1RewardManagerCaller}
	var incomeResult []interface{}
	err := rawCaller.Call(&bind.CallOpts{Context: ctx,
		From: common.HexToAddress(stakerAddress),
	}, &incomeResult, "stakerRewardsAmount", strategiesParam)
	if err != nil {
		log.Error("appchain rpc server ", "method", "L1StakerRewardsAmount", "err", err)
		return &pb.L1StakerRewardsAmountResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	log.Info("appchain rpc server ", "method", "L1StakerRewardsAmount", "incomeResult", incomeResult)
	return &pb.L1StakerRewardsAmountResponse{
		Success: true,
		Message: "success",
		Income:  incomeResult[0].(*big.Int).String(),
	}, nil
}

func (s *RpcServer) L2StakerRewardsAmount(ctx context.Context, request *pb.L2StakerRewardsAmountRequest) (*pb.L2StakerRewardsAmountResponse, error) {
	chainId := request.ChainId
	chainIdBig, _ := big.NewInt(0).SetString(chainId, 10)
	chainIdUint := chainIdBig.Uint64()
	// get call instance
	l2RewardManager := s.l2RewardManagers[chainIdUint]
	if l2RewardManager == nil {
		return &pb.L2StakerRewardsAmountResponse{
			Success: false,
			Message: "this chainId is unsupported",
		}, nil
	}
	// create req params
	stakerAddress := request.StakerAddress
	strategy := request.Strategy
	paramAddress := make([]common.Address, 0)
	paramAddress = append(paramAddress, common.HexToAddress(stakerAddress))
	l2RewardManagerCaller := l2RewardManager.L2RewardManagerCaller
	rawCaller := &l2_reward_manager_bindings.L2RewardManagerCallerRaw{Contract: &l2RewardManagerCaller}
	var incomeResult []interface{}
	err := rawCaller.Call(&bind.CallOpts{Context: ctx,
		From: common.HexToAddress(stakerAddress),
	}, &incomeResult, "stakerRewardsAmount",
		common.HexToAddress(strategy),
	)
	if err != nil {
		log.Error("appchain rpc server ", "method", "L2StakerRewardsAmount", "err", err)
		return &pb.L2StakerRewardsAmountResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	log.Info("appchain rpc server ", "method", "L2StakerRewardsAmount", "incomeResult", incomeResult)
	return &pb.L2StakerRewardsAmountResponse{
		Success: true,
		Message: "success",
		Income:  incomeResult[0].(*big.Int).String(),
	}, nil
}

func (s *RpcServer) L2UnStakeRecord(ctx context.Context, request *pb.L2UnStakeRecordRequest) (*pb.L2UnStakeRecordResponse, error) {
	pageSize := request.PageSize
	page := request.Page
	strategy := request.Strategy
	staker := request.StakerAddress
	if page == 0 {
		page = 1
	}
	if pageSize <= 10 {
		pageSize = 10
	}
	records := make([]*pb.L2UnStakeRecord, 0)
	unStakeList, total, currentPage := s.db.AppChainUnStake.ListAppChainUnStake(page, pageSize, staker, strategy)
	if len(unStakeList) > 0 {
		for _, unStake := range unStakeList {
			record := &pb.L2UnStakeRecord{
				Guid:          unStake.GUID.String(),
				BlockNumber:   unStake.BlockNumber.Int64(),
				TxHash:        unStake.TxHash.String(),
				EthAmount:     unStake.EthAmount.String(),
				LockedAmount:  unStake.DETHLocked.String(),
				ClaimTxHash:   unStake.ClaimTxHash.String(),
				L2Strategy:    unStake.L2Strategy.String(),
				Staker:        unStake.Staker.String(),
				Bridge:        unStake.Bridge.String(),
				SourceChainId: unStake.SourceChainId,
				DestChainId:   unStake.DestChainId,
				Status:        int32(unStake.Status),
				NotifyRelayer: unStake.NotifyRelayer,
				Created:       unStake.Created,
				Updated:       unStake.Updated,
			}
			records = append(records, record)
		}
	}
	resultPage := &pb.L2UnStakeRecordResponse_L2UnStakePage{
		CurrentPage: uint32(currentPage),
		PageSize:    pageSize,
		Total:       total,
		Records:     records,
	}
	return &pb.L2UnStakeRecordResponse{
		Success: true,
		Message: "success",
		Page:    resultPage,
	}, nil
}

func (s *RpcServer) L2StakeRecord(ctx context.Context, request *pb.L2StakeRecordRequest) (*pb.L2StakeRecordResponse, error) {
	page := request.Page
	pageSize := request.PageSize
	strategy := request.Strategy
	staker := request.StakerAddress
	records := make([]*pb.L2StakeRecord, 0)
	if page == 0 {
		page = 1
	}
	if pageSize <= 10 {
		pageSize = 10
	}
	stakeList, total, currentPage := s.db.AppChainStake.ListAppChainStake(page, pageSize, staker, strategy)

	if len(stakeList) > 0 {
		for _, stake := range stakeList {
			stakeHash := stake.TxHash
			increasedInfo := s.db.AppChainOperatorSharesIncreased.GetOperatorSharesIncreasedByTxHash(stakeHash)
			record := &pb.L2StakeRecord{
				Staker:       stake.Staker.String(),
				Strategy:     stake.StrategyAddress.String(),
				Shares:       stake.Shares.String(),
				TxHash:       stakeHash.String(),
				TokenAddress: stake.TokenAddress.String(),
				BlockNumber:  stake.BlockNumber.Int64(),
				Guid:         stake.GUID.String(),
				ChainId:      stake.ChainId,
				Status:       int32(increasedInfo.Status),
				UseShares:    increasedInfo.UseShares.String(),
				Created:      stake.Created,
			}
			records = append(records, record)
		}
	}

	resultPage := &pb.L2StakeRecordResponse_L2StakePage{
		CurrentPage: uint32(currentPage),
		PageSize:    pageSize,
		Total:       total,
		Records:     records,
	}
	return &pb.L2StakeRecordResponse{
		Success: true,
		Message: "success",
		Page:    resultPage,
	}, nil
}

func (s *RpcServer) L2WithdrawRecord(ctx context.Context, request *pb.L2WithdrawRecordRequest) (*pb.L2WithdrawRecordResponse, error) {
	page := request.Page
	pageSize := request.PageSize
	strategy := request.Strategy
	address := request.Address
	records := make([]*pb.L2WithdrawRecord, 0)
	if page == 0 {
		page = 1
	}
	if pageSize <= 10 {
		pageSize = 10
	}
	withdrawList, total, currentPage := s.db.AppChainWithdraw.ListAppChainWithdraw(page, pageSize, address, strategy)

	if len(withdrawList) > 0 {
		for _, withdraw := range withdrawList {
			stakeHash := withdraw.TxHash
			record := &pb.L2WithdrawRecord{
				Staker:      withdraw.Staker.String(),
				Strategy:    withdraw.Strategy.String(),
				Shares:      withdraw.Shares.String(),
				Operator:    withdraw.Operator.String(),
				TxHash:      stakeHash.String(),
				BlockNumber: withdraw.BlockNumber.Int64(),
				Guid:        withdraw.GUID.String(),
				ChainId:     withdraw.ChainId,
				Created:     withdraw.Created,
			}
			records = append(records, record)
		}
	}

	resultPage := &pb.L2WithdrawRecordResponse_L2WithdrawPage{
		CurrentPage: uint32(currentPage),
		PageSize:    pageSize,
		Total:       total,
		Records:     records,
	}
	return &pb.L2WithdrawRecordResponse{
		Success: true,
		Message: "success",
		Page:    resultPage,
	}, nil
}
