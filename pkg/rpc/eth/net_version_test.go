package eth

import (
	"fmt"
	"testing"
)

func TestNetVersion(t *testing.T) {
	ethClient := GetBaseSetupTest()
	netVersion, err := GetNetVersion(ethClient)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(netVersion)
}
