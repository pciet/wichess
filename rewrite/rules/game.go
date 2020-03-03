package rules

// Both a board and previous move represents a game because of en passant.
type Game struct {
	Board
	Previous Move
}

// TODO: should Board be a pointer? measure the cost of copying it for methods
