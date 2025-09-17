package model

import (
	"fmt"
	"time"

	"database/sql/driver"

	"github.com/coin50etf/coin-market/internal/pkg/constant"
	"github.com/coin50etf/coin-market/internal/pkg/json"
	"github.com/shopspring/decimal"
)

type ProtocolPools []*ProtocolPool

func (p ProtocolPools) Value() (driver.Value, error) {
	if p == nil {
		return "[]", nil
	}
	return json.Marshal(p)
}

func (p *ProtocolPools) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal ProtocolPools: value is not []byte")
	}
	return json.Unmarshal(bytes, p)
}

type ProtocolPosition struct {
	ID                   int64     `gorm:"primaryKey;autoIncrement"`
	InternalProtocolName string    `gorm:"size:128;index;not null"`
	InternalProtocolID   string    `gorm:"size:128;index;not null"`
	Address              string    `gorm:"size:128;index;not null"`
	ChainID              string    `gorm:"size:64;index;not null"`
	ProtocolID           string    `gorm:"size:64;index;not null"`
	PoolID               string    `gorm:"size:64;index;not null"`
	CustomID             string    `gorm:"size:1024;not null"`
	AssetTokens          string    `gorm:"type:text;not null;default:''"`
	SyncTime             time.Time `gorm:"not null"`
	CreatedAt            time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime;not null"`
}

func (s *ProtocolPosition) TableName() string {
	return constant.TableNameProtocolPosition
}

type ProtocolMapping struct {
	ID                   int64         `gorm:"primaryKey;autoIncrement"`
	InternalProtocolName string        `gorm:"size:128;index;not null"`
	InternalProtocolID   string        `gorm:"size:64;unique;not null"`
	ProtocolPools        ProtocolPools `gorm:"type:json" json:"protocol_pools"`
	CreatedAt            time.Time     `gorm:"autoCreateTime;not null"`
	UpdatedAt            time.Time     `gorm:"autoUpdateTime;not null"`
}

func (s *ProtocolMapping) TableName() string {
	return constant.TableNameProtocolMapping
}

type ProtocolPool struct {
	Address    string `json:"address"`
	ChainID    string `json:"chain_id"`
	ProtocolID string `json:"protocol_id"`
	PoolID     string `json:"pool_id"`
	Status     int    `json:"status"`
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
