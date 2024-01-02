package handler

import (
	"amongis/database"
	"amongis/model"

	"github.com/gofiber/fiber/v2"
)

func PostLocation(c *fiber.Ctx) error {
	// parse body and validate basic input
	location := new(model.Location)
	if inputErr := c.BodyParser(location); inputErr != nil {
		return c.
			Status(500).
			JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": inputErr.Error()})
	}

	// validate business logic
	logicErr := location.Validate()
	if logicErr != nil {
		return c.
			Status(500).
			JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": logicErr.Error()})
	}

	// save location to db
	dbErr := database.CreateLocation(*location)
	if dbErr != nil {
		return c.
			Status(500).
			JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": dbErr.Error()})

	}

	// compute location actionable items
	// TODO

	return c.Status(200).SendString("location logged 👍")
}

func PostCheckExit(c *fiber.Ctx) error {
	body := c.Body()
	exitStatus := database.CheckExit(string(body))

	if exitStatus {
		return c.Status(200).SendString("exit is opened in: " + string(body))
	} else {
		return c.Status(500).SendString("exit is closed in: " + string(body))
	}	
}
