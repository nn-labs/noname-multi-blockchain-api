package wallet

import (
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	pb "nn-blockchain-api/pkg/grpc_client/proto/wallet"
	grpc_mock "nn-blockchain-api/pkg/grpc_client/proto/wallet/mocks"
	"testing"
)

func TestNewService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	tests := []struct {
		name         string
		walletClient pb.WalletServiceClient
		log          *logrus.Logger
		expect       func(*testing.T, Service, error)
	}{
		{
			name:         "should return wallet service",
			walletClient: grpc_mock.NewMockWalletServiceClient(controller),
			log:          logrus.New(),
			expect: func(t *testing.T, s Service, err error) {
				assert.NotNil(t, s)
				assert.Nil(t, err)
			},
		},
		{
			name:         "should return invalid wallet client",
			walletClient: nil,
			log:          logrus.New(),
			expect: func(t *testing.T, s Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "code: 500; status: internal_error; message: invalid wallet client")
			},
		},
		{
			name:         "should return invalid logger",
			walletClient: grpc_mock.NewMockWalletServiceClient(controller),
			log:          nil,
			expect: func(t *testing.T, s Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "code: 500; status: internal_error; message: invalid logger")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := NewService(tc.walletClient, tc.log)
			tc.expect(t, svc, err)
		})
	}
}
