package polygon

import (
	"fmt"
	"github.com/cornerstone-labs/acorus/event/polygon/utils"
	"testing"
)

func TestAbi_contract(t *testing.T) {
	s := utils.DepositEventSignatureHash.String()
	fmt.Println(s)

	s2 := utils.ClaimEventSignatureHash.String()
	fmt.Println(s2)

	s3 := utils.OldClaimEventSignatureHash.String()
	fmt.Println(s3)
}
