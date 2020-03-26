package main

import "net/http"

func main() {
	// A PostgreSQL database stores information that persists between reboots.
	InitializeDatabaseConnection()

	// HTTP handlers are how people get the game interface in their web
	// browser and how game interactions are done with this host program.
	InitializeHTTP()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		Panic(err)
	}
}
