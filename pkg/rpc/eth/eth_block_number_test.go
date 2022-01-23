package eth

import (
	"fmt"
	"testing"
)

func TestGetEthBlockNumber(t *testing.T) {
	ethClient := GetBaseSetupTest()
	baseResponse, err := GetEthBlockNumber(ethClient)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
