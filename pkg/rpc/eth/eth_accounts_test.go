package eth

import (
	"fmt"
	"testing"
)

func TestGetEthAccounts(t *testing.T) {
	ethClient := GetBaseSetupTest()
	baseResponse, err := GetEthAccounts(ethClient)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
