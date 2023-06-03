package main

import (
	"log"

	"github.com/carrngai/z2-cdp-server/pkg/database"
	"github.com/carrngai/z2-cdp-server/pkg/gateway"
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
