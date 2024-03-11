package main

import (
	"focus/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// get environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get host name, user, dbname, sslmode, password from dotenv
	hostName := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")
	password := os.Getenv("DB_PASSWORD")

	// Connect to the database
	dsn := "host=" + hostName + " user=" + user + " dbname=" + dbName + " sslmode=" + sslMode + " password=" + password
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Automigrate will check the defined model and attempt to automatically migrate the schema
	err = db.AutoMigrate(&models.User{}, &models.Habit{}, &models.HabitCompletion{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed successfully")
}
