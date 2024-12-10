package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase подключается к базе данных PostgreSQL
func ConnectDatabase() {
	// Строка подключения (замените на свои данные)
	dsn := "user=postgres password=Na260206 dbname=dotnet host=localhost port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Successfully connected to the database!")
}
