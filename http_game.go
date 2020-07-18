package main

// The different game mode handlers use the same game template.
const GameHTMLTemplate = "html/game.tmpl"

func init() { ParseHTMLTemplate(GameHTMLTemplate) }

type GameHTMLTemplateData struct {
	Name string
	GameHeader
}
