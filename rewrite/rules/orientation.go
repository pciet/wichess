package rules

// Orientation is which player owns the piece and indicates the direction of movement.
type Orientation int

const (
	White Orientation = 0
	Black Orientation = 1
)

func (an Orientation) String() string {
	if an == White {
		return "white"
	} else if an == Black {
		return "black"
	}
	return "undefined"
}
