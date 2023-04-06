package users_controller

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	db_connection "github.com/ssssshel/ms_aster_user_data_go/src/db"
	"github.com/ssssshel/ms_aster_user_data_go/src/models"
	associations_models "github.com/ssssshel/ms_aster_user_data_go/src/models/associations"
	"github.com/ssssshel/restponses-go"
	"github.com/ssssshel/restponses-go/utils"
	"gorm.io/gorm"
)

func HandleGetUserData(c *fiber.Ctx) error {
	db := db_connection.DBConn

	userId := c.Params("userId")

	userChan := make(chan associations_models.UserData)
	errChan := make(chan error)

	go func() {
		defer close(userChan)
		defer close(errChan)
		userModel := associations_models.UserData{}

		if err := db.Preload("USER_ROLE").WithContext(c.Context()).First(&userModel, &userId).Error; err != nil {
			errChan <- err
			return
		}
		userChan <- userModel
	}()

	select {
	case user := <-userChan:
		return c.Status(fiber.StatusOK).JSON(restponses.Response2xxSuccessfull(restponses.Status200Ok, &restponses.BaseSuccessfulInput{Data: user}))

	case err := <-errChan:
		if errors.Is(err, gorm.ErrRecordNotFound) {
			nfUser := fmt.Sprintf("User with id %s", userId)

			return c.Status(fiber.StatusNotFound).JSON(restponses.Response4xxClientError(restponses.Status404NotFound, &restponses.BaseClientErrorInput{Detail: err.Error(), StatusOptions: utils.Status404Opt(nfUser)}))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(restponses.Response5xxServerError(restponses.Status500InternalServerError, &restponses.BaseServerErrorInput{Detail: err.Error()}))
		}
	}

}

// func HandleGetUserDataNoThreads(c *fiber.Ctx) error {
// 	db := db_connection.DBConn

// 	userId := c.Params("userId")

// 	userModel := associations_models.UserData{}

// 	if err := db.Preload("USER_ROLE").WithContext(c.Context()).First(&userModel, &userId).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			nfUser := fmt.Sprintf("User with id %s", userId)

// 			return c.Status(fiber.StatusNotFound).JSON(restponses.Response4xxClientError(restponses.Status404NotFound, &restponses.BaseClientErrorInput{Detail: err.Error(), StatusOptions: utils.Status404Opt(nfUser)}))
// 		} else {
// 			return c.Status(fiber.StatusInternalServerError).JSON(restponses.Response5xxServerError(restponses.Status500InternalServerError, &restponses.BaseServerErrorInput{Detail: err.Error()}))
// 		}
// 	} else {
// 		return c.Status(fiber.StatusOK).JSON(restponses.Response2xxSuccessfull(restponses.Status200Ok, &restponses.BaseSuccessfulInput{Data: userModel}))
// 	}

// }

func HandleUpdateUserData(c *fiber.Ctx) error {
	db := db_connection.DBConn

	userId := c.Params("userId")
	payload := models.UserData{}

	userChan := make(chan models.UserData)
	validationsErrorsChan := make(chan error)
	errChan := make(chan error)

	v := validator.New()

	go func() {
		defer close(userChan)
		defer close(validationsErrorsChan)
		defer close(errChan)

		fmt.Println("payload: ", &payload)
		if err := c.BodyParser(&payload); err != nil {
			validationsErrorsChan <- err
			return
		}

		if err := v.Struct(payload); err != nil {
			validationsErrorsChan <- err
			return
		}

		userModel := models.UserData{}
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			nfUser := fmt.Sprintf("User with id %s", userId)

			return c.Status(fiber.StatusNotFound).JSON(restponses.Response4xxClientError(restponses.Status404NotFound, &restponses.BaseClientErrorInput{Detail: err.Error(), StatusOptions: utils.Status404Opt(nfUser)}))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(restponses.Response5xxServerError(restponses.Status500InternalServerError, &restponses.BaseServerErrorInput{Detail: err.Error()}))
	}

}
