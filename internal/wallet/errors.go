package wallet

import (
	"nn-blockchain-api/pkg/codes"
	"nn-blockchain-api/pkg/errors"
)

const (
	StatusInvalidRequest    errors.Status = "invalid_request"
	StatusInvalidPayload    errors.Status = "invalid_payload"
	StatusInvalidWalletType errors.Status = "invalid_wallet_type"
	StatusInternalError     errors.Status = "internal_error"
)

var (
	ErrInvalidRequest    = errors.New(codes.BadRequest, StatusInvalidRequest)
	ErrInvalidWalletType = errors.New(codes.NotFound, StatusInvalidWalletType)
	ErrInternal          = errors.New(codes.InternalError, StatusInternalError)
	ErrInvalidPayload    = errors.New(codes.InternalError, StatusInvalidPayload)
)
