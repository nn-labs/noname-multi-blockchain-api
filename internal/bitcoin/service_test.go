package bitcoin

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"nn-blockchain-api/pkg/errors"
	"nn-blockchain-api/pkg/rpc/bitcoin"
	mock_bitcoin "nn-blockchain-api/pkg/rpc/bitcoin/mocks"
	"testing"
)

func TestNewService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	tests := []struct {
		name         string
		log          *logrus.Logger
		btcTxSvc     bitcoin.TransactionService
		btcWalletSvc bitcoin.WalletService
		btcHealthSvc bitcoin.HealthService
		expect       func(*testing.T, Service, error)
	}{
		{
			name:         "should return bitcoin service",
			log:          logrus.New(),
			btcTxSvc:     mock_bitcoin.NewMockTransactionService(controller),
			btcWalletSvc: mock_bitcoin.NewMockWalletService(controller),
			btcHealthSvc: mock_bitcoin.NewMockHealthService(controller),
			expect: func(t *testing.T, s Service, err error) {
				assert.NotNil(t, s)
				assert.Nil(t, err)
			},
		},
		{
			name:         "should return invalid logger",
			log:          nil,
			btcTxSvc:     mock_bitcoin.NewMockTransactionService(controller),
			btcWalletSvc: mock_bitcoin.NewMockWalletService(controller),
			btcHealthSvc: mock_bitcoin.NewMockHealthService(controller),
			expect: func(t *testing.T, s Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "code: 500; status: internal_error; message: invalid logger")
			},
		},
		{
			name:         "should return invalid btc transaction service",
			log:          logrus.New(),
			btcTxSvc:     nil,
			btcWalletSvc: mock_bitcoin.NewMockWalletService(controller),
			btcHealthSvc: mock_bitcoin.NewMockHealthService(controller),
			expect: func(t *testing.T, s Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "code: 500; status: internal_error; message: invalid btc transaction service")
			},
		},
		{
			name:         "should return invalid btc wallet service",
			log:          logrus.New(),
			btcTxSvc:     mock_bitcoin.NewMockTransactionService(controller),
			btcWalletSvc: nil,
			btcHealthSvc: mock_bitcoin.NewMockHealthService(controller),
			expect: func(t *testing.T, s Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "code: 500; status: internal_error; message: invalid btc wallet service")
			},
		},
		{
			name:         "should return invalid btc health service",
			log:          logrus.New(),
			btcTxSvc:     mock_bitcoin.NewMockTransactionService(controller),
			btcWalletSvc: mock_bitcoin.NewMockWalletService(controller),
			btcHealthSvc: nil,
			expect: func(t *testing.T, s Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "code: 500; status: internal_error; message: invalid btc health service")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := NewService(tc.log, tc.btcTxSvc, tc.btcWalletSvc, tc.btcHealthSvc)
			tc.expect(t, svc, err)
		})
	}
}

func TestService_StatusNode(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcClient := mock_bitcoin.NewMockIBtcClient(controller)
	btcTxSvc := mock_bitcoin.NewMockTransactionService(controller)
	btcWalletSvc := mock_bitcoin.NewMockWalletService(controller)
	btcHealthSvc := mock_bitcoin.NewMockHealthService(controller)
	network := "test"

	service, _ := NewService(&logrus.Logger{}, btcTxSvc, btcWalletSvc, btcHealthSvc)

	status := bitcoin.StatusNode{
		Chain:                "test",
		Blocks:               2138184,
		Headers:              2138184,
		Verificationprogress: 0.9999989599050734,
		Softforks: struct {
			Bip34 struct {
				Type   string      `json:"type"`
				Active bool        `json:"active"`
				Height interface{} `json:"height"`
			} `json:"bip34"`
			Bip66 struct {
				Type   string      `json:"type"`
				Active bool        `json:"active"`
				Height interface{} `json:"height"`
			} `json:"bip66"`
			Bip65 struct {
				Type   string      `json:"type"`
				Active bool        `json:"active"`
				Height interface{} `json:"height"`
			} `json:"bip65"`
			Csv struct {
				Type   string      `json:"type"`
				Active bool        `json:"active"`
				Height interface{} `json:"height"`
			} `json:"csv"`
			Segwit struct {
				Type   string      `json:"type"`
				Active bool        `json:"active"`
				Height interface{} `json:"height"`
			} `json:"segwit"`
			Taproot struct {
				Type string `json:"type"`
				Bip9 struct {
					Status              string      `json:"status"`
					StartTime           interface{} `json:"start_time"`
					Timeout             interface{} `json:"timeout"`
					Since               interface{} `json:"since"`
					MinActivationHeight int         `json:"min_activation_height"`
				} `json:"bip9"`
				Active bool `json:"active"`
			} `json:"taproot"`
		}{
			Bip34: struct {
				Type   string      `json:"type"`
				Active bool        `json:"active"`
				Height interface{} `json:"height"`
			}{
				Type:   "buried",
				Active: true,
				Height: 21111,
			},
			Bip66: struct {
				Type   string      `json:"type"`
				Active bool        `json:"active"`
				Height interface{} `json:"height"`
			}{
				Type:   "buried",
				Active: true,
				Height: 330776,
			},
			Bip65: struct {
				Type   string      `json:"type"`
				Active bool        `json:"active"`
				Height interface{} `json:"height"`
			}{
				Type:   "buried",
				Active: true,
				Height: 581885,
			},
			Csv: struct {
				Type   string      `json:"type"`
				Active bool        `json:"active"`
				Height interface{} `json:"height"`
			}{
				Type:   "buried",
				Active: true,
				Height: 770112,
			},
			Segwit: struct {
				Type   string      `json:"type"`
				Active bool        `json:"active"`
				Height interface{} `json:"height"`
			}{
				Type:   "buried",
				Active: true,
				Height: 834624,
			},
			Taproot: struct {
				Type string `json:"type"`
				Bip9 struct {
					Status              string      `json:"status"`
					StartTime           interface{} `json:"start_time"`
					Timeout             interface{} `json:"timeout"`
					Since               interface{} `json:"since"`
					MinActivationHeight int         `json:"min_activation_height"`
				} `json:"bip9"`
				Active bool `json:"active"`
			}{
				Type: "bip9",
				Bip9: struct {
					Status              string      `json:"status"`
					StartTime           interface{} `json:"start_time"`
					Timeout             interface{} `json:"timeout"`
					Since               interface{} `json:"since"`
					MinActivationHeight int         `json:"min_activation_height"`
				}{
					Status:              "active",
					StartTime:           1619222400,
					Timeout:             1628640000,
					Since:               2011968,
					MinActivationHeight: 0,
				},
				Active: true,
			},
		},
		Warnings: "Unknown new rules activated (versionbit 28)",
	}

	dto := &StatusNodeDTO{Network: network}

	tests := []struct {
		name      string
		ctx       context.Context
		btcClient bitcoin.IBtcClient
		dto       *StatusNodeDTO
		setup     func(dto *StatusNodeDTO)
		expect    func(t *testing.T, status *StatusNodeInfoDTO, err error)
	}{
		{
			name:      "should return ok",
			ctx:       context.Background(),
			btcClient: btcClient,
			dto:       dto,
			setup: func(dto *StatusNodeDTO) {
				btcHealthSvc.EXPECT().Status(dto.Network).Return(&status, nil)
			},
			expect: func(t *testing.T, status *StatusNodeInfoDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, status.Chain, network)
			},
		},
		{
			name:      "should return failed check node status",
			ctx:       context.Background(),
			btcClient: btcClient,
			dto:       dto,
			setup: func(dto *StatusNodeDTO) {
				btcHealthSvc.EXPECT().Status(dto.Network).Return(nil, errors.NewInternal("failed check node status"))
			},
			expect: func(t *testing.T, status *StatusNodeInfoDTO, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, err, errors.NewInternal("failed check node status"))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.dto)
			w, err := service.StatusNode(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}
