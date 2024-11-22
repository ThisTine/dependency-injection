package di

import (
	"github.com/google/wire"
	"simpleGinAPI/internal/service"
)

var serviceSet = wire.NewSet(
	service.ProvideAuthService,
	service.ProvideHashService,
)
