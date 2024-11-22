package di

import (
	"github.com/google/wire"
	"simpleGinAPI/internal/routes"
)

var routesSet = wire.NewSet(
	routes.ProvideAuthRouter,
	routes.ProvideRouter,
)
