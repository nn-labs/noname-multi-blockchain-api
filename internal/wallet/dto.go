package wallet

import (
	"github.com/go-playground/validator/v10"
	"nn-blockchain-api/pkg/errors"
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

type CoinNameDto struct {
	Name string `json:"name" validate:"required"`
}

type MnemonicDTO struct {
	Length   string `json:"length" validate:"required"`
	Language string `json:"language" validate:"required"`
}
