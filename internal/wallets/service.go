package wallets

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "nn-blockchain-api/internal/wallets/proto"
	"nn-blockchain-api/pkg/errors"
)

type Service interface {
	CreateWallet(ctx context.Context, walletName string) (*Wallet, error)
}

type service struct {
	grpcHost string
	log      *logrus.Logger
}

func NewService(grpcHost string, log *logrus.Logger) (Service, error) {
	if grpcHost == "" {
		return nil, errors.NewInternal("invalid grpc host")
	}
	if log == nil {
		return nil, errors.NewInternal("invalid logger")
	}
	return &service{grpcHost: grpcHost, log: log}, nil
}

func (svc *service) CreateWallet(ctx context.Context, walletName string) (*Wallet, error) {
	conn, err := grpc.Dial(svc.grpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewWalletsServiceClient(conn)

	response, err := client.CreateWallet(ctx, &pb.CoinName{Name: walletName})
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
