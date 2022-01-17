package bitcoin

import (
	"github.com/go-playground/validator/v10"
	"math/big"
	"nn-blockchain-api/pkg/errors"
	"nn-blockchain-api/pkg/rpc/bitcoin"
)

func Validate(dto interface{}) error {
	validate := validator.New()

	if err := validate.Struct(dto); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errors.WithMessage(ErrInvalidRequest, err.Error())
		}

		validationErr := ErrInvalidRequest
		for _, err := range err.(validator.ValidationErrors) {
			validationErr = errors.WithMessage(validationErr, err.Error())
		}
		return validationErr
	}
	return nil
}

type RawTransactionDTO struct {
	Utxo        []*bitcoin.UTXO `json:"utxo" validate:"required"`
	FromAddress string          `json:"from_address" validate:"required"`
	ToAddress   string          `json:"to_address" validate:"required"`
	Amount      *big.Int        `json:"amount" validate:"required"`
}
