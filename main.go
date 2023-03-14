package main

import (
	"log"

	"superindo/database"
	"superindo/routes"
)

func main() {
	log.Println("Initiate Application")
	r := routes.Init()

	database.StartDB()

	r.Run(":4123")
}
