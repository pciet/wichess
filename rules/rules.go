// Package rules is named as a metaphor of the rulebook included with tabletop games; this package
// is where the rules of Wisconsin Chess are implemented. Most importantly these are the
// calculation of moves available for a player given a board position and the results of making
// a move.
//
// Defined here is the board and addressing squares on the board, lists of moves in terms of that
// addressing, the orientation of a player or piece (white or black), and a representation of a
// piece on the board.
//
// Package piece is imported by this package and separately defines the variety of pieces with
// individual moves and characteristics.
package rules
