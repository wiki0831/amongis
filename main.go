package main

import (
	"amongis/config"
	"amongis/database"
	"amongis/router"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	//Launch fiber framework
	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))

	//Add logging
	os.Mkdir("logs", 0755)
	file, err := os.Create("./logs/server.log")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app.Use(logger.New(logger.Config{
		Next:         nil,
		Format:       "${time} ${status} - ${latency} ${method} ${path} -- req:${body} ; res:${resBody}\n",
		TimeFormat:   "2006/01/02 15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       file,
	}))

	//DB Connection
	database.ConnectDB(os.Getenv("DATABASE_URL"))
	
	//Close db connection upon return
	defer database.DB.Close()

	//Set up API routes
	router.SetupRoutes(app)

	//Listens to port
	log.Fatal(app.Listen(fmt.Sprintf(":%s", *config.PORT)))
}
