package services

import (
	"context"
	"gorten/internal/gorten/models"
	"gorten/internal/gorten/repositories"
	pkgerr "gorten/pkg/errors"
	"gorten/pkg/logs"
	"gorten/pkg/utils"
)

type CategoryServiceImpl interface {
	List(ctx context.Context, skip, limit int, sort string) ([]models.Category, error)
	GetByID(ctx context.Context, id string) (*models.Category, error)
	Create(ctx context.Context, category *models.Category) error
	UpdateByID(ctx context.Context, id string, category *models.Category) error
}

type CategoryService struct {
	categoryRepo repositories.CategoryRepositoryImpl
}

func CategoryServiceInit(repo repositories.CategoryRepositoryImpl) *CategoryService {
	return &CategoryService{categoryRepo: repo}
}

func (s *CategoryService) List(ctx context.Context, skip, limit int, sort string) ([]models.Category, error) {
	categories, err := s.categoryRepo.GetAll(ctx, skip, limit, sort)

	if err != nil {
		logs.Logger.Printf("Error on CategoryService::List. Reason: %v", err)
		return nil, pkgerr.ErrInternalServerError
	}
	return categories, nil
}

func (s *CategoryService) GetByID(ctx context.Context, id string) (*models.Category, error) {
	category, err := s.categoryRepo.GetByID(ctx, id)
	if err != nil {
		logs.Logger.Printf("Error on CategoryService::GetByID. Reason: %v", err)
		return nil, pkgerr.ErrCategoryNotFound
	}
	return category, nil
}

func (s *CategoryService) Create(ctx context.Context, category *models.Category) error {
	categoryID, _ := utils.GenerateUUID()
	category.CategoryID = categoryID

	err := s.categoryRepo.Create(ctx, category)
	if err != nil {
		logs.Logger.Printf("Error on CategoryService::Create. Reason: %v", err)
		return pkgerr.ErrInternalServerError
	}
	return nil
}

func (s *CategoryService) UpdateByID(ctx context.Context, id string, category *models.Category) error {
	existingCategory, err := s.categoryRepo.GetByID(ctx, id)
	if err != nil {
		logs.Logger.Printf("Error on retrieve GetByID into CategoryService::UpdateByID. Reason: %v", err)
		return pkgerr.ErrCategoryNotFound
	}

	existingCategory.Name = category.Name

	if err := s.categoryRepo.Update(ctx, existingCategory); err != nil {
		logs.Logger.Printf("Error on CategoryService::UpdateByID. Reason: %v", err)
		return pkgerr.ErrInternalServerError
	}
	return nil
}
