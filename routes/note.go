package routes

import (
	"github.com/Assis-Mohanty/notes/handlers"
	"github.com/Assis-Mohanty/notes/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NotesRoutes(app *fiber.App, db *gorm.DB) {
	notes := app.Group("/notes", middleware.JWTProtected())
	notes.Post("/", handlers.CreateNote(db))
	notes.Get("/", handlers.GetAllNotes(db))
	notes.Get("/:id", handlers.GetNote(db))
	notes.Put("/:id", handlers.UpdateNote(db))
	notes.Delete("/:id", handlers.DeleteNote(db))
}
