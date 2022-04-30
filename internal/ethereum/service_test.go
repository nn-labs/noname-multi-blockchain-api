package ethereum_test

import (
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"nn-blockchain-api/internal/ethereum"
	ethereum_rpc "nn-blockchain-api/pkg/rpc/ethereum"
	mock_ethereum_rpc "nn-blockchain-api/pkg/rpc/ethereum/mocks"
	"testing"
)

func TestNewService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	tests := []struct {
		name      string
		ethRpcSvc ethereum_rpc.Service
		log       *logrus.Logger
		expect    func(*testing.T, ethereum.Service, error)
	}{
		{
			name:      "should return ethereum service",
			ethRpcSvc: mock_ethereum_rpc.NewMockService(controller),
			log:       logrus.New(),
			expect: func(t *testing.T, s ethereum.Service, err error) {
				assert.NotNil(t, s)
				assert.Nil(t, err)
			},
		},
		{
			name:      "should return invalid ethereum rpc service",
			ethRpcSvc: nil,
			log:       logrus.New(),
			expect: func(t *testing.T, s ethereum.Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "invalid ethereum rpc service")
			},
		},
		{
			name:      "should return invalid logger",
			ethRpcSvc: mock_ethereum_rpc.NewMockService(controller),
			log:       nil,
			expect: func(t *testing.T, s ethereum.Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "invalid logger")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := ethereum.NewService(tc.ethRpcSvc, tc.log)
			tc.expect(t, svc, err)
		})
	}
}
