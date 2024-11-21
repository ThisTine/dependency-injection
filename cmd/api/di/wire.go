//go:build wireinject
// +build wireinject

package di

import "github.com/google/wire"

//go:generate go run -mod=mod github.com/google/wire/cmd/wire

func InitializedApp() (IPillar, error) {
	wire.Build(
		//serviceSet,
		appSet,
	)

	return &Pillar{}, nil
}
