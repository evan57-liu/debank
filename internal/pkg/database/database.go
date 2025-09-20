package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/coin50etf/coin-market/internal/model"
	"github.com/coin50etf/coin-market/internal/pkg/config"
)

func init() {
	time.Local = time.UTC
}

type PostgresDB struct {
	DB *gorm.DB
}

// NewPostgresDB 连接 PostgresDB
func NewPostgresDB() (*PostgresDB, error) {
	dbConfig := config.Conf.PostgresDB

	gormConfig := &gorm.Config{
		PrepareStmt: true, // 使用预编译 SQL 语句，加速数据库操作
	}
	if dbConfig.EnableSQLLog {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dbConfig.DSN), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect TimescaleDB: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConn)                                      // 设置最大连接数
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConn)                                      // 设置最大空闲连接数
	sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifeInMin) * time.Minute) // 限制单个连接最大存活时间
	sqlDB.SetConnMaxIdleTime(time.Duration(dbConfig.ConnMaxIdleInMin) * time.Minute) // 限制空闲连接存活时间

	if dbConfig.EnableMigration {
		if err := db.AutoMigrate(&model.ProtocolPosition{}, &model.UserToken{}, &model.WalletAddress{}, &model.WalletAssetSnapshot{}, &model.Transaction{}); err != nil {
			return nil, err
		}
	}

	return &PostgresDB{DB: db}, nil
}

// Close 关闭数据库连接池
func (m *PostgresDB) Close() error {
	sqlDB, err := m.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
