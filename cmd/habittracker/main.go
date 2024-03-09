package main

import (
	"focus/api"
	"focus/config"
	"log"
)

func main() {
	db := config.ConnectDatabase()

	router := api.SetupRouter(db)

	// Start serving the application
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run the server: ", err)
	}
}
