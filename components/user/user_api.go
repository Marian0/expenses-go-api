package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marian0/expenses-go-api/common"
	"gorm.io/gorm"
)

//UserAPI - s
type UserAPI struct {
	DB             *gorm.DB
	UserService    *UserService
	UserRepository *UserRepository
}

//NewUserAPI -
func NewUserAPI(DB *gorm.DB) *UserAPI {
	userRepository := NewUserRepository(DB)
	userService := NewUserService(userRepository)

	return &UserAPI{
		DB:             DB,
		UserService:    &userService,
		UserRepository: &userRepository,
	}
}

//FindAll -
func (p *UserAPI) FindAll(c *gin.Context) {
	users := p.UserService.FindAll()

	common.RespondSuccess(c, http.StatusOK, gin.H{"users": ToUserDTOs(users)})
}

//FindByID -
func (p *UserAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := p.UserService.FindByID(uint(id))

	common.RespondSuccess(c, http.StatusOK, gin.H{"user": ToUserDTO(&user)})
}

//Create -
func (p *UserAPI) Create(c *gin.Context) {
	//Logged User
	// uuid, _ := c.Get(common.CONST_UUID_IDENTIFIER)

	var userDTO UserCreateDTO

	err := c.ShouldBindJSON(&userDTO)
	if err != nil {
		common.RespondError(c, http.StatusBadRequest, err.Error(), "There was a problem trying to create a User")
		return
	}

	createdUser := p.UserService.Create(&User{
		Name:  userDTO.Name,
		Email: userDTO.Email,
	})

	common.RespondSuccess(c, http.StatusCreated, gin.H{"user": ToUserDTO(createdUser)})
}

//Update -
func (p *UserAPI) Update(c *gin.Context) {
	var userDTO UserUpdateDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		common.RespondError(c, http.StatusBadRequest, err.Error(), "There was a problem trying to update a User")
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user := p.UserService.FindByID(uint(id))
	if user == (User{}) {
		common.RespondError(c, http.StatusNotFound, "User not found "+c.Param("id"), "User not found")
		return
	}

	user.ID = uint(id)
	user.Name = userDTO.Name
	if userDTO.Email != "" {
		user.Email = userDTO.Email
	}
	p.UserService.Save(user)

	common.RespondSuccess(c, http.StatusOK, gin.H{"user": ToUserDTO(&user)})
}

//Delete -
func (p *UserAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := p.UserService.FindByID(uint(id))
	if user == (User{}) {
		common.RespondError(c, http.StatusNotFound, "User not found "+c.Param("id"), "User not found")
		return
	}

	p.UserService.Delete(user)

	common.RespondSuccess(c, http.StatusOK, gin.H{"user": ToUserDTO(&user)})
}
