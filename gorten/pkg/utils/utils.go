package utils

import (
	pkgerr "gorten/pkg/errors"

	"github.com/google/uuid"
)

func GenerateUUID() (string, error) {
	id, err := uuid.NewUUID()

	if err != nil {
		return "Failed to generate UUID", pkgerr.ErrFailedGenerateUUID
	}
	return id.String(), nil
}
