package rules

// Piece ownership and direction of travel is shown by orientation.
type Orientation int

const (
	White Orientation = 0
	Black Orientation = 1
)

func (an Orientation) Opponent() Orientation {
	switch an {
	case White:
		return Black
	case Black:
		return White
	}
	Panic("unknown orientation", an)
	return White
}

func (an Orientation) String() string {
	switch an {
	case White:
		return "white"
	case Black:
		return "black"
	}
	Panic("unknown orientation", an)
	return ""
}

func BoolToOrientation(a bool) Orientation {
	if a == false {
		return White
	}
	return Black
}
