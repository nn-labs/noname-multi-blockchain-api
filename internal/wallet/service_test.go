package wallet

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"nn-blockchain-api/pkg/errors"
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
		logger       *zap.SugaredLogger
		expect       func(*testing.T, Service, error)
	}{
		{
			name:         "should return wallet service",
			walletClient: grpc_mock.NewMockWalletServiceClient(controller),
			logger:       &zap.SugaredLogger{},
			expect: func(t *testing.T, s Service, err error) {
				assert.NotNil(t, s)
				assert.Nil(t, err)
			},
		},
		{
			name:         "should return invalid wallet client",
			walletClient: nil,
			logger:       &zap.SugaredLogger{},
			expect: func(t *testing.T, s Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "code: 500; status: internal_error; message: invalid wallet client")
			},
		},
		{
			name:         "should return invalid logger",
			walletClient: grpc_mock.NewMockWalletServiceClient(controller),
			logger:       nil,
			expect: func(t *testing.T, s Service, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, s)
				assert.EqualError(t, err, "code: 500; status: internal_error; message: invalid logger")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := NewService(tc.walletClient, tc.logger)
			tc.expect(t, svc, err)
		})
	}
}

func TestService_CreateWallet(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockWalletClient := grpc_mock.NewMockWalletServiceClient(controller)

	service, _ := NewService(mockWalletClient, &zap.SugaredLogger{})

	dto := &CoinNameDTO{
		Name:     "BTC",
		Mnemonic: "mnemonic",
	}

	invalidWalletNameDto := &CoinNameDTO{
		Name:     "ASD",
		Mnemonic: "mnemonic",
	}

	walletData := &pb.Wallet{
		Mnemonic: "mnemonic",
		CoinName: "BTC",
		Address:  "address",
		Private:  "privateKey",
	}

	tests := []struct {
		name         string
		ctx          context.Context
		walletClient pb.WalletServiceClient
		dto          *CoinNameDTO
		setup        func(context.Context, *CoinNameDTO)
		expect       func(*testing.T, *DTO, error)
	}{
		{
			name:         "should return status ok",
			ctx:          context.Background(),
			walletClient: mockWalletClient,
			dto:          dto,
			setup: func(ctx context.Context, dto *CoinNameDTO) {
				mockWalletClient.EXPECT().CreateWallet(ctx, &pb.CreateWalletData{
					WalletName: dto.Name,
					Mnemonic:   &dto.Mnemonic,
				}).Return(&pb.WalletInfo{Wallet: walletData}, nil)
			},
			expect: func(t *testing.T, w *DTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, w.CoinName, dto.Name)
			},
		},
		{
			name:         "should return invalid wallet type",
			ctx:          context.Background(),
			walletClient: mockWalletClient,
			dto:          invalidWalletNameDto,
			setup: func(ctx context.Context, dto *CoinNameDTO) {
				mockWalletClient.EXPECT().CreateWallet(ctx, &pb.CreateWalletData{
					WalletName: dto.Name,
					Mnemonic:   &dto.Mnemonic,
				}).Return(nil, ErrInvalidWalletType)
			},
			expect: func(t *testing.T, w *DTO, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, errors.WithMessage(ErrInvalidWalletType, "code: 404; status: invalid_wallet_type"), err)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.ctx, tc.dto)
			w, err := service.CreateWallet(tc.ctx, tc.dto.Name, &tc.dto.Mnemonic)
			tc.expect(t, w, err)
		})
	}
}

func TestService_CreateMnemonic(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockWalletClient := grpc_mock.NewMockWalletServiceClient(controller)

	service, _ := NewService(mockWalletClient, &zap.SugaredLogger{})

	mnemonicReturns := "mnemonic"

	dto := &MnemonicDTO{
		Length:   "12",
		Language: "english",
	}

	invalidDto := &MnemonicDTO{
		Length:   "123123",
		Language: "english",
	}

	tests := []struct {
		name         string
		ctx          context.Context
		walletClient pb.WalletServiceClient
		dto          *MnemonicDTO
		setup        func(context.Context, *MnemonicDTO)
		expect       func(*testing.T, *CreatedMnemonicDTO, error)
	}{
		{
			name:         "should return status ok",
			ctx:          context.Background(),
			walletClient: mockWalletClient,
			dto:          dto,
			setup: func(ctx context.Context, dto *MnemonicDTO) {
				mockWalletClient.EXPECT().CreateMnemonic(ctx, &pb.CreateMnemonicData{
					MnemonicLength: dto.Length,
					Language:       dto.Language,
				}).Return(&pb.MnemonicInfo{Mnemonic: mnemonicReturns}, nil)
			},
			expect: func(t *testing.T, mnemonic *CreatedMnemonicDTO, err error) {
				assert.Nil(t, err)
				assert.Equal(t, mnemonic.Mnemonic, mnemonicReturns)
			},
		},
		{
			name:         "should return failed to create mnemonic",
			ctx:          context.Background(),
			walletClient: mockWalletClient,
			dto:          invalidDto,
			setup: func(ctx context.Context, dto *MnemonicDTO) {
				mockWalletClient.EXPECT().CreateMnemonic(ctx, &pb.CreateMnemonicData{
					MnemonicLength: dto.Length,
					Language:       dto.Language,
				}).Return(nil, ErrCreateMnemonic)
			},
			expect: func(t *testing.T, mnemonic *CreatedMnemonicDTO, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, errors.WithMessage(ErrCreateMnemonic, "code: 500; status: error_create_mnemonic"), err)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(tc.ctx, tc.dto)
			w, err := service.CreateMnemonic(tc.ctx, tc.dto.Length, tc.dto.Language)
			tc.expect(t, w, err)
		})
	}
}
