package main

import (
	"github.com/pciet/wichess/rules"
)

type MoveSet struct {
	From  BoadAddress  `json:"from"`
	Moves []rules.Move `json:"moves"`
}
