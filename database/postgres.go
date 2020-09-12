package database

import (
	"log"
	"os"
	"time"

	"github.com/marian0/expenses-go-api/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//ConnectToPosgres - Connect to posrtgres DB
func ConnectToPosgres() *gorm.DB {

	//Define logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	//Connect and init database
	dsn := "host=localhost user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=5432 sslmode=disable TimeZone=" + os.Getenv("POSTGRES_TIMEZONE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	common.LogFatal(err)
	return db
}
