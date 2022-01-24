package eth

import (
	"fmt"
	"testing"
)

func TestGetEthBlockTransactionCountByHash(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{"0x4141D0Eb7252905bA4d98C0d330D16B9d49368fA", "latest"}
	baseResponse, err := GetEthBlockTransactionCountByHash(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
