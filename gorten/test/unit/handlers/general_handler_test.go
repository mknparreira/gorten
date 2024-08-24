package handlers_test

import (
	"context"
	"gorten/internal/gorten/api/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouterGeneral(handler handlers.GeneralHandlerImpl) *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/ping", handler.PingHandler)
	return router
}

func TestPingHandler_List(t *testing.T) {
	generalHandler := &handlers.GeneralHandler{}
	router := setupRouterGeneral(generalHandler)

	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/api/v1/ping", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}
