//go:build wireinject
// +build wireinject

package repository

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewWarRepository,
)
