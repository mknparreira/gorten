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

func TestCompanyRepository_GetAll(t *testing.T) {
	ctx := context.Background()
	companyA := factories.CompanyFactory()
	companyB := factories.CompanyFactory(func(u *models.Company) {
		u.Name = "CompanyB"
	})

	expectedCompanies := []models.Company{*companyA, *companyB}
	mockRepo := new(mocks.MockCompanyRepository)
	mockRepo.On("GetAll", ctx, 0, 10, "desc").Return(expectedCompanies, nil)
	companies, err := mockRepo.GetAll(ctx, 0, 10, "desc")

	require.NoError(t, err)
	assert.Len(t, companies, 2)
	assert.Equal(t, companyA.Name, companies[0].Name)
	assert.Equal(t, companyB.Name, companies[1].Name)
	mockRepo.AssertExpectations(t)
}

func TestCompanyRepository_GetByID(t *testing.T) {
	ctx := context.Background()
	newCompany := factories.CompanyFactory()
	mockRepo := new(mocks.MockCompanyRepository)

	mockRepo.On("GetByID", ctx, newCompany.CompanyID).Return(newCompany, nil)
	company, err := mockRepo.GetByID(ctx, newCompany.CompanyID)

	require.NoError(t, err)
	assert.NotNil(t, company)
	assert.Equal(t, newCompany.Name, company.Name)
	mockRepo.AssertExpectations(t)
}

func TestCompanyRepository_Create(t *testing.T) {
	ctx := context.Background()
	newCompany := factories.CompanyFactory()
	mockRepo := new(mocks.MockCompanyRepository)

	mockRepo.On("Create", ctx, newCompany).Return(nil)
	err := mockRepo.Create(ctx, newCompany)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCompanyRepository_Update(t *testing.T) {
	ctx := context.Background()
	company := factories.CompanyFactory()
	mockRepo := new(mocks.MockCompanyRepository)

	mockRepo.On("Update", ctx, company).Return(nil)
	err := mockRepo.Update(ctx, company)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
