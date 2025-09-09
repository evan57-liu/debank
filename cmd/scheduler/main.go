package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "ariga.io/atlas-provider-gorm/gormschema"

	"github.com/coin50etf/coin-market/internal/pkg/config"
	"github.com/coin50etf/coin-market/internal/pkg/logger"
)

func main() {
	log.Println("Starting market-scheduler...")

	if err := config.InitConfig("/scheduler"); err != nil {
		log.Fatalf("init config failed: %v", err)
	}
	logger.InitLogger()

	s, cleanup, err := initScheduler()
	if err != nil {
		log.Fatalf("init scheduler failed: %v", err)
	}
	defer cleanup()
	s.Start()

	// 优雅退出
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	<-sigs

	s.Stop()
	log.Println("market-scheduler stopped")
}
