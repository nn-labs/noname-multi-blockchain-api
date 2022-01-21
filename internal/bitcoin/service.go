package bitcoin

import (
	"context"
	"github.com/sirupsen/logrus"
	"nn-blockchain-api/pkg/errors"
	"nn-blockchain-api/pkg/rpc/bitcoin"
)

type Service interface {
	StatusNode(ctx context.Context) (*StatusNodeDTO, error)
	CreateTransaction(ctx context.Context, dto *RawTransactionDTO) (*CreatedRawTransactionDTO, error)
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

func (svc *service) StatusNode(ctx context.Context) (*StatusNodeDTO, error) {
	status, err := bitcoin.Status()
	if err != nil {
		svc.log.WithContext(ctx).Errorf("failed check node status")
		return nil, errors.NewInternal("failed check node status")
	}

	return &StatusNodeDTO{
		Chain:                status.Chain,
		Blocks:               status.Blocks,
		Headers:              status.Headers,
		Verificationprogress: status.Verificationprogress,
		Softforks:            status.Softforks,
		Warnings:             status.Warnings,
	}, nil
}

func (svc *service) CreateTransaction(ctx context.Context, dto *RawTransactionDTO) (*CreatedRawTransactionDTO, error) {
	tx, err := bitcoin.CreateTransaction(dto.Utxo, dto.ToAddress, dto.Amount)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.NewInternal(err.Error())
	}

	return &CreatedRawTransactionDTO{
		Tx: tx,
	}, nil
}
