package handler

import (
	"amongis/config"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("You have reached amongis")
}

// Health Check
func HealthCheck(c *fiber.Ctx) error {
	ch3 := make(chan string)
	go internal_HealthCheck(ch3)

	return c.JSON(fiber.Map{
		"Internal": <-ch3,
	})
}

// Health Check
func Pong(c *fiber.Ctx) error {
	return c.SendString("pong")
}

func internal_HealthCheck(ch chan string) {
	start := time.Now()
	url := fmt.Sprintf("http://127.0.0.1:%s", *config.PORT)
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		ch <- err.Error()
	} else {
		elapsed := time.Since(start).Milliseconds()
		ch <- fmt.Sprintf("status: %d, ping: %vms", res.StatusCode, elapsed)
	}
}
