package database

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/marian0/expenses-go-api/database/migrations"
	"gorm.io/gorm"
)

//MigrateDatabase - Run the defined migrations
func MigrateDatabase(DB *gorm.DB) {

	m := gormigrate.New(DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		&migrations.M202009120000,
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
