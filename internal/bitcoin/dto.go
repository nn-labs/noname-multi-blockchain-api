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
	Network string `json:"network" validate:"required"`
}

type StatusNodeInfoDTO struct {
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
	Network     string `json:"network" validate:"required"`
}

type DecodeRawTransactionDTO struct {
	Tx      string `json:"tx" validate:"required"`
	Network string `json:"network" validate:"required"`
}

type DecodedRawTransactionDTO struct {
	Txid     string `json:"txid"`
	Hash     string `json:"hash"`
	Version  int    `json:"version"`
	Size     int    `json:"size"`
	Vsize    int    `json:"vsize"`
	Weight   int    `json:"weight"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Txid      string `json:"txid"`
		Vout      int    `json:"vout"`
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`

		Sequence int64 `json:"sequence"`
	} `json:"vin"`
	Vout []struct {
		Value        float64 `json:"value"`
		N            int     `json:"n"`
		ScriptPubKey struct {
			Asm     string `json:"asm"`
			Hex     string `json:"hex"`
			Address string `json:"address"`
			Type    string `json:"type"`
		} `json:"scriptPubKey"`
	} `json:"vout"`
}

type FundForRawTransactionDTO struct {
	CreatedTxHex  string `json:"created_tx_hex" validate:"required"`
	ChangeAddress string `json:"change_address" validate:"required"`
	Network       string `json:"network" validate:"required"`
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
	Network string `json:"network" validate:"required"`
}

type SignedRawTransactionDTO struct {
	Hash string `json:"hash"`
}

type SendRawTransactionDTO struct {
	SignedTx string `json:"signed_tx" validate:"required"`
	Network  string `json:"network" validate:"required"`
}

type SentRawTransactionDTO struct {
	TxId string `json:"tx_id"`
}

type ImportAddressDTO struct {
	Address string `json:"address" validate:"required"`
	Network string `json:"network" validate:"required"`
}

type ImportAddressInfoDTO struct {
	Message string `json:"message"`
}

type CreateWalletDTO struct {
	//Password string `json:"password" validate:"required"`
	Network string `json:"network" validate:"required"`
}

type CreatedWalletInfoDTO struct {
	WalletId string `json:"wallet_id"`
	Password string `json:"password"`
}

type LoadWalletDTO struct {
	Network string `json:"network" validate:"required"`
}

type LoadWalletInfoDTO struct {
	Message string `json:"message"`
}
