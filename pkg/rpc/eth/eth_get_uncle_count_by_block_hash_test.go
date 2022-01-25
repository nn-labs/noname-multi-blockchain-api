package eth

import (
	"fmt"
	"testing"
)

func TestGetEthUncleCountByBlockHash(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{"0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238"}
	baseResponse, err := GetEthUncleCountByBlockHash(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
