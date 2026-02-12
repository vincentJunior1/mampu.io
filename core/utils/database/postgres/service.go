package postgres

import (
	"core-system/core/utils/runtime"
	"fmt"

	logs "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var logMode = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"error":  logger.Error,
	"warn":   logger.Warn,
	"info":   logger.Info,
}

// ConnectPostgres connection ...
func ConnectPostgres(cfg ConfigPostgres) *gorm.DB {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	fmt.Println("dsn:", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logMode[cfg.Mode]),
	})
	if err != nil {
		logs.WithFields(logs.Fields{"Message": err}).Error(runtime.GetCaller())
		panic("Error open database connection")
	}

	logs.Info("database connected successfully")
	if cfg.Debug == "true" {
		return db.Debug()
	}

	return db
}
