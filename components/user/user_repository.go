package user

import (
	"gorm.io/gorm"
)

//UserRepository -
type UserRepository struct {
	DB *gorm.DB
}

//NewUserRepository -
func NewUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

// FindAll -
func (p *UserRepository) FindAll() []User {
	var users []User
	p.DB.Find(&users)

	return users
}

//FindByID -
func (p *UserRepository) FindByID(id uint) User {
	var user User
	p.DB.First(&user, id)

	return user
}

//Create -
func (p *UserRepository) Create(user *User) *User {
	p.DB.Create(user)
	return user
}

//Save -
func (p *UserRepository) Save(user User) User {
	p.DB.Save(&user)

	return user
}

//Delete -
func (p *UserRepository) Delete(user User) {
	p.DB.Delete(&user)
}
