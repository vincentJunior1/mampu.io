package repository

import (
	"context"
	"core-system/service/core-banking/repository/entities"

	"gorm.io/gorm"
)

type RepoInterface interface {
	GetBalanceByUserID(ctx context.Context, userID int) (entities.Wallet, error)
}

type repo struct {
	db *gorm.DB
}
