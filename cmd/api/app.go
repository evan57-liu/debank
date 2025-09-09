//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/coin50etf/coin-market/internal"
	"github.com/coin50etf/coin-market/internal/handler"
	"github.com/coin50etf/coin-market/internal/pkg"
	"github.com/coin50etf/coin-market/internal/pkg/database"
	"github.com/coin50etf/coin-market/internal/repo"
	"github.com/coin50etf/coin-market/internal/service"
)

var providerSet = wire.NewSet(
	pkg.ProviderSet,
	repo.ProviderSet,
	service.ProviderSet,
	handler.ProviderSet,
	wire.NewSet(registerCleanup),
	internal.RegisterRoutes,
)

// initApp 使用 wire 进行依赖注入
func initApp() (*gin.Engine, func(), error) {
	wire.Build(providerSet)

	return nil, nil, nil
}

func registerCleanup(postgresDB *database.PostgresDB) func() {
	return func() {
		_ = postgresDB.Close()
	}
}
