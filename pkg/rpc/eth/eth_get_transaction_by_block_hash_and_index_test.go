package eth

import (
	"fmt"
	"testing"
)

func TestGetEthTransactionByBlockHashAndIndex(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{
		"0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331",
		"0x0",
	}
	baseResponse, err := getEthTransactionByBlockHashAndIndex(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
