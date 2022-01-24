package eth

import (
	"fmt"
	"testing"
)

func TestGetEthBlockTransactionCountByNumber(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{"latest"}
	baseResponse, err := GetEthBlockTransactionCountByNumber(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
