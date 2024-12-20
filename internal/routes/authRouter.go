package routes

import (
	"github.com/gin-gonic/gin"
	"simpleGinAPI/internal/controller"
)

type IRouterGroup interface {
	InitializedGroup(app *gin.Engine)
}

type AuthRouter struct {
	controller controller.IAuthController
}

func (r *AuthRouter) InitializedGroup(app *gin.Engine) {
	userAPI := app.Group("/auth")
	{
		userAPI.POST("/login", r.controller.Login)
		userAPI.POST("/register", r.controller.Register)
	}
}
