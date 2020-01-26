package rules

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

type PieceKind int

type Piece struct {
	PieceKind
	Orientation
}
