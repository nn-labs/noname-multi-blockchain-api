package wallet

import (
	"nn-blockchain-api/pkg/codes"
	"nn-blockchain-api/pkg/errors"
)

const (
	StatusInvalidRequest    errors.Status = "invalid_request"
	StatusInvalidWalletType errors.Status = "invalid_wallet_type"
)

var (
	ErrInvalidRequest    = errors.New(codes.BadRequest, StatusInvalidRequest)
	ErrInvalidWalletType = errors.New(codes.NotFound, StatusInvalidWalletType)
)
