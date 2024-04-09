package main

import (
	"fmt"
	common2 "github.com/cornerstone-labs/acorus/common"
)

func main() {
	batchId := common2.GetBatchId()
	fmt.Println(batchId)
}
