package repositories_test

import (
	"testing"

	"gorten/internal/gorten/models"
	"gorten/test/integration/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAll(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	mockRepo.On("GetAll").Return([]models.User{
		{UserID: "1", Name: "John Doe"},
		{UserID: "2", Name: "Jane Doe"},
	}, nil)
	users, err := mockRepo.GetAll()

	require.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "John Doe", users[0].Name)
	assert.Equal(t, "Jane Doe", users[1].Name)
	mockRepo.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	mockRepo.On("GetByID", "1").Return(&models.User{UserID: "1", Name: "John Doe"}, nil)
	user, err := mockRepo.GetByID("1")

	require.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "John Doe", user.Name)
	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	mockRepo.On("Create", &models.User{UserID: "1", Name: "John Doe"}).Return(nil)
	err := mockRepo.Create(&models.User{UserID: "1", Name: "John Doe"})

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	mockRepo.On("Update", &models.User{UserID: "1", Name: "John Doe"}).Return(nil)
	err := mockRepo.Update(&models.User{UserID: "1", Name: "John Doe"})

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
