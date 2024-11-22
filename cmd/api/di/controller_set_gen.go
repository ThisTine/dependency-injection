package di

import (
	"github.com/google/wire"
	"simpleGinAPI/internal/controller"
)

var controllerSet = wire.NewSet(
	controller.ProvideAuthController,
)
