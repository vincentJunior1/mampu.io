package repository

import (
	"core-system/core/utils/getenv"
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

func InitRepo() RepoInterface {
	return &repo{
		db: connectDB(),
	}
}

// ConnectMysql connection ...
func connectDB() *gorm.DB {
	username := getenv.Getenv[string]("USER_DATABASE")
	password := getenv.Getenv[string]("PASS_DATABASE")
	host := getenv.Getenv[string]("HOST_DATABASE")
	port := getenv.Getenv[string]("PORT_DATABASE")
	dbName := getenv.Getenv[string]("DB_DATABASE")
	debug := getenv.Getenv[string]("DEBUG_DATABASE")
	mode := getenv.Getenv[string]("LOG_MODE_DATABASE")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbName)
	fmt.Println("dsn:", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logMode[mode]),
	})
	if err != nil {
		logs.WithFields(logs.Fields{"Message": err}).Error(runtime.GetCaller())
		panic("Error open database connection")
	}

	logs.Info("database connected successfully")
	if debug == "true" {
		return db.Debug()
	}

	return db
}
