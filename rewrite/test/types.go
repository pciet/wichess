package test

import "github.com/pciet/wichess/rules"

type (
	AddressedSquare = rules.AddressedSquare
	Square          = rules.Square
	Address         = rules.Address
	Piece           = rules.Piece
	Move            = rules.Move
	MoveSet         = rules.MoveSet
	Board           = rules.Board
	Orientation     = rules.Orientation
	State           = rules.State
)

const (
	White = rules.White
	Black = rules.Black

	Pawn   = rules.Pawn
	Knight = rules.Knight
	Bishop = rules.Bishop
	Rook   = rules.Rook
	Queen  = rules.Queen
	King   = rules.King

	Normal    = rules.Normal
	Promotion = rules.Promotion
	Check     = rules.Check
	Checkmate = rules.Checkmate
	Draw      = rules.Draw
)
