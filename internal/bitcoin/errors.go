package bitcoin

import (
	"nn-blockchain-api/pkg/codes"
	"nn-blockchain-api/pkg/errors"
)

const (
	StatusInvalidRequest errors.Status = "invalid_request"
)

var (
	ErrInvalidRequest = errors.New(codes.BadRequest, StatusInvalidRequest)
)
