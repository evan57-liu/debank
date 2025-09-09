package repo

import (
	"github.com/coin50etf/coin-market/internal/model"
	"github.com/coin50etf/coin-market/internal/pkg/database"
)

type ProtocolPositionRepository struct {
	BaseRepository[model.ProtocolPosition]
}

func NewProtocolPositionRepository(postgresDB *database.PostgresDB) *ProtocolPositionRepository {
	return &ProtocolPositionRepository{
		BaseRepository: BaseRepository[model.ProtocolPosition]{PostgresDB: postgresDB},
	}
}
