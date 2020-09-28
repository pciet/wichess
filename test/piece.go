package test

import (
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

type testPiece struct {
	Address     rules.Address     `json:"addr"`
	Kind        piece.Kind        `json:"k"`
	Orientation rules.Orientation `json:"o"`
	Moved       bool              `json:"m"`
	Start       rules.Address     `json:"s"`
}
