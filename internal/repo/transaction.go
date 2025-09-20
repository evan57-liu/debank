package repo

import (
	"github.com/coin50etf/coin-market/internal/model"
	"github.com/coin50etf/coin-market/internal/pkg/database"
)

type TransactionRepository struct {
	BaseRepository[model.Transaction]
}

func NewTransactionRepository(postgresDB *database.PostgresDB) *TransactionRepository {
	return &TransactionRepository{
		BaseRepository: BaseRepository[model.Transaction]{PostgresDB: postgresDB},
	}
}
