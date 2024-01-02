package handler

import (
	"amongis/database"
	"amongis/logic"
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
	player := new(model.Player)
	if inputErr := c.BodyParser(player); inputErr != nil {
		return c.
			Status(500).
			JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": inputErr.Error()})
	}

	// validate business logic
	validtionErr := player.Validate()
	if validtionErr != nil {
		return c.
			Status(500).
			JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": validtionErr.Error()})
	}

	// save player telemetry to db
	dbErr := database.CreatePlayerTelemetry(*player)
	if dbErr != nil {
		return c.
			Status(500).
			JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": dbErr.Error()})

	}

	// compute user actionable items
	userActions, logicErr := logic.GetAvailableAction(player)
	if logicErr != nil {
		return c.
			Status(500).
			JSON(fiber.Map{"status": "error", "message": "unable to compute user actions", "errors": logicErr.Error()})
	}
	// perform action items
	if player.Status == "alive" {
		logic.PerformActionItem(player, userActions)
	}
	newAvailableAction, _ := logic.GetAvailableAction(player)
	return c.Status(200).JSON(newAvailableAction)
}
