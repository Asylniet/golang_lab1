package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	handlers2 "lab1/handlers"
	"os"
)

var DB *gorm.DB

func Init() {
	handlers2.LoadEnvVariables()
	var err error
	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to DB")

	handlers2.UNUSED(DB)
}
