package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/coin50etf/coin-market/internal/pkg/app"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Ping(c *gin.Context) {
	app.Success(c, nil)
}
