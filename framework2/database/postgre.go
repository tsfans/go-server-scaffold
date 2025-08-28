package database

import (
	"fmt"
	"time"

	"github.com/tsfans/go/framework/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

type PostgreDBConfig struct {
	Host          string `json:"host"`
	Port          int    `json:"port"`
	User          string `json:"user"`
	Password      string `json:"password"`
	DB            string `json:"db"`
	SslMode       string `json:"ssl_mode"`
	Timezone      string `json:"timezone"`
	SlowThreshold int    `json:"slow_threshold"`
	MaxIdleCount  int    `json:"max_idle_count"`
	MaxOpen       int    `json:"max_open"`
	MaxLifetime   int    `json:"max_lifetime"`
	MaxIdleTime   int    `json:"max_idle_time"`
}

func initPostgreDB() {
	log.Debug("initializing postgre ...")
	cfg := PostgreDBConfig{
		Host:          config.GetString("postgre.host", "127.0.0.1"),
		Port:          config.GetInt("postgre.port", 5432),
		User:          config.GetString("postgre.user", "root"),
		Password:      config.GetString("postgre.password", "123456"),
		DB:            config.GetString("postgre.db", "test"),
		SslMode:       config.GetString("postgre.ssl_mode", "disable"),
		Timezone:      config.GetString("postgre.timezone", "Asia/Shanghai"),
		SlowThreshold: config.GetInt("postgre.slow_threshold", 100),
		MaxIdleCount:  config.GetInt("postgre.max_idle_count", 10),
		MaxOpen:       config.GetInt("postgre.max_open", 100),
		MaxLifetime:   config.GetInt("postgre.max_lifetime", 3600),
		MaxIdleTime:   config.GetInt("postgre.max_idle_time", 300),
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DB, cfg.SslMode, cfg.Timezone,
	)

	// 自定义日志，打印成业务日志格式
	newLogger := logger.New(
		log,
		logger.Config{
			SlowThreshold: time.Duration(cfg.SlowThreshold) * time.Millisecond,
			LogLevel:      logger.Info,
			Colorful:      false,
		},
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Errorf("connect postgre error: %s", err)
		return
	}

	// 设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Errorf("get postgre connection error: %s", err)
		return
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleCount)
	sqlDB.SetMaxOpenConns(cfg.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(cfg.MaxIdleTime) * time.Second)
}
