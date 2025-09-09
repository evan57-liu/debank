package repo

import (
	"github.com/coin50etf/coin-market/internal/model"
	"github.com/coin50etf/coin-market/internal/pkg/database"
)

type UserTokenRepository struct {
	BaseRepository[model.UserToken]
}

func NewUserTokenRepository(postgresDB *database.PostgresDB) *UserTokenRepository {
	return &UserTokenRepository{
		BaseRepository: BaseRepository[model.UserToken]{PostgresDB: postgresDB},
	}
}
