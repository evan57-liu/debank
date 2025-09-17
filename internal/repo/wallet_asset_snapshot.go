package repo

import (
	"github.com/coin50etf/coin-market/internal/model"
	"github.com/coin50etf/coin-market/internal/pkg/database"
)

type WalletAssetSnapshotRepository struct {
	BaseRepository[model.WalletAssetSnapshot]
}

func NewWalletAssetSnapshotRepository(postgresDB *database.PostgresDB) *WalletAssetSnapshotRepository {
	return &WalletAssetSnapshotRepository{
		BaseRepository: BaseRepository[model.WalletAssetSnapshot]{PostgresDB: postgresDB},
	}
}
