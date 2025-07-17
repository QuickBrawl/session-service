package handles

import (
	"net/http"

	"github.com/QuickBrawl/session-service/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

func CreateSession(c *fiber.Ctx) error {
	db := database.New()
	defer db.Close()

	newUUID := uuid.New().String()
	err := db.StoreSessionID(newUUID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Status": "Error", "Message": err.Error()})
	}

	log.Info("Created session with id: " + newUUID)
	return c.JSON(fiber.Map{"Status": "OK", "Data": newUUID})
}
