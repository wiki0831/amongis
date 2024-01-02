package router

import (
	"amongis/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.Welcome)

	api := app.Group("/api", logger.New())
	api.Get("/", handler.Welcome)
	api.Get("/ping", handler.Pong)
	api.Get("/health", handler.HealthCheck)

	player := api.Group("/player", logger.New())
	player.Get("/", handler.GetAllPlayers)
	player.Post("/", handler.PostPlayerTelem)
	player.Get("/:userName", handler.GetPlayer)


	location := api.Group("/location",logger.New())
	location.Post("/",handler.PostLocation)
	location.Post("/checkExit",handler.PostCheckExit)
}
