package factories

import (
	"gorten/internal/gorten/models"
	"time"

	"github.com/google/uuid"
)

func CompanyFactory(overrides ...func(*models.Company)) *models.Company {
	company := &models.Company{
		CompanyID: uuid.New().String(),
		Name:      "New company",
		Address:   "Address",
		Contact:   "924439738",
		Email:     "company.email@gmail.com",
		CreatedAt: time.Now(),
	}

	for _, override := range overrides {
		override(company)
	}

	return company
}
