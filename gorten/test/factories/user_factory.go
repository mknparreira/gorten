package factories

import (
	"gorten/internal/gorten/models"

	"github.com/google/uuid"
)

func UserFactory(overrides ...func(*models.User)) *models.User {
	user := &models.User{
		UserID:   uuid.New().String(),
		Name:     "John Grand Doe",
		Email:    "john.doe@example.com",
		Password: "123456",
	}

	for _, override := range overrides {
		override(user)
	}

	return user
}
