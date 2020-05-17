package main

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/pciet/wichess/rules"
)

func LoadGame(tx *sql.Tx, id GameIdentifier, forUpdate bool) Game {
	return Game{
		Header: LoadGameHeader(tx, id, forUpdate),
		Board:  LoadGameBoard(tx, id, forUpdate),
	}
}

// LoadGameBoard loads from the database and prepares a Board.
// If the game doesn't exist then the PieceIdentifiers field is nil.
func LoadGameBoard(tx *sql.Tx, id GameIdentifier, forUpdate bool) Board {
	var query string
	if forUpdate {
		query = GamesBoardForUpdateQuery
	} else {
		query = GamesBoardQuery
	}

	var values []sql.NullInt64
	err := tx.QueryRow(query, id).Scan(pq.Array(&values))
	if err == sql.ErrNoRows {
		return Board{}
	} else if err != nil {
		Panic(err)
	}

	if len(values) != 64 {
		Panic(id, "read board with length", len(values))
	}

	b := Board{CollectionPieces: make([]AddressedCollectionSlot, 0, 8)}

	for i, v := range values {
		if v.Valid == false {
			Panic(id, "sql null at", i)
		}
		p := EncodedPiece(v.Int64).Decode()
		b.Board[i] = rules.Square(p.Piece.ApplyCharacteristics())
		if p.Slot != 0 {
			b.CollectionPieces = append(b.CollectionPieces, AddressedCollectionSlot{
				Slot:    p.Slot,
				Address: rules.AddressIndex(i).Address(),
			})
		}
	}

	return b
}

// LoadGameHeader gets the header from the database.
// If the header isn't found then the ID field is 0.
func LoadGameHeader(tx *sql.Tx, id GameIdentifier, forUpdate bool) GameHeader {
	var query string
	if forUpdate {
		query = GamesHeaderForUpdateQuery
	} else {
		query = GamesHeaderQuery
	}

	h := GameHeader{ID: id}
	var active, previousActive bool
	err := tx.QueryRow(query, id).Scan(
		&h.Conceded,
		&h.White.Name,
		&h.White.Acknowledge,
		&h.White.Left,
		&h.White.Right,
		&h.Black.Name,
		&h.Black.Acknowledge,
		&h.Black.Left,
		&h.Black.Right,
		&active,
		&previousActive,
		&h.From,
		&h.To,
		&h.DrawTurns,
		&h.Turn)
	if err == sql.ErrNoRows {
		DebugPrintln("found no games with id", id)
		h.ID = 0
	} else if err != nil {
		Panic("failed to query database row:", err)
	}
	h.Active = rules.BoolToOrientation(active)
	h.PreviousActive = rules.BoolToOrientation(previousActive)
	return h
}

func (a GameIdentifier) Int() int { return int(a) }
