package rules

// Both a board and previous move represents a game because of en passant.
type Game struct {
	Board
	Previous Move
}

var NoPreviousMove = Move{Address{0, 8}, Address{0, 8}}

// TODO: should Board be a pointer? measure the cost of copying it for methods
