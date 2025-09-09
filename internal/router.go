package internal

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/coin50etf/coin-market/internal/handler"
	"github.com/coin50etf/coin-market/internal/pkg/middleware"
)

// RegisterRoutes 负责注册所有 API 路由
func RegisterRoutes(
	healthHandler *handler.HealthHandler,
	protocolHandler *handler.ProtocolHandler,
) *gin.Engine {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.BestSpeed))
	r.Use(middleware.ErrorHandlerMiddleware())

	cmGroup := r.Group("debank")
	apiGroup := cmGroup.Group("/api")
	v1Group := apiGroup.Group("/v1")
	v1Group.GET("/ping", healthHandler.Ping)

	symbolGroup := v1Group.Group("/protocols")
	{
		symbolGroup.POST("process-protocol", protocolHandler.ProcessProtocol)
		symbolGroup.POST("process-user-tokens", protocolHandler.ProcessUserTokens)
	}

	return r
}
