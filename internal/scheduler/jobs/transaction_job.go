package jobs

import (
	"context"
	"time"

	"github.com/coin50etf/coin-market/internal/pkg/logger"
	"github.com/coin50etf/coin-market/internal/pkg/utils/ctxutils"
	"github.com/coin50etf/coin-market/internal/service"
)

type TransactionJob struct {
	TransactionService *service.TransactionService
}

func NewTransactionJob(TransactionService *service.TransactionService) *TransactionJob {
	return &TransactionJob{TransactionService: TransactionService}
}

func (j *TransactionJob) Run() {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(context.Background(), "Transaction job panic", "err", r)
		}
	}()

	now := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()
	ctx = ctxutils.SetTraceID(ctx, "Transaction_job")
	ctx = ctxutils.SetUserID(ctx, "system")

	if err := j.TransactionService.ProcessTransaction(ctx); err != nil {
		logger.Error(ctx, "failed to process Transaction", "err", err)
	}

	logger.Warn(ctx, "Transaction job finished", "cost", time.Since(now))
}
