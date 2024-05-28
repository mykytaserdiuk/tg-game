package soap

import (
	"context"

	"github.com/mykytaserdiuk/souptgbot/pkg/models"
	"gorm.io/gorm"
)

type WalletService interface {
	Insert(ctx context.Context, userID string) (*int, error)
	Get(ctx context.Context, walletID string) (*models.Wallet, error)
	IsExists(ctx context.Context, walletID string) error
	Admin(ctx context.Context) ([]*models.Wallet, error)
}
type UserService interface {
	Get(ctx context.Context, userID string) (*models.User, error)
}

type WalletRepo interface {
	Insert(db *gorm.DB, wallet *models.Wallet) (*int, error)
	Get(db *gorm.DB, userID uint) (*models.Wallet, error)
	GetAll(db *gorm.DB) ([]*models.Wallet, error)
	IsExists(db *gorm.DB, walletID uint) (bool, error)
}
type UserRepo interface {
	Insert(db *gorm.DB, user *models.User) (*int, error)
	Get(db *gorm.DB, userID uint) (*models.User, error)
	IsExists(db *gorm.DB, userID uint) (bool, error)
}
