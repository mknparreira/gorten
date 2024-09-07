package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gorten/internal/gorten/api/handlers"
	"gorten/internal/gorten/models"
	"gorten/test/factories"
	"gorten/test/integration/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter(handler handlers.UserHandlerImpl) *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/users", handler.List)
	router.GET("/api/v1/users/:id", handler.UserByID)
	router.POST("/api/v1/users", handler.Create)
	router.PUT("/api/v1/users/:id", handler.UpdateByID)
	return router
}

func TestUserHandler_List(t *testing.T) {
	user := factories.UserFactory()
	expectedUsers := []models.User{*user}
	mockUserService := new(mocks.MockUserService)
	mockUserService.On("List", mock.Anything, 0, 10, "desc").Return(expectedUsers, nil)

	userHandler := handlers.User(mockUserService)
	router := setupRouter(userHandler)

	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/api/v1/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Grand Doe")
	mockUserService.AssertCalled(t, "List", mock.Anything, 0, 10, "desc")
}

func TestUserHandler_UserByID(t *testing.T) {
	user := factories.UserFactory()
	mockUserService := new(mocks.MockUserService)
	mockUserService.On("GetByID", mock.Anything, user.UserID).Return(user, nil)

	userHandler := handlers.User(mockUserService)
	router := setupRouter(userHandler)

	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/api/v1/users/"+user.UserID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), user.Name)
	mockUserService.AssertCalled(t, "GetByID", mock.Anything, user.UserID)
}

func TestUserHandler_Create(t *testing.T) {
	newUser := factories.UserFactory(func(u *models.User) {
		u.UserID = ""
		u.CreatedAt = time.Time{}
	})
	ctx := context.Background()
	mockUserService := new(mocks.MockUserService)
	mockUserService.On("Create", mock.Anything, newUser).Return(nil)

	userHandler := handlers.User(mockUserService)
	router := setupRouter(userHandler)

	body := map[string]string{
		"name":     newUser.Name,
		"email":    newUser.Email,
		"password": newUser.Password,
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req, _ := http.NewRequestWithContext(ctx, "POST", "/api/v1/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockUserService.AssertCalled(t, "Create", mock.Anything, newUser)
}

func TestUserHandler_UpdateByID(t *testing.T) {
	user := factories.UserFactory()
	ctx := context.Background()
	mockUserService := new(mocks.MockUserService)
	mockUserService.On("UpdateByID", mock.Anything, user.UserID, mock.AnythingOfType("*models.User")).Return(nil)

	userHandler := handlers.User(mockUserService)
	router := setupRouter(userHandler)

	body := map[string]string{
		"userID":   user.UserID,
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req, _ := http.NewRequestWithContext(ctx, "PUT", "/api/v1/users/"+user.UserID, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockUserService.AssertCalled(t, "UpdateByID", mock.Anything, user.UserID, mock.AnythingOfType("*models.User"))
}
