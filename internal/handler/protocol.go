package handler

import (
	"github.com/coin50etf/coin-market/internal/pkg/app"
	"github.com/gin-gonic/gin"

	"github.com/coin50etf/coin-market/internal/service"
)

type ProtocolHandler struct {
	service *service.ProtocolService
}

func NewProtocolHandler(svc *service.ProtocolService) *ProtocolHandler {
	return &ProtocolHandler{service: svc}
}

func (s *ProtocolHandler) ProcessProtocol(c *gin.Context) {
	if err := s.service.ProcessProtocol(c.Request.Context()); err != nil {
		_ = c.Error(err)
		return

	}

	app.Success(c, nil)
}

func (s *ProtocolHandler) ProcessUserTokens(c *gin.Context) {
	if err := s.service.ProcessUserTokens(c.Request.Context()); err != nil {
		_ = c.Error(err)
		return

	}

	app.Success(c, nil)
}
