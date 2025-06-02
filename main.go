package main

import (
	"github.com/Assis-Mohanty/notes/config"
	"github.com/Assis-Mohanty/notes/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	
	app := fiber.New()

	config.ConnectDB()
	routes.AuthRoutes(app, config.DB)
	
	app.Listen(":3000")
}
