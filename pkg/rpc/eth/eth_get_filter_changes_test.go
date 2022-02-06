package eth

import (
	"fmt"
	"testing"
)

func TestGetEthFilterChanges(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{
		"0x16",
	}
	baseResponse, err := GetEthFilterChanges(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
