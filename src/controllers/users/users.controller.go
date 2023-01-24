package users_controller

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	db_connection "github.com/ssssshel/ms_aster_user_data_go/src/db"
	"github.com/ssssshel/ms_aster_user_data_go/src/models"
	"github.com/ssssshel/restponses-go"
	"github.com/ssssshel/restponses-go/utils"
	"gorm.io/gorm"
)

func HandleGetUserData(c *fiber.Ctx) error {
	userId := c.Params("userId")
	fmt.Println(userId)

	db := db_connection.DBConn

	userData := []models.UserData{}

	dbResponse := db.First(&userData, &userId)

	if dbResponse.Error != nil {
		err := dbResponse.Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("ERROR: ", dbResponse)
			return c.Status(fiber.StatusNotFound).JSON(restponses.Response2xxSuccessfull(restponses.Status200, "", "", "UsersData", "", utils.Status201Opt("nn")))
		}

		fmt.Println("ERROR: ", dbResponse)
		return c.Status(fiber.StatusInternalServerError).JSON("gern")
	}

	fmt.Println("RES: ", userData)
	return c.Status(fiber.StatusOK).JSON(userData)
}
