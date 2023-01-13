package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	db_connection "github.com/ssssshel/ms_aster_user_data_go/src/db"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Env error")
	}

	db_connection.PostgresConnection()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("MS User Data")
	})

	// v1 := app.Group("/v1")

	app.Listen(":3000")
}
