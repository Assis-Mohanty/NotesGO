package handlers

import (
	"net/http"

	"github.com/Assis-Mohanty/notes/models"
	"github.com/Assis-Mohanty/notes/utils"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func Register(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		var body request
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}
		if body.Name == "" || body.Email == "" || body.Password == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "All fields are required"})
		}

		//checking if user all ready exists
		var existing models.User
		if err := db.Where("email = ?", body.Email).First(&existing).Error; err == nil {
			return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "user already exists"})
		}
		hashedPassword, err := utils.HashPassword(body.Password)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to hash the password",
			})
		}
		user := models.User{
			Name:     body.Name,
			Email:    body.Email,
			Password: hashedPassword,
		}

		if err := db.Create(&user).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to Create User",
			})
		}

		return c.Status(http.StatusCreated).JSON(fiber.Map{
			"message": "Created a User successfully"})

	}

}


func Login(db *gorm.DB) fiber.Handler{
	return func (c *fiber.Ctx) error {
		type request struct {
			Email string `json:"email"`
			Password string `json:"password"`
		}
		var body request
		if err:=c.BodyParser(&body);err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Request",
			})
		}
		var user models.User
		if err:=db.Where("email=?",body.Email).First(&user).Error;err!= nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		if err:=utils.CompareHashPassword(user.Password,body.Password);err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Password ",
			})
		}

		accessToken,err :=utils.GenerateJWT(user.ID,false)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Access Token generation failed",
			})
		}
		refreshToken,err :=utils.GenerateJWT(user.ID,false)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Refresh Token generation failed",
			})
		}
		return c.JSON(fiber.Map{"accessToken":accessToken,"refreshToken":refreshToken})

	}
}


