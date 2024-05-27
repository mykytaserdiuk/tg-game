package postgres

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewPool(cfg *Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = migrate(db)
	if err != nil {
		return nil, err
	}
	return db, err
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate()
	return err
}
