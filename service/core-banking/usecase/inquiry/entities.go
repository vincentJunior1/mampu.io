package usecase

import (
	"core-system/service/core-banking/repository"
)

type UsecaseInterface interface {
}

type UsecaseConfig struct {
	Repo repository.RepoInterface
}

type usecase struct {
	repo repository.RepoInterface
}
