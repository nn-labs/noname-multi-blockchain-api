package eth

import (
	"fmt"
	"testing"
)

func GetBaseSetupTest() IEthClient {
	ethClient := NewEthClient("http://localhost:123")

	return ethClient
}

func TestClient(t *testing.T) {
	ethClient := GetBaseSetupTest()
	web3ClientVersion, err := ethClient.GetWeb3ClientVersion()
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(web3ClientVersion)
}
