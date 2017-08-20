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
	http.HandleFunc("/request", requestMatchHandler)

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
