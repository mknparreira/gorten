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

func setupProductRouter(handler handlers.ProductHandlerImpl) *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/products", handler.List)
	router.GET("/api/v1/products/:id", handler.GetByID)
	router.POST("/api/v1/products", handler.Create)
	router.PUT("/api/v1/products/:id", handler.UpdateByID)
	return router
}

func TestProductHandler_List(t *testing.T) {
	product := factories.ProductFactory()
	expectedProducts := []models.Product{*product}
	mockProductService := new(mocks.MockProductService)
	mockProductService.On("List", mock.Anything, 0, 10, "desc").Return(expectedProducts, nil)

	productHandler := handlers.Product(mockProductService)
	router := setupProductRouter(productHandler)

	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/api/v1/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Product")
	mockProductService.AssertCalled(t, "List", mock.Anything, 0, 10, "desc")
}

func TestProductHandler_ProductByID(t *testing.T) {
	product := factories.ProductFactory()
	mockProductService := new(mocks.MockProductService)
	mockProductService.On("GetByID", mock.Anything, product.ProductID).Return(product, nil)

	productHandler := handlers.Product(mockProductService)
	router := setupProductRouter(productHandler)

	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", "/api/v1/products/"+product.ProductID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), product.Name)
	mockProductService.AssertCalled(t, "GetByID", mock.Anything, product.ProductID)
}

func TestProductHandler_Create(t *testing.T) {
	newProduct := factories.ProductFactory(func(u *models.Product) {
		u.ProductID = ""
		u.CreatedAt = time.Time{}
	})
	ctx := context.Background()
	mockProductService := new(mocks.MockProductService)
	mockProductService.On("Create", mock.Anything, newProduct).Return(nil)

	productHandler := handlers.Product(mockProductService)
	router := setupProductRouter(productHandler)

	body := map[string]string{
		"name": newProduct.Name,
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req, _ := http.NewRequestWithContext(ctx, "POST", "/api/v1/products", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockProductService.AssertCalled(t, "Create", mock.Anything, newProduct)
}

func TestProductHandler_UpdateByID(t *testing.T) {
	product := factories.ProductFactory()
	ctx := context.Background()
	mockProductService := new(mocks.MockProductService)
	mockProductService.On("UpdateByID", mock.Anything, product.ProductID, mock.AnythingOfType("*models.Product")).Return(nil)

	productHandler := handlers.Product(mockProductService)
	router := setupProductRouter(productHandler)

	body := map[string]string{
		"productID":   product.ProductID,
		"name":        product.Name,
		"Description": product.Description,
	}
	reqBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req, _ := http.NewRequestWithContext(ctx, "PUT", "/api/v1/products/"+product.ProductID, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockProductService.AssertCalled(t, "UpdateByID", mock.Anything, product.ProductID, mock.AnythingOfType("*models.Product"))
}
