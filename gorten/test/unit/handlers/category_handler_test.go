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

func setupCategoryRouter(handler handlers.CategoryHandlerImpl) *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/categories", handler.List)
	router.GET("/api/v1/categories/:id", handler.GetByID)
	router.POST("/api/v1/categories", handler.Create)
	router.PUT("/api/v1/categories/:id", handler.UpdateByID)
	return router
}

func TestCategoryHandler_List(t *testing.T) {
	category := factories.CategoryFactory()
	expectedCategories := []models.Category{*category}
	mockCategoryService := new(mocks.MockCategoryService)
	mockCategoryService.On("List", mock.Anything, 0, 10, "desc").Return(expectedCategories, nil)

	categoryHandler := handlers.Category(mockCategoryService)
	router := setupCategoryRouter(categoryHandler)

	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/api/v1/categories", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Category")
	mockCategoryService.AssertCalled(t, "List", mock.Anything, 0, 10, "desc")
}

func TestCategoryHandler_CategoryByID(t *testing.T) {
	category := factories.CategoryFactory()
	mockCategoryService := new(mocks.MockCategoryService)
	mockCategoryService.On("GetByID", mock.Anything, category.CategoryID).Return(category, nil)

	categoryHandler := handlers.Category(mockCategoryService)
	router := setupCategoryRouter(categoryHandler)

	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/api/v1/categories/"+category.CategoryID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), category.Name)
	mockCategoryService.AssertCalled(t, "GetByID", mock.Anything, category.CategoryID)
}

func TestCategoryHandler_Create(t *testing.T) {
	newCategory := factories.CategoryFactory(func(u *models.Category) {
		u.CategoryID = ""
		u.CreatedAt = time.Time{}
	})
	ctx := context.Background()
	mockCategoryService := new(mocks.MockCategoryService)
	mockCategoryService.On("Create", mock.Anything, newCategory).Return(nil)

	categoryHandler := handlers.Category(mockCategoryService)
	router := setupCategoryRouter(categoryHandler)

	body := map[string]string{
		"name": newCategory.Name,
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req, _ := http.NewRequestWithContext(ctx, "POST", "/api/v1/categories", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockCategoryService.AssertCalled(t, "Create", mock.Anything, newCategory)
}

func TestCategoryHandler_UpdateByID(t *testing.T) {
	category := factories.CategoryFactory()
	ctx := context.Background()
	mockCategoryService := new(mocks.MockCategoryService)
	mockCategoryService.On("UpdateByID", mock.Anything, category.CategoryID, mock.AnythingOfType("*models.Category")).Return(nil)

	categoryHandler := handlers.Category(mockCategoryService)
	router := setupCategoryRouter(categoryHandler)

	body := map[string]string{
		"categoryID": category.CategoryID,
		"name":       category.Name,
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req, _ := http.NewRequestWithContext(ctx, "PUT", "/api/v1/categories/"+category.CategoryID, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockCategoryService.AssertCalled(t, "UpdateByID", mock.Anything, category.CategoryID, mock.AnythingOfType("*models.Category"))
}
