package database

import (
	"fmt"
	"log"
	"time"

	"github.com/dendik-creation/task-control/internal/config"
	"github.com/dendik-creation/task-control/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Println("PostgreSQL connected successfully")

	err = db.AutoMigrate(
		&domain.User{},
		&domain.Column{},
		&domain.Task{},
	)
	if err != nil {
		log.Fatal("Failed run migration:", err)
	}
	log.Println("Successfull run migration")

	return db, nil
}
