package di

import (
	"github.com/gin-gonic/gin"
)

type IPillar interface {
	RunApp() error
}

type Pillar struct {
	app *gin.Engine
}

func (p *Pillar) RunApp() error {
	p.app.Use(gin.Recovery())
	p.app.Use(gin.Logger())

	return p.app.Run(":8000")
}

func ProvideGinEngine() IPillar {
	app := gin.New()

	//route.InitializedRouter(app)

	return &Pillar{
		app: app,
	}
}
