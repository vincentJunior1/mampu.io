package repository

import (
	"context"
	"core-system/service/core-banking/repository/entities"
	usecaseEntities "core-system/service/core-banking/usecase/entities"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func (r *repo) GetBalanceByUserID(ctx *gin.Context, userID int) (entities.Wallet, error) {
	var entity entities.Wallet
	query := r.db.WithContext(ctx).Model(entity)
	query = query.Where("user_id = ?", userID)
	query.First(&entity)

	return entity, query.Error
}

func (r *repo) WithdrawByUserID(ctx context.Context, payload usecaseEntities.WithdrawRequest) (*entities.Wallet, error) {
	var entity entities.Wallet

	tx := r.db.Begin()

	if err := tx.WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).Where("user_id = ?", payload.UserID).First(&entity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if entity.Balance < payload.Amount {
		tx.Rollback()
		return nil, errors.New("Insufficient Balance")
	}
	entity.Balance -= payload.Amount

	if err := tx.WithContext(ctx).Save(&entity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &entity, nil
}
