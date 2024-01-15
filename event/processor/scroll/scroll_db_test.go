package scroll

import (
	"context"
	"flag"
	"fmt"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestDB_contract(t *testing.T) {
	var f = flag.String("c", "../../../acorus.yaml", "config path")
	flag.Parse()

	conf, err := config.New(*f)
	if err != nil {
		panic(err)
	}
	fmt.Println(conf.RPCs)
	ctx := context.Background()
	db, _ := database.NewDB(ctx, conf.MasterDb)
	contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(conf.RPCs[1].EventContracts[0])}

	db.ContractEvents.ContractEventsWithFilter("534352", contractEventFilter, big.NewInt(0), big.NewInt(100))

}
