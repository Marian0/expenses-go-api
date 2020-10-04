package user

//UserDTO - Generic Data Transfer Object
type UserDTO struct {
	ID    uint   `json:"id,string,omitempty"`
	Name  string `json:"name" binding:"required,lte=100"`
	Email string `json:"email" binding:"required,email,lte=100"`
}

//UserCreateDTO - Create Data Transfer Object
type UserCreateDTO struct {
	Name  string `json:"name" binding:"required,lte=100"`
	Email string `json:"email" binding:"required,email,lte=100"`
}

//UserUpdateDTO - Update Data Transfer Object
type UserUpdateDTO struct {
	Name  string `json:"name" binding:"lte=100"`
	Email string `json:"email" binding:"omitempty,email,lte=100"`
}
