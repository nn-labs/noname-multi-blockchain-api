package eth

import (
	"fmt"
	"testing"
)

func TestGetEthGasPrice(t *testing.T) {
	ethClient := GetBaseSetupTest()
	baseResponse, err := GetEthGasPrice(ethClient)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
