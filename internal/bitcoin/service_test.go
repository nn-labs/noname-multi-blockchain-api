package bitcoin_test

import (
	"context"
	"nn-blockchain-api/internal/bitcoin"
	"nn-blockchain-api/pkg/errors"
	bitcoin_rpc "nn-blockchain-api/pkg/rpc/bitcoin"

	mock_bitcoin_rpc "nn-blockchain-api/pkg/rpc/bitcoin/mocks"

	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	tests := []struct {
		name      string
		log       *logrus.Logger
		btcRpcSvc bitcoin_rpc.Service
		expect    func(*testing.T, bitcoin.Service, error)
	}{
		{
			name:      "should return bitcoin service",
			log:       logrus.New(),
			btcRpcSvc: mock_bitcoin_rpc.NewMockService(controller),
			expect: func(t *testing.T, s bitcoin.Service, err error) {
				assert.NotNil(t, s)
				assert.Nil(t, err)
			},
		},
		{
			name:      "should return invalid btc rpc service",
			btcRpcSvc: nil,
			log:       logrus.New(),
			expect: func(t *testing.T, s bitcoin.Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "invalid btc rpc service")
			},
		},
		{
			name:      "should return invalid logger",
			btcRpcSvc: mock_bitcoin_rpc.NewMockService(controller),
			log:       nil,
			expect: func(t *testing.T, s bitcoin.Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "invalid logger")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := bitcoin.NewService(tc.btcRpcSvc, tc.log)
			tc.expect(t, svc, err)
		})
	}
}

func TestService_StatusNode(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcRpcSvc := mock_bitcoin_rpc.NewMockService(controller)
	network := "test"

	service, _ := bitcoin.NewService(btcRpcSvc, &logrus.Logger{})

	status := bitcoin_rpc.StatusNode{
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

	dto := &bitcoin.StatusNodeDTO{Network: network}

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *bitcoin.StatusNodeDTO
		setup  func(ctx context.Context, dto *bitcoin.StatusNodeDTO)
		expect func(t *testing.T, status *bitcoin.StatusNodeInfoDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.StatusNodeDTO) {
				btcRpcSvc.EXPECT().Status(ctx, dto.Network).Return(&status, nil)
			},
			expect: func(t *testing.T, status *bitcoin.StatusNodeInfoDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, status.Chain, network)
			},
		},
		{
			name: "should return failed check node status",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.StatusNodeDTO) {
				btcRpcSvc.EXPECT().Status(ctx, dto.Network).Return(nil, bitcoin.ErrFailedGetStatusNode)
			},
			expect: func(t *testing.T, status *bitcoin.StatusNodeInfoDTO, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, err, errors.WithMessage(bitcoin.ErrFailedGetStatusNode, bitcoin.ErrFailedGetStatusNode.Error()))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.ctx, tc.dto)
			w, err := service.StatusNode(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}

func TestService_CreateTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcRpcSvc := mock_bitcoin_rpc.NewMockService(controller)

	service, _ := bitcoin.NewService(btcRpcSvc, &logrus.Logger{})

	tx := "transaction"
	fee := 0.0000259

	dto := &bitcoin.CreateRawTransactionDTO{
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
		dto    *bitcoin.CreateRawTransactionDTO
		setup  func(ctx context.Context, dto *bitcoin.CreateRawTransactionDTO)
		expect func(t *testing.T, createdTx *bitcoin.CreatedRawTransactionDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.CreateRawTransactionDTO) {
				btcRpcSvc.EXPECT().CreateTransaction(ctx, bitcoin_rpc.UTXO(dto.Utxo), dto.FromAddress, dto.ToAddress, dto.Amount, dto.Network).Return(&tx, &fee, nil)
			},
			expect: func(t *testing.T, createdTx *bitcoin.CreatedRawTransactionDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, createdTx.Tx, tx)
				assert.Equal(t, createdTx.Fee, fee)
			},
		},
		{
			name: "should return error",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.CreateRawTransactionDTO) {
				btcRpcSvc.EXPECT().CreateTransaction(ctx, bitcoin_rpc.UTXO(dto.Utxo), dto.FromAddress, dto.ToAddress, dto.Amount, dto.Network).Return(nil, nil, bitcoin.ErrFailedCreateTx)
			},
			expect: func(t *testing.T, createdTx *bitcoin.CreatedRawTransactionDTO, err error) {
				assert.Nil(t, createdTx)
				assert.Equal(t, err, errors.WithMessage(bitcoin.ErrFailedCreateTx, bitcoin.ErrFailedCreateTx.Error()))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.ctx, tc.dto)
			w, err := service.CreateTransaction(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}

func TestService_DecodeTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcRpcSvc := mock_bitcoin_rpc.NewMockService(controller)

	service, _ := bitcoin.NewService(btcRpcSvc, &logrus.Logger{})

	dto := &bitcoin.DecodeRawTransactionDTO{
		Tx:      "transaction",
		Network: "test",
	}

	decodeTx := &bitcoin_rpc.DecodedTx{
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
		dto    *bitcoin.DecodeRawTransactionDTO
		setup  func(ctx context.Context, dto *bitcoin.DecodeRawTransactionDTO)
		expect func(t *testing.T, decodeTx *bitcoin.DecodedRawTransactionDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.DecodeRawTransactionDTO) {
				btcRpcSvc.EXPECT().DecodeTransaction(ctx, dto.Tx, dto.Network).Return(decodeTx, nil)
			},
			expect: func(t *testing.T, decodeTx *bitcoin.DecodedRawTransactionDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, decodeTx.Txid, "tx")
			},
		},
		{
			name: "should return error",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.DecodeRawTransactionDTO) {
				btcRpcSvc.EXPECT().DecodeTransaction(ctx, dto.Tx, dto.Network).Return(nil, bitcoin.ErrFailedDecodeTx)
			},
			expect: func(t *testing.T, decodeTx *bitcoin.DecodedRawTransactionDTO, err error) {
				assert.Nil(t, decodeTx)
				assert.Equal(t, err, errors.WithMessage(bitcoin.ErrFailedDecodeTx, bitcoin.ErrFailedDecodeTx.Error()))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.ctx, tc.dto)
			w, err := service.DecodeTransaction(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}

func TestService_FoundForRawTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcRpcSvc := mock_bitcoin_rpc.NewMockService(controller)

	service, _ := bitcoin.NewService(btcRpcSvc, &logrus.Logger{})

	dto := &bitcoin.FundForRawTransactionDTO{
		CreatedTxHex:  "tx",
		ChangeAddress: "address",
		Network:       "test",
	}

	tx := "transaction"
	fee := 0.0000259

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *bitcoin.FundForRawTransactionDTO
		setup  func(ctx context.Context, dto *bitcoin.FundForRawTransactionDTO)
		expect func(t *testing.T, fundedTx *bitcoin.FundedRawTransactionDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.FundForRawTransactionDTO) {
				btcRpcSvc.EXPECT().FundForTransaction(ctx, dto.CreatedTxHex, dto.ChangeAddress, dto.Network).Return(tx, &fee, nil)
			},
			expect: func(t *testing.T, fundedTx *bitcoin.FundedRawTransactionDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, fundedTx.Tx, tx)
				assert.Equal(t, fundedTx.Fee, fee)
			},
		},
		{
			name: "should return error",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.FundForRawTransactionDTO) {
				btcRpcSvc.EXPECT().FundForTransaction(ctx, dto.CreatedTxHex, dto.ChangeAddress, dto.Network).Return("", nil, bitcoin.ErrFailedFundForTx)
			},
			expect: func(t *testing.T, fundedTx *bitcoin.FundedRawTransactionDTO, err error) {
				assert.Nil(t, fundedTx)
				assert.Equal(t, err, errors.WithMessage(bitcoin.ErrFailedFundForTx, bitcoin.ErrFailedFundForTx.Error()))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.ctx, tc.dto)
			w, err := service.FoundForRawTransaction(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}

func TestService_SignTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcRpcSvc := mock_bitcoin_rpc.NewMockService(controller)

	service, _ := bitcoin.NewService(btcRpcSvc, &logrus.Logger{})

	dto := &bitcoin.SignRawTransactionDTO{
		Tx:         "tx",
		PrivateKey: "private",
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
		Network: "test",
	}

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *bitcoin.SignRawTransactionDTO
		setup  func(ctx context.Context, dto *bitcoin.SignRawTransactionDTO)
		expect func(t *testing.T, signedTx *bitcoin.SignedRawTransactionDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.SignRawTransactionDTO) {
				btcRpcSvc.EXPECT().SignTransaction(ctx, dto.Tx, dto.PrivateKey, bitcoin_rpc.UTXO(dto.Utxo), dto.Network).Return("hash", nil)
			},
			expect: func(t *testing.T, signedTx *bitcoin.SignedRawTransactionDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, signedTx.Hash, "hash")
			},
		},
		{
			name: "should return error",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.SignRawTransactionDTO) {
				btcRpcSvc.EXPECT().SignTransaction(ctx, dto.Tx, dto.PrivateKey, bitcoin_rpc.UTXO(dto.Utxo), dto.Network).Return("", bitcoin.ErrFailedSignTx)
			},
			expect: func(t *testing.T, signedTx *bitcoin.SignedRawTransactionDTO, err error) {
				assert.Nil(t, signedTx)
				assert.Equal(t, err, errors.WithMessage(bitcoin.ErrFailedSignTx, bitcoin.ErrFailedSignTx.Error()))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.ctx, tc.dto)
			w, err := service.SignTransaction(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}

func TestService_SendTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcRpcSvc := mock_bitcoin_rpc.NewMockService(controller)

	service, _ := bitcoin.NewService(btcRpcSvc, &logrus.Logger{})

	dto := &bitcoin.SendRawTransactionDTO{
		SignedTx: "hash",
		Network:  "test",
	}

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *bitcoin.SendRawTransactionDTO
		setup  func(ctx context.Context, dto *bitcoin.SendRawTransactionDTO)
		expect func(t *testing.T, sentTx *bitcoin.SentRawTransactionDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.SendRawTransactionDTO) {
				btcRpcSvc.EXPECT().SendTransaction(ctx, dto.SignedTx, dto.Network).Return("tx_id", nil)
			},
			expect: func(t *testing.T, sentTx *bitcoin.SentRawTransactionDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, sentTx.TxId, "tx_id")
			},
		},
		{
			name: "should return error",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.SendRawTransactionDTO) {
				btcRpcSvc.EXPECT().SendTransaction(ctx, dto.SignedTx, dto.Network).Return("", bitcoin.ErrFailedSendTx)
			},
			expect: func(t *testing.T, sentTx *bitcoin.SentRawTransactionDTO, err error) {
				assert.Nil(t, sentTx)
				assert.Equal(t, err, errors.WithMessage(bitcoin.ErrFailedSendTx, bitcoin.ErrFailedSendTx.Error()))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.ctx, tc.dto)
			w, err := service.SendTransaction(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}

func TestService_WalletInfo(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcRpcSvc := mock_bitcoin_rpc.NewMockService(controller)

	service, _ := bitcoin.NewService(btcRpcSvc, &logrus.Logger{})

	dto := &bitcoin.WalletDTO{
		WalletId: "wallet_id",
		Network:  "test",
	}

	info := &bitcoin_rpc.Info{
		Walletname:            "wallet_id",
		Walletversion:         0,
		Format:                "format",
		Balance:               0,
		UnconfirmedBalance:    0,
		ImmatureBalance:       0,
		Txcount:               0,
		Keypoololdest:         0,
		Keypoolsize:           0,
		Hdseedid:              "seed",
		KeypoolsizeHdInternal: 0,
		Paytxfee:              0,
		PrivateKeysEnabled:    false,
		AvoidReuse:            false,
		Scanning:              false,
		Descriptors:           false,
	}

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *bitcoin.WalletDTO
		setup  func(ctx context.Context, dto *bitcoin.WalletDTO)
		expect func(t *testing.T, walletInfo *bitcoin.WalletInfoDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.WalletDTO) {
				btcRpcSvc.EXPECT().WalletInfo(ctx, dto.WalletId, dto.Network).Return(info, nil)
			},
			expect: func(t *testing.T, walletInfo *bitcoin.WalletInfoDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, walletInfo.Walletname, dto.WalletId)
			},
		},
		{
			name: "should return error",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.WalletDTO) {
				btcRpcSvc.EXPECT().WalletInfo(ctx, dto.WalletId, dto.Network).Return(nil, bitcoin.ErrFailedGetWalletInfo)
			},
			expect: func(t *testing.T, walletInfo *bitcoin.WalletInfoDTO, err error) {
				assert.Nil(t, walletInfo)
				assert.Equal(t, err, errors.WithMessage(bitcoin.ErrFailedGetWalletInfo, bitcoin.ErrFailedGetWalletInfo.Error()))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.ctx, tc.dto)
			w, err := service.WalletInfo(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}

func TestService_CreateWallet(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcRpcSvc := mock_bitcoin_rpc.NewMockService(controller)

	service, _ := bitcoin.NewService(btcRpcSvc, &logrus.Logger{})

	dto := &bitcoin.CreateWalletDTO{
		Network: "test",
	}

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *bitcoin.CreateWalletDTO
		setup  func(ctx context.Context, dto *bitcoin.CreateWalletDTO)
		expect func(t *testing.T, createdWallet *bitcoin.CreatedWalletInfoDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.CreateWalletDTO) {
				btcRpcSvc.EXPECT().CreateWallet(ctx, dto.Network).Return("wallet_id", nil)
			},
			expect: func(t *testing.T, createdWallet *bitcoin.CreatedWalletInfoDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, createdWallet.WalletId, "wallet_id")
			},
		},
		{
			name: "should return error",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.CreateWalletDTO) {
				btcRpcSvc.EXPECT().CreateWallet(ctx, dto.Network).Return("", bitcoin.ErrFailedCreateWallet)
			},
			expect: func(t *testing.T, createdWallet *bitcoin.CreatedWalletInfoDTO, err error) {
				assert.Nil(t, createdWallet)
				assert.Equal(t, err, errors.WithMessage(bitcoin.ErrFailedCreateWallet, bitcoin.ErrFailedCreateWallet.Error()))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.ctx, tc.dto)
			w, err := service.CreateWallet(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}

func TestService_LoadWaller(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcRpcSvc := mock_bitcoin_rpc.NewMockService(controller)

	service, _ := bitcoin.NewService(btcRpcSvc, &logrus.Logger{})

	dto := &bitcoin.LoadWalletDTO{
		WalletId: "wallet_id",
		Network:  "test",
	}

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *bitcoin.LoadWalletDTO
		setup  func(ctx context.Context, dto *bitcoin.LoadWalletDTO)
		expect func(t *testing.T, loadedWallet *bitcoin.LoadWalletInfoDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.LoadWalletDTO) {
				btcRpcSvc.EXPECT().LoadWallet(ctx, dto.WalletId, dto.Network).Return(nil)
			},
			expect: func(t *testing.T, loadedWallet *bitcoin.LoadWalletInfoDTO, err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "should return error",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.LoadWalletDTO) {
				btcRpcSvc.EXPECT().LoadWallet(ctx, dto.WalletId, dto.Network).Return(bitcoin.ErrFailedLoadWallet)
			},
			expect: func(t *testing.T, loadedWallet *bitcoin.LoadWalletInfoDTO, err error) {
				assert.Nil(t, loadedWallet)
				assert.Equal(t, err, errors.WithMessage(bitcoin.ErrFailedLoadWallet, bitcoin.ErrFailedLoadWallet.Error()))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.ctx, tc.dto)
			w, err := service.LoadWaller(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}

func TestService_ImportAddress(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcRpcSvc := mock_bitcoin_rpc.NewMockService(controller)

	service, _ := bitcoin.NewService(btcRpcSvc, &logrus.Logger{})

	dto := &bitcoin.ImportAddressDTO{
		Address:  "address",
		WalletId: "wallet_id",
		Network:  "test",
	}

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *bitcoin.ImportAddressDTO
		setup  func(ctx context.Context, dto *bitcoin.ImportAddressDTO)
		expect func(t *testing.T, importedAddress *bitcoin.ImportAddressInfoDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.ImportAddressDTO) {
				btcRpcSvc.EXPECT().ImportAddress(ctx, dto.Address, dto.WalletId, dto.Network).Return(nil)
			},
			expect: func(t *testing.T, importedAddress *bitcoin.ImportAddressInfoDTO, err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "should return error",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.ImportAddressDTO) {
				btcRpcSvc.EXPECT().ImportAddress(ctx, dto.Address, dto.WalletId, dto.Network).Return(bitcoin.ErrFailedImportAddress)
			},
			expect: func(t *testing.T, importedAddress *bitcoin.ImportAddressInfoDTO, err error) {
				assert.Nil(t, importedAddress)
				assert.Equal(t, err, errors.WithMessage(bitcoin.ErrFailedImportAddress, bitcoin.ErrFailedImportAddress.Error()))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.ctx, tc.dto)
			w, err := service.ImportAddress(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}

func TestService_RescanWallet(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcRpcSvc := mock_bitcoin_rpc.NewMockService(controller)

	service, _ := bitcoin.NewService(btcRpcSvc, &logrus.Logger{})

	dto := &bitcoin.RescanWalletDTO{
		WalletId: "wallet_id",
		Network:  "test",
	}

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *bitcoin.RescanWalletDTO
		setup  func(ctx context.Context, dto *bitcoin.RescanWalletDTO)
		expect func(t *testing.T, rescanInfo *bitcoin.RescanWalletInfoDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.RescanWalletDTO) {
				btcRpcSvc.EXPECT().RescanWallet(ctx, dto.WalletId, dto.Network).Return(nil)
			},
			expect: func(t *testing.T, rescanInfo *bitcoin.RescanWalletInfoDTO, err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "should return error",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.RescanWalletDTO) {
				btcRpcSvc.EXPECT().RescanWallet(ctx, dto.WalletId, dto.Network).Return(bitcoin.ErrFailedRescanWallet)
			},
			expect: func(t *testing.T, rescanInfo *bitcoin.RescanWalletInfoDTO, err error) {
				assert.Nil(t, rescanInfo)
				assert.Equal(t, err, errors.WithMessage(bitcoin.ErrFailedRescanWallet, bitcoin.ErrFailedRescanWallet.Error()))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.ctx, tc.dto)
			w, err := service.RescanWallet(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}

func TestService_ListUnspent(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	btcRpcSvc := mock_bitcoin_rpc.NewMockService(controller)

	service, _ := bitcoin.NewService(btcRpcSvc, &logrus.Logger{})

	dto := &bitcoin.ListUnspentDTO{
		Address:  "address",
		WalletId: "wallet_id",
		Network:  "test",
	}

	listUTXO := []*bitcoin_rpc.Unspent{
		{
			Txid:          "tx_id",
			Vout:          1,
			Address:       "address",
			Label:         "label",
			ScriptPubKey:  "pkScript",
			Amount:        0.001,
			Confirmations: 4,
			Spendable:     false,
			Solvable:      false,
			Safe:          true,
		},
	}

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *bitcoin.ListUnspentDTO
		setup  func(ctx context.Context, dto *bitcoin.ListUnspentDTO)
		expect func(t *testing.T, listUTXO *bitcoin.ListUnspentInfoDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.ListUnspentDTO) {
				btcRpcSvc.EXPECT().ListUnspent(ctx, dto.Address, dto.WalletId, dto.Network).Return(listUTXO, nil)
			},
			expect: func(t *testing.T, utxoInfo *bitcoin.ListUnspentInfoDTO, err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "should return error",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *bitcoin.ListUnspentDTO) {
				btcRpcSvc.EXPECT().ListUnspent(ctx, dto.Address, dto.WalletId, dto.Network).Return(nil, bitcoin.ErrFailedGetUnspent)
			},
			expect: func(t *testing.T, utxoInfo *bitcoin.ListUnspentInfoDTO, err error) {
				assert.Nil(t, utxoInfo)
				assert.Equal(t, err, errors.WithMessage(bitcoin.ErrFailedGetUnspent, bitcoin.ErrFailedGetUnspent.Error()))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.ctx, tc.dto)
			w, err := service.ListUnspent(tc.ctx, tc.dto)
			tc.expect(t, w, err)
		})
	}
}
