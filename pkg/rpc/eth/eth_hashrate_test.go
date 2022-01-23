package eth

import (
	"fmt"
	"testing"
)

func TestGetEthHashrate(t *testing.T) {
	ethClient := GetBaseSetupTest()
	baseResponse, err := GetEthHashrate(ethClient)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
