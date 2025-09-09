package pkg

import (
	"github.com/coin50etf/coin-market/internal/pkg/database"
	"github.com/coin50etf/coin-market/internal/pkg/third_party/debank"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	database.NewPostgresDB,
	debank.NewClient,
)
