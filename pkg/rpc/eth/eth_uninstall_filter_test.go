package eth

import (
	"fmt"
	"testing"
)

func TestGetEthUninstallFilter(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{"0xb"}
	baseResponse, err := GetEthUninstallFilter(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
