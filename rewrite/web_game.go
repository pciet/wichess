package main

import (
	"net/http"
)

const game_web_template = "web/html/game.html"

func init() { ParseHTMLTemplate(game_web_template) }

type GameWebTemplateData struct {
	Name string
	GameHeader
}

func WriteGameWebTemplate(w http.ResponseWriter, g GameWebTemplateData) {
	WriteWebTemplate(w, game_web_template, g)
}
