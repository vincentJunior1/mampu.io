package repository

import (
	"core-system/core/utils/database/postgres"
)

func InitRepo(cfg postgres.ConfigPostgres) RepoInterface {
	return &repo{
		db: postgres.ConnectPostgres(cfg),
	}
}
