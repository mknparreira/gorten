package services

import (
	"context"
	"gorten/internal/gorten/models"
	"gorten/internal/gorten/repositories"
	pkgerr "gorten/pkg/errors"
	"gorten/pkg/logs"
	"gorten/pkg/utils"
)

type CompanyServiceImpl interface {
	List(ctx context.Context, skip, limit int, sort string) ([]models.Company, error)
	GetByID(ctx context.Context, id string) (*models.Company, error)
	Create(ctx context.Context, company *models.Company) error
	UpdateByID(ctx context.Context, id string, company *models.Company) error
}

type CompanyService struct {
	companyRepo repositories.CompanyRepositoryImpl
}

func CompanyServiceInit(repo repositories.CompanyRepositoryImpl) *CompanyService {
	return &CompanyService{companyRepo: repo}
}

func (s *CompanyService) List(ctx context.Context, skip, limit int, sort string) ([]models.Company, error) {
	companies, err := s.companyRepo.GetAll(ctx, skip, limit, sort)

	if err != nil {
		logs.Logger.Printf("Error on CompanyService::List. Reason: %v", err)
		return nil, pkgerr.ErrInternalServerError
	}
	return companies, nil
}

func (s *CompanyService) GetByID(ctx context.Context, id string) (*models.Company, error) {
	company, err := s.companyRepo.GetByID(ctx, id)
	if err != nil {
		logs.Logger.Printf("Error on CompanyService::GetByID. Reason: %v", err)
		return nil, pkgerr.ErrCompanyNotFound
	}
	return company, nil
}

func (s *CompanyService) Create(ctx context.Context, company *models.Company) error {
	companyID, _ := utils.GenerateUUID()
	company.CompanyID = companyID

	err := s.companyRepo.Create(ctx, company)
	if err != nil {
		logs.Logger.Printf("Error on CompanyService::Create. Reason: %v", err)
		return pkgerr.ErrInternalServerError
	}
	return nil
}

func (s *CompanyService) UpdateByID(ctx context.Context, id string, company *models.Company) error {
	existingCompany, err := s.companyRepo.GetByID(ctx, id)
	if err != nil {
		logs.Logger.Printf("Error on retrieve GetByID into CompanyService::UpdateByID. Reason: %v", err)
		return pkgerr.ErrCompanyNotFound
	}

	existingCompany.Name = company.Name
	existingCompany.Address = company.Address
	existingCompany.Contact = company.Contact
	existingCompany.Email = company.Email

	if err := s.companyRepo.Update(ctx, existingCompany); err != nil {
		logs.Logger.Printf("Error on CompanyService::UpdateByID. Reason: %v", err)
		return pkgerr.ErrInternalServerError
	}
	return nil
}
