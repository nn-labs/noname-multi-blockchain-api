package eth

import (
	"fmt"
	"testing"
)

func TestGetEthNewBlockFilter(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{}
	baseResponse, err := GetEthNewBlockFilter(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
