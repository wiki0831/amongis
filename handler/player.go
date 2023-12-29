package handler

import (
	"amongis/database"
	"amongis/model"

	"github.com/gofiber/fiber/v2"
)

func GetPlayer(c *fiber.Ctx) error {
	userName := c.Params("userName")

	// check out from DB
	playerInfo, dbErr := database.GetPlayerInfo(userName)
	if dbErr != nil {
		return c.
			Status(500).
			JSON(fiber.Map{"status": "error", "message": "issue with DB checkout", "errors": dbErr.Error()})
	}

	return c.Status(200).JSON(playerInfo)
}

func GetAllPlayers(c *fiber.Ctx) error {
	// check out from DB
	playerInfo, dbErr := database.GetAllPlayerInfo()
	if dbErr != nil {
		return c.
			Status(500).
			JSON(fiber.Map{"status": "error", "message": "issue with DB checkout", "errors": dbErr.Error()})
	}

	return c.Status(200).JSON(playerInfo)
}

func PostPlayerTelem(c *fiber.Ctx) error {
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

	// save player telemetry to db
	dbErr := database.CreatePlayerTelemetry(*user)
	if dbErr != nil {
		return c.
			Status(500).
			JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": dbErr.Error()})

	}

	// compute user actionable items
	// TODO

	return c.Status(200).SendString("telmetry logged 👍")
}
