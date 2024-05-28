package service

import (
	"context"
	"strconv"

	"github.com/mykytaserdiuk/souptgbot/internal/soap"
	"github.com/mykytaserdiuk/souptgbot/pkg/models"
	"gorm.io/gorm"
)

type WalletService struct {
	db         *gorm.DB
	walletRepo soap.WalletRepo
	userRepo   soap.UserRepo
}

func NewWalletService(db *gorm.DB, walletRepo soap.WalletRepo, userRepo soap.UserRepo) soap.WalletService {
	wallet := &WalletService{db: db, walletRepo: walletRepo, userRepo: userRepo}
	return wallet
}
func (s *WalletService) Insert(ctx context.Context, userID string) (*int, error) {
	uuserID, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return nil, err
	}
	ok, err := s.userRepo.IsExists(s.db, uint(uuserID))
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, models.ErrorWalletNotFound
	}

	wallet := models.Wallet{
		UserID:  uint(uuserID),
		Ammount: 0,
	}
	id, err := s.walletRepo.Insert(s.db, &wallet)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *WalletService) Get(ctx context.Context, walletID string) (*models.Wallet, error) {
	uwalletID, err := strconv.ParseUint(walletID, 10, 32)
	if err != nil {
		return nil, err
	}
	ok, err := s.walletRepo.IsExists(s.db, uint(uwalletID))
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, models.ErrorWalletNotFound
	}

	wallet, err := s.walletRepo.Get(s.db, uint(uwalletID))
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (s *WalletService) IsExists(ctx context.Context, walletID string) error {
	uwalletID, err := strconv.ParseUint(walletID, 10, 32)
	if err != nil {
		return err
	}
	ok, err := s.walletRepo.IsExists(s.db, uint(uwalletID))
	if err != nil {
		return err
	} else if !ok {
		return models.ErrorWalletNotFound
	}

	return nil
}

func (s *WalletService) Admin(ctx context.Context) ([]*models.Wallet, error) {

	wallets, err := s.walletRepo.GetAll(s.db)
	if err != nil {
		return nil, err
	}
	return wallets, nil
}
