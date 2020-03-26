package main

import (
	"database/sql"

	"github.com/pciet/wichess/rules"
)

const NewPlayerPieceCount = 3

func NewPlayer(tx *sql.Tx, name, crypt string) {
	result, err := tx.Exec(PlayerNewInsert,
		name,
		crypt,
		0, 0, 0,
		InitialRating,
		0, 0, 0, 0, 0, 0, 0, 0)
	if err != nil {
		Panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		Panic(err)
	}
	if count != 1 {
		Panic(count, "rows affected by new player insert for", name)
	}

	for i := 0; i < NewPlayerPieceCount; i++ {
		InsertNewPiece(tx, rules.RandomSpecialPieceKind(), name)
	}
}
