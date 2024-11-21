package repository

import "simpleGinAPI/internal/models"

//go:generate mockgen -destination=mocks/mockUserRepository.go -package=mocks simpleGinAPI/internal/repository IUserRepository
type IUserRepository interface {
	VerifyUniqueUsername(username string) (bool, error)
	GetUserByUsername(username string) (*models.User, error)
}
