package wichess

import "github.com/pciet/wichess/rules"

func PlayerWithOrientation(o rules.Orientation, white, black string) string {
	if o == rules.White {
		return white
	} else if o == rules.Black {
		return black
	}
	Panic("bad orientation", o)
	return ""
}

func OrientationOf(player, white, black string) rules.Orientation {
	if player == white {
		return rules.White
	} else if player == black {
		return rules.Black
	}
	Panic(player, "not", white, black)
	return rules.White
}

func Opponent(of rules.Orientation, white, black string) string {
	if of == rules.White {
		return black
	} else if of == rules.Black {
		return white
	}
	Panic(of, "not white", white, "or black", black)
	return ""
}
