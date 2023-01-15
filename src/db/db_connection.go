package db_connection

import (
	"fmt"
	"log"

	"github.com/ssssshel/ms_aster_user_data_go/src/utils/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func PostgresConnection() {

	config := config.PostgresConfig()

	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=%s sslmode=require ", config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT, config.DB_SCHEMA)
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error trying to connect to the database")
	} else {
		fmt.Println("Successfully connected with db")
	}

}
