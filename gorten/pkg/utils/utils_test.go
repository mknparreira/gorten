package utils

import (
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUtils_GenerateUUID(t *testing.T) {
	id, err := GenerateUUID()
	validateRegex := "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"

	require.NoError(t, err)
	assert.NotEmpty(t, id)
	assert.Regexp(t, validateRegex, id)
}

func TestUtils_ValidationErrors(t *testing.T) {
	validate := validator.New()
	expectedMessage := "validation failed: [Field 'Name' validate failed on 'required' Field 'Age' validate failed on 'gte']"

	type User struct {
		Name string `validate:"required"`
		Age  int    `validate:"gte=18"`
	}

	user := &User{}
	err := validate.Struct(user)

	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		t.Fatalf("expected validation errors, got: %v", err)
	}
	resultErr := ValidationErrors(validationErrors)

	require.Error(t, resultErr)
	assert.Contains(t, resultErr.Error(), expectedMessage)
}

func TestUtils_ConvertStringSortToInteger(t *testing.T) {
	tests := []struct {
		name     string
		sort     string
		expected int
	}{
		{"Ascending", "asc", 1},
		{"Descending", "desc", -1},
		{"Empty String", "", -1},
		{"Invalid String", "invalid", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertStringSortforInteger(tt.sort)
			assert.Equal(t, tt.expected, result)
		})
	}
}
