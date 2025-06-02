package main

import (
	"fmt"
	"log"

	"github.com/Assis-Mohanty/notes/config"
	"github.com/Assis-Mohanty/notes/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	config.ConnectDB()

	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), 14)

	user := models.User{
		Name:     "Casdwasd",
		Email:    "sdasd@example.com",
		Password: string(password),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		log.Fatalf("failed to seed user: %v", err)
	}

	for i := 1; i <= 10; i++ {
		note := models.Note{
			Title:   fmt.Sprintf("Casdwasd Note %d", i),
			Content: "hi askdnalsdwad ",
			UserID:  user.ID,
		}
		if err := config.DB.Create(&note).Error; err != nil {
			log.Printf("failed to create note %d: %v", i, err)
		}
	}

	fmt.Println("seed cli user and notes successfully")
}
