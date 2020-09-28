package wichess

import (
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/rules"
)

type GamePlayerHTMLTemplateData struct {
	memory.PlayerName
	memory.Captures
}

type PreviousMoveHTMLTemplateData struct {
	From, To int // to be sure it's not a string the address index is set as an int
}

type GameHTMLTemplateData struct {
	memory.GameIdentifier
	Conceded     bool
	White, Black GamePlayerHTMLTemplateData
	Active       rules.Orientation
	PreviousMove PreviousMoveHTMLTemplateData
	Player       memory.PlayerName
}
