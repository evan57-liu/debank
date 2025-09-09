//go:build wireinject
// +build wireinject

package main

import (
	"log"

	"github.com/google/wire"

	"github.com/coin50etf/coin-market/internal/handler"
	"github.com/coin50etf/coin-market/internal/pkg"
	"github.com/coin50etf/coin-market/internal/pkg/database"
	"github.com/coin50etf/coin-market/internal/repo"
	"github.com/coin50etf/coin-market/internal/scheduler"
	"github.com/coin50etf/coin-market/internal/service"
)

var providerSet = wire.NewSet(
	pkg.ProviderSet,
	repo.ProviderSet,
	service.ProviderSet,
	handler.ProviderSet,
	wire.NewSet(registerCleanup),
	scheduler.ProviderSet,
)

// initScheduler 使用 wire 进行依赖注入
func initScheduler() (*scheduler.Scheduler, func(), error) {
	wire.Build(providerSet)

	return nil, nil, nil
}

func registerCleanup(postgresDB *database.PostgresDB) func() {
	return func() {
		_ = postgresDB.Close()

		log.Println("Closed mysql, timescale and redis connections")
	}
}
