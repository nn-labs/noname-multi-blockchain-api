package eth

import (
	"fmt"
	"testing"
)

func TestEthSign(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{"0x9b2055d370f73ec7d8a03e965129118dc8f5bf83", "0xdeadbeaf"}
	ethProtocolVersion, err := GetEthSign(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(ethProtocolVersion)
}
