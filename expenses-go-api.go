package main

import (
	"log"

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
}
