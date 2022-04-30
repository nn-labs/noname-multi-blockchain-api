package ethereum

import (
	"nn-blockchain-api/pkg/codes"
	"nn-blockchain-api/pkg/errors"
)

const (
	StatusInvalidRequest      errors.Status = "invalid_request"
	StatusFailedGetStatusNode errors.Status = "failed_get_status_node"
	StatusFailedCreateTx      errors.Status = "failed_create_tx"
	StatusFailedSignTx        errors.Status = "failed_sign_tx"
	StatusFailedSendTx        errors.Status = "failed_send_tx"
)

var (
	ErrInvalidRequest      = errors.New(codes.BadRequest, StatusInvalidRequest)
	ErrFailedGetStatusNode = errors.New(codes.BadRequest, StatusFailedGetStatusNode)
	ErrFailedCreateTx      = errors.New(codes.BadRequest, StatusFailedCreateTx)
	ErrFailedSignTx        = errors.New(codes.BadRequest, StatusFailedSignTx)
	ErrFailedSendTx        = errors.New(codes.BadRequest, StatusFailedSendTx)
)
