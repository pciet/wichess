package wichess

import (
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/rules"
)

type GamePlayerHTMLTemplateData struct {
	memory.PlayerName
	memory.Captures
}

type GameHTMLTemplateData struct {
	memory.GameIdentifier
	Conceded     bool
	White, Black GamePlayerHTMLTemplateData
	Active       rules.Orientation
	PreviousMove rules.Move
	Player       rules.Orientation
}
