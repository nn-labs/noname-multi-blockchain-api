package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	PORT               string `mapstructure:"PORT"`
	Environment        string `mapstructure:"APP_ENV"`
	GRpcHost           string `mapstructure:"GRPC_HOST"`
	BtcRpcEndpointTest string `mapstructure:"BTC_RPC_ENDPOINT_TEST"`
	BtcRpcEndpointMain string `mapstructure:"BTC_RPC_ENDPOINT_MAIN"`
	BtcRpcUser         string `mapstructure:"BTC_RPC_USER"`
	BtcRpcPassword     string `mapstructure:"BTC_RPC_PASSWORD"`
}

func Get(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var configuration Config
	err = viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return &configuration, nil
}
