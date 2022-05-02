package wallet

import (
	"context"
	"go.uber.org/zap"
	"nn-blockchain-api/pkg/errors"
	pb "nn-blockchain-api/pkg/grpc_client/proto/wallet"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go
type Service interface {
	CreateWallet(ctx context.Context, walletName string, mnemonic *string) (*DTO, error)
	CreateMnemonic(ctx context.Context, length, language string) (*CreatedMnemonicDTO, error)
}

type service struct {
	walletClient pb.WalletServiceClient
	logger       *zap.SugaredLogger
}

func NewService(walletClient pb.WalletServiceClient, logger *zap.SugaredLogger) (Service, error) {
	if walletClient == nil {
		return nil, errors.NewInternal("invalid wallet client")
	}
	if logger == nil {
		return nil, errors.NewInternal("invalid logger")
	}
	return &service{walletClient: walletClient, logger: logger}, nil
}

func (s *service) CreateWallet(ctx context.Context, walletName string, mnemonic *string) (*DTO, error) {
	response, err := s.walletClient.CreateWallet(ctx, &pb.CreateWalletData{WalletName: walletName, Mnemonic: mnemonic})
	if err != nil {
		s.logger.Errorf("failed to create wallet: %v", err)
		return nil, errors.WithMessage(ErrInvalidWalletType, err.Error())
	}

	return &DTO{
		Mnemonic: response.Wallet.Mnemonic,
		CoinName: response.Wallet.CoinName,
		Address:  response.Wallet.Address,
		Private:  response.Wallet.Private,
	}, nil
}

func (s *service) CreateMnemonic(ctx context.Context, length, language string) (*CreatedMnemonicDTO, error) {
	response, err := s.walletClient.CreateMnemonic(ctx, &pb.CreateMnemonicData{MnemonicLength: length, Language: language})
	if err != nil {
		s.logger.Errorf("failed to create mnemonic: %v", err)
		return nil, errors.WithMessage(ErrCreateMnemonic, err.Error())
	}

	return &CreatedMnemonicDTO{Mnemonic: response.Mnemonic}, nil
}
