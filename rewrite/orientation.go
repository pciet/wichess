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

func PlayerWithOrientation(o rules.Orientation, white, black string) string {
	if o == rules.White {
		return white
	} else if o == rules.Black {
		return black
	}
	Panic("bad orientation", o)
	return ""
}

func Opponent(of, white, black string) string {
	if of == white {
		return black
	} else if of == black {
		return white
	}
	Panic(of, "not white", white, "or black", black)
	return of
}
