package bitcoin_rpc_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	bitcoin_rpc "nn-blockchain-api/pkg/rpc/bitcoin"
	mock_bitcoin_rpc "nn-blockchain-api/pkg/rpc/bitcoin/mocks"
	"testing"
)

func TestNewService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	tests := []struct {
		name      string
		btcClient bitcoin_rpc.Client
		expect    func(*testing.T, bitcoin_rpc.Service, error)
	}{
		{
			name:      "should return bitcoin rpc service",
			btcClient: mock_bitcoin_rpc.NewMockClient(controller),
			expect: func(t *testing.T, s bitcoin_rpc.Service, err error) {
				assert.NotNil(t, s)
				assert.Nil(t, err)
			},
		},
		{
			name:      "should return invalid bitcoin rpc client",
			btcClient: nil,
			expect: func(t *testing.T, s bitcoin_rpc.Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "invalid bitcoin rpc client")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := bitcoin_rpc.NewService(tc.btcClient)
			tc.expect(t, svc, err)
		})
	}
}
