package eth

import (
	"fmt"
	"testing"
)

func TestGetEthTransactionByBlockNumberAndIndex(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{
		"0x29c", // 668
		"0x0",   // 0
	}
	baseResponse, err := getEthTransactionByBlockNumberAndIndex(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
