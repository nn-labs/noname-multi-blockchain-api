package ethereum_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"nn-blockchain-api/internal/ethereum"
	"nn-blockchain-api/pkg/errors"
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
		logger    *zap.SugaredLogger
		expect    func(*testing.T, ethereum.Service, error)
	}{
		{
			name:      "should return ethereum service",
			ethRpcSvc: mock_ethereum_rpc.NewMockService(controller),
			logger:    &zap.SugaredLogger{},
			expect: func(t *testing.T, s ethereum.Service, err error) {
				assert.NotNil(t, s)
				assert.Nil(t, err)
			},
		},
		{
			name:      "should return invalid ethereum rpc service",
			ethRpcSvc: nil,
			logger:    &zap.SugaredLogger{},
			expect: func(t *testing.T, s ethereum.Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "invalid ethereum rpc service")
			},
		},
		{
			name:      "should return invalid logger",
			ethRpcSvc: mock_ethereum_rpc.NewMockService(controller),
			logger:    nil,
			expect: func(t *testing.T, s ethereum.Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "invalid logger")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := ethereum.NewService(tc.ethRpcSvc, tc.logger)
			tc.expect(t, svc, err)
		})
	}
}

func TestService_StatusNode(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	ethRpcSvc := mock_ethereum_rpc.NewMockService(controller)

	service, _ := ethereum.NewService(ethRpcSvc, &zap.SugaredLogger{})

	statusInfo := ethereum_rpc.StatusNodeResponse{
		CurrentBlock:        "0x321",
		HealedBytecodeBytes: "0x321",
		HealedBytecodes:     "0x321",
		HealedTrienodeBytes: "0x321",
		HealedTrienodes:     "0x321",
		HealingBytecode:     "0x321",
		HealingTrienodes:    "0x321",
		HighestBlock:        "0x321",
		StartingBlock:       "0x321",
		SyncedAccountBytes:  "0x321",
		SyncedAccounts:      "0x321",
		SyncedBytecodeBytes: "0x321",
		SyncedBytecodes:     "0x321",
		SyncedStorage:       "0x321",
		SyncedStorageBytes:  "0x321",
		SyncMessage:         "synced",
	}

	network := "test"
	dto := &ethereum.StatusNodeDTO{Network: network}

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *ethereum.StatusNodeDTO
		setup  func(ctx context.Context, dto *ethereum.StatusNodeDTO)
		expect func(t *testing.T, status *ethereum.NodeInfoDTO, err error)
	}{
		{
			name: "should return ok",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *ethereum.StatusNodeDTO) {
				ethRpcSvc.EXPECT().Status(ctx, dto.Network).Return(&statusInfo, nil)
			},
			expect: func(t *testing.T, status *ethereum.NodeInfoDTO, err error) {
				assert.Nil(t, err)
				//assert.Equal(t, s, network)
			},
		},
		{
			name: "should return failed check node status",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *ethereum.StatusNodeDTO) {
				ethRpcSvc.EXPECT().Status(ctx, dto.Network).Return(nil, ethereum.ErrFailedGetStatusNode)
			},
			expect: func(t *testing.T, status *ethereum.NodeInfoDTO, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, err, errors.WithMessage(ethereum.ErrFailedGetStatusNode, ethereum.ErrFailedGetStatusNode.Error()))
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

	ethRpcSvc := mock_ethereum_rpc.NewMockService(controller)

	service, _ := ethereum.NewService(ethRpcSvc, &zap.SugaredLogger{})

	tx := "transaction"
	fee := 0.000528288415914

	dto := &ethereum.CreateRawTransactionDTO{
		FromAddress: "from",
		ToAddress:   "to",
		Amount:      0.1,
		Network:     "test",
	}

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *ethereum.CreateRawTransactionDTO
		setup  func(ctx context.Context, dto *ethereum.CreateRawTransactionDTO)
		expect func(t *testing.T, createdTxDto *ethereum.CreatedRawTransactionDTO, err error)
	}{
		{
			name: "should return created transaction",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *ethereum.CreateRawTransactionDTO) {
				ethRpcSvc.EXPECT().CreateTransaction(ctx, dto.FromAddress, dto.ToAddress, dto.Amount, dto.Network).Return(&tx, &fee, nil)
			},
			expect: func(t *testing.T, createdTxDto *ethereum.CreatedRawTransactionDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, createdTxDto.Tx, tx)
				assert.Equal(t, createdTxDto.Fee, fee)
			},
		},
		{
			name: "should return failed to create transaction",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *ethereum.CreateRawTransactionDTO) {
				ethRpcSvc.EXPECT().CreateTransaction(ctx, dto.FromAddress, dto.ToAddress, dto.Amount, dto.Network).Return(nil, nil, ethereum.ErrFailedCreateTx)
			},
			expect: func(t *testing.T, createdTxDto *ethereum.CreatedRawTransactionDTO, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, err, errors.WithMessage(ethereum.ErrFailedCreateTx, ethereum.ErrFailedCreateTx.Error()))
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

func TestService_SignTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	ethRpcSvc := mock_ethereum_rpc.NewMockService(controller)

	service, _ := ethereum.NewService(ethRpcSvc, &zap.SugaredLogger{})

	signedTx := "signed_transaction"

	dto := &ethereum.SignRawTransactionDTO{
		Tx:         "transaction",
		PrivateKey: "private_key",
		Network:    "test",
	}

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *ethereum.SignRawTransactionDTO
		setup  func(ctx context.Context, dto *ethereum.SignRawTransactionDTO)
		expect func(t *testing.T, signedTxDto *ethereum.SignedRawTransactionDTO, err error)
	}{
		{
			name: "should return signed transaction",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *ethereum.SignRawTransactionDTO) {
				ethRpcSvc.EXPECT().SignTransaction(ctx, dto.Tx, dto.PrivateKey, dto.Network).Return(&signedTx, nil)
			},
			expect: func(t *testing.T, signedTxDto *ethereum.SignedRawTransactionDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, signedTxDto.SignedTx, signedTx)
			},
		},
		{
			name: "should return failed to signed transaction",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *ethereum.SignRawTransactionDTO) {
				ethRpcSvc.EXPECT().SignTransaction(ctx, dto.Tx, dto.PrivateKey, dto.Network).Return(nil, ethereum.ErrFailedSignTx)
			},
			expect: func(t *testing.T, signedTxDto *ethereum.SignedRawTransactionDTO, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, err, errors.WithMessage(ethereum.ErrFailedSignTx, ethereum.ErrFailedSignTx.Error()))
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

	ethRpcSvc := mock_ethereum_rpc.NewMockService(controller)

	service, _ := ethereum.NewService(ethRpcSvc, &zap.SugaredLogger{})

	dto := &ethereum.SendRawTransactionDTO{
		SignedTx: "signed_tx",
		Network:  "test",
	}

	txId := "some_tx_id"

	tests := []struct {
		name   string
		ctx    context.Context
		dto    *ethereum.SendRawTransactionDTO
		setup  func(ctx context.Context, dto *ethereum.SendRawTransactionDTO)
		expect func(t *testing.T, sentTxId *ethereum.SentRawTransactionDTO, err error)
	}{
		{
			name: "should return sent transaction id",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *ethereum.SendRawTransactionDTO) {
				ethRpcSvc.EXPECT().SendTransaction(ctx, dto.SignedTx, dto.Network).Return(&txId, nil)
			},
			expect: func(t *testing.T, sentTxDto *ethereum.SentRawTransactionDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, sentTxDto.TxId, txId)
			},
		},
		{
			name: "should return failed to send transaction",
			ctx:  context.Background(),
			dto:  dto,
			setup: func(ctx context.Context, dto *ethereum.SendRawTransactionDTO) {
				ethRpcSvc.EXPECT().SendTransaction(ctx, dto.SignedTx, dto.Network).Return(nil, ethereum.ErrFailedSendTx)
			},
			expect: func(t *testing.T, sentTxDto *ethereum.SentRawTransactionDTO, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, err, errors.WithMessage(ethereum.ErrFailedSendTx, ethereum.ErrFailedSendTx.Error()))
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
