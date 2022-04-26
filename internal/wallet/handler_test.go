package wallet_test

import (
	"nn-blockchain-api/internal/wallet"
	mock_wallet "nn-blockchain-api/internal/wallet/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	tests := []struct {
		name      string
		walletSvc wallet.Service
		expect    func(*testing.T, *wallet.Handler, error)
	}{
		{
			name:      "should return service",
			walletSvc: mock_wallet.NewMockService(controller),
			expect: func(t *testing.T, s *wallet.Handler, err error) {
				assert.NotNil(t, s)
				assert.Nil(t, err)
			},
		},
		{
			name:      "should return invalid wallet service",
			walletSvc: nil,
			expect: func(t *testing.T, s *wallet.Handler, err error) {
				assert.Nil(t, s)
				assert.NotNil(t, err)
				assert.EqualError(t, err, "invalid wallet service")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := wallet.NewHandler(tc.walletSvc)
			tc.expect(t, svc, err)
		})
	}
}
