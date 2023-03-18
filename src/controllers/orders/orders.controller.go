package orders_controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	db_connection "github.com/ssssshel/ms_aster_user_data_go/src/db"
	"github.com/ssssshel/ms_aster_user_data_go/src/models"
	"github.com/ssssshel/restponses-go"
	"gorm.io/gorm"
)

func HandleGetUserOrders(c *fiber.Ctx) error {
	db := db_connection.DBConn

	userId := c.Query("userId", "")
	orderid := c.Query("orderId", "")
	orderModel := []models.UserOrders{}

	fmt.Println(userId, orderid)

	ordersChan := make(chan []models.UserOrders)
	errChan := make(chan error)
	validationErrChan := make(chan error)

	go func() {
		defer close(ordersChan)
		defer close(errChan)

		if userId != "" {
			parsedUserId, _ := strconv.ParseUint(userId, 10, 64)
			if orderid != "" {
				parsedOrderId, _ := strconv.ParseUint(orderid, 10, 64)
				if err := db.Where(&models.UserOrders{ID_USER_ORDER: parsedOrderId, ID_USER: parsedUserId}).Find(&orderModel).Error; err != nil {
					errChan <- err
					return
				}
				ordersChan <- orderModel
				return
			} else {
				if err := db.Where(&models.UserOrders{ID_USER: parsedUserId}).Find(&orderModel).Error; err != nil {
					errChan <- err
					return
				}
				ordersChan <- orderModel
				return
			}
		} else {
			validationErrChan <- errors.New("userId is required")
			return
		}

	}()

	select {
	case orders := <-ordersChan:
		if len(orders) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(restponses.Response4xxClientError(restponses.Status404NotFound, &restponses.BaseClientErrorInput{Detail: "Orders not found"}))
		}
		return c.Status(fiber.StatusOK).JSON(restponses.Response2xxSuccessfull(restponses.Status200Ok, &restponses.BaseSuccessfulInput{Data: orders}))
	case err := <-validationErrChan:
		return c.Status(fiber.StatusBadRequest).JSON(restponses.Response4xxClientError(restponses.Status400BadRequest, &restponses.BaseClientErrorInput{Detail: err.Error()}))
	case err := <-errChan:
		// errString := fmt.Sprintf("%s", err)
		errString := err.Error()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(restponses.Response4xxClientError(restponses.Status404NotFound, &restponses.BaseClientErrorInput{Detail: errString}))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(restponses.Response5xxServerError(restponses.Status500InternalServerError, &restponses.BaseServerErrorInput{Detail: errString}))
	}
}
