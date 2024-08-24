package middlewares_test

import (
	"context"
	"encoding/json"

	"gorten/internal/gorten/api/middlewares"
	"gorten/internal/gorten/models"
	pkgerr "gorten/pkg/errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(middlewares.ErrorMiddleware())
	return router
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}

func TestErrorMiddleware_Success(t *testing.T) {
	router := setupRouter()
	router.GET("/success", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/success", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	expectedResponse := `{"message": "success"}`

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestErrorMiddleware_WithError(t *testing.T) {
	router := setupRouter()
	router.GET("/with-error", func(c *gin.Context) {
		_ = c.Error(pkgerr.ErrSomethingWentWrong)
		c.Status(http.StatusBadRequest)
	})

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/with-error", nil)
	if err != nil {
		t.Fatalf("Failed to create a request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var resp models.ErrorResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("Failed to Unmarshal actual response: %v", err)
	}

	expectedResponse := models.ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: pkgerr.ErrSomethingWentWrong.Error(),
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expectedResponse, resp)
}

func TestErrorMiddleware_Panic(t *testing.T) {
	router := setupRouter()
	router.GET("/panic", func(_ *gin.Context) {
		panic("unexpected error")
	})

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/panic", nil)
	if err != nil {
		t.Fatalf("Failed to create a request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var resp models.ErrorResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)

	if err != nil {
		t.Fatalf("Failed to Unmarshal actual response: %v", err)
	}

	expectedResponse := models.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: pkgerr.ErrInternalServerError.Error(),
	}

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expectedResponse, resp)
}
