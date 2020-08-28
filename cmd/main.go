// The wichess command line program is separated from package wichess to share application symbols
// with the test package.
package main

import (
	"net/http"

	"github.com/pciet/wichess"
)

func main() {
	wichess.LoadHTMLTemplates()

	// A PostgreSQL database stores information that persists between reboots.
	wichess.InitializeDatabaseConnection()

	// HTTP handlers are how people get the game interface in their web browser and how game
	// interactions are done with this host program.
	wichess.InitializeHTTP()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		wichess.Panic(err)
	}
}
