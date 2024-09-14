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

const companyName = "CompanyA"

func TestCompanyService_List(t *testing.T) {
	ctx := context.Background()
	company := factories.CompanyFactory()
	mockRepo := new(mocks.MockCompanyRepository)
	service := services.CompanyServiceInit(mockRepo)

	expectedCompanies := []models.Company{*company}
	mockRepo.On("GetAll", ctx, 0, 10, "desc").Return(expectedCompanies, nil)

	companies, err := service.List(ctx, 0, 10, "desc")

	require.NoError(t, err)
	assert.Equal(t, expectedCompanies, companies)
	mockRepo.AssertExpectations(t)
}

func TestCompanyService_GetByID(t *testing.T) {
	ctx := context.Background()
	expectedCompany := factories.CompanyFactory()
	mockRepo := new(mocks.MockCompanyRepository)
	service := services.CompanyServiceInit(mockRepo)

	mockRepo.On("GetByID", ctx, expectedCompany.CompanyID).Return(expectedCompany, nil)
	company, err := service.GetByID(ctx, expectedCompany.CompanyID)

	require.NoError(t, err)
	assert.Equal(t, expectedCompany, company)
	mockRepo.AssertExpectations(t)
}

func TestCompanyService_Create(t *testing.T) {
	ctx := context.Background()
	newCompany := factories.CompanyFactory()
	mockRepo := new(mocks.MockCompanyRepository)
	service := services.CompanyServiceInit(mockRepo)

	mockRepo.On("Create", ctx, newCompany).Return(nil)
	err := service.Create(ctx, newCompany)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCompanyService_UpdateByID(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(mocks.MockCompanyRepository)
	service := services.CompanyServiceInit(mockRepo)

	existingCompany := factories.CompanyFactory()
	updatedCompany := factories.CompanyFactory(func(u *models.Company) {
		u.Name = janeDoeName
	})

	mockRepo.On("GetByID", ctx, existingCompany.CompanyID).Return(existingCompany, nil)
	mockRepo.On("Update", ctx, existingCompany).Return(nil)

	err := service.UpdateByID(ctx, existingCompany.CompanyID, updatedCompany)

	require.NoError(t, err)
	assert.Equal(t, updatedCompany.Name, existingCompany.Name)
	mockRepo.AssertExpectations(t)
}

func TestCompanyService_UpdateByID_CompanyNotFound(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(mocks.MockCompanyRepository)
	service := services.CompanyServiceInit(mockRepo)
	company := factories.CompanyFactory()
	updatedCompany := factories.CompanyFactory(func(u *models.Company) {
		u.Name = companyName
	})

	mockRepo.On("GetByID", ctx, company.CompanyID).Return(nil, errors.ErrCompanyNotFound)
	err := service.UpdateByID(ctx, company.CompanyID, updatedCompany)

	require.Error(t, err)
	mockRepo.AssertExpectations(t)
}
