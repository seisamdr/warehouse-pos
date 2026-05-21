package database

import (
	"fmt"
	"micro-warehouse/user-service/config"
	"micro-warehouse/user-service/models"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func ConnectionPostgres(cfg config.Config) (*Postgres, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.SqlDB.User, cfg.SqlDB.Password, cfg.SqlDB.Host, cfg.SqlDB.Port, cfg.SqlDB.DBName)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Errorf("[Postgres] ConnectionPostgres - 1: %v", err)
		return nil, err
	}

	db.AutoMigrate(&models.User{}, &models.Role{}, &models.UserRole{})

	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("[Postgres] ConnectionPostgres - 2: %v", err)
		return nil, err
	}

	SeedRole(db)
	SeedManager(db)

	sqlDB.SetMaxIdleConns(cfg.SqlDB.DBMaxOpenConns)
	sqlDB.SetMaxOpenConns(cfg.SqlDB.DBMaxIdleConns)

	return &Postgres{DB: db}, nil
}
