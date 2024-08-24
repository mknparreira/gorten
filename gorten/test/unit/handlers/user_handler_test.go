package handlers_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorten/internal/gorten/api/handlers"
	"gorten/internal/gorten/models"
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
	expectedUsers := []models.User{
		{UserID: "51a9cbf0-efe9-4331-ae74-8805f813c6e8", Name: "John Doe", Email: "john.doe@example.com"},
	}
	mockUserService := new(mocks.MockUserService)
	mockUserService.On("List", mock.Anything).Return(expectedUsers, nil)

	userHandler := handlers.User(mockUserService)
	router := setupRouter(userHandler)

	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/api/v1/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
	mockUserService.AssertCalled(t, "List", mock.Anything)
}

func TestUserHandler_UserByID(t *testing.T) {
	user := &models.User{UserID: "51a9cbf0-efe9-4331-ae74-8805f813c6e8", Name: "John Doe", Email: "john.doe@example.com"}
	mockUserService := new(mocks.MockUserService)
	mockUserService.On("GetByID", mock.Anything, "51a9cbf0-efe9-4331-ae74-8805f813c6e8").Return(user, nil)

	userHandler := handlers.User(mockUserService)
	router := setupRouter(userHandler)

	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/api/v1/users/51a9cbf0-efe9-4331-ae74-8805f813c6e8", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
	mockUserService.AssertCalled(t, "GetByID", mock.Anything, "51a9cbf0-efe9-4331-ae74-8805f813c6e8")
}

func TestUserHandler_Create(t *testing.T) {
	newUser := &models.User{
		UserID:   "51a9cbf0-efe9-4331-ae74-8805f813c6e8",
		Name:     "John Brolie Doe",
		Email:    "john.doe@example.com",
		Password: "123456"}
	ctx := context.Background()
	mockUserService := new(mocks.MockUserService)
	mockUserService.On("Create", mock.Anything, newUser).Return(nil)

	userHandler := handlers.User(mockUserService)
	router := setupRouter(userHandler)

	reqBody := `{"userID": "51a9cbf0-efe9-4331-ae74-8805f813c6e8", "name": "John Brolie Doe", "email": "john.doe@example.com", "Password": "123456"}`
	req, _ := http.NewRequestWithContext(ctx, "POST", "/api/v1/users", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockUserService.AssertCalled(t, "Create", mock.Anything, newUser)
}

func TestUserHandler_UpdateByID(t *testing.T) {
	ctx := context.Background()
	mockUserService := new(mocks.MockUserService)
	mockUserService.On("UpdateByID", mock.Anything, "51a9cbf0-efe9-4331-ae74-8805f813c6e8", mock.AnythingOfType("*models.User")).Return(nil)

	userHandler := handlers.User(mockUserService)
	router := setupRouter(userHandler)

	reqBody := `{"UserID": "51a9cbf0-efe9-4331-ae74-8805f813c6e8", "Name": "John Broli Doe", "Email": "john.doe@example.com", "Password": "123456"}`
	req, _ := http.NewRequestWithContext(ctx, "PUT", "/api/v1/users/51a9cbf0-efe9-4331-ae74-8805f813c6e8", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockUserService.AssertCalled(t, "UpdateByID", mock.Anything, "51a9cbf0-efe9-4331-ae74-8805f813c6e8", mock.AnythingOfType("*models.User"))
}
