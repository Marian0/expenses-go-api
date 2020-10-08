package expense

import (
	"gorm.io/gorm"
)

//Repository -
type Repository struct {
	DB *gorm.DB
}

//NewRepository -
func NewRepository(DB *gorm.DB) *Repository {
	return &Repository{DB: DB}
}

//FindAll -
func (eRepo *Repository) FindAll() []Expense {
	var expenses []Expense
	eRepo.DB.Find(&expenses)

	return expenses
}

//FindByID -
func (eRepo *Repository) FindByID(id uint) Expense {
	var expense Expense
	eRepo.DB.First(&expense, id)

	return expense
}

//Create -
func (eRepo *Repository) Create(expense *Expense) *Expense {
	eRepo.DB.Create(expense)
	return expense
}

//Save -
func (eRepo *Repository) Save(expense Expense) Expense {
	eRepo.DB.Save(&expense)

	return expense
}

//Delete -
func (eRepo *Repository) Delete(expense Expense) {
	eRepo.DB.Delete(&expense)
}
