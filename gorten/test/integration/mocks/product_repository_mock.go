package mocks

import (
	"context"
	"fmt"
	"gorten/internal/gorten/models"

	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetAll(ctx context.Context, skip int, limit int, sort string) ([]models.Product, error) {
	args := m.Called(ctx, skip, limit, sort)
	products, _ := args.Get(0).([]models.Product)
	err, _ := args.Get(1).(error)
	return products, err
}

func (m *MockProductRepository) GetByID(ctx context.Context, productID string) (*models.Product, error) {
	args := m.Called(ctx, productID)
	product, _ := args.Get(0).(*models.Product)
	err, _ := args.Get(1).(error)
	return product, err
}

func (m *MockProductRepository) Create(ctx context.Context, product *models.Product) error {
	args := m.Called(ctx, product)
	if err, ok := args.Get(0).(error); ok {
		return err
	}
	return nil
}

func (m *MockProductRepository) Update(ctx context.Context, product *models.Product) error {
	args := m.Called(ctx, product)
	if err := args.Error(0); err != nil {
		return fmt.Errorf("MockProductRepository.Update: %w", err)
	}
	return nil
}
