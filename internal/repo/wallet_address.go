package repo

import (
	"github.com/coin50etf/coin-market/internal/model"
	"github.com/coin50etf/coin-market/internal/pkg/database"
)

type WalletAddressRepository struct {
	BaseRepository[model.WalletAddress]
}

func NewWalletAddressRepository(postgresDB *database.PostgresDB) *WalletAddressRepository {
	return &WalletAddressRepository{
		BaseRepository: BaseRepository[model.WalletAddress]{PostgresDB: postgresDB},
	}
}
