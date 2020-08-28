package wichess

// The different game mode handlers use the same game template.
const GameHTMLTemplate = "html/game.tmpl"

type GameHTMLTemplateData struct {
	Name string
	GameHeader
}
