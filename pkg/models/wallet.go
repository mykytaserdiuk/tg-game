package models

import (
	"time"
)

type Wallet struct {
	ID           uint      `gorm:"autoIncrement:true"`
	CreationTime time.Time `gorm:"autoCreateTime:true"`
	Ammount      int
	UserID       uint `gorm:"foreignKey:ID"`
	DeletionTime *time.Time
}
type User struct {
	ID     uint   `gorm:"autoIncrement:false"`
	Wallet Wallet `gorm:"foreignKey:ID"`
}
