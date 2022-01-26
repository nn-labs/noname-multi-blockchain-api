package eth

import (
	"fmt"
	"testing"
)

func TestGetEthCall(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := map[string]string{
		"": "{see above}",
	}
	ethProtocolVersion, err := GetEthCall(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(ethProtocolVersion)
}
