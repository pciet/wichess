package main

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/pciet/wichess/rules"
)

// NewPlayer inserts a database row then returns the player's ID.
func NewPlayer(tx *sql.Tx, name, crypt string) int {
	var id int
	err := tx.QueryRow(PlayersNewInsert,
		name, crypt,
		rules.RandomSpecialPieceKind(), rules.RandomSpecialPieceKind(),
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
