package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/marian0/expenses-go-api/components/expense"
	"gorm.io/gorm"
)

//M202009120001 -
var M202009120001 = gormigrate.Migration{
	ID: "202009120001",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&expense.Expense{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("expenses")
	},
}
