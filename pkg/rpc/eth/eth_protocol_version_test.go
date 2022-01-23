package eth

import (
	"fmt"
	"testing"
)

func TestGetEthProtocolVersion(t *testing.T) {
	ethClient := GetBaseSetupTest()
	ethProtocolVersion, err := GetEthProtocolVersion(ethClient)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(ethProtocolVersion)
}
