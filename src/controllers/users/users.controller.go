package users_controller

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	db_connection "github.com/ssssshel/ms_aster_user_data_go/src/db"
	"github.com/ssssshel/ms_aster_user_data_go/src/models"
	"gorm.io/gorm"
)

func HandleGetUserData(c *fiber.Ctx) error {
	userId := c.Params("userId")
	fmt.Println(userId)

	db := db_connection.DBConn

	userData := []models.UserData{}

	// type res struct {
	// 	Data string
	// }

	// resp := res{
	// 	Data: userId,
	// }

	dbResponse := db.First(&userData, &userId)

	if dbResponse.Error != nil {
		err := dbResponse.Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("ERROR: ", dbResponse)
			return c.Status(fiber.StatusNotFound).JSON("no existe")
		}

		fmt.Println("ERROR: ", dbResponse)
		return c.Status(fiber.StatusInternalServerError).JSON("Error generico")
	}

	fmt.Println("RES: ", userData)
	return c.Status(fiber.StatusOK).JSON(userData)
}
