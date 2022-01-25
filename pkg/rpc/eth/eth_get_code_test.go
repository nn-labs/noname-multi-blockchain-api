package eth

import (
	"fmt"
	"testing"
)

func TestGetEthCode(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{"0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b", "0x2"}
	baseResponse, err := GetEthCode(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
