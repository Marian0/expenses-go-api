package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marian0/expenses-go-api/common"
)

func init() {
	//Load environment variables
	err := godotenv.Load()
	common.LogFatal(err)
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
