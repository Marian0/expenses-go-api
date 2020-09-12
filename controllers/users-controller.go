package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marian0/expenses-go-api/models"
	"github.com/marian0/expenses-go-api/repositories"
	"gorm.io/gorm"
)

//UsersController -
type UsersController struct {
	userRepository *repositories.UserRepository
}

//NewUsersController -
func NewUsersController(DB *gorm.DB) *UsersController {
	return &UsersController{
		userRepository: &repositories.UserRepository{
			DB: DB,
		},
	}
}

//GetUsers ... Get all users
func (ctrl *UsersController) GetUsers(c *gin.Context) {
	var users []models.User
	err := ctrl.userRepository.GetAllUsers(&users)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

//CreateUser ... Create User
func (ctrl *UsersController) CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := ctrl.userRepository.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
