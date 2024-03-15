package main

import (
	"fmt"
	"github.com/cornerstone-labs/acorus/rpc/airdrop"
)

func main() {
	airDropGRpcUrl := "k8s-testnetl-hailston-bf4544780c-1757724324.ap-southeast-1.elb.amazonaws.com:80"
	airDropRpcService, err := airdrop.NewAirDropRpcService(airDropGRpcUrl)
	if err != nil {
		fmt.Println("airdrop rpc did not connect: ", err)
		return
	}
	points, err := airDropRpcService.SubmitDppLinkPoints("", "1", "address")
	fmt.Println(err)
	fmt.Println(points)
}
