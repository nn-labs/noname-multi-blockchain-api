package bitcoin_test

import (
	"nn-blockchain-api/internal/bitcoin"
	mock_bitcoin "nn-blockchain-api/internal/bitcoin/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	tests := []struct {
		name   string
		btcSvc bitcoin.Service
		expect func(*testing.T, *bitcoin.Handler, error)
	}{
		{
			name:   "should return service",
			btcSvc: mock_bitcoin.NewMockService(controller),
			expect: func(t *testing.T, s *bitcoin.Handler, err error) {
				assert.NotNil(t, s)
				assert.Nil(t, err)
			},
		},
		{
			name:   "should return invalid bitcoin service",
			btcSvc: nil,
			expect: func(t *testing.T, s *bitcoin.Handler, err error) {
				assert.Nil(t, s)
				assert.NotNil(t, err)
				assert.EqualError(t, err, "invalid bitcoin service")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := bitcoin.NewHandler(tc.btcSvc)
			tc.expect(t, svc, err)
		})
	}
}
