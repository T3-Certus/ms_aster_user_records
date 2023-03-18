package middlewares

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/ssssshel/ms_aster_user_data_go/src/utils/config"
	"github.com/ssssshel/restponses-go"
)

func VerifyAccessToken(c *fiber.Ctx) error {
	accessToken := c.Request().Header.Peek("authorization")

	if len(accessToken) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(restponses.Response4xxClientError(restponses.Status401Unauthorized, &restponses.BaseClientErrorInput{Detail: "Access token is required"}))
	}

	tok := []byte(strings.Replace(string(accessToken), "Bearer ", "", 1))
	authChan := make(chan jwt.Token)
	errChan := make(chan error)

	go func() {
		defer close(authChan)
		defer close(errChan)

		withVerificationKey := config.ATKey()

		verifiedToken, err := jwt.Parse(tok, jwt.WithVerify(jwa.HS256, []byte(withVerificationKey)))
		if err != nil {
			errChan <- err
			return
		}

		tokenExp := verifiedToken.Expiration().Unix()
		now := time.Now().Unix()

		if tokenExp < now {
			errChan <- errors.New("token expired")
			return
		}

		rol, ok := verifiedToken.Get("rol")
		if !ok || (rol != "admin" && rol != "user" && rol != "superadmin") {
			errChan <- errors.New("forbidden")
			return
		}

		authChan <- verifiedToken
	}()

	select {
	case verifiedToken := <-authChan:
		fmt.Printf("Verified token: %v", verifiedToken)
		return c.Next()
	case err := <-errChan:
		if err.Error() == "forbidden" {
			return c.Status(fiber.StatusForbidden).JSON(restponses.Response4xxClientError(restponses.Status403Forbidden, &restponses.BaseClientErrorInput{Detail: err.Error()}))
		}
		return c.Status(fiber.StatusUnauthorized).JSON(restponses.Response4xxClientError(restponses.Status401Unauthorized, &restponses.BaseClientErrorInput{Detail: err.Error()}))
	}
}
