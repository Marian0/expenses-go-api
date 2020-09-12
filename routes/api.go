package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marian0/expenses-go-api/controllers"
	"gorm.io/gorm"
)

//DefineRoutes -
func DefineRoutes(DB *gorm.DB) *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api")
	{

		usersController := controllers.NewUsersController(DB)

		grp1.GET("users", usersController.GetUsers)
		grp1.POST("users", usersController.CreateUser)
		// grp1.GET("user/:id", Controllers.GetUserByID)
		// grp1.PUT("user/:id", Controllers.UpdateUser)
		// grp1.DELETE("user/:id", Controllers.DeleteUser)
	}
	return r

}
