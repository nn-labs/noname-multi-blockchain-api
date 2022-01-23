package eth

import (
	"fmt"
	"testing"
)

func TestEthCoinbase(t *testing.T) {
	ethClient := GetBaseSetupTest()
	baseResponse, err := GetEthCoinbase(ethClient)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
