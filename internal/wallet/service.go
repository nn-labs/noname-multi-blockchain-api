package wallet

import (
	"context"
	"github.com/sirupsen/logrus"
	"nn-blockchain-api/pkg/errors"
	"nn-blockchain-api/pkg/grpc_client/proto/wallet"
)

type Service interface {
	CreateWallet(ctx context.Context, walletName string) (*Wallet, error)
}

type service struct {
	walletClient wallet.WalletServiceClient
	log          *logrus.Logger
}

func NewService(walletClient wallet.WalletServiceClient, log *logrus.Logger) (Service, error) {
	if walletClient == nil {
		return nil, errors.NewInternal("invalid wallet client")
	}
	if log == nil {
		return nil, errors.NewInternal("invalid logger")
	}
	return &service{walletClient: walletClient, log: log}, nil
}

func (svc *service) CreateWallet(ctx context.Context, walletName string) (*Wallet, error) {
	response, err := svc.walletClient.CreateWallet(ctx, &wallet.CoinName{Name: walletName})
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
