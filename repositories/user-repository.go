package repositories

import (
	"github.com/marian0/expenses-go-api/models"
	"gorm.io/gorm"
)

//UserRepository -
type UserRepository struct {
	DB *gorm.DB
}

// User Repository

//GetAllUsers Fetch all user data
func (r *UserRepository) GetAllUsers(users *[]models.User) (err error) {
	if err = r.DB.Find(users).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func (r *UserRepository) CreateUser(user *models.User) (err error) {
	if err = r.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
