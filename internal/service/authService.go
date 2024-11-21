package service

import (
	"simpleGinAPI/internal/models"
	"simpleGinAPI/internal/repository"
)

type IAuthService interface {
	Register(username string, password string, email string) (models.User, error)
	Login(username string, password string) (models.User, error)
	GetToken(username string) (string, error)
}

type AuthService struct {
	hashService IHashService
	userRepo    repository.IUserRepository
}

func (a AuthService) Register(username string, password string, email string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthService) Login(username string, password string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthService) GetToken(username string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func ProvideAuthService(hs IHashService, ur repository.IUserRepository) IAuthService {
	return &AuthService{
		hashService: hs,
		userRepo:    ur,
	}
}
