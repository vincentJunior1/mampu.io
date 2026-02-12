package repository

import (
	"core-system/migration/seeder"

	logs "github.com/sirupsen/logrus"
)

func (r *repo) SeedData_Test() {
	if err := seeder.Seed(r.db); err != nil {
		logs.Fatalf("[!] Error Seeding: %+v", err)
	}
}
