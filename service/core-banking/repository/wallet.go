package repository

import (
	"context"
	"core-system/service/core-banking/repository/entities"
)

func (r *repo) GetBalanceByUserID(ctx context.Context, userID int) (entities.Wallet, error) {
	var entity entities.Wallet
	query := r.db.WithContext(ctx).Model(&entity)
	query = query.Where("user_id = ?", userID)
	query.First(&entity)

	return entity, query.Error
}
