package eth

import (
	"fmt"
	"testing"
)

func TestGetEthUncleCountByBlockNumber(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{"0xe8"}
	baseResponse, err := GetEthUncleCountByBlockNumber(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
