package service

import (
	"context"
	"strconv"

	"github.com/mykytaserdiuk/souptgbot/internal/soap"
	"github.com/mykytaserdiuk/souptgbot/pkg/models"
	"gorm.io/gorm"
)

type UserService struct {
	db         *gorm.DB
	walletRepo soap.WalletRepo
	userRepo   soap.UserRepo
}

func NewUserService(db *gorm.DB, walletRepo soap.WalletRepo, userRepo soap.UserRepo) soap.UserService {
	user := &UserService{db: db, walletRepo: walletRepo, userRepo: userRepo}
	return user
}

func (s *UserService) Get(ctx context.Context, userID string) (*models.User, error) {
	uuserID, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.Get(s.db, uint(uuserID))
	if err != nil {
		return nil, err
	}
	return user, nil
}
