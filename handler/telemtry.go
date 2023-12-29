package handler

import (
	"amongis/database"
	"amongis/model"

	"github.com/gofiber/fiber/v2"
)

func LogTelemetry(c *fiber.Ctx) error {
	// parse body and validate basic input
	user := new(model.PlayerModel)
	if inputErr := c.BodyParser(user); inputErr != nil {
		return c.
			Status(500).
			JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": inputErr.Error()})
	}

	// validate business logic
	logicErr := user.Validate()
	if logicErr != nil {
		return c.
			Status(500).
			JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": logicErr.Error()})
	}

	// save telemetry to db
	dbErr := database.CreateTelemetry(*user)
	if dbErr != nil {
		return c.
			Status(500).
			JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": dbErr.Error()})

	}

	// compute user actionable items
	// TODO

	return c.Status(200).SendString("telmetry logged 👍")
}
