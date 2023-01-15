package users_routes

import (
	"github.com/gofiber/fiber/v2"
	users_controller "github.com/ssssshel/ms_aster_user_data_go/src/controllers/users"
)

func Routes(parent fiber.Router) fiber.Router {

	routesGroup := parent.Group("/users")

	routesGroup.Get("/user-data/:userId", users_controller.HandleGetUserData)
	// routesGroup.Put("/update-user/:userId")

	return routesGroup

}
