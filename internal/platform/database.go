package platform

import (
	"context"
	"fmt"

	"github.com/icakces/order-api/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(ctx context.Context, cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	_ = ctx
	return db, nil
}
