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
