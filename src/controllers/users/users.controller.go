package users_controller

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	db_connection "github.com/ssssshel/ms_aster_user_data_go/src/db"
	"github.com/ssssshel/ms_aster_user_data_go/src/models"
	"github.com/ssssshel/restponses-go"
	"github.com/ssssshel/restponses-go/utils"
	"gorm.io/gorm"
)

func HandleGetUserData(c *fiber.Ctx) error {
	db := db_connection.DBConn

	userId := c.Params("userId")
	userModel := []models.UserData{}

	userChan := make(chan []models.UserData)
	errChan := make(chan error)

	go func() {
		if err := db.First(&userModel, &userId).Error; err != nil {
			errChan <- err
			return
		}
		userChan <- userModel
	}()

	select {
	case user := <-userChan:
		return c.Status(fiber.StatusOK).JSON(restponses.Response2xxSuccessfull(restponses.Status200Ok, &restponses.BaseSuccessfulInput{Data: user}))

	case err := <-errChan:
		errString := fmt.Sprintf("%s", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			nfUser := fmt.Sprintf("User with id %s", userId)

			return c.Status(fiber.StatusNotFound).JSON(restponses.Response4xxClientError(restponses.Status404NotFound, &restponses.BaseClientErrorInput{Detail: errString, StatusOptions: utils.Status404Opt(nfUser)}))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(restponses.Response5xxServerError(restponses.Status500InternalServerError, &restponses.BaseServerErrorInput{Detail: errString}))
	}

}

func HandleUpdateUserData(c *fiber.Ctx) error {
	db := db_connection.DBConn

	userId := c.Params("userId")

	payload := models.UserData{}

	// payload.id_user_rol = uint(c.Body()["id_user_rol"])
	// payload.user_cellphone = string(c.Body()["user_cellphone"])
	// payload.user_document_number = string(c.Body()["user_document_number"])
	// payload.user_document_type = string(c.Body()["user_document_type"])
	// payload.user_email = string(c.Body()["user_email"])
	// payload.user_name = string(c.Body()["user_name"])
	// payload.user_surname = string(c.Body()["user_surname"])

	userModel := []models.UserData{}

	userChan := make(chan []models.UserData)
	validationsErrorsChan := make(chan error)
	errChan := make(chan error)

	v := validator.New()

	go func() {
		fmt.Println("payload: ", &payload)
		if err := c.BodyParser(&payload); err != nil {
			validationsErrorsChan <- err
			return
		}

		if err := v.Struct(payload); err != nil {
			validationsErrorsChan <- err
			return
		}

		if err := db.First(&userModel, &userId).Updates(&payload).Error; err != nil {
			errChan <- err
			return
		}
		userChan <- userModel
	}()

	select {
	case user := <-userChan:
		return c.Status(fiber.StatusOK).JSON(restponses.Response2xxSuccessfull(restponses.Status200Ok, &restponses.BaseSuccessfulInput{Data: user}))
	case err := <-validationsErrorsChan:
		return c.Status(fiber.StatusBadRequest).JSON(restponses.Response4xxClientError(restponses.Status400BadRequest, &restponses.BaseClientErrorInput{Detail: err.Error()}))
	case err := <-errChan:
		errString := fmt.Sprintf("%s", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			nfUser := fmt.Sprintf("User with id %s", userId)

			return c.Status(fiber.StatusNotFound).JSON(restponses.Response4xxClientError(restponses.Status404NotFound, &restponses.BaseClientErrorInput{Detail: errString, StatusOptions: utils.Status404Opt(nfUser)}))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(restponses.Response5xxServerError(restponses.Status500InternalServerError, &restponses.BaseServerErrorInput{Detail: errString}))
	}

}
