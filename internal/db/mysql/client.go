package mysql

import (
	"fmt"
	"log"

	"github.com/mykytaserdiuk/souptgbot/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewPool(cfg *Config) (*gorm.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	log.Printf("%s", connStr)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
		DSN:                       connStr,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}
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
