package jobs

import (
	"context"
	"time"

	"github.com/coin50etf/coin-market/internal/pkg/logger"
	"github.com/coin50etf/coin-market/internal/pkg/utils/ctxutils"
	"github.com/coin50etf/coin-market/internal/service"
)

type ProtocolJob struct {
	protocolService *service.ProtocolService
}

func NewProtocolJob(protocolService *service.ProtocolService) *ProtocolJob {
	return &ProtocolJob{protocolService: protocolService}
}

func (j *ProtocolJob) Run() {
	defer func() {
		if r := recover(); r != nil {
			logger.Error(context.Background(), "protocol job panic", "err", r)
		}
	}()

	now := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	ctx = ctxutils.SetTraceID(ctx, "protocol_job")
	ctx = ctxutils.SetUserID(ctx, "system")

	if err := j.protocolService.ProcessUserTokens(ctx); err != nil {
		logger.Error(ctx, "failed to process user tokens", "err", err)
	}
	if err := j.protocolService.ProcessProtocol(ctx); err != nil {
		logger.Error(ctx, "failed to process protocol", "err", err)
	}

	logger.Warn(ctx, "protocol job finished", "cost", time.Since(now))
}
