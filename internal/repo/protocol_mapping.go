package repo

import (
	"github.com/coin50etf/coin-market/internal/model"
	"github.com/coin50etf/coin-market/internal/pkg/database"
)

type ProtocolMappingRepository struct {
	BaseRepository[model.ProtocolMapping]
}

func NewProtocolMappingRepository(postgresDB *database.PostgresDB) *ProtocolMappingRepository {
	return &ProtocolMappingRepository{
		BaseRepository: BaseRepository[model.ProtocolMapping]{PostgresDB: postgresDB},
	}
}
