package config

import (
	"os"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	type env struct {
		port               string
		appEnv             string
		gRpcHost           string
		btcRpcEndpointTest string
		btcRpcEndpointMain string
		btcRpcUser         string
		btcRpcPassword     string
		ethRpcEndpointTest string
		ethRpcEndpointMain string
	}

	type args struct {
		env env
	}

	setEnv := func(env env) {
		os.Setenv("PORT", env.port)
		os.Setenv("APP_ENV", env.appEnv)
		os.Setenv("GRPC_HOST", env.gRpcHost)
		os.Setenv("BTC_RPC_ENDPOINT_TEST", env.btcRpcEndpointTest)
		os.Setenv("BTC_RPC_ENDPOINT_MAIN", env.btcRpcEndpointMain)
		os.Setenv("BTC_RPC_USER", env.btcRpcUser)
		os.Setenv("BTC_RPC_PASSWORD", env.btcRpcPassword)
		os.Setenv("ETH_RPC_ENDPOINT_TEST", env.ethRpcEndpointTest)
		os.Setenv("ETH_RPC_ENDPOINT_MAIN", env.ethRpcEndpointMain)
	}

	tests := []struct {
		name      string
		args      args
		want      *Config
		wantError bool
	}{
		{
			name: "Test config file!",
			args: args{
				env: env{
					port:               ":5000",
					appEnv:             "development",
					gRpcHost:           "localhost:123321",
					btcRpcEndpointTest: "http://localhost",
					btcRpcEndpointMain: "http://localhost",
					btcRpcUser:         "user",
					btcRpcPassword:     "password",
					ethRpcEndpointTest: "http://localhost",
					ethRpcEndpointMain: "http://localhost",
				},
			},
			want: &Config{
				PORT:   ":5000",
				AppEnv: "development",
				GRps: GRps{
					GRpcHost: "localhost:123321",
				},
				BtcRpc: BtcRpc{
					BtcRpcEndpointTest: "http://localhost",
					BtcRpcEndpointMain: "http://localhost",
					BtcRpcUser:         "user",
					BtcRpcPassword:     "password",
				},
				EthRpc: EthRpc{
					EthRpcEndpointTest: "http://localhost",
					EthRpcEndpointMain: "http://localhost",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			setEnv(test.args.env)

			got, err := Get()
			if (err != nil) != test.wantError {
				t.Errorf("Init() error = %v, wantErr %v", err, test.wantError)

				return
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Init() got = %v, want %v", got, test.want)
			}
		})
	}
}
