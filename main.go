package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marian0/expenses-go-api/common"
	"github.com/marian0/expenses-go-api/components/user"
	"github.com/marian0/expenses-go-api/components/expense"
	"github.com/marian0/expenses-go-api/config"
	"github.com/marian0/expenses-go-api/database"
	"github.com/marian0/expenses-go-api/middlewares"
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

	expenseRepository := expense.NewRepository(DB)
	expenseService := expense.NewService(expenseRepository)
	expenseAPI := expense.NewAPI(expenseService)

	router := gin.Default()

	// configure firebase
	firebaseAuth := config.SetupFirebase()

	// set firebase auth to gin context with a middleware to all incoming request
	router.Use(func(c *gin.Context) {
		c.Set("firebaseAuth", firebaseAuth)
	})

	router.GET("/login", func(c *gin.Context) {
		common.RespondSuccess(c, http.StatusOK, gin.H{"message": "This endpoint should be public!"})
	})

	// using the auth middleware to validate api requests
	router.Use(middlewares.AuthMiddleware)
	{
		router.GET("/users", userAPI.FindAll)
		router.GET("/users/:id", userAPI.FindByID)
		router.POST("/users", userAPI.Create)
		router.PUT("/users/:id", userAPI.Update)
		router.DELETE("/users/:id", userAPI.Delete)
		
		router.GET("/expenses", expenseAPI.FindAll)
		router.GET("/expenses/:id", expenseAPI.FindByID)
		router.POST("/expenses", expenseAPI.Create)
		router.PUT("/expenses/:id", expenseAPI.Update)
		router.DELETE("/expenses/:id", expenseAPI.Delete)
	}

	err := router.Run(":" + os.Getenv("APP_PORT"))

	if err != nil {
		panic(err)
	}
}
