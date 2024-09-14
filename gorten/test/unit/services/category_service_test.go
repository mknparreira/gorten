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

const categoryName = "CategoryA"

func TestCategoryService_List(t *testing.T) {
	ctx := context.Background()
	category := factories.CategoryFactory()
	mockRepo := new(mocks.MockCategoryRepository)
	service := services.CategoryServiceInit(mockRepo)

	expectedCategories := []models.Category{*category}
	mockRepo.On("GetAll", ctx, 0, 10, "desc").Return(expectedCategories, nil)

	categories, err := service.List(ctx, 0, 10, "desc")

	require.NoError(t, err)
	assert.Equal(t, expectedCategories, categories)
	mockRepo.AssertExpectations(t)
}

func TestCategoryService_GetByID(t *testing.T) {
	ctx := context.Background()
	expectedCategory := factories.CategoryFactory()
	mockRepo := new(mocks.MockCategoryRepository)
	service := services.CategoryServiceInit(mockRepo)

	mockRepo.On("GetByID", ctx, expectedCategory.CategoryID).Return(expectedCategory, nil)
	category, err := service.GetByID(ctx, expectedCategory.CategoryID)

	require.NoError(t, err)
	assert.Equal(t, expectedCategory, category)
	mockRepo.AssertExpectations(t)
}

func TestCategoryService_Create(t *testing.T) {
	ctx := context.Background()
	newCategory := factories.CategoryFactory()
	mockRepo := new(mocks.MockCategoryRepository)
	service := services.CategoryServiceInit(mockRepo)

	mockRepo.On("Create", ctx, newCategory).Return(nil)
	err := service.Create(ctx, newCategory)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryService_UpdateByID(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(mocks.MockCategoryRepository)
	service := services.CategoryServiceInit(mockRepo)

	existingCategory := factories.CategoryFactory()
	updatedCategory := factories.CategoryFactory(func(u *models.Category) {
		u.Name = janeDoeName
	})

	mockRepo.On("GetByID", ctx, existingCategory.CategoryID).Return(existingCategory, nil)
	mockRepo.On("Update", ctx, existingCategory).Return(nil)

	err := service.UpdateByID(ctx, existingCategory.CategoryID, updatedCategory)

	require.NoError(t, err)
	assert.Equal(t, updatedCategory.Name, existingCategory.Name)
	mockRepo.AssertExpectations(t)
}

func TestCategoryService_UpdateByID_CategoryNotFound(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(mocks.MockCategoryRepository)
	service := services.CategoryServiceInit(mockRepo)
	category := factories.CategoryFactory()
	updatedCategory := factories.CategoryFactory(func(u *models.Category) {
		u.Name = categoryName
	})

	mockRepo.On("GetByID", ctx, category.CategoryID).Return(nil, errors.ErrCategoryNotFound)
	err := service.UpdateByID(ctx, category.CategoryID, updatedCategory)

	require.Error(t, err)
	mockRepo.AssertExpectations(t)
}
