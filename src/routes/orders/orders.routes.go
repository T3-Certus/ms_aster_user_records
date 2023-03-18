package orders_routes

import (
	"github.com/gofiber/fiber/v2"
	orders_controller "github.com/ssssshel/ms_aster_user_data_go/src/controllers/orders"
)

func Routes(parent fiber.Router) fiber.Router {
	routesGroup := parent.Group("/orders")

	routesGroup.Get("/get-orders", orders_controller.HandleGetUserOrders)

	return routesGroup
}
