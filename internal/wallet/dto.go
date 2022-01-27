package wallet

import (
	"github.com/go-playground/validator/v10"
	"nn-blockchain-api/pkg/errors"
	"strings"
)

func Validate(dto interface{}) error {
	validate := validator.New()

	_ = validate.RegisterValidation("mnemonic", func(fl validator.FieldLevel) bool {
		mnemonic := strings.Split(fl.Field().String(), " ")

		if len(mnemonic) == 12 || len(mnemonic) == 24 {
			return true
		}

		return false
	})

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

type CoinNameDTO struct {
	Name     string `json:"name" validate:"required"`
	Mnemonic string `json:"mnemonic" validate:"mnemonic"`
}

type MnemonicDTO struct {
	Length   string `json:"length" validate:"required"`
	Language string `json:"language" validate:"required"`
}

type DTO struct {
	Mnemonic string
	CoinName string
	Address  string
	Private  string
}

type CreatedMnemonicDTO struct {
	Mnemonic string
}
