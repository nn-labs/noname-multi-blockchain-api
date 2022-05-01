package ethereum_rpc_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	ethereum_rpc "nn-blockchain-api/pkg/rpc/ethereum"
	mock_ethereum_rpc "nn-blockchain-api/pkg/rpc/ethereum/mocks"
	"testing"
)

func TestNewService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	tests := []struct {
		name      string
		ethClient ethereum_rpc.Client
		expect    func(*testing.T, ethereum_rpc.Service, error)
	}{
		{
			name:      "should return ethereum rpc service",
			ethClient: mock_ethereum_rpc.NewMockClient(controller),
			expect: func(t *testing.T, s ethereum_rpc.Service, err error) {
				assert.NotNil(t, s)
				assert.Nil(t, err)
			},
		},
		{
			name:      "should return invalid ethereum rpc client",
			ethClient: nil,
			expect: func(t *testing.T, s ethereum_rpc.Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "invalid ethereum rpc client")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := ethereum_rpc.NewService(tc.ethClient)
			tc.expect(t, svc, err)
		})
	}
}
