package seeder

import (
	"core-system/service/core-banking/repository/entities"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	users := []entities.User{
		{Username: "admin", Password: string(password)},
		{Username: "vincent", Password: string(password)},
	}

	for _, user := range users {
		if err := db.FirstOrCreate(&user, entities.User{Username: user.Username}).Error; err != nil {
			return err
		}

		wallet := entities.Wallet{
			UserID:  user.ID,
			Balance: 1000000,
		}

		if err := db.FirstOrCreate(&wallet, entities.Wallet{UserID: user.ID}).Error; err != nil {
			return err
		}
	}

	return nil
}
