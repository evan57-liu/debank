package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "ariga.io/atlas-provider-gorm/gormschema"

	"github.com/coin50etf/coin-market/internal/pkg/config"
	"github.com/coin50etf/coin-market/internal/pkg/logger"
)

func main() {
	if err := config.InitConfig("/api"); err != nil {
		log.Fatalf("init config failed: %v", err)
	}
	logger.InitLogger()

	ctx := context.Background()
	r, cleanup, err := initApp()
	if err != nil {
		logger.Fatal(ctx, "init app failed", "error", err)
	}
	defer cleanup()

	logger.Debug(ctx, "config", "config", config.Conf)

	// 创建 http.Server
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Conf.Server.Host, config.Conf.Server.Port),
		Handler: r,
	}

	// 启动服务器（在 goroutine 中运行，不会阻塞 main）
	go func() {
		logger.Info(ctx, "server is running", "host", config.Conf.Server.Host, "port", config.Conf.Server.Port)

		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(ctx, "server is down", "error", err)
		}
	}()

	// 监听系统信号（例如 Ctrl+C、终止信号等）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit // 阻塞等待信号

	logger.Info(ctx, "shutting down server")

	// 5 秒超时上下文，确保正在处理的请求能有时间完成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅地关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal(ctx, "server forced to shutdown", "error", err)
	}

	logger.Info(ctx, "server exited")
}
