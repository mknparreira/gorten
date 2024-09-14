package mocks

import (
	"context"
	"fmt"
	"gorten/internal/gorten/models"

	"github.com/stretchr/testify/mock"
)

type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) List(ctx context.Context, skip int, limit int, sort string) ([]models.Category, error) {
	args := m.Called(ctx, skip, limit, sort)

	categories, _ := args.Get(0).([]models.Category)
	err, _ := args.Get(1).(error)
	return categories, err
}

func (m *MockCategoryService) GetByID(ctx context.Context, id string) (*models.Category, error) {
	args := m.Called(ctx, id)
	category, _ := args.Get(0).(*models.Category)
	err, _ := args.Get(1).(error)
	return category, err
}

func (m *MockCategoryService) Create(ctx context.Context, category *models.Category) error {
	args := m.Called(ctx, category)
	if err, ok := args.Get(0).(error); ok {
		return err
	}
	return nil
}

func (m *MockCategoryService) UpdateByID(ctx context.Context, id string, category *models.Category) error {
	args := m.Called(ctx, id, category)

	if err := args.Error(0); err != nil {
		return fmt.Errorf("MockCategoryService.Update: %w", err)
	}
	return nil
}
