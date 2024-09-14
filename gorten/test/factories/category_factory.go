package factories

import (
	"gorten/internal/gorten/models"
	"time"

	"github.com/google/uuid"
)

func CategoryFactory(overrides ...func(*models.Category)) *models.Category {
	category := &models.Category{
		CategoryID: uuid.New().String(),
		Name:       "Category",
		CreatedAt:  time.Now(),
	}

	for _, override := range overrides {
		override(category)
	}

	return category
}
