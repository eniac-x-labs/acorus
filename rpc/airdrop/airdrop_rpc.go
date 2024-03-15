package airdrop

import (
	"context"
	"github.com/cornerstone-labs/acorus/rpc/airdrop/protobuf/airdrop"
	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AirDropRpcService interface {
	SubmitDppLinkPoints(consumerToken, pointType, address string) (*airdrop.DppLinkPointsRep, error)
}

type airDropRpcService struct {
	rpcService airdrop.AirdropServiceClient
}

func NewAirDropRpcService(rpcUrl string) (AirDropRpcService, error) {
	conn, err := grpc.Dial(rpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("airdrop rpc did not connect: ", err)
		return nil, err
	}
	aDServiceClient := airdrop.NewAirdropServiceClient(conn)
	rpcService := &airDropRpcService{aDServiceClient}
	return rpcService, nil
}

func (a airDropRpcService) SubmitDppLinkPoints(consumerToken, pointType, address string) (*airdrop.DppLinkPointsRep, error) {
	ctx := context.Background()
	dappLinkPointsReq := &airdrop.DppLinkPointsReq{
		ConsumerToken: consumerToken,
		Type:          pointType,
		Address:       address,
	}
	submitStatus, err := a.rpcService.SubmitDppLinkPoints(ctx, dappLinkPointsReq)
	return submitStatus, err
}
