package user

//ToUser -
func ToUser(userDTO UserDTO) *User {
	return &User{Name: userDTO.Name, Email: userDTO.Email}
}

//ToUserDTO -
func ToUserDTO(user *User) UserDTO {
	return UserDTO{ID: user.ID, Name: user.Name, Email: user.Email}
}

//ToUserDTOs -
func ToUserDTOs(users []User) []UserDTO {
	userdtos := make([]UserDTO, len(users))

	for i, itm := range users {
		userdtos[i] = ToUserDTO(&itm)
	}

	return userdtos
}
