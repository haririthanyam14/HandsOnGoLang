package db

import (
	// "fmt"
	"HandsOnGoLang/pkg/config"
	// "payment-svc/pkg/liberror"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBHandler interface {
	GetDB() (*gorm.DB, error)
}

type gormDBHandler struct {
	config config.DBConfig
}

func (dbHandler *gormDBHandler) GetDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbHandler.config.Address()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewDBHandler(config config.DBConfig) DBHandler {
	return &gormDBHandler{
		config: config,
	}
}
