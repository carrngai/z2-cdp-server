package main

import (
	"log"

	"./pkg/database"
	"./pkg/gateway"
)

func main() {
	// Initialize the database
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the gateway
	gw := gateway.NewGateway()

	// Start the gateway
	gw.Start()
}
