package main

import (
	"database/sql"

	"github.com/pciet/wichess/rules"
)

func LoadGame(tx *sql.Tx, id GameIdentifier) Game {
	return Game{
		Header: LoadGameHeader(tx, id),
		Board:  LoadGameBoard(tx, id),
	}
}

// LoadGameBoard loads from the database and prepares a Board.
// If the game doesn't exist then the PieceIdentifiers field is nil.
func LoadGameBoard(tx *sql.Tx, id GameIdentifier) Board {
	var ep [64]EncodedPiece
	epp := make([]interface{}, 64)
	for i, _ := range ep {
		epp[i] = &(ep[i])
	}

	err := tx.QueryRow(GamesBoardQuery, id).Scan(epp...)
	if err == sql.ErrNoRows {
		return Board{PieceIdentifiers: nil}
	} else if err != nil {
		Panic(err)
	}

	b := Board{PieceIdentifiers: make([]AddressedPieceIdentifier, 0, 8)}

	for i, v := range ep {
		p := v.Decode()
		b.Board[i] = rules.Square(p.Piece.ApplyCharacteristics())
		if p.ID != 0 {
			b.PieceIdentifiers = append(b.PieceIdentifiers,
				AddressedPieceIdentifier{
					ID:      p.ID,
					Address: rules.AddressIndex(i).Address(),
				})
		}
	}

	return b
}

// LoadGameHeader gets the header from the database. If the header isn't found
// then the ID field is 0.
func LoadGameHeader(tx *sql.Tx, id GameIdentifier) GameHeader {
	h := GameHeader{ID: id}
	err := tx.QueryRow(GamesHeaderQuery, id).Scan(
		&h.PrizePiece,
		&h.Competitive,
		&h.Recorded,
		&h.Conceded,
		&h.White.Name,
		&h.White.Acknowledge,
		&h.White.LatestMove,
		&h.White.Elapsed,
		&h.White.ElapsedUpdated,
		&h.Black.Name,
		&h.Black.Acknowledge,
		&h.Black.LatestMove,
		&h.Black.Elapsed,
		&h.Black.ElapsedUpdated,
		&h.Active,
		&h.PreviousActive,
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
	return h
}

func (a GameIdentifier) Int() int { return int(a) }
