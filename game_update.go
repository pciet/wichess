package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/pciet/wichess/rules"
)

// UpdateGame puts board changes into the database for a game, updates the latest move, sets
// the draw turn count, sets the active player, adds any taken pieces to the captured pieces lists,
// and increments the turn number.
// If the previous move shouldn't be updated, such as when the promotion pick is done, then m is
// set to rules.NoMove.
func UpdateGame(tx *sql.Tx, id GameIdentifier, active, previousActive rules.Orientation,
	drawTurns int, turn int, m rules.Move, with []AddressedPiece, takes []CapturedPiece) {
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
		s.WriteString(GamesBoard)
		s.WriteRune('[')
		// Postgres arrays are indexed 1-(n+1) instead of 0-n
		s.WriteString(strconv.Itoa(p.Address.Index().Int() + 1))
		s.WriteRune(']')
		placeholder(false)
	}

	for _, p := range takes {
		args = append(args, p.Kind)
		if p.Orientation == rules.Black {
			s.WriteString(GamesWhiteCaptures)
		} else if p.Orientation == rules.White {
			s.WriteString(GamesBlackCaptures)
		} else {
			Panic("orientation", p.Orientation, "not white or black")
		}
		s.WriteRune('[')
		s.WriteString(strconv.Itoa(p.CaptureSlot))
		s.WriteRune(']')
		placeholder(false)
	}

	args = append(args, previousActive)
	s.WriteString(GamesPreviousActive)
	placeholder(false)

	args = append(args, active)
	s.WriteString(GamesActive)
	placeholder(false)

	if m != rules.NoMove {
		// the move is recorded for future en passant calculation
		args = append(args, m.From.Index())
		s.WriteString(GamesMoveFrom)
		placeholder(false)

		args = append(args, m.To.Index())
		s.WriteString(GamesMoveTo)
		placeholder(false)
	}

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
		fmt.Println(args)
	}

	SQLExecRow(tx, s.String(), args...)
}
