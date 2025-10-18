//go:build wireinject
// +build wireinject

package main

import (
	"t06/handler"
	"t06/repository"
	"t06/service"

	"github.com/google/wire"
)

func InitializeApp() handler.WarHandler {
	wire.Build(
		repository.ProviderSet,
		service.ProviderSet,
		handler.ProviderSet,
	)

	return handler.WarHandler{}
}
