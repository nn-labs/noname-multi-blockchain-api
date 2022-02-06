package eth

import (
	"fmt"
	"testing"
)

func TestGetEthNewPendingFilter(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{}
	baseResponse, err := GetEthNewPendingTransactionFilter(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
