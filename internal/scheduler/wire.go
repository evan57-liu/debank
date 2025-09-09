package scheduler

import (
	"github.com/coin50etf/coin-market/internal/scheduler/jobs"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	jobs.NewProtocolJob,
	NewScheduler,
)
