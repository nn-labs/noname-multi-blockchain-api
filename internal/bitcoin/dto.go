package bitcoin

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

type StatusNodeDTO struct {
	Chain                string      `json:"chain"`
	Blocks               interface{} `json:"blocks"`
	Headers              interface{} `json:"headers"`
	Verificationprogress interface{} `json:"verificationprogress"`
	Softforks            struct {
		Bip34 struct {
			Type   string      `json:"type"`
			Active bool        `json:"active"`
			Height interface{} `json:"height"`
		} `json:"bip34"`
		Bip66 struct {
			Type   string      `json:"type"`
			Active bool        `json:"active"`
			Height interface{} `json:"height"`
		} `json:"bip66"`
		Bip65 struct {
			Type   string      `json:"type"`
			Active bool        `json:"active"`
			Height interface{} `json:"height"`
		} `json:"bip65"`
		Csv struct {
			Type   string      `json:"type"`
			Active bool        `json:"active"`
			Height interface{} `json:"height"`
		} `json:"csv"`
		Segwit struct {
			Type   string      `json:"type"`
			Active bool        `json:"active"`
			Height interface{} `json:"height"`
		} `json:"segwit"`
		Taproot struct {
			Type string `json:"type"`
			Bip9 struct {
				Status              string      `json:"status"`
				StartTime           interface{} `json:"start_time"`
				Timeout             interface{} `json:"timeout"`
				Since               interface{} `json:"since"`
				MinActivationHeight int         `json:"min_activation_height"`
			} `json:"bip9"`

			Active bool `json:"active"`
		} `json:"taproot"`
	} `json:"softforks"`

	Warnings string `json:"warnings"`
}

type CreatedRawTransactionDTO struct {
	Tx  string  `json:"tx"`
	Fee float64 `json:"fee"`
}

type CreateRawTransactionDTO struct {
	Utxo []struct {
		TxId     string `json:"txid" validate:"required"`
		Vout     int64  `json:"vout" validate:"required"`
		Amount   int64  `json:"amount" validate:"required"`
		PKScript string `json:"pk_script" validate:"required"`
	} `json:"utxo" validate:"dive"`
	FromAddress string `json:"from_address" validate:"required"`
	ToAddress   string `json:"to_address" validate:"required"`
	Amount      int64  `json:"amount" validate:"required"`
}

type FundForRawTransactionDTO struct {
	CreatedTxHex  string `json:"created_tx_hex" validate:"required"`
	ChangeAddress string `json:"change_address" validate:"required"`
}

type FundedRawTransactionDTO struct {
	Tx  string  `json:"tx"`
	Fee float64 `json:"fee"`
}

type SignRawTransactionDTO struct {
	Tx         string `json:"tx" validate:"required"`
	PrivateKey string `json:"privateKey" validate:"required"`
	Utxo       []struct {
		TxId     string `json:"txid" validate:"required"`
		Vout     int64  `json:"vout" validate:"required"`
		Amount   int64  `json:"amount" validate:"required"`
		PKScript string `json:"pk_script" validate:"required"`
	} `json:"utxo" validate:"dive"`
}

type SignedRawTransactionDTO struct {
	Hash string `json:"hash"`
}

type SendRawTransactionDTO struct {
	SignedTx string `json:"signed_tx"`
}
