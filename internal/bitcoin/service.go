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
	log       *logrus.Logger
	btcClient bitcoin.IBtcClient
}

func NewService(log *logrus.Logger, btcClient bitcoin.IBtcClient) (Service, error) {
	if log == nil {
		return nil, errors.NewInternal("invalid logger")
	}
	if btcClient == nil {
		return nil, errors.NewInternal("invalid btc client")
	}
	return &service{log: log, btcClient: btcClient}, nil
}

func (svc *service) StatusNode(ctx context.Context) (*StatusNodeDTO, error) {
	status, err := bitcoin.Status(svc.btcClient)
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
	tx, err := bitcoin.CreateTransaction(svc.btcClient, dto.Utxo, dto.ToAddress, dto.Amount)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.NewInternal(err.Error())
	}

	return &CreatedRawTransactionDTO{
		Tx: tx,
	}, nil
}
