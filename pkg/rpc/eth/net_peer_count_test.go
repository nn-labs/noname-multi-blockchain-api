package eth

import (
	"fmt"
	"testing"
)

func TestGetNetPeerCount(t *testing.T) {
	ethClient := GetBaseSetupTest()
	netPeerCount, err := GetNetPeerCount(ethClient)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(netPeerCount)
}
