package main

import (
	"github.com/QuickBrawl/session-service/handles"
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App) {
	sessionGroup := app.Group("/session")
	registerSessionRoutes(sessionGroup)
}

func registerSessionRoutes(group fiber.Router) {
	group.Get("/register", handles.CreateSession)
}
