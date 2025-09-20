package model

import (
	"time"

	"github.com/coin50etf/coin-market/internal/pkg/constant"
)

type Transaction struct {
	ID            int64  `gorm:"primaryKey;autoIncrement"`
	WalletAddress string `gorm:"size:128;index"`
	ChainID       string `gorm:"size:64;index;not null"`
	TxHash        string `gorm:"size:128;index;not null;unique"`
	Name          string `gorm:"size:128;index"`
	FromAddress   string `gorm:"size:128;index;not null"`
	ToAddress     string `gorm:"size:128;index;not null"`
	Detail        string `gorm:"type:text;not null;default:''"`

	SyncTime  time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
}

func (s *Transaction) TableName() string {
	return constant.TableNameTransaction
}
