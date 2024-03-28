package bridge

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cornerstone-labs/acorus/rpc/bridge/protobuf/pb"
)

type BridgeRpcService interface {
	ChangeTransferStatus(sourceChainId, destChainId, txHash string) (*pb.CrossChainTransferStatusResponse, error)
	CrossChainTransfer(sourceChainId, destChainId, amount,
		receiveAddress, tokenAddress, fee, nonce, sourceHash string) (*pb.CrossChainTransferResponse, error)
	UpdateWithdrawFundingPoolBalance(sourceChainId, destChainId, amount,
		receiveAddress, tokenAddress, hash string) (*pb.UpdateWithdrawFundingPoolBalanceResponse, error)
	UpdateDepositFundingPoolBalance(sourceChainId, destChainId, amount,
		receiveAddress, tokenAddress, hash string) (*pb.UpdateDepositFundingPoolBalanceResponse, error)
	UnstakeBatch(sourceHash, batchId, bridgeAddress, sourceChainId, destChainId string) (*pb.UnstakeBatchResponse, error)
	UnstakeSingle(stakerAddress, strategyAddress, sharesAmount, chainId, sourceHash string) (*pb.UnstakeSingleResponse, error)
}

type bridgeRpcService struct {
	bRpcService pb.BridgeServiceClient
}

func NewBridgeRpcService(rpcUrl string) (BridgeRpcService, error) {
	conn, err := grpc.Dial(rpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("bridge rpc did not connect: ", err)
		return nil, err
	}
	bridgeServiceClient := pb.NewBridgeServiceClient(conn)
	brService := &bridgeRpcService{bridgeServiceClient}
	return brService, nil
}

func (r *bridgeRpcService) ChangeTransferStatus(sourceChainId, destChainId, txHash string) (*pb.CrossChainTransferStatusResponse, error) {
	ctx := context.Background()
	transferRequest := &pb.CrossChainTransferStatusRequest{SourceChainId: sourceChainId,
		DestChainId: destChainId, TxHash: txHash,
	}
	log.Info("ChangeTransferStatusRpc", "transferRequest", transferRequest)
	transferStatus, err := r.bRpcService.ChangeTransferStatus(ctx, transferRequest)
	return transferStatus, err
}

func (r *bridgeRpcService) CrossChainTransfer(sourceChainId, destChainId, amount,
	receiveAddress, tokenAddress, fee, nonce, sourceHash string) (*pb.CrossChainTransferResponse, error) {
	ctx := context.Background()
	crossChainReq := &pb.CrossChainTransferRequest{
		SourceChainId:  sourceChainId,
		DestChainId:    destChainId,
		Amount:         amount,
		ReceiveAddress: receiveAddress,
		TokenAddress:   tokenAddress,
		Fee:            fee,
		Nonce:          nonce,
		SourceHash:     sourceHash,
	}
	log.Info("CrossChainTransferRpc", "crossChainReq", crossChainReq)
	crossChainTransfer, err := r.bRpcService.CrossChainTransfer(ctx, crossChainReq)
	return crossChainTransfer, err
}

func (r *bridgeRpcService) UpdateDepositFundingPoolBalance(sourceChainId, destChainId, amount,
	receiveAddress, tokenAddress, hash string) (*pb.UpdateDepositFundingPoolBalanceResponse, error) {
	ctx := context.Background()
	updateFundingPoolReq := &pb.UpdateDepositFundingPoolBalanceRequest{
		SourceChainId:  sourceChainId,
		DestChainId:    destChainId,
		Amount:         amount,
		ReceiveAddress: receiveAddress,
		TokenAddress:   tokenAddress,
		SourceHash:     hash,
	}
	log.Info("UpdateDepositFundingPoolBalanceRpc", "updateFundingPoolReq", updateFundingPoolReq)
	poolBalanceResponse, err := r.bRpcService.UpdateDepositFundingPoolBalance(ctx, updateFundingPoolReq)
	return poolBalanceResponse, err
}

func (r *bridgeRpcService) UpdateWithdrawFundingPoolBalance(sourceChainId, destChainId, amount,
	receiveAddress, tokenAddress, hash string) (*pb.UpdateWithdrawFundingPoolBalanceResponse, error) {
	ctx := context.Background()
	updateFundingPoolReq := &pb.UpdateWithdrawFundingPoolBalanceRequest{
		SourceChainId:  sourceChainId,
		DestChainId:    destChainId,
		Amount:         amount,
		ReceiveAddress: receiveAddress,
		TokenAddress:   tokenAddress,
		SourceHash:     hash,
	}
	log.Info("UpdateWithdrawFundingPoolBalanceRpc", "updateFundingPoolReq", updateFundingPoolReq)
	poolBalanceResponse, err := r.bRpcService.UpdateWithdrawFundingPoolBalance(ctx, updateFundingPoolReq)
	return poolBalanceResponse, err
}

func (r *bridgeRpcService) UnstakeBatch(sourceHash, batchId, bridgeAddress, sourceChainId, destChainId string) (*pb.UnstakeBatchResponse, error) {
	ctx := context.Background()
	upstakeBatchReq := &pb.UnstakeBatchRequest{
		UnstakeRequestId: batchId,
		BridgeAddress:    bridgeAddress,
		SourceChainId:    sourceChainId,
		DestChainId:      destChainId,
		SourceHash:       sourceHash,
	}
	log.Info("UnstakeBatchRpc", "upstakeBatchReq", upstakeBatchReq)
	return r.bRpcService.UnstakeBatch(ctx, upstakeBatchReq)
}

func (r *bridgeRpcService) UnstakeSingle(stakerAddress, strategyAddress, sharesAmount, chainId, sourceHash string) (*pb.UnstakeSingleResponse, error) {
	ctx := context.Background()
	unstakeSingleReq := &pb.UnstakeSingleRequest{
		StakerAddress:   stakerAddress,
		StrategyAddress: strategyAddress,
		SharesAmount:    sharesAmount,
		ChainId:         chainId,
		SourceHash:      sourceHash,
	}
	log.Info("UnstakeSingleRpc", "UnstakeSingle", unstakeSingleReq)
	return r.bRpcService.UnstakeSingle(ctx, unstakeSingleReq)
}
