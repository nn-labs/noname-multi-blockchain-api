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
		name   string
		ctx    context.Context
		dto    *StatusNodeDTO
		setup  func(dto *StatusNodeDTO)
		expect func(t *testing.T, status *StatusNodeInfoDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(dto *StatusNodeDTO) {
				btcHealthSvc.EXPECT().Status(dto.Network).Return(&status, nil)
			},
			expect: func(t *testing.T, status *StatusNodeInfoDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, status.Chain, network)
			},
		},
		{
			name: "should return failed check node status",
			ctx:  context.Background(),
			dto:  dto,
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

func TestService_CreateTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcTxSvc := mock_bitcoin.NewMockTransactionService(controller)
	btcWalletSvc := mock_bitcoin.NewMockWalletService(controller)
	btcHealthSvc := mock_bitcoin.NewMockHealthService(controller)

	service, _ := NewService(&logrus.Logger{}, btcTxSvc, btcWalletSvc, btcHealthSvc)

	tx := "transaction"
	fee := 0.0000259

	dto := &CreateRawTransactionDTO{
		Utxo: []struct {
			TxId     string `json:"txid" validate:"required"`
			Vout     int64  `json:"vout" validate:"required"`
			Amount   int64  `json:"amount" validate:"required"`
			PKScript string `json:"pk_script" validate:"required"`
		}{
			{
				TxId:     "989d301c546841d0ac5c8354c7d78079e3603b089682d1639b2ee1c1a8010c6a",
				Vout:     1,
				Amount:   1045428,
				PKScript: "76a914690cd6356789d30b99063632e0651a8d0c206c7f88ac",
			},
		},
		FromAddress: "mq6Qd7JJKsgBYkMFsGCk24MHMxUkuyTnkU",
		ToAddress:   "mmfbzo2533SFa34ErmYNY4RdVtfw5XYK1u",
		Amount:      10000,
		Network:     "test",
	}

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *CreateRawTransactionDTO
		setup  func(dto *CreateRawTransactionDTO)
		expect func(t *testing.T, createdTx *CreatedRawTransactionDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(dto *CreateRawTransactionDTO) {
				btcTxSvc.EXPECT().CreateTransaction(bitcoin.UTXO(dto.Utxo), dto.FromAddress, dto.ToAddress, dto.Amount, dto.Network).Return(&tx, &fee, nil)
			},
			expect: func(t *testing.T, createdTx *CreatedRawTransactionDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, createdTx.Tx, tx)
				assert.Equal(t, createdTx.Fee, fee)
			},
		},
		{
			name: "should return error",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(dto *CreateRawTransactionDTO) {
				btcTxSvc.EXPECT().CreateTransaction(bitcoin.UTXO(dto.Utxo), dto.FromAddress, dto.ToAddress, dto.Amount, dto.Network).Return(nil, nil, errors.NewInternal("failed to create transaction"))
			},
			expect: func(t *testing.T, createdTx *CreatedRawTransactionDTO, err error) {
				assert.Nil(t, createdTx)
				assert.Equal(t, err, errors.WithMessage(err, "code: 500; status: internal_error; message: failed to create transaction"))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.dto)
			w, err := service.CreateTransaction(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}

func TestService_DecodeTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcTxSvc := mock_bitcoin.NewMockTransactionService(controller)
	btcWalletSvc := mock_bitcoin.NewMockWalletService(controller)
	btcHealthSvc := mock_bitcoin.NewMockHealthService(controller)

	service, _ := NewService(&logrus.Logger{}, btcTxSvc, btcWalletSvc, btcHealthSvc)

	dto := &DecodeRawTransactionDTO{
		Tx:      "transaction",
		Network: "test",
	}

	decodeTx := &bitcoin.DecodedTx{
		Txid:     "tx",
		Hash:     "hash",
		Version:  0,
		Size:     0,
		Vsize:    0,
		Weight:   0,
		Locktime: 0,
		Vin: []struct {
			Txid      string `json:"txid"`
			Vout      int    `json:"vout"`
			ScriptSig struct {
				Asm string `json:"asm"`
				Hex string `json:"hex"`
			} `json:"scriptSig"`
			Sequence int64 `json:"sequence"`
		}{
			{
				Txid: "tx",
				Vout: 0,
				ScriptSig: struct {
					Asm string `json:"asm"`
					Hex string `json:"hex"`
				}{
					Asm: "asm",
					Hex: "hex",
				},
				Sequence: 0,
			},
		},
		Vout: []struct {
			Value        float64 `json:"value"`
			N            int     `json:"n"`
			ScriptPubKey struct {
				Asm     string `json:"asm"`
				Hex     string `json:"hex"`
				Address string `json:"address"`
				Type    string `json:"type"`
			} `json:"scriptPubKey"`
		}{
			{
				Value: 0,
				N:     0,
				ScriptPubKey: struct {
					Asm     string `json:"asm"`
					Hex     string `json:"hex"`
					Address string `json:"address"`
					Type    string `json:"type"`
				}{
					Asm:     "asm",
					Hex:     "hex",
					Address: "address",
					Type:    "type",
				},
			},
		},
	}

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *DecodeRawTransactionDTO
		setup  func(dto *DecodeRawTransactionDTO)
		expect func(t *testing.T, createdTx *DecodedRawTransactionDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(dto *DecodeRawTransactionDTO) {
				btcTxSvc.EXPECT().DecodeTransaction(dto.Tx, dto.Network).Return(decodeTx, nil)
			},
			expect: func(t *testing.T, decodeTx *DecodedRawTransactionDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, decodeTx.Txid, "tx")
			},
		},
		{
			name: "should return error",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(dto *DecodeRawTransactionDTO) {
				btcTxSvc.EXPECT().DecodeTransaction(dto.Tx, dto.Network).Return(nil, errors.NewInternal("failed to decode transaction"))
			},
			expect: func(t *testing.T, decodeTx *DecodedRawTransactionDTO, err error) {
				assert.Nil(t, decodeTx)
				assert.Equal(t, err, errors.WithMessage(err, "code: 500; status: internal_error; message: failed to decode transaction"))

			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.dto)
			w, err := service.DecodeTransaction(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}

func TestService_FoundForRawTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcTxSvc := mock_bitcoin.NewMockTransactionService(controller)
	btcWalletSvc := mock_bitcoin.NewMockWalletService(controller)
	btcHealthSvc := mock_bitcoin.NewMockHealthService(controller)

	service, _ := NewService(&logrus.Logger{}, btcTxSvc, btcWalletSvc, btcHealthSvc)

	dto := &FundForRawTransactionDTO{
		CreatedTxHex:  "tx",
		ChangeAddress: "address",
		Network:       "test",
	}

	tx := "transaction"
	fee := 0.0000259

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *FundForRawTransactionDTO
		setup  func(dto *FundForRawTransactionDTO)
		expect func(t *testing.T, createdTx *FundedRawTransactionDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(dto *FundForRawTransactionDTO) {
				btcTxSvc.EXPECT().FundForTransaction(dto.CreatedTxHex, dto.ChangeAddress, dto.Network).Return(tx, &fee, nil)
			},
			expect: func(t *testing.T, fundedTx *FundedRawTransactionDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, fundedTx.Tx, tx)
				assert.Equal(t, fundedTx.Fee, fee)
			},
		},
		{
			name: "should return error",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(dto *FundForRawTransactionDTO) {
				btcTxSvc.EXPECT().FundForTransaction(dto.CreatedTxHex, dto.ChangeAddress, dto.Network).Return("", nil, errors.NewInternal("failed to fund for transaction"))
			},
			expect: func(t *testing.T, fundedTx *FundedRawTransactionDTO, err error) {
				assert.Nil(t, fundedTx)
				assert.Equal(t, err, errors.WithMessage(err, "code: 500; status: internal_error; message: failed to fund for transaction"))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.dto)
			w, err := service.FoundForRawTransaction(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}
