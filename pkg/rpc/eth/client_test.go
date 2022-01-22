package eth

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"testing"
)
import "nn-blockchain-api/config"

func GetBaseSetupTest() IEthClient {
	logger := logrus.New()

	cfg, err := config.Get("../../../")
	if err != nil {
		logger.Fatalf("failed to load config: %v", err)
	}

	ethClient := NewEthClient(cfg.EthEndpoint)

	return ethClient
}

func TestClient(t *testing.T) {
	ethClient := GetBaseSetupTest()
	web3ClientVersion, err := ethClient.GetWeb3ClientVersion()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(web3ClientVersion)
}
