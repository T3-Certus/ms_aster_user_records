package config

import (
	"os"

	utils_interfaces "github.com/ssssshel/ms_aster_user_data_go/src/utils/interfaces"
)

func PostgresConfig() *utils_interfaces.TPostgresConfig {
	values := &utils_interfaces.TPostgresConfig{
		DB_NAME:     os.Getenv("POSTGRESQL_DB"),
		DB_PORT:     os.Getenv("POSTGRESQL_DB_PORT"),
		DB_HOST:     os.Getenv("POSTGRESQL_DB_HOST"),
		DB_PASSWORD: os.Getenv("POSTGRESQL_DB_PASSWORD"),
		DB_SCHEMA:   os.Getenv("POSTGRESQL_DB_SCHEMA"),
		DB_USER:     os.Getenv("POSTGRESQL_DB_USER"),
	}

	return values
}

func ATKey() string {
	key := os.Getenv("ACCESS_TOKEN_PRIVATE_KEY")
	return key
}

type APIVersion uint

const (
	V1 APIVersion = 1
	V2 APIVersion = 2
)
