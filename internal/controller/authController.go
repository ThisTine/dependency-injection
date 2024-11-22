package controller

import "github.com/gin-gonic/gin"

type IAuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	GetToken(ctx *gin.Context) string
}

type AuthController struct {
}

func ProvideAuthController() IAuthController {
	return AuthController{}
}

func (a AuthController) Login(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a AuthController) Register(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a AuthController) GetToken(ctx *gin.Context) string {
	//TODO implement me
	panic("implement me")
}
