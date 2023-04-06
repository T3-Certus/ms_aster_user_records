package users_routes

import (
	"github.com/gofiber/fiber/v2"
	users_controller "github.com/ssssshel/ms_aster_user_data_go/src/controllers/users"
	config "github.com/ssssshel/ms_aster_user_data_go/src/utils/config"
)

func createRoutes(router fiber.Router, version config.APIVersion) fiber.Router {

	routesGroup := router.Group("/users")

	switch version {
	case config.V1:
		routesGroup.Get("/user-data/:userId", users_controller.HandleGetUserData)
		routesGroup.Put("/update-user/:userId", users_controller.HandleUpdateUserData)
		return routesGroup

	case config.V2:
		routesGroup.Get("/user/:userId", users_controller.HandleGetUserDataNoThreads)
		routesGroup.Put("/user/:userId", users_controller.HandleUpdateUserData)
		return routesGroup

	default:
		return nil
	}
}

func RouteBuilderV1(router fiber.Router) fiber.Router {
	return createRoutes(router, config.V1)
}

func RouteBuilderV2(router fiber.Router) fiber.Router {
	return createRoutes(router, config.V2)
}
