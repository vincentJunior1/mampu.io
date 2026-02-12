package repository

import (
	"core-system/core/utils/database/postgres"

	"gorm.io/gorm/logger"
)

var logMode = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"error":  logger.Error,
	"warn":   logger.Warn,
	"info":   logger.Info,
}

func InitRepo(cfg postgres.ConfigPostgres) RepoInterface {
	return &repo{
		db: postgres.ConnectPostgres(cfg),
	}
}
