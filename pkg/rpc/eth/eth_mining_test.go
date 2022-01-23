package eth

import (
	"fmt"
	"testing"
)

func TestGetEthMining(t *testing.T) {
	ethClient := GetBaseSetupTest()
	baseResponse, err := GetEthMining(ethClient)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
