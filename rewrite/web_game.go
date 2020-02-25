package main

import (
	"net/http"
)

// This file doesn't represent an HTTP path.
// The different game mode handlers use the same game template.

const game_web_template = "web/html/game.tmpl"

func init() { ParseHTMLTemplate(game_web_template) }

type GameWebTemplateData struct {
	Name string
	GameHeader
}

func WriteGameWebTemplate(w http.ResponseWriter, g GameWebTemplateData) {
	WriteWebTemplate(w, game_web_template, g)
}
