package main

import (
	"core-system/core/utils/database/postgres"
	"core-system/core/utils/getenv"
	"core-system/migration/seeder"
	"log"

	logs "github.com/sirupsen/logrus"
)

func main() {
	if err := getenv.LoadEnv(".env"); err != nil {
		logs.Fatal(err)
	}
	repo := postgres.ConnectPostgres(postgres.ConfigPostgres{
		Username: getenv.Getenv[string]("USER_DATABASE"),
		Password: getenv.Getenv[string]("PASS_DATABASE"),
		Host:     getenv.Getenv[string]("HOST_DATABASE"),
		Port:     getenv.Getenv[string]("PORT_DATABASE"),
		DbName:   getenv.Getenv[string]("DB_DATABASE"),
		Debug:    getenv.Getenv[string]("DEBUG_DATABASE"),
		Mode:     getenv.Getenv[string]("LOG_MODE_DATABASE"),
	})

	if err := seeder.Seed(repo); err != nil {
		log.Fatalf("seeding failed: %+v", err)
	}

	log.Println("seeding success")
}
