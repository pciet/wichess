package main

import (
	"database/sql"

	"github.com/pciet/wichess/rules"
)

// LoadPiece is both a helper function for making a new game and
// the function to get piece information from the database.
// Correct initial square placement is checked for by comparing
// basicPiece to the basic kind of the piece in the database.
// If id is set to 0 then a piece of the basicPiece kind is returned
// and the database isn't queried.
// A Kind of NoKind in the returned piece indicates an invalid request.
func LoadPiece(tx *sql.Tx, id PieceIdentifier,
	basicPiece rules.PieceKind, o rules.Orientation, owner string) Piece {
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
	err := tx.QueryRow(PieceQuery, id).Scan(&player, &p.Piece.Kind, &taken)
	if (err == sql.ErrNoRows) || (owner != player) || taken {
		DebugPrintln("invalid request by", owner, "for piece with ID", id)
		return Piece{}
	} else if err != nil {
		Panic("failed to query database row:", err)
	}

	if basicPiece != rules.BasicKind(p.Kind) {
		DebugPrintln("bad kind", p.Kind, "for", basicPiece, "square")
		return Piece{}
	}

	return p
}

func InsertNewPiece(tx *sql.Tx, k rules.PieceKind, owner string) {
	result, err := tx.Exec(PieceInsert, k, owner, 0, false)
	if err != nil {
		Panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		Panic(err)
	}
	if count != 1 {
		Panic(count, "rows affected by new piece insert for", owner)
	}
}
