package scheduler

import (
	"log"

	"github.com/robfig/cron/v3"

	"github.com/coin50etf/coin-market/internal/scheduler/jobs"
)

type Scheduler struct {
	cron           *cron.Cron
	protocolJob    *jobs.ProtocolJob
	transactionJob *jobs.TransactionJob
}

func NewScheduler(
	protocolJob *jobs.ProtocolJob,
	transactionJob *jobs.TransactionJob,
) *Scheduler {
	return &Scheduler{
		cron:           cron.New(cron.WithSeconds()),
		protocolJob:    protocolJob,
		transactionJob: transactionJob,
	}
}

// RegisterJobs 注册定时任务
func (s *Scheduler) RegisterJobs() {
	log.Println("Registering scheduled jobs...")
	// 每天的凌晨1点30分执行协议数据处理任务
	if _, err := s.cron.AddJob("0 30 1 * * ?", s.protocolJob); err != nil {
		log.Fatal("Failed to schedule symbol job:", err)
	}

	// 每周一的凌晨12点执行协议数据处理任务
	if _, err := s.cron.AddJob("0 0 0 * * 1", s.transactionJob); err != nil {
		log.Fatal("Failed to schedule symbol job:", err)
	}

	log.Println("All scheduled jobs are registered successfully.")
}

// Start 启动定时任务
func (s *Scheduler) Start() {
	log.Println("Starting scheduler...")
	s.RegisterJobs()
	s.cron.Start()
}

// Stop 停止所有任务
func (s *Scheduler) Stop() {
	log.Println("Stopping scheduler...")
	s.cron.Stop()
}
