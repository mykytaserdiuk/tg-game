package mysql

import (
	"log"
	"time"

	"github.com/mykytaserdiuk/souptgbot/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPool(cfg *Config) (*gorm.DB, error) {
	log.Printf("%s", cfg.DBurl)
	db, err := gorm.Open(postgres.Open(cfg.DBurl), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxIdleTime(5 * time.Second)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	sqlDB.SetMaxIdleConns(0)
	sqlDB.SetMaxOpenConns(100)
	err = migrate(db)
	if err != nil {
		return nil, err
	}
	return db, err
}

func migrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&models.Wallet{})
	if err != nil {
		return
	}

	return
}
