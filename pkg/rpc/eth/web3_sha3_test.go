package eth

import (
	"fmt"
	"testing"
)

func TestWeb3Sha3(t *testing.T) {
	ethClient := GetBaseSetupTest()
	baseResponse, err := GetWeb3Sha3(ethClient, []string{"0x68656c6c6f20776f726c64"})
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
