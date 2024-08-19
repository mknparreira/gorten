package repositories_test

import (
	"context"
	"testing"

	"gorten/internal/gorten/models"
	"gorten/test/integration/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAll(t *testing.T) {
	ctx := context.Background()
	expectedUsers := []models.User{
		{UserID: "a0563613-4d29-4fb0-82ac-b22551eaef14", Name: "John Doe"},
		{UserID: "84114466-d801-4fd7-aef9-dc728da219b5", Name: "Jane Doe"},
	}
	mockRepo := new(mocks.MockUserRepository)

	mockRepo.On("GetAll", ctx).Return(expectedUsers, nil)
	users, err := mockRepo.GetAll(ctx)

	require.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "John Doe", users[0].Name)
	assert.Equal(t, "Jane Doe", users[1].Name)
	mockRepo.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(mocks.MockUserRepository)

	user := &models.User{UserID: "675ff6a9-516a-479a-b895-dcf89fd7654e", Name: "John Doe"}
	mockRepo.On("GetByID", ctx, "675ff6a9-516a-479a-b895-dcf89fd7654e").Return(user, nil)
	user, err := mockRepo.GetByID(ctx, "675ff6a9-516a-479a-b895-dcf89fd7654e")

	require.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "John Doe", user.Name)
	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(mocks.MockUserRepository)

	user := &models.User{UserID: "51a9cbf0-efe9-4331-ae74-8805f813c6e8", Name: "John Doe"}
	mockRepo.On("Create", ctx, user).Return(nil)
	err := mockRepo.Create(ctx, user)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(mocks.MockUserRepository)

	user := &models.User{UserID: "51a9cbf0-efe9-4331-ae74-8805f813c6e8", Name: "John Doe"}
	mockRepo.On("Update", ctx, user).Return(nil)
	err := mockRepo.Update(ctx, user)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
