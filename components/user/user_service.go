package user

//UserService -
type UserService struct {
	UserRepository UserRepository
}

//NewUserService -
func NewUserService(p UserRepository) UserService {
	return UserService{UserRepository: p}
}

//FindAll -
func (p *UserService) FindAll() []User {
	return p.UserRepository.FindAll()
}

// FindByID -
func (p *UserService) FindByID(id uint) User {
	return p.UserRepository.FindByID(id)
}

// Create -
func (p *UserService) Create(user *User) *User {
	return p.UserRepository.Create(user)
}

// Save -
func (p *UserService) Save(user User) User {
	p.UserRepository.Save(user)

	return user
}

//Delete -
func (p *UserService) Delete(user User) {
	p.UserRepository.Delete(user)
}
