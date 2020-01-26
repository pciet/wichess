package main

import (
	"database/sql"

	"github.com/pciet/wichess/rules"
)

// A 0 id returns a piece of the basicPiece kind but doesn't query the database.
// A returned Piece of NoKind indicates an invalid request.
// The correct initial square placement is checked for here by comparing basicPiece to the basic kind of the piece in the database.
func LoadPiece(tx *sql.Tx, id PieceIdentifier, basicPiece rules.PieceKind, o rules.Orientation, owner string) Piece {
	p := Piece{
		ID: id,
		Piece: rules.Piece{
			Kind:        basicPiece,
			Orientation: o,
		},
	}

	if id == 0 {
		return p
	}

	var player string
	var taken bool
	err := tx.QueryRow(piece_query, id).Scan(&player, &p.Piece.Kind, &taken)
	if err == sql.ErrNoRows || owner != player || taken {
		DebugPrintln("invalid request by", owner, "for piece with ID", id)
		return Piece{}
	} else if err != nil {
		panic("failed to query database row:", err)
	}

	if basicPiece != rules.BasicKind(p.Kind) {
		DebugPrintln("requested piece with wrong basic kind", rules.BasicKind(p.Kind), "for square with base", basicPiece)
		return Piece{}
	}

	return p
}
