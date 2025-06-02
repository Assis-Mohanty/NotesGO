package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Assis-Mohanty/notes/models"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set")
	}
	
	log.Printf("Attempting to connect to database...")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	log.Println("Connected to DB successfully")
	DB = db
	db.AutoMigrate(&models.User{}, &models.Note{})
}

func ce(err error, msg ...string) {
	if err != nil {
		if len(msg) > 0 {
			log.Fatalf("%s: %v", msg[0], err)
		} else {
			log.Fatal(err)
		}
	}
}
