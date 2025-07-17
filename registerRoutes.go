package main

import "github.com/gofiber/fiber/v2"

func registerRoutes(app *fiber.App) {
	sessionGroup := app.Group("/session")
	registerSessionRoutes(sessionGroup)
}

func registerSessionRoutes(group fiber.Router) {

}
