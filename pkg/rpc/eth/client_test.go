package eth

import (
	"fmt"
	"nn-blockchain-api/config"
	"testing"

	"github.com/sirupsen/logrus"
)

func GetBaseSetupTest() IEthClient {
	logger := logrus.New()

	cfg, err := config.Get()
	if err != nil {
		logger.Fatalf("failed to load config: %v", err)
	}

	ethClient := NewEthClient(cfg.EthRpc.EthRpcEndpointTest)

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
