package models

import (
	"gorm.io/gorm"
)

//Expense - Model
type Expense struct {
	gorm.Model
	Date    string `json:"date"`
	Details string `json:"details"`
	Amount  int    `json:"amount"`
	UserID  int    `json:"user_id"`
	User    User
}
