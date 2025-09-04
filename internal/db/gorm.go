package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormConfig interface {
	DSN() string
}

func Open(cfg GormConfig) (*gorm.DB, error) {
	gcfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN()), gcfg)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(19)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("gorm connected to postgres")
	return db, nil
}
