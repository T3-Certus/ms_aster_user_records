package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type env uint

const (
	development              env = 1
	testing                  env = 2
	production               env = 3
	developmentWithoutTokens env = 4 // DO NOT USE THIS IN PRODUCTION
)

func Environments(app *fiber.App, env env) {

	if env > 0 && env < 5 {
		switch env {
		case development:
			fmt.Println("Running in development mode")
			defaultInitConf(app, true)
		case testing:
			fmt.Println("Running in testing mode")
			defaultInitConf(app, true)
		case production:
			fmt.Println("Running in production mode")
			defaultInitConf(app, true)
		case developmentWithoutTokens:
			fmt.Println("Running in development mode without tokens")
			defaultInitConf(app, false)
		}
	}

	log.Fatal("Invalid environment configuration")

}
