package rules

import (
	"log"
)

// Piece ownership and direction of travel is shown by orientation.
type Orientation int

const (
	White Orientation = 0
	Black Orientation = 1
)

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
