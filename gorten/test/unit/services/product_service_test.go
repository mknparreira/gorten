package services_test

import (
	"context"
	"gorten/internal/gorten/models"
	"gorten/internal/gorten/services"
	"gorten/pkg/errors"
	"gorten/test/factories"
	"gorten/test/integration/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const productName = "ProductA"

func TestProductService_List(t *testing.T) {
	ctx := context.Background()
	product := factories.ProductFactory()
	mockRepo := new(mocks.MockProductRepository)
	service := services.ProductServiceInit(mockRepo)

	expectedProducts := []models.Product{*product}
	mockRepo.On("GetAll", ctx, 0, 10, "desc").Return(expectedProducts, nil)

	products, err := service.List(ctx, 0, 10, "desc")

	require.NoError(t, err)
	assert.Equal(t, expectedProducts, products)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetByID(t *testing.T) {
	ctx := context.Background()
	expectedProduct := factories.ProductFactory()
	mockRepo := new(mocks.MockProductRepository)
	service := services.ProductServiceInit(mockRepo)

	mockRepo.On("GetByID", ctx, expectedProduct.ProductID).Return(expectedProduct, nil)
	product, err := service.GetByID(ctx, expectedProduct.ProductID)

	require.NoError(t, err)
	assert.Equal(t, expectedProduct, product)
	mockRepo.AssertExpectations(t)
}

func TestProductService_Create(t *testing.T) {
	ctx := context.Background()
	newProduct := factories.ProductFactory()
	mockRepo := new(mocks.MockProductRepository)
	service := services.ProductServiceInit(mockRepo)

	mockRepo.On("Create", ctx, newProduct).Return(nil)
	err := service.Create(ctx, newProduct)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_UpdateByID(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(mocks.MockProductRepository)
	service := services.ProductServiceInit(mockRepo)

	existingProduct := factories.ProductFactory()
	updatedProduct := factories.ProductFactory(func(u *models.Product) {
		u.Name = janeDoeName
	})

	mockRepo.On("GetByID", ctx, existingProduct.ProductID).Return(existingProduct, nil)
	mockRepo.On("Update", ctx, existingProduct).Return(nil)

	err := service.UpdateByID(ctx, existingProduct.ProductID, updatedProduct)

	require.NoError(t, err)
	assert.Equal(t, updatedProduct.Name, existingProduct.Name)
	mockRepo.AssertExpectations(t)
}

func TestProductService_UpdateByID_ProductNotFound(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(mocks.MockProductRepository)
	service := services.ProductServiceInit(mockRepo)
	product := factories.ProductFactory()
	updatedProduct := factories.ProductFactory(func(u *models.Product) {
		u.Name = productName
	})

	mockRepo.On("GetByID", ctx, product.ProductID).Return(nil, errors.ErrProductNotFound)
	err := service.UpdateByID(ctx, product.ProductID, updatedProduct)

	require.Error(t, err)
	mockRepo.AssertExpectations(t)
}
