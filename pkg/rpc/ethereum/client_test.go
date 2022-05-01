package ethereum_rpc_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	ethereum_rpc "nn-blockchain-api/pkg/rpc/ethereum"
	"testing"
)

func TestNewClient(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	ethRpcTest := "localhost:1234"
	ethRpcMain := "localhost:4321"

	tests := []struct {
		name                  string
		ethRpcEndpointTestNet string
		ethRpcEndpointMainNet string
		expect                func(*testing.T, ethereum_rpc.Client, error)
	}{
		{
			name:                  "should return ethereum rpc client",
			ethRpcEndpointTestNet: ethRpcTest,
			ethRpcEndpointMainNet: ethRpcMain,
			expect: func(t *testing.T, c ethereum_rpc.Client, err error) {
				assert.NotNil(t, c)
				assert.Nil(t, err)
			},
		},
		{
			name:                  "should return invalid ethereum rpc testnet endpoint",
			ethRpcEndpointTestNet: "",
			ethRpcEndpointMainNet: ethRpcMain,
			expect: func(t *testing.T, c ethereum_rpc.Client, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, c)
				assert.EqualError(t, err, "invalid ethereum rpc testnet endpoint")
			},
		},
		{
			name:                  "should return invalid ethereum rpc mainnet endpoint",
			ethRpcEndpointTestNet: ethRpcTest,
			ethRpcEndpointMainNet: "",
			expect: func(t *testing.T, c ethereum_rpc.Client, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, c)
				assert.EqualError(t, err, "invalid ethereum rpc mainnet endpoint")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := ethereum_rpc.NewClient(tc.ethRpcEndpointTestNet, tc.ethRpcEndpointMainNet)
			tc.expect(t, svc, err)
		})
	}
}
