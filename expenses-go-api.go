package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/marian0/expenses-go-api/common"
	"github.com/marian0/expenses-go-api/database"
	"github.com/marian0/expenses-go-api/routes"
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
	routes.DefineRoutes(DB).Run(":" + os.Getenv("APP_PORT"))
}
