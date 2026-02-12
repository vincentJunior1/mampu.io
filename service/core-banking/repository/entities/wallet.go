package entities

import "time"

type Wallet struct {
	ID        int        `gorm:"column:id" json:"ID"`
	UserID    int        `gorm:"column:user_id" json:"userID"`
	Balance   float64    `gorm:"column:balance" json:"balance"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (Wallet) TableName() string {
	return "wallets"
}
