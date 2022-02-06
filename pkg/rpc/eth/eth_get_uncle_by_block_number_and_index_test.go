package eth

import (
	"fmt"
	"testing"
)

func TestGetEthUncleByBlockNumberAndIndex(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{
		"0x29c", // 668
		"0x0",   // 0
	}
	baseResponse, err := GetEthUncleByBlockNumberAndIndex(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
