package wichess

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/pciet/wichess/piece"
)

// TODO: change players table database access to use ID when available

type PlayerIdentifier int

type Player struct {
	Name string
	ID   PlayerIdentifier
}

// NewPlayer inserts a database row then returns the player's ID.
func NewPlayer(tx *sql.Tx, name, crypt string) PlayerIdentifier {
	left, right := piece.TwoDifferentSpecialKinds()
	var id PlayerIdentifier
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

func PlayerName(tx *sql.Tx, id PlayerIdentifier) string {
	var name string
	err := tx.QueryRow(PlayersNameQuery, id).Scan(&name)
	if err == sql.ErrNoRows {
		return ""
	} else if err != nil {
		Panic(err)
	}
	return name
}

// PlayerID returns -1 if the name doesn't match a database row.
func PlayerID(tx *sql.Tx, name string) PlayerIdentifier {
	var id PlayerIdentifier
	err := tx.QueryRow(PlayersIdentifierQuery, name).Scan(&id)
	if err == sql.ErrNoRows {
		return -1
	} else if err != nil {
		Panic(err)
	}
	return id
}

func (a PlayerIdentifier) Int() int { return int(a) }
