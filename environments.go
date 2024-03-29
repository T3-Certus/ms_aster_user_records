package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type env uint

const (
	development              env = 1
	testing                  env = 2
	production               env = 3
	developmentWithoutTokens env = 4 // DO NOT USE THIS IN PRODUCTION
)

func initEnvs() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func EnvironmentsManager(app *fiber.App, env env) {

	if env > 0 && env < 5 {
		switch env {
		case development:
			fmt.Println("Running in development mode")
			initEnvs()
			defaultInitConf(app, true)
		case testing:
			fmt.Println("Running in testing mode")
			initEnvs()
			defaultInitConf(app, true)
		case production:
			fmt.Println("Running in production mode")
			defaultInitConf(app, true)
		case developmentWithoutTokens:
			fmt.Println("Running in development mode without tokens")
			initEnvs()
			defaultInitConf(app, false)
		}
	}

	log.Fatal("Invalid environment configuration")

}
