package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	db_connection "github.com/ssssshel/ms_aster_user_data_go/src/db"
	"github.com/ssssshel/ms_aster_user_data_go/src/middlewares"
	users_routes "github.com/ssssshel/ms_aster_user_data_go/src/routes/users"
)

func main() {
	app := fiber.New()

	EnvironmentsManager(app, production)
}

func defaultInitConf(app *fiber.App, tokenization bool) {
	db_connection.PostgresConnection()

	app.Use(logger.New())

	app.Use(cors.New())

	if tokenization {
		app.Use(middlewares.VerifyAccessToken)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("MS User Records")
	})

	v1 := app.Group("/v1")
	v2 := app.Group("/v2")

	users_routes.RouteBuilderV1(v1)
	users_routes.RouteBuilderV2(v2)
	// orders_routes.Routes(v)

	app.Listen(":3000")

}
