package bitcoin

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"nn-blockchain-api/pkg/rpc/bitcoin"
	mock_bitcoin "nn-blockchain-api/pkg/rpc/bitcoin/mocks"
	"testing"
)

func TestNewService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	tests := []struct {
		name      string
		log       *logrus.Logger
		btcClient bitcoin.IBtcClient
		expect    func(*testing.T, Service, error)
	}{
		{
			name:      "should return bitcoin service",
			log:       logrus.New(),
			btcClient: mock_bitcoin.NewMockIBtcClient(controller),
			expect: func(t *testing.T, s Service, err error) {
				assert.NotNil(t, s)
				assert.Nil(t, err)
			},
		},
		{
			name:      "should return invalid logger",
			log:       nil,
			btcClient: mock_bitcoin.NewMockIBtcClient(controller),
			expect: func(t *testing.T, s Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "code: 500; status: internal_error; message: invalid logger")
			},
		},
		{
			name:      "should return invalid btc client",
			log:       logrus.New(),
			btcClient: nil,
			expect: func(t *testing.T, s Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "code: 500; status: internal_error; message: invalid btc client")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := NewService(tc.log, tc.btcClient)
			tc.expect(t, svc, err)
		})
	}
}

func TestService_StatusNode(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcClient := mock_bitcoin.NewMockIBtcClient(controller)

	service, _ := NewService(&logrus.Logger{}, btcClient)

	tests := []struct {
		name      string
		ctx       context.Context
		btcClient bitcoin.IBtcClient
		setup     func()
		expect    func(t *testing.T, status *StatusNodeDTO, err error)
	}{
		{
			name:      "should return status ok",
			ctx:       context.Background(),
			btcClient: btcClient,
			setup: func() {

			},
			expect: func(t *testing.T, status *StatusNodeDTO, err error) {
				assert.Nil(t, err)
				//assert.Equal(t, w.CoinName, dto.Name)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup()
			w, err := service.StatusNode(tc.ctx)
			tc.expect(t, w, err)
		})
	}
}
