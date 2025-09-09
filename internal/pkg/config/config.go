package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"

	"github.com/coin50etf/coin-market/internal/pkg/constant"
)

const (
	evnPrefix = "DEBANK"
)

var Conf *Config

type Config struct {
	AppName    string           `mapstructure:"app_name"`
	Env        string           `mapstructure:"env"`
	Server     ServerConfig     `mapstructure:"server"`
	PostgresDB DBConfig         `mapstructure:"postgres_db"`
	Log        LogConfig        `mapstructure:"log"`
	ThirdParty ThirdPartyConfig `mapstructure:"third_party"`
}

type ThirdPartyConfig struct {
	Debank DebankConfig `mapstructure:"debank"`
}

type DebankConfig struct {
	BaseUrl   string `mapstructure:"base_url"`
	AccessKey string `mapstructure:"access_key"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type DBConfig struct {
	DSN              string `mapstructure:"dsn"`
	MaxOpenConn      int    `mapstructure:"max_open_conn"`
	MaxIdleConn      int    `mapstructure:"max_idle_conn"`
	ConnMaxLifeInMin int    `mapstructure:"conn_max_life_in_min"`
	ConnMaxIdleInMin int    `mapstructure:"conn_max_idle_in_min"`
	EnableSQLLog     bool   `mapstructure:"enable_sql_log"`
	EnableMigration  bool   `mapstructure:"enable_migration"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	LogPath    string `mapstructure:"log_path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
}

// InitConfig 初始化配置
func InitConfig(path string) error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs" + path)
	v.AddConfigPath(".")

	v.SetEnvPrefix(evnPrefix)
	// 允许使用环境变量覆盖配置值
	v.AutomaticEnv()

	// 允许使用 "DATABASE_HOST" 等环境变量覆盖配置
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := bindEnvs(v); err != nil {
		return fmt.Errorf("unable to bind env vars: %w", err)
	}

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("error loading base config file: %w", err)
	}
	log.Println("Base config loaded from:", v.ConfigFileUsed())

	// 读取环境变量 `DEBANK_ENV` 确定加载哪个环境文件
	env := v.GetString("ENV") // 例如：`expotirt DEBANK_ENV = dev`
	if env == "" {
		env = constant.EvnDev
	}
	envConfigFile := fmt.Sprintf("config_%s.yaml", env)
	// 读取环境配置
	v.SetConfigName(envConfigFile) // 覆盖默认配置的值
	if err := v.MergeInConfig(); err != nil {
		log.Printf("No specific env config found for %s, using defaults\n", env)
		return fmt.Errorf("error loading env config file: %w", err)
	}
	log.Println("Loaded env config:", v.ConfigFileUsed())

	Conf = new(Config)
	if err := v.Unmarshal(Conf); err != nil {
		return fmt.Errorf("unable to decode into Config struct: %w", err)
	}
	Conf.Env = env

	return nil
}

func bindEnvs(v *viper.Viper) error {
	var err error
	err = v.BindEnv("postgres_db.dsn")
	err = v.BindEnv("third_party.debank.access_key")

	return err
}
