package model

import (
	"time"

	"github.com/coin50etf/coin-market/internal/pkg/constant"
	"github.com/shopspring/decimal"
)

type ProtocolPosition struct {
	ID              int64  `gorm:"primaryKey;autoIncrement"`
	Address         string `gorm:"size:128;index;not null"`
	ChainID         string `gorm:"size:64;index;not null"`
	ProtocolID      string `gorm:"size:64;index;not null"`
	ProtocolName    string `gorm:"size:128;not null"`
	PoolID          string `gorm:"size:64;index;not null"`
	PoolName        string `gorm:"size:128;not null"`
	PoolDescription string `gorm:"size:256;not null;default:''"`
	AssetTokens     string `gorm:"type:text;not null;default:''"`

	SyncTime  time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
}

func (s *ProtocolPosition) TableName() string {
	return constant.TableNameProtocolPosition
}

type UserToken struct {
	ID             int64            `gorm:"primaryKey;autoIncrement"`
	Address        string           `gorm:"size:128;index;not null"`
	ContractID     string           `gorm:"size:128;index;not null"`
	ChainID        string           `gorm:"size:64;index;not null"`
	Symbol         string           `gorm:"size:64;index;not null"`
	Decimals       int              `gorm:"not null"`
	LogoUrl        *string          `json:"logo_url"`
	Price          decimal.Decimal  `gorm:"type:decimal(40,30);index;not null;default:'0''"`
	Price24HChange *decimal.Decimal `gorm:"type:decimal(40,30);index;default:null"`
	TimeAt         float64          `gorm:"not null"`
	Amount         decimal.Decimal  `gorm:"type:decimal(40,30);index;not null;default:'0''"`
	SyncTime       time.Time        `gorm:"not null"`
	CreatedAt      time.Time        `gorm:"autoCreateTime;not null"`
	UpdatedAt      time.Time        `gorm:"autoUpdateTime;not null"`
}

func (s *UserToken) TableName() string {
	return constant.TableNameUserToken
}

type WalletAddress struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Address   string    `gorm:"size:128;index;not null;unique"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
}

func (s *WalletAddress) TableName() string {
	return constant.TableNameWalletAddress
}

type WalletAssetSnapshot struct {
	ID            int64     `gorm:"primaryKey;autoIncrement"`
	WalletAddress string    `gorm:"size:128;index"`
	SyncTime      time.Time `gorm:"not null"`

	TotalUSDValue decimal.Decimal `gorm:"type:decimal(40,30);index;not null;default:'0''"`

	CreatedAt time.Time
}

func (s *WalletAssetSnapshot) TableName() string {
	return constant.TableNameWalletAssetSnapshot
}
