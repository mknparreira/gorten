package services

import (
	"context"
	"gorten/internal/gorten/models"
	"gorten/internal/gorten/repositories"
	pkgerr "gorten/pkg/errors"
	"gorten/pkg/logs"
	"gorten/pkg/utils"
)

type ProductServiceImpl interface {
	List(ctx context.Context, skip, limit int, sort string) ([]models.Product, error)
	GetByID(ctx context.Context, id string) (*models.Product, error)
	Create(ctx context.Context, product *models.Product) error
	UpdateByID(ctx context.Context, id string, product *models.Product) error
}

type ProductService struct {
	productRepo repositories.ProductRepositoryImpl
}

func ProductServiceInit(repo repositories.ProductRepositoryImpl) *ProductService {
	return &ProductService{productRepo: repo}
}

func (s *ProductService) List(ctx context.Context, skip, limit int, sort string) ([]models.Product, error) {
	products, err := s.productRepo.GetAll(ctx, skip, limit, sort)

	if err != nil {
		logs.Logger.Printf("Error on ProductService::List. Reason: %v", err)
		return nil, pkgerr.ErrInternalServerError
	}
	return products, nil
}

func (s *ProductService) GetByID(ctx context.Context, id string) (*models.Product, error) {
	product, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		logs.Logger.Printf("Error on ProductService::GetByID. Reason: %v", err)
		return nil, pkgerr.ErrProductNotFound
	}
	return product, nil
}

func (s *ProductService) Create(ctx context.Context, product *models.Product) error {
	productID, _ := utils.GenerateUUID()
	product.ProductID = productID

	err := s.productRepo.Create(ctx, product)
	if err != nil {
		logs.Logger.Printf("Error on ProductService::Create. Reason: %v", err)
		return pkgerr.ErrInternalServerError
	}
	return nil
}

func (s *ProductService) UpdateByID(ctx context.Context, id string, product *models.Product) error {
	existingProduct, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		logs.Logger.Printf("Error on retrieve GetByID into ProductService::UpdateByID. Reason: %v", err)
		return pkgerr.ErrProductNotFound
	}

	existingProduct.Name = product.Name
	existingProduct.Description = product.Description
	existingProduct.Price = product.Price
	existingProduct.CategoryID = product.CategoryID

	if err := s.productRepo.Update(ctx, existingProduct); err != nil {
		logs.Logger.Printf("Error on ProductService::UpdateByID. Reason: %v", err)
		return pkgerr.ErrInternalServerError
	}
	return nil
}
