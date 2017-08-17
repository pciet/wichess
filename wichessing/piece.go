// Copyright 2017 Matthew Juran
// All Rights Reserved

package wichessing

import ()

type Piece struct {
	Kind
	Takes int
}

func CopyFromPiece(from Piece, to *Piece) {
	to.Kind = from.Kind
	to.Takes = from.Takes
}

type Kind int

const (
	King Kind = iota + 1
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

func NameForKind(piece Kind) string {
	switch piece {
	case King:
		return "King"
	case Queen:
		return "Queen"
	case Rook:
		return "Rook"
	case Bishop:
		return "Bishop"
	case Knight:
		return "Knight"
	case Pawn:
		return "Pawn"
	}
	return ""
}
