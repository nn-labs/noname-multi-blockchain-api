package ethereum

import (
	"fmt"
	"nn-blockchain-api/pkg/errors"
	"strings"

	"github.com/go-playground/validator/v10"
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

type NodeInfoDTO struct {
	CurrentBlock        string `json:"currentBlock,omitempty"`
	HealedBytecodeBytes string `json:"healedBytecodeBytes,omitempty"`
	HealedBytecodes     string `json:"healedBytecodes,omitempty"`
	HealedTrienodeBytes string `json:"healedTrienodeBytes,omitempty"`
	HealedTrienodes     string `json:"healedTrienodes,omitempty"`
	HealingBytecode     string `json:"healingBytecode,omitempty"`
	HealingTrienodes    string `json:"healingTrienodes,omitempty"`
	HighestBlock        string `json:"highestBlock,omitempty"`
	StartingBlock       string `json:"startingBlock,omitempty"`
	SyncedAccountBytes  string `json:"syncedAccountBytes,omitempty"`
	SyncedAccounts      string `json:"syncedAccounts,omitempty"`
	SyncedBytecodeBytes string `json:"syncedBytecodeBytes,omitempty"`
	SyncedBytecodes     string `json:"syncedBytecodes,omitempty"`
	SyncedStorage       string `json:"syncedStorage,omitempty"`
	SyncedStorageBytes  string `json:"syncedStorageBytes,omitempty"`
	SyncMessage         string `json:"sync_message,omitempty"`
}

type CreateRawTransactionDTO struct {
	FromAddress string  `json:"from_address" validate:"required"`
	ToAddress   string  `json:"to_address" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
	Network     string  `json:"network" validate:"required"`
}

type CreatedRawTransactionDTO struct {
	Tx  string  `json:"tx"`
	Fee float64 `json:"fee"`
}

type SignRawTransactionDTO struct {
	Tx         string `json:"tx" validate:"required"`
	PrivateKey string `json:"privateKey" validate:"required"`
	Network    string `json:"network" validate:"required"`
}

type SignedRawTransactionDTO struct {
	SignedTx string `json:"signed_tx"`
}

type SendRawTransactionDTO struct {
	SignedTx string `json:"signed_tx" validate:"required"`
	Network  string `json:"network" validate:"required"`
}

type SentRawTransactionDTO struct {
	TxId string `json:"tx_id"`
}
