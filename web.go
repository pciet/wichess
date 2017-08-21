// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func executeWebTemplate(w http.ResponseWriter, file string, data interface{}) {
	t, err := template.ParseFiles(file)
	if err != nil {
		panicExit(fmt.Sprintf("failed to parse %v: %v", file, err.Error()))
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		panicExit(fmt.Sprintf("failed to execute template %v: %v", file, err.Error()))
		return
	}
}

func webError(w http.ResponseWriter, r *http.Request, message string, the error) {
	if the != nil {
		log.Printf("%v (%v)", message, the.Error())
	} else {
		log.Println(message)
	}
	http.NotFound(w, r)
}
