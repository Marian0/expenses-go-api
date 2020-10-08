package expense

import (
	"github.com/marian0/expenses-go-api/components/user"

	"gorm.io/gorm"
)

//Expense - Model
type Expense struct {
	gorm.Model
	Date    string `json:"date"`
	Details string `json:"details"`
	Amount  int    `json:"amount"`
	UserID  int    `json:"user_id"`
	User    user.User
}
