package handler

import (
	"github.com/coin50etf/coin-market/internal/pkg/app"
	"github.com/gin-gonic/gin"

	"github.com/coin50etf/coin-market/internal/service"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(svc *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: svc}
}

func (s *TransactionHandler) ProcessTransaction(c *gin.Context) {
	if err := s.service.ProcessTransaction(c.Request.Context()); err != nil {
		_ = c.Error(err)
		return

	}

	app.Success(c, nil)
}
