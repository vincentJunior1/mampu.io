package repository

import (
	"context"
	"core-system/service/core-banking/repository/entities"
	usecaseEntities "core-system/service/core-banking/usecase/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepoInterface interface {
	GetBalanceByUserID(ctx *gin.Context, userID int) (entities.Wallet, error)
	WithdrawByUserID(ctx context.Context, payload usecaseEntities.WithdrawRequest) (*entities.Wallet, error)
}

type repo struct {
	db *gorm.DB
}
