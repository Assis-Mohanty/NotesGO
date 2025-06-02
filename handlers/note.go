package handlers

import (
	"fmt"
	"net/http"

	"github.com/Assis-Mohanty/notes/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateNote(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(uint)

		var note models.Note

		if err := c.BodyParser(&note); err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Input"})
		}

		if note.Title == "" || note.Content == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Feilds cant be empty"})
		}

		note.UserID = userID

		if err := db.Create(&note).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong while creating a note"})
		}
		fmt.Println("Parsed Title:", note.Title)
		fmt.Println("Parsed Content:", note.Content)
		fmt.Println("User ID:", userID)
		return c.JSON(note)
	}
}

func GetAllNotes(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(uint)
		search := c.Query("search")

		var notes []models.Note
		query := db.Where("user_id = ?", userID)
		if search != "" {
			like := "%" + search + "%"
			query = query.Where("title LIKE ? OR content LIKE ?", like, like)
		}

		page := c.QueryInt("page", 1)
		limit := c.QueryInt("limit", 10)
		offset := (page - 1) * limit

		if err := query.Limit(limit).Offset(offset).Find(&notes).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch notes",
			})
		}

		return c.JSON(notes)
	}
}

func GetNote(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(uint)
		id := c.Params("id")
		var note models.Note

		if err := db.First(&note, "id=? AND user_id =?", id, userID).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Note not found"})
		}
		return c.JSON(note)
	}
}

func UpdateNote(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(uint)
		id := c.Params("id")

		var note models.Note
		if err := db.First(&note, "id = ? AND user_id = ?", id, userID).Error; err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "seomthing went wrong "})
		}
		var update models.Note

		if err := c.BodyParser(&update); err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Input"})
		}
		note.Title = update.Title
		note.Content = update.Content

		if err := db.Save(&note).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "cannot update "})
		}
		return c.JSON(note)

	}
}

func DeleteNote(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(uint)
		id := c.Params("id")

		var note models.Note

		if err := db.First(&note, "id= ? AND user_id = ?", id, userID).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Note not found "})
		}

		if err := db.Delete(&note).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "delete failed"})
		}
		return c.SendStatus(http.StatusNoContent)
	}
}
