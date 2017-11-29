// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type gameTemplate struct {
	GameInfo
	Name      string
	TotalTime time.Duration
	NowTime   time.Time
}

const request_assignments = "assignments[]"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var parsedTemplates = map[string]*template.Template{}
var parsedTemplatesLock = sync.Mutex{}

func executeWebTemplate(w http.ResponseWriter, file string, data interface{}) {
	var t *template.Template
	var err error
	parsedTemplatesLock.Lock()
	t, has := parsedTemplates[file]
	if has == false {
		t, err = template.ParseFiles(file)
		if err != nil {
			panicExit(fmt.Sprintf("failed to parse %v: %v", file, err.Error()))
		}
		parsedTemplates[file] = t
	}
	parsedTemplatesLock.Unlock()
	err = t.Execute(w, data)
	if err != nil {
		log.Printf("failed to execute template %v: %v\n", file, err.Error())
	}
}

type BoardAssignments struct {
	Assignments []int `json:"assignments"`
	Index       int   `json:"index"`
}

func gameSetupFromRequest(array []int) (gameSetup, error) {
	var s gameSetup
	if len(array) != 16 {
		return s, errors.New("web: array to read is not length 16")
	}
	for i, id := range array {
		s[i] = id
	}
	return s, nil
}
