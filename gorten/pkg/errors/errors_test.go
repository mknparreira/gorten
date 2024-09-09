package errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var customErrorMessage = "custom error message"

func TestNewCustomError_WithCorrectStatusCodeAndMessage(t *testing.T) {
	statusCode := http.StatusNotFound

	customErr := NewCustomError(statusCode, customErrorMessage)

	assert.Equal(t, statusCode, customErr.StatusCode)
	assert.Equal(t, customErrorMessage, customErr.Message)
}

func TestCustomError_Error(t *testing.T) {
	statusCode := http.StatusNotFound

	customErr := NewCustomError(statusCode, customErrorMessage)

	assert.Equal(t, customErrorMessage, customErr.Error())
}

func TestCustomError_WithMessage(t *testing.T) {
	statusCode := http.StatusNotFound
	originalMessage := "original message"
	newMessage := "new message"

	customErr := NewCustomError(statusCode, originalMessage)
	updatedErr := customErr.WithMessage(newMessage)

	assert.Equal(t, statusCode, updatedErr.StatusCode)
	assert.Equal(t, newMessage, updatedErr.Message)
	assert.NotEqual(t, originalMessage, updatedErr.Message)
}
