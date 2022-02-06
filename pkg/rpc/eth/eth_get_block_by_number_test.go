package eth

import (
	"fmt"
	"testing"
)

func TestGetEthBlockByNumber(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{
		"0x1b4", // 436
		"true",
	}
	ethGas, err := GetEthBlockByNumber(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(ethGas)
}
