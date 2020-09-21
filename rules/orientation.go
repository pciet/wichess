package rules

import "log"

// Orientation indicates piece ownership and direction of movement.
type Orientation int

const (
	White Orientation = iota
	Black
)

func (an Orientation) Opponent() Orientation {
	switch an {
	case White:
		return Black
	case Black:
		return White
	}
	log.Panicln("unknown orientation", an)
	return White
}

func (an Orientation) String() string {
	switch an {
	case White:
		return "white"
	case Black:
		return "black"
	}
	log.Panicln("unknown orientation", an)
	return ""
}
