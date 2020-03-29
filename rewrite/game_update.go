package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/pciet/wichess/rules"
)

// UpdateGame puts board changes into the database for a game,
// updates the latest move, sets the draw turn count,
// swaps the active player, and increments the turn number.
func UpdateGame(tx *sql.Tx, id GameIdentifier,
	white string, black string, active string,
	drawTurns int, turn int, m rules.Move, with []AddressedPiece) {
	var s strings.Builder
	s.WriteString("UPDATE ")
	s.WriteString(GamesTable)
	s.WriteString(" SET ")

	i := 1

	placeholder := func(last bool) {
		s.WriteString("=$")
		s.WriteString(strconv.Itoa(i))
		if last == false {
			s.WriteString(", ")
		}
		i++
	}

	args := make([]interface{}, 0, 4)

	for _, p := range with {
		args = append(args, p.Piece.Encode())
		s.WriteRune('s')
		s.WriteString(strconv.Itoa(p.Address.Index().Int()))
		placeholder(false)
	}

	switch active {
	case white:
		args = append(args, black)
	case black:
		args = append(args, white)
	default:
		Panic("game", id, "active player",
			active, "not the white or black player")
	}
	s.WriteString(GamesActive)
	placeholder(false)

	// the move is recorded for future en passant calculation
	args = append(args, m.From.Index())
	s.WriteString(GamesFrom)
	placeholder(false)

	args = append(args, m.To.Index())
	s.WriteString(GamesTo)
	placeholder(false)

	// draw turns are reset or incremented depending on the move made
	args = append(args, drawTurns)
	s.WriteString(GamesDrawTurns)
	placeholder(false)

	args = append(args, turn+1)
	s.WriteString(GamesTurn)
	placeholder(true)

	args = append(args, strconv.Itoa(id.Int()))
	s.WriteString(" WHERE ")
	s.WriteString(GamesIdentifier)
	s.WriteString("=$")
	s.WriteString(strconv.Itoa(i))
	s.WriteRune(';')

	if DebugSQL {
		fmt.Println(s.String())
	}

	r, err := tx.Exec(s.String(), args...)
	if err != nil {
		Panic(err)
	}

	c, err := r.RowsAffected()
	if err != nil {
		Panic(err)
	}
	if c != 1 {
		Panic(c, "rows affected by", s.String())
	}
}
