package ethereum

import (
	"context"
	gErrors "errors"
	"go.uber.org/zap"
	"nn-blockchain-api/pkg/errors"
	ethereum_rpc "nn-blockchain-api/pkg/rpc/ethereum"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go

type Service interface {
	StatusNode(ctx context.Context, dto *StatusNodeDTO) (*NodeInfoDTO, error)

	CreateTransaction(ctx context.Context, dto *CreateRawTransactionDTO) (*CreatedRawTransactionDTO, error)
	SignTransaction(ctx context.Context, dto *SignRawTransactionDTO) (*SignedRawTransactionDTO, error)
	SendTransaction(ctx context.Context, dto *SendRawTransactionDTO) (*SentRawTransactionDTO, error)
}

type service struct {
	ethRpcSvc ethereum_rpc.Service
	logger    *zap.SugaredLogger
}

func NewService(ethRpcSvc ethereum_rpc.Service, logger *zap.SugaredLogger) (Service, error) {
	if ethRpcSvc == nil {
		return nil, gErrors.New("invalid ethereum rpc service")
	}
	if logger == nil {
		return nil, gErrors.New("invalid logger")
	}
	return &service{ethRpcSvc: ethRpcSvc, logger: logger}, nil
}

func (s *service) StatusNode(ctx context.Context, dto *StatusNodeDTO) (*NodeInfoDTO, error) {
	status, err := s.ethRpcSvc.Status(ctx, dto.Network)
	if err != nil {
		s.logger.Errorf("failed check node status: %v", err)
		return nil, errors.WithMessage(ErrFailedGetStatusNode, err.Error())
	}

	return &NodeInfoDTO{
		CurrentBlock:        status.CurrentBlock,
		HealedBytecodeBytes: status.HealedBytecodeBytes,
		HealedBytecodes:     status.HealedBytecodes,
		HealedTrienodeBytes: status.HealedTrienodeBytes,
		HealedTrienodes:     status.HealedTrienodes,
		HealingBytecode:     status.HealingBytecode,
		HealingTrienodes:    status.HealingTrienodes,
		HighestBlock:        status.HighestBlock,
		StartingBlock:       status.StartingBlock,
		SyncedAccountBytes:  status.SyncedAccountBytes,
		SyncedAccounts:      status.SyncedAccounts,
		SyncedBytecodeBytes: status.SyncedBytecodeBytes,
		SyncedBytecodes:     status.SyncedBytecodes,
		SyncedStorage:       status.SyncedStorage,
		SyncedStorageBytes:  status.SyncedStorageBytes,
		SyncMessage:         status.SyncMessage,
	}, nil
}

func (s *service) CreateTransaction(ctx context.Context, dto *CreateRawTransactionDTO) (*CreatedRawTransactionDTO, error) {
	tx, fee, err := s.ethRpcSvc.CreateTransaction(ctx, dto.FromAddress, dto.ToAddress, dto.Amount, dto.Network)
	if err != nil {
		s.logger.Errorf("failed create transaction: %v", err)
		return nil, errors.WithMessage(ErrFailedCreateTx, err.Error())
		//return nil, ErrFailedCreateTx
	}

	return &CreatedRawTransactionDTO{
		Tx:  *tx,
		Fee: *fee,
	}, nil
}

func (s *service) SignTransaction(ctx context.Context, dto *SignRawTransactionDTO) (*SignedRawTransactionDTO, error) {
	signedTx, err := s.ethRpcSvc.SignTransaction(ctx, dto.Tx, dto.PrivateKey, dto.Network)
	if err != nil {
		s.logger.Errorf("failed sign transaction: %v", err)
		return nil, errors.WithMessage(ErrFailedSignTx, err.Error())
		//return nil, ErrFailedSignTx
	}

	return &SignedRawTransactionDTO{
		SignedTx: *signedTx,
	}, nil
}

func (s *service) SendTransaction(ctx context.Context, dto *SendRawTransactionDTO) (*SentRawTransactionDTO, error) {
	txId, err := s.ethRpcSvc.SendTransaction(ctx, dto.SignedTx, dto.Network)
	if err != nil {
		s.logger.Errorf("failed send transaction: %v", err)
		return nil, errors.WithMessage(ErrFailedSendTx, err.Error())
		//return nil, ErrFailedSendTx
	}

	return &SentRawTransactionDTO{
		TxId: *txId,
	}, nil
}
