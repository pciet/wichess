package main

import "net/http"

// The different game mode handlers use the same game template.
const GameWebTemplate = "web/html/game.tmpl"

func init() { ParseHTMLTemplate(GameWebTemplate) }

type GameWebTemplateData struct {
	Name string
	GameHeader
}

func WriteGameWebTemplate(w http.ResponseWriter, g GameWebTemplateData) {
	WriteWebTemplate(w, GameWebTemplate, g)
}
