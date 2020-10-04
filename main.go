package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marian0/expenses-go-api/common"
	"github.com/marian0/expenses-go-api/components/user"
	"github.com/marian0/expenses-go-api/database"
	"gorm.io/gorm"
)

//DB - global variable to represent the DB connection
var DB *gorm.DB

func init() {
	//Load environment variables
	err := godotenv.Load()
	common.LogFatal(err)

	//Connect to postgres
	DB = database.ConnectToPosgres()

	//Run migrations engine to keep DB up to date
	database.MigrateDatabase(DB)
}

func main() {
	log.Println("Hello Expeses Api")

	userAPI := user.NewUserAPI(DB)

	r := gin.Default()

	r.GET("/users", userAPI.FindAll)
	r.GET("/users/:id", userAPI.FindByID)
	r.POST("/users", userAPI.Create)
	r.PUT("/users/:id", userAPI.Update)
	r.DELETE("/users/:id", userAPI.Delete)

	err := r.Run(":" + os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}

}
