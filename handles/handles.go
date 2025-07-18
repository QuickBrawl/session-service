package handles

import (
	"net/http"

	"github.com/QuickBrawl/session-service/database"
	_ "github.com/gofiber/contrib/websocket"
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

type SessionJoinData struct {
	userId string
}

func JoinSession(c *fiber.Ctx) error {
	db := database.New()
	defer db.Close()

	// TODO:
	// Maybe get this from the authentication service? cookie id?
	var joinData SessionJoinData
	err := c.BodyParser(joinData)
	if err != nil {
		log.Error(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Status": "Error", "Message": err.Error()})
	}

	sessionId := c.Params("sessionId")
	exists, err := db.SessionExists(sessionId)
	if exists {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Status": "Error", "Message": "Session already exists."})
	}
	if err != nil {
		log.Error(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Status": "Error", "Message": err.Error()})
	}

	err = db.AddUserToSession(joinData.userId, sessionId)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(fiber.Map{"Status": "Error", "Message": err.Error()})
	}

	return c.SendStatus(http.StatusOK)
}
