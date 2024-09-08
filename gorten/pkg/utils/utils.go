package utils

import (
	"fmt"
	pkgerr "gorten/pkg/errors"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func GenerateUUID() (string, error) {
	id, err := uuid.NewUUID()

	if err != nil {
		return "Failed to generate UUID", pkgerr.ErrFailedGenerateUUID
	}
	return id.String(), nil
}

func ValidationErrors(ve validator.ValidationErrors) error {
	out := make([]string, len(ve))
	for i, fe := range ve {
		out[i] = fmt.Sprintf("Field '%s' validate failed on '%s'", fe.Field(), fe.Tag())
	}
	return pkgerr.ErrValidationFailed.WithMessage(fmt.Sprintf("validation failed: %v", out))
}

func ConvertStringSortforInteger(sort string) int {
	if sort == "asc" {
		return 1
	}
	return -1
}
