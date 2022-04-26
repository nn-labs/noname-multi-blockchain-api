package bitcoin

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"nn-blockchain-api/pkg/errors"
	"strings"
)

type ApiError struct {
	Field string
	Msg   string
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "is required"
	}
	return ""
}

func Validate(dto interface{}) error {
	validate := validator.New()

	if err := validate.Struct(dto); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errors.WithMessage(ErrInvalidRequest, err.Error())
		}

		var out []string
		for _, err := range err.(validator.ValidationErrors) {
			//out = append(out, ApiError{
			//	Field: err.Field(),
			//	Msg:   msgForTag(err.Tag()),
			//})
			out = append(out, fmt.Sprintf("%v - %v", err.Field(), msgForTag(err.Tag())))
		}
		return errors.WithMessage(ErrInvalidRequest, strings.Join(out, ", "))
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
	Address  string `json:"address" validate:"required"`
	WalletId string `json:"wallet_id" validate:"required"`
	Network  string `json:"network" validate:"required"`
}

type ImportAddressInfoDTO struct {
	Message string `json:"message"`
}

type WalletDTO struct {
	WalletId string `json:"wallet_id" validate:"required"`
	Network  string `json:"network" validate:"required"`
}

type WalletInfoDTO struct {
	Walletname            string      `json:"walletname"`
	Walletversion         int         `json:"walletversion"`
	Format                string      `json:"format"`
	Balance               float64     `json:"balance"`
	UnconfirmedBalance    float64     `json:"unconfirmed_balance"`
	ImmatureBalance       float64     `json:"immature_balance"`
	Txcount               int         `json:"txcount"`
	Keypoololdest         int         `json:"keypoololdest"`
	Keypoolsize           int         `json:"keypoolsize"`
	Hdseedid              string      `json:"hdseedid"`
	KeypoolsizeHdInternal int         `json:"keypoolsize_hd_internal"`
	Paytxfee              float64     `json:"paytxfee"`
	PrivateKeysEnabled    bool        `json:"private_keys_enabled"`
	AvoidReuse            bool        `json:"avoid_reuse"`
	Scanning              interface{} `json:"scanning"`
	Descriptors           bool        `json:"descriptors"`
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
	WalletId string `json:"wallet_id" validate:"required"`
	Network  string `json:"network" validate:"required"`
}

type LoadWalletInfoDTO struct {
	Message string `json:"message"`
}

type RescanWalletDTO struct {
	WalletId string `json:"wallet_id" validate:"required"`
	Network  string `json:"network" validate:"required"`
}

type RescanWalletInfoDTO struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ListUnspentDTO struct {
	Address  string `json:"address" validate:"required"`
	WalletId string `json:"wallet_id" validate:"required"`
	Network  string `json:"network" validate:"required"`
}

type UnspentInfoDTO struct {
	Txid          string  `json:"txid"`
	Vout          int     `json:"vout"`
	Address       string  `json:"address"`
	Label         string  `json:"label"`
	ScriptPubKey  string  `json:"scriptPubKey"`
	Amount        float64 `json:"amount"`
	Confirmations int     `json:"confirmations"`
	Spendable     bool    `json:"spendable"`
	Solvable      bool    `json:"solvable"`
	Safe          bool    `json:"safe"`
}

type ListUnspentInfoDTO struct {
	Result []*UnspentInfoDTO
}
