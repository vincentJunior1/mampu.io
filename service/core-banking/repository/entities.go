package repository

import "gorm.io/gorm"

type RepoInterface interface {
}

type repo struct {
	db *gorm.DB
}
