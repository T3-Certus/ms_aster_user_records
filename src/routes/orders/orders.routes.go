package orders_routes

import "github.com/gofiber/fiber/v2"

func Routes(parent fiber.Router) fiber.Router {
	routesGroup := parent.Group("/orders")

	routesGroup.Get("/get-orders")

	return routesGroup
}
