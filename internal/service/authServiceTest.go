package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"simpleGinAPI/internal/models"
	"simpleGinAPI/internal/repository/mocks"
	"simpleGinAPI/internal/service/mock"
)

func TestAuthService_GetToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockIUserRepository(ctrl)
	mockHashService := mock.NewMockIHashService(ctrl)

	authService := ProvideAuthService(mockHashService, mockUserRepo)

	t.Run("User exists and encoding is successful", func(t *testing.T) {
		username := "testuser"
		password := "testpassword"

		mockUser := &models.User{
			Username: username,
			Password: password,
		}

		mockUserRepo.EXPECT().GetUserByUsername(username).Return(mockUser, nil)
		mockHashService.EXPECT().Base64Encode(username).Return("encodedUsername", nil)

		token, err := authService.GetToken(username)

		assert.NoError(t, err)
		assert.NotNil(t, token)
		assert.Equal(t, "encodedUsername", token)
	})

	t.Run("User does not exist", func(t *testing.T) {
		username := "nonexistentuser"

		mockUserRepo.EXPECT().GetUserByUsername(username).Return(nil, nil)

		token, err := authService.GetToken(username)

		assert.Error(t, err)
		assert.Nil(t, token)
	})

	t.Run("Error in fetching user", func(t *testing.T) {
		username := "testuser"

		mockUserRepo.EXPECT().GetUserByUsername(username).Return(nil, errors.New("database error"))

		token, err := authService.GetToken(username)

		assert.Error(t, err)
		assert.Nil(t, token)
	})

	t.Run("Error in encoding username", func(t *testing.T) {
		username := "testuser"
		password := "testpassword"

		mockUser := &models.User{
			Username: username,
			Password: password,
		}

		mockUserRepo.EXPECT().GetUserByUsername(username).Return(mockUser, nil)
		mockHashService.EXPECT().Base64Encode(username).Return("", errors.New("encoding error"))

		token, err := authService.GetToken(username)

		assert.Error(t, err)
		assert.Nil(t, token)
	})
}
