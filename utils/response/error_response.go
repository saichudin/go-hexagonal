package response

import (
	"e-menu-tentakel/utils/validation"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Message          string            `json:"message"`
	ErrorValidations map[string]string `json:"error_validation,omitempty"`
}

func GenerateResponseError(err error, data ...map[string]string) interface{} {
	var ve validator.ValidationErrors
	var message string

	if errors.As(err, &ve) {
		var out = make(map[string]string)

		for _, fe := range ve {
			field := fe.Namespace()
			if field == "" {
				field = "field"
			}

			out[field] = fe.Translate(validation.Translate)

			if message == "" {
				message = fe.Translate(validation.Translate)
			}
		}
		return &ErrorResponse{
			Message:          message,
			ErrorValidations: removeFirstStruct(out),
		}
	}

	var errorVal map[string]string
	if len(data) > 0 {
		errorVal = data[0]
	}

	return &ErrorResponse{
		Message:          err.Error(),
		ErrorValidations: errorVal,
	}
}

func removeFirstStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
