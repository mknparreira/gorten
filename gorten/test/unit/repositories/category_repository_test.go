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

func TestCategoryRepository_GetAll(t *testing.T) {
	ctx := context.Background()
	categoryA := factories.CategoryFactory()
	categoryB := factories.CategoryFactory(func(u *models.Category) {
		u.Name = "CategoryB"
	})

	expectedCategories := []models.Category{*categoryA, *categoryB}
	mockRepo := new(mocks.MockCategoryRepository)
	mockRepo.On("GetAll", ctx, 0, 10, "desc").Return(expectedCategories, nil)
	categories, err := mockRepo.GetAll(ctx, 0, 10, "desc")

	require.NoError(t, err)
	assert.Len(t, categories, 2)
	assert.Equal(t, categoryA.Name, categories[0].Name)
	assert.Equal(t, categoryB.Name, categories[1].Name)
	mockRepo.AssertExpectations(t)
}

func TestCategoryRepository_GetByID(t *testing.T) {
	ctx := context.Background()
	newCategory := factories.CategoryFactory()
	mockRepo := new(mocks.MockCategoryRepository)

	mockRepo.On("GetByID", ctx, newCategory.CategoryID).Return(newCategory, nil)
	category, err := mockRepo.GetByID(ctx, newCategory.CategoryID)

	require.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, newCategory.Name, category.Name)
	mockRepo.AssertExpectations(t)
}

func TestCategoryRepository_Create(t *testing.T) {
	ctx := context.Background()
	newCategory := factories.CategoryFactory()
	mockRepo := new(mocks.MockCategoryRepository)

	mockRepo.On("Create", ctx, newCategory).Return(nil)
	err := mockRepo.Create(ctx, newCategory)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryRepository_Update(t *testing.T) {
	ctx := context.Background()
	category := factories.CategoryFactory()
	mockRepo := new(mocks.MockCategoryRepository)

	mockRepo.On("Update", ctx, category).Return(nil)
	err := mockRepo.Update(ctx, category)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
