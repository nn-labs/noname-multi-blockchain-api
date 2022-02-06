package eth

import (
	"fmt"
	"testing"
)

func TestGetEthNewFilter(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{}
	baseResponse, err := GetEthNewFilter(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
