package factories

import (
	"gorten/internal/gorten/models"
	"time"

	"github.com/google/uuid"
)

func ProductFactory(overrides ...func(*models.Product)) *models.Product {
	product := &models.Product{
		ProductID:   uuid.New().String(),
		Name:        "Product",
		Description: "Description",
		Price:       2.99,
		CategoryID:  CategoryFactory().CategoryID,
		CompanyID:   CompanyFactory().CompanyID,
		CreatedAt:   time.Now(),
	}

	for _, override := range overrides {
		override(product)
	}

	return product
}
