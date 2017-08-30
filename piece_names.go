// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"github.com/pciet/wichess/wichessing"
)

func nameForKind(the wichessing.Kind) string {
	switch the {
	case wichessing.King:
		return "King"
	case wichessing.Queen:
		return "Queen"
	case wichessing.Rook:
		return "Rook"
	case wichessing.Bishop:
		return "Bishop"
	case wichessing.Knight:
		return "Knight"
	case wichessing.Pawn:
		return "Pawn"
	case wichessing.Swap:
		return "Swap"
	case wichessing.Lock:
		return "Lock"
	case wichessing.Recon:
		return "Recon"
	case wichessing.Detonate:
		return "Detonate"
	case wichessing.Ghost:
		return "Ghost"
	case wichessing.Steal:
		return "Steal"
	case wichessing.Guard:
		return "Guard"
	case wichessing.Rally:
		return "Rally"
	case wichessing.Fortify:
		return "Fortify"
	}
	return ""
}
