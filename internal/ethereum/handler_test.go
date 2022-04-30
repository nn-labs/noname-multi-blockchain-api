package ethereum_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"nn-blockchain-api/internal/ethereum"
	mock_ethereum "nn-blockchain-api/internal/ethereum/mocks"
	"testing"
)

func TestNewHandler(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	tests := []struct {
		name   string
		ethSvc ethereum.Service
		expect func(*testing.T, *ethereum.Handler, error)
	}{
		{
			name:   "should return service",
			ethSvc: mock_ethereum.NewMockService(controller),
			expect: func(t *testing.T, s *ethereum.Handler, err error) {
				assert.NotNil(t, s)
				assert.Nil(t, err)
			},
		},
		{
			name:   "should return invalid bitcoin service",
			ethSvc: nil,
			expect: func(t *testing.T, s *ethereum.Handler, err error) {
				assert.Nil(t, s)
				assert.NotNil(t, err)
				assert.EqualError(t, err, "invalid bitcoin service")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := ethereum.NewHandler(tc.ethSvc)
			tc.expect(t, svc, err)
		})
	}
}
