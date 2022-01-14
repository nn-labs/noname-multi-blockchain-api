package wallet

import (
	"context"
	"github.com/sirupsen/logrus"
	"nn-blockchain-api/pkg/errors"
	pb "nn-blockchain-api/pkg/grpc_client/proto/wallet"
)

type Service interface {
	CreateWallet(ctx context.Context, walletName string) (*Wallet, error)
	CreateMnemonic(ctx context.Context, length, language string) (*Mnemonic, error)
}

type service struct {
	walletClient pb.WalletServiceClient
	log          *logrus.Logger
}

func NewService(walletClient pb.WalletServiceClient, log *logrus.Logger) (Service, error) {
	if walletClient == nil {
		return nil, errors.NewInternal("invalid wallet client")
	}
	if log == nil {
		return nil, errors.NewInternal("invalid logger")
	}
	return &service{walletClient: walletClient, log: log}, nil
}

func (svc *service) CreateWallet(ctx context.Context, walletName string) (*Wallet, error) {
	response, err := svc.walletClient.CreateWallet(ctx, &pb.CreateWalletData{WalletName: walletName})
	if err != nil {
		svc.log.WithContext(ctx).Errorf("failed to create btc wallet: %v", err)
		return nil, errors.WithMessage(ErrInvalidWalletType, err.Error())
	}

	return &Wallet{
		Mnemonic: response.Wallet.Mnemonic,
		CoinName: response.Wallet.CoinName,
		Address:  response.Wallet.Address,
		Private:  response.Wallet.Private,
	}, nil
}

func (svc *service) CreateMnemonic(ctx context.Context, length, language string) (*Mnemonic, error) {
	response, err := svc.walletClient.CreateMnemonic(ctx, &pb.CreateMnemonicData{MnemonicLength: length, Language: language})
	if err != nil {
		svc.log.WithContext(ctx).Errorf("failed to create mnemonic: %v", err)
		return nil, errors.WithMessage(ErrInternal, err.Error())
	}

	return &Mnemonic{Mnemonic: response.Mnemonic}, nil
}
