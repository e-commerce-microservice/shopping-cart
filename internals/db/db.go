package db

import (
	"fmt"
	"time"

	"github.com/e-commerce-microservice/shopping-cart/internals/config"
	"github.com/e-commerce-microservice/shopping-cart/internals/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDB(config *config.Config) (*gorm.DB, error) {
	dial := postgres.Open(config.DSN)

	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	// Verify the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("db handle: %w", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}

	if err := db.AutoMigrate(
		&repo.Cart{},
		&repo.CartItem{},
	); err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}

	return db, nil
}
