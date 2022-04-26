package config

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PORT   string `required:"true" default:"5000" envconfig:"PORT"`
	AppEnv string `required:"true" envconfig:"APP_ENV"`

	GRps
	BtcRpc
}

type GRps struct {
	GRpcHost string `required:"true" envconfig:"GRPC_HOST"`
}

type BtcRpc struct {
	BtcRpcEndpointTest string `required:"true" envconfig:"BTC_RPC_ENDPOINT_TEST"`
	BtcRpcEndpointMain string `required:"true" envconfig:"BTC_RPC_ENDPOINT_MAIN"`
	BtcRpcUser         string `required:"true" envconfig:"BTC_RPC_USER"`
	BtcRpcPassword     string `required:"true" envconfig:"BTC_RPC_PASSWORD"`
}

var (
	once   sync.Once
	config *Config
)

func Get() (*Config, error) {
	var err error
	once.Do(func() {
		var cfg Config
		// If you run it locally and through terminal please set up this in Load function (../.env)
		_ = godotenv.Load(".env")

		if err = envconfig.Process("", &cfg); err != nil {
			return
		}

		config = &cfg
	})

	return config, err
}
