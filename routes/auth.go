package routes

import (
	"github.com/Assis-Mohanty/notes/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRoutes(app *fiber.App, db *gorm.DB) {
	auth := app.Group("/auth")
	auth.Post("/register", handlers.Register(db))
	auth.Post("/login", handlers.Login(db))
}
