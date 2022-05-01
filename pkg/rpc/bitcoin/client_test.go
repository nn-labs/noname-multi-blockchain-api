package bitcoin_rpc_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	bitcoin_rpc "nn-blockchain-api/pkg/rpc/bitcoin"
	"testing"
)

func TestNewClient(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcRpcTest := "localhost:1234"
	btcRpcMain := "localhost:4321"
	btcRpcUser := "user"
	btcRpcPassword := "password"

	tests := []struct {
		name                  string
		btcRpcEndpointTestNet string
		btcRpcEndpointMainNet string
		btcUser               string
		btcPassword           string
		expect                func(*testing.T, bitcoin_rpc.Client, error)
	}{
		{
			name:                  "should return bitcoin rpc client",
			btcRpcEndpointTestNet: btcRpcTest,
			btcRpcEndpointMainNet: btcRpcMain,
			btcUser:               btcRpcUser,
			btcPassword:           btcRpcPassword,
			expect: func(t *testing.T, c bitcoin_rpc.Client, err error) {
				assert.NotNil(t, c)
				assert.Nil(t, err)
			},
		},
		{
			name:                  "should return invalid bitcoin rpc testnet endpoint",
			btcRpcEndpointTestNet: "",
			btcRpcEndpointMainNet: btcRpcMain,
			btcUser:               btcRpcUser,
			btcPassword:           btcRpcPassword,
			expect: func(t *testing.T, c bitcoin_rpc.Client, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, c)
				assert.EqualError(t, err, "invalid bitcoin rpc testnet endpoint")
			},
		},
		{
			name:                  "should return invalid bitcoin rpc mainnet endpoint",
			btcRpcEndpointTestNet: btcRpcTest,
			btcRpcEndpointMainNet: "",
			btcUser:               btcRpcUser,
			btcPassword:           btcRpcPassword,
			expect: func(t *testing.T, c bitcoin_rpc.Client, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, c)
				assert.EqualError(t, err, "invalid bitcoin rpc mainnet endpoint")
			},
		},
		{
			name:                  "should return invalid bitcoin rpc user",
			btcRpcEndpointTestNet: btcRpcTest,
			btcRpcEndpointMainNet: btcRpcMain,
			btcUser:               "",
			btcPassword:           btcRpcPassword,
			expect: func(t *testing.T, c bitcoin_rpc.Client, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, c)
				assert.EqualError(t, err, "invalid bitcoin rpc user")
			},
		},
		{
			name:                  "should return invalid bitcoin rpc password",
			btcRpcEndpointTestNet: btcRpcTest,
			btcRpcEndpointMainNet: btcRpcMain,
			btcUser:               btcRpcUser,
			btcPassword:           "",
			expect: func(t *testing.T, c bitcoin_rpc.Client, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, c)
				assert.EqualError(t, err, "invalid bitcoin rpc password")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := bitcoin_rpc.NewClient(tc.btcRpcEndpointTestNet, tc.btcRpcEndpointMainNet, tc.btcUser, tc.btcPassword)
			tc.expect(t, svc, err)
		})
	}
}
