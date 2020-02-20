package main

import (
	"database/sql"
	"log"

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
		log.Panicln("failed to query database row:", err)
	}

	if basicPiece != rules.BasicKind(p.Kind) {
		DebugPrintln("requested piece with wrong basic kind", rules.BasicKind(p.Kind), "for square with base", basicPiece)
		return Piece{}
	}

	return p
}

func InsertNewPiece(tx *sql.Tx, k rules.PieceKind, owner string) {
	result, err := tx.Exec(piece_insert, k, owner, 0, false)
	if err != nil {
		log.Panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Panic(err)
	}
	if count != 1 {
		log.Panicln(count, "rows affected by new piece insert for", owner)
	}
}

func (p Piece) Encode() EncodedPiece {
	var enc uint64
	enc |= (uint64(p.ID) & encoded_piece_identifier_mask) << encoded_piece_identifier_bit
	enc |= (uint64(p.Orientation) & encoded_piece_orientation_mask) << encoded_piece_orientation_bit
	enc |= (uint64(btoi(p.Moved)) & encoded_piece_moved_mask) << encoded_piece_moved_bit
	enc |= (uint64(p.Kind) & encoded_piece_kind_mask) << encoded_piece_kind_bit
	return EncodedPiece(enc)
}

func (e EncodedPiece) Decode() Piece {
	return Piece{
		ID: PieceIdentifier((e >> encoded_piece_identifier_bit) & encoded_piece_identifier_mask),
		Piece: rules.Piece{
			Orientation: rules.Orientation((e >> encoded_piece_orientation_bit) & encoded_piece_orientation_mask),
			Moved:       itob(int((e >> encoded_piece_moved_bit) & encoded_piece_moved_mask)),
			Kind:        rules.PieceKind((e >> encoded_piece_kind_bit) & encoded_piece_kind_mask),
		},
	}
}

func btoi(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func itob(i int) bool {
	if i == 0 {
		return false
	} else {
		return true
	}
}
