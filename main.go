// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"net/http"
)

func main() {
	initializeDatabaseConnection()

	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/login", loginHandler)

	http.HandleFunc("/pieces", freePiecesHandler)

	http.HandleFunc("/games/", gamesHandler)
	http.HandleFunc("/moves/", movesHandler)
	http.HandleFunc("/move/", moveRequestHandler)
	http.HandleFunc("/moven/", moveNotificationWebsocketHandler)
	http.HandleFunc("/acknowledge", acknowledgeGameCompletionHandler)

	http.HandleFunc("/easycomputerrequest", easyComputerRequestHandler)
	http.HandleFunc("/easycomputer", easyComputerHandler)

	http.HandleFunc("/competitive48", competitive48RequestHandler)
	http.HandleFunc("/competitive48n", competitive48NotificationWebsocketHandler)
	http.HandleFunc("/competitive48/", competitive48Handler)

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("web/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("web/css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("web/img"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panicExit(err.Error())
	}
}

func panicExit(message string) {
	panic(message)
}
