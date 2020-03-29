package main

import "github.com/pciet/wichess/rules"

func ActiveOrientation(active, white, black string) rules.Orientation {
	if active == white {
		return rules.White
	} else if active == black {
		return rules.Black
	}
	Panic("active", active, "not white", white, "or black", black)
	return rules.White
}
