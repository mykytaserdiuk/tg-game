package mysql

import (
	"github.com/mykytaserdiuk/souptgbot/internal/soap"
	"github.com/mykytaserdiuk/souptgbot/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletRepo struct {
}

func NewWalletRepo() soap.WalletRepo {
	return &WalletRepo{}
}

func (r *WalletRepo) Insert(db *gorm.DB, wallet *models.Wallet) (*int, error) {

	res := db.Model(&models.Wallet{}).Clauses(clause.Returning{}).Create(&wallet)
	if res.Error != nil {
		return nil, res.Error
	}
	i := int(wallet.ID)
	return &i, nil
}
func (r *WalletRepo) Get(db *gorm.DB, walletID uint) (*models.Wallet, error) {
	var wallet models.Wallet
	res := db.Model(&models.Wallet{}).Take(&wallet).Where("id = ?", walletID)

	if res.Error != nil {
		return nil, res.Error
	}

	return &wallet, nil
}
func (r *WalletRepo) IsExists(db *gorm.DB, walletID uint) (bool, error) {
	var exists bool
	err := db.Model(models.Wallet{}).
		Select("count(*) > 0").
		Where("id = ?", walletID).
		Find(&exists).
		Error
	return exists, err
}
func (r *WalletRepo) GetAll(db *gorm.DB) ([]*models.Wallet, error) {
	var users []*models.Wallet
	err := db.Model(models.User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
