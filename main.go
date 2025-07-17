package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})
	log.SetLevel(log.LevelInfo)

	registerRoutes(app)

	app.Listen(":3000")
}
