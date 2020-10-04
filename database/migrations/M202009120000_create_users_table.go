package migrations

import (
	"os/user"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

//M202009120000 -
var M202009120000 = gormigrate.Migration{
	ID: "202009120000",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&user.User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("users")
	},
}
