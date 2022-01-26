package eth

import (
	"fmt"
	"testing"
)

func TestGetEthBlockByHash(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{
		"0xdc0818cf78f21a8e70579cb46a43643f78291264dda342ae31049421c82d21ae",
		"false",
	}
	ethGas, err := GetEthBlockByHash(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(ethGas)
}
