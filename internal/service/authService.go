package service

import "simpleGinAPI/internal/models"

type IAuthService interface {
	Register(username string, password string, email string) models.User
	Login(username string, password string) models.User
	GetToken() string
}
