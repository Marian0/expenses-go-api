package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marian0/expenses-go-api/common"
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

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":" + os.Getenv("APP_PORT"))
}
