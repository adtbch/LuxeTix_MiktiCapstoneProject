package database

import (
	"fmt"
	"luxe/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase(pgConfig config.PGConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pgConfig.Host,
		pgConfig.Port,
		pgConfig.User,
		pgConfig.Password,
		pgConfig.Database,
		pgConfig.SSLMode,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}