package repositories_test

import (
	"context"
	"testing"

	"gorten/internal/gorten/models"
	"gorten/test/factories"
	"gorten/test/integration/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductRepository_GetAll(t *testing.T) {
	ctx := context.Background()
	productA := factories.ProductFactory()
	productB := factories.ProductFactory(func(u *models.Product) {
		u.Name = "ProductB"
		u.Description = "DescriptionB"
	})

	expectedProducts := []models.Product{*productA, *productB}
	mockRepo := new(mocks.MockProductRepository)
	mockRepo.On("GetAll", ctx, 0, 10, "desc").Return(expectedProducts, nil)
	products, err := mockRepo.GetAll(ctx, 0, 10, "desc")

	require.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, productA.Name, products[0].Name)
	assert.Equal(t, productB.Name, products[1].Name)
	mockRepo.AssertExpectations(t)
}

func TestProductRepository_GetByID(t *testing.T) {
	ctx := context.Background()
	newProduct := factories.ProductFactory()
	mockRepo := new(mocks.MockProductRepository)

	mockRepo.On("GetByID", ctx, newProduct.ProductID).Return(newProduct, nil)
	product, err := mockRepo.GetByID(ctx, newProduct.ProductID)

	require.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, newProduct.Name, product.Name)
	mockRepo.AssertExpectations(t)
}

func TestProductRepository_Create(t *testing.T) {
	ctx := context.Background()
	newProduct := factories.ProductFactory()
	mockRepo := new(mocks.MockProductRepository)

	mockRepo.On("Create", ctx, newProduct).Return(nil)
	err := mockRepo.Create(ctx, newProduct)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductRepository_Update(t *testing.T) {
	ctx := context.Background()
	product := factories.ProductFactory()
	mockRepo := new(mocks.MockProductRepository)

	mockRepo.On("Update", ctx, product).Return(nil)
	err := mockRepo.Update(ctx, product)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
