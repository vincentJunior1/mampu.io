package entities

import "time"

type User struct {
	ID        int       `gorm:"column:id" json:"ID"`
	Username  string    `gorm:"column:username" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (User) TableName() string {
	return "users"
}
