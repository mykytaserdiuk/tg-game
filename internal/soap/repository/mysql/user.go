package mysql

import (
	"github.com/mykytaserdiuk/souptgbot/internal/soap"
	"github.com/mykytaserdiuk/souptgbot/pkg/models"
	"gorm.io/gorm"
)

type UserRepo struct {
}

func NewUserRepo() soap.UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) Insert(db *gorm.DB, user *models.User) (*int, error) {
	return nil, nil
}
func (r *UserRepo) Get(db *gorm.DB, userID uint) (*models.User, error) {
	return nil, nil
}
func (r *UserRepo) IsExists(db *gorm.DB, userID uint) (bool, error) {
	var exists bool
	err := db.Model(models.User{}).
		Select("count(*) > 0").
		Where("id = ?", userID).
		Find(&exists).
		Error

	return exists, err
}
