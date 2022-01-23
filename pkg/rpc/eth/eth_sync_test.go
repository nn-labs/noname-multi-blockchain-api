package eth

import (
	"fmt"
	"testing"
)

func TestEthSync(t *testing.T) {
	ethClient := GetBaseSetupTest()
	baseResponse, ethSyncResponse, err := GetEthSync(ethClient)
	if err != nil {
		t.Errorf(err.Error())
	}

	if baseResponse != nil {
		fmt.Println(baseResponse)
	}

	if ethSyncResponse != nil {
		fmt.Println(ethSyncResponse)
	}
}
