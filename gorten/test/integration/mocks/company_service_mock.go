package mocks

import (
	"context"
	"fmt"
	"gorten/internal/gorten/models"

	"github.com/stretchr/testify/mock"
)

type MockCompanyService struct {
	mock.Mock
}

func (m *MockCompanyService) List(ctx context.Context, skip int, limit int, sort string) ([]models.Company, error) {
	args := m.Called(ctx, skip, limit, sort)

	companies, _ := args.Get(0).([]models.Company)
	err, _ := args.Get(1).(error)
	return companies, err
}

func (m *MockCompanyService) GetByID(ctx context.Context, id string) (*models.Company, error) {
	args := m.Called(ctx, id)
	company, _ := args.Get(0).(*models.Company)
	err, _ := args.Get(1).(error)
	return company, err
}

func (m *MockCompanyService) Create(ctx context.Context, company *models.Company) error {
	args := m.Called(ctx, company)
	if err, ok := args.Get(0).(error); ok {
		return err
	}
	return nil
}

func (m *MockCompanyService) UpdateByID(ctx context.Context, id string, company *models.Company) error {
	args := m.Called(ctx, id, company)

	if err := args.Error(0); err != nil {
		return fmt.Errorf("MockCompanyService.Update: %w", err)
	}
	return nil
}
