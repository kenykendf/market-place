package database

import (
	"fmt"
	"log"
	"superindo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "root"
	password = "root"
	port     = 5432
	dbname   = "superindo_databases"
	DB       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	DB, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	DB.Debug().AutoMigrate(
		models.User{},
		models.Categories{},
		models.Item{},
		models.Order{},
		models.Carts{},
	)

	fmt.Println("Success To Connect Docker Postgresql")
}

func GetDB() *gorm.DB {
	return DB
}
