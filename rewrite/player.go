package main

import (
	"database/sql"

	"github.com/pciet/wichess/rules"
)

func NewPlayer(tx *sql.Tx, name, crypt string) {
	result, err := tx.Exec(PlayerNewInsert,
		name, crypt,
		rules.RandomSpecialPieceKind(), rules.RandomSpecialPieceKind())
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
}

func PlayerPiecePicks(tx *sql.Tx, name string) (left, right rules.PieceKind) {
	err := tx.QueryRow(PlayerPiecePicksQuery, name).Scan(&left, &right)
	if err != nil {
		Panic(err)
	}
	return
}
