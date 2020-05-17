package main

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/pciet/wichess/rules"
)

// TODO: change players table database access to use ID when available

type Player struct {
	Name string
	ID   int
}

// NewPlayer inserts a database row then returns the player's ID.
func NewPlayer(tx *sql.Tx, name, crypt string) int {
	left, right := rules.TwoDifferentSpecialPieces()
	var id int
	err := tx.QueryRow(PlayersNewInsert,
		name, crypt, left, right,
		pq.Array([CollectionCount]EncodedPiece{}),
		0, 0,
	).Scan(&id)
	if err != nil {
		Panic(err)
	}
	return id
}

func PlayerName(tx *sql.Tx, playerID int) string {
	var name string
	err := tx.QueryRow(PlayersNameQuery, playerID).Scan(&name)
	if err == sql.ErrNoRows {
		return ""
	} else if err != nil {
		Panic(err)
	}
	return name
}

// PlayerID returns -1 if the name doesn't match a database row.
func PlayerID(tx *sql.Tx, name string) int {
	var id int
	err := tx.QueryRow(PlayersIdentifierQuery, name).Scan(&id)
	if err == sql.ErrNoRows {
		DebugPrintln(PlayersIdentifierQuery, name)
		return -1
	} else if err != nil {
		Panic(err)
	}
	return id
}

func PlayerPiecePicks(tx *sql.Tx, name string) (left, right rules.PieceKind) {
	err := tx.QueryRow(PlayersPiecePicksQuery, name).Scan(&left, &right)
	if err != nil {
		Panic(err)
	}
	return
}
