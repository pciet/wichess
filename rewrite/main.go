package main

import (
	"log"
	"net/http"
)

func main() {
	// A PostgreSQL database is used to store all information that needs to persist between reboots.
	InitializeDatabaseConnection()

	// HTTP handlers are how people get the game interface in their web browser and how game interactions are done with this program.
	InitializeHTTP()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
