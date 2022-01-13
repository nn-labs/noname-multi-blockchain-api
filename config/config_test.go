package config

import (
	"os"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	type env struct {
		port        string
		environment string
		gRpcHost    string
	}

	type args struct {
		env env
	}

	setEnv := func(env env) {
		os.Setenv("PORT", env.port)
		os.Setenv("ENVIRONMENT", env.environment)
		os.Setenv("GRPC_HOST", env.gRpcHost)
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
					port:        ":5000",
					environment: "development",
					gRpcHost:    "localhost:123321",
				},
			},
			want: &Config{
				PORT:        ":5000",
				Environment: "development",
				GRpcHost:    "localhost:123321",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			setEnv(test.args.env)

			got, err := Get("..")
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
