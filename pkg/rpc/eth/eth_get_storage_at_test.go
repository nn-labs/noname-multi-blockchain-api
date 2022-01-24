package eth

import (
	"fmt"
	"testing"
)

func TestGetEthStorageAt(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{"0x4141D0Eb7252905bA4d98C0d330D16B9d49368fA", "0x0", "latest"}
	baseResponse, err := GetEthStorageAt(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
