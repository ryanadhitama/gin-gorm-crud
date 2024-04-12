package config

import (
	"fmt"
	"os"
	"strconv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gin-gorm-crud/utils"
	"github.com/joho/godotenv"
)

func DatabaseConnection() *gorm.DB {
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }

	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	var (
		host     = os.Getenv("DATABASE_HOST")
		user     = os.Getenv("DATABASE_USERNAME")
		password = os.Getenv("DATABASE_PASSWORD")
		dbName   = os.Getenv("DATABASE_NAME")
	)

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	utils.ErrorPanic(err)

	return db
}