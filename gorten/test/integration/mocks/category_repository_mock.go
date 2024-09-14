package mocks

import (
	"context"
	"fmt"
	"gorten/internal/gorten/models"

	"github.com/stretchr/testify/mock"
)

type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) GetAll(ctx context.Context, skip int, limit int, sort string) ([]models.Category, error) {
	args := m.Called(ctx, skip, limit, sort)
	categories, _ := args.Get(0).([]models.Category)
	err, _ := args.Get(1).(error)
	return categories, err
}

func (m *MockCategoryRepository) GetByID(ctx context.Context, categoryID string) (*models.Category, error) {
	args := m.Called(ctx, categoryID)
	category, _ := args.Get(0).(*models.Category)
	err, _ := args.Get(1).(error)
	return category, err
}

func (m *MockCategoryRepository) Create(ctx context.Context, category *models.Category) error {
	args := m.Called(ctx, category)
	if err, ok := args.Get(0).(error); ok {
		return err
	}
	return nil
}

func (m *MockCategoryRepository) Update(ctx context.Context, category *models.Category) error {
	args := m.Called(ctx, category)
	if err := args.Error(0); err != nil {
		return fmt.Errorf("MockCategoryRepository.Update: %w", err)
	}
	return nil
}
