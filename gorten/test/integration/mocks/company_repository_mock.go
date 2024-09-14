package mocks

import (
	"context"
	"fmt"
	"gorten/internal/gorten/models"

	"github.com/stretchr/testify/mock"
)

type MockCompanyRepository struct {
	mock.Mock
}

func (m *MockCompanyRepository) GetAll(ctx context.Context, skip int, limit int, sort string) ([]models.Company, error) {
	args := m.Called(ctx, skip, limit, sort)
	companies, _ := args.Get(0).([]models.Company)
	err, _ := args.Get(1).(error)
	return companies, err
}

func (m *MockCompanyRepository) GetByID(ctx context.Context, companyID string) (*models.Company, error) {
	args := m.Called(ctx, companyID)
	company, _ := args.Get(0).(*models.Company)
	err, _ := args.Get(1).(error)
	return company, err
}

func (m *MockCompanyRepository) Create(ctx context.Context, company *models.Company) error {
	args := m.Called(ctx, company)
	if err, ok := args.Get(0).(error); ok {
		return err
	}
	return nil
}

func (m *MockCompanyRepository) Update(ctx context.Context, company *models.Company) error {
	args := m.Called(ctx, company)
	if err := args.Error(0); err != nil {
		return fmt.Errorf("MockCompanyRepository.Update: %w", err)
	}
	return nil
}
