package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/marian0/expenses-go-api/models"
	"gorm.io/gorm"
)

var M202009120000 = gormigrate.Migration{
	ID: "202009120000",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&models.User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("users")
	},
}
