package bitcoin

import (
	"context"
	"github.com/sirupsen/logrus"
	"nn-blockchain-api/pkg/errors"
	"nn-blockchain-api/pkg/rpc/bitcoin"
)

type Service interface {
	StatusNode(ctx context.Context) (*bitcoin.StatusNodeDTO, error)
	CreateTransaction(ctx context.Context, dto *RawTransactionDTO) (string, error)
}

type service struct {
	log *logrus.Logger
}

func NewService(log *logrus.Logger) (Service, error) {
	if log == nil {
		return nil, errors.NewInternal("invalid logger")
	}
	return &service{log: log}, nil
}

func (svc *service) StatusNode(ctx context.Context) (*bitcoin.StatusNodeDTO, error) {
	sn, err := bitcoin.StatusNode()
	if err != nil {
		svc.log.WithContext(ctx).Errorf("failed check node status")
		return nil, errors.NewInternal("failed check node status")
	}

	return sn, nil
}

func (svc *service) CreateTransaction(ctx context.Context, dto *RawTransactionDTO) (string, error) {
	tx, err := bitcoin.CreateTransaction(dto.Utxo, dto.ToAddress, dto.Amount)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return "", errors.NewInternal(err.Error())
	}

	return tx, nil
}
