// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

type gameTemplate struct {
	White  string
	Black  string
	GameID int
	Name   string
}

const request_assignments = "assignments[]"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func executeWebTemplate(w http.ResponseWriter, file string, data interface{}) {
	t, err := template.ParseFiles(file)
	if err != nil {
		panicExit(fmt.Sprintf("failed to parse %v: %v", file, err.Error()))
	}
	err = t.Execute(w, data)
	if err != nil {
		panicExit(fmt.Sprintf("failed to execute template %v: %v", file, err.Error()))
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

func gameSetupFromForm(array []string) (gameSetup, error) {
	var s gameSetup
	if len(array) != 16 {
		return s, errors.New("web: array to parse is not length 16")
	}
	var err error
	for i, str := range array {
		s[i], err = strconv.Atoi(str)
		if err != nil {
			return s, err
		}
	}
	return s, nil
}
