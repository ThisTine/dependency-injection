package repository

import "simpleGinAPI/internal/models"

type IUserRepository interface {
	VerifyUniqueUsername(username string) (bool, error)
	GetUserByUsername(username string) (*models.User, error)
}
