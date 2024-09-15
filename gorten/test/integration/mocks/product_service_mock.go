package mocks

import (
	"context"
	"fmt"
	"gorten/internal/gorten/models"

	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) List(ctx context.Context, skip int, limit int, sort string) ([]models.Product, error) {
	args := m.Called(ctx, skip, limit, sort)

	products, _ := args.Get(0).([]models.Product)
	err, _ := args.Get(1).(error)
	return products, err
}

func (m *MockProductService) GetByID(ctx context.Context, id string) (*models.Product, error) {
	args := m.Called(ctx, id)
	product, _ := args.Get(0).(*models.Product)
	err, _ := args.Get(1).(error)
	return product, err
}

func (m *MockProductService) Create(ctx context.Context, product *models.Product) error {
	args := m.Called(ctx, product)
	if err, ok := args.Get(0).(error); ok {
		return err
	}
	return nil
}

func (m *MockProductService) UpdateByID(ctx context.Context, id string, product *models.Product) error {
	args := m.Called(ctx, id, product)

	if err := args.Error(0); err != nil {
		return fmt.Errorf("MockProductService.Update: %w", err)
	}
	return nil
}
