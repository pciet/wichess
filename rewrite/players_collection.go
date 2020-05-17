package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/lib/pq"
	"github.com/pciet/wichess/rules"
)

func PlayerCollection(tx *sql.Tx, playerID int) Collection {
	var values []sql.NullInt64
	err := tx.QueryRow(PlayersCollectionQuery, playerID).Scan(pq.Array(&values))
	if err != nil {
		DebugPrintln(PlayersCollectionQuery, playerID)
		Panic(err)
	}

	if len(values) != CollectionCount {
		Panic(playerID, "bad collection length", len(values))
	}

	var c Collection
	for i, v := range values {
		if v.Valid == false {
			Panic(playerID, "sql null at", i)
		}
		c[i] = EncodedPiece(v.Int64).Decode()
	}
	return c
}

// PlayersSelectedCollectionPieces returns a decoded subset of all collection pieces for a player.
// The kinds of the left and right picks are always returned. If a requested collection slot is
// outside the array range then a panic occurs.
func PlayerSelectedCollectionPieces(tx *sql.Tx, playerID int,
	slots []CollectionSlot) ([]Piece, rules.PieceKind, rules.PieceKind) {

	var s strings.Builder
	s.WriteString("SELECT ")
	for _, slot := range slots {
		s.WriteString(PlayersCollection)
		s.WriteRune('[')
		s.WriteString(strconv.Itoa(slot.Int()))
		s.WriteRune(']')
		s.WriteString(", ")
	}

	s.WriteString(PlayersLeftKind)
	s.WriteString(", ")
	s.WriteString(PlayersRightKind)

	s.WriteString(" FROM ")
	s.WriteString(PlayersTable)
	s.WriteString(" WHERE ")
	s.WriteString(PlayersIdentifier)
	s.WriteString("=$1;")

	if DebugSQL {
		DebugPrintln(s.String())
	}

	var collectionValues []sql.NullInt64
	var left, right rules.PieceKind
	values := make([]interface{}, 0, 3)
	if len(slots) != 0 {
		values = append(values, pq.Array(&collectionValues))
	}
	values = append(values, &left)
	values = append(values, &right)

	err := tx.QueryRow(s.String(), playerID).Scan(values...)
	if err != nil {
		DebugPrintln(s.String())
		DebugPrintln("player ID =", playerID)
		Panic(err)
	}

	out := make([]Piece, len(slots))
	for i, p := range collectionValues {
		if p.Valid == false {
			Panic("invalid collection slot")
		}
		out[i] = EncodedPiece(p.Int64).Decode()
	}

	return out, left, right
}

// PlayerCollectionFlagInUse flags pieces in the player's collection to be in use, and it updates
// the pick piece slots if one or both slot is indicated to be used.
func PlayerCollectionFlagInUse(tx *sql.Tx, playerID int,
	slots []CollectionSlot, kinds []rules.PieceKind,
	left, right rules.PieceKind, replaceLeft, replaceRight bool) {

	updateLeft, updateRight := rules.NoKind, rules.NoKind
	if replaceLeft && replaceRight {
		updateLeft, updateRight = rules.TwoDifferentSpecialPieces()
	} else if replaceLeft {
		updateLeft = rules.DifferentSpecialPiece(right)
	} else if replaceRight {
		updateRight = rules.DifferentSpecialPiece(left)
	} else {
		if len(slots) == 0 {
			return
		}
	}

	var s strings.Builder
	s.WriteString("UPDATE ")
	s.WriteString(PlayersTable)
	s.WriteString(" SET ")

	i := 1

	// TODO: overlap with UpdateGame, have a generalized SQL func for dynamic updates

	placeholder := func(last bool) {
		s.WriteString("=$")
		s.WriteString(strconv.Itoa(i))
		if last == false {
			s.WriteString(", ")
		}
		i++
	}

	args := make([]interface{}, 0, 4)

	if updateLeft != rules.NoKind {
		args = append(args, updateLeft)
		s.WriteString(PlayersLeftKind)
		if (updateRight == rules.NoKind) && (len(slots) == 0) {
			placeholder(true)
		} else {
			placeholder(false)
		}
	}

	if updateRight != rules.NoKind {
		args = append(args, updateRight)
		s.WriteString(PlayersRightKind)
		if len(slots) == 0 {
			placeholder(true)
		} else {
			placeholder(false)
		}
	}

	for j, slot := range slots {
		args = append(args, Piece{
			InUse: true,
			Piece: rules.Piece{
				Kind: kinds[j],
			},
		}.Encode())
		s.WriteString(PlayersCollection)
		s.WriteRune('[')
		s.WriteString(strconv.Itoa(slot.Int()))
		s.WriteRune(']')
		if j == (len(slots) - 1) {
			placeholder(true)
		} else {
			placeholder(false)
		}
	}

	s.WriteString(" WHERE ")
	s.WriteString(PlayersIdentifier)
	s.WriteString("=$")
	s.WriteString(strconv.Itoa(i))
	s.WriteRune(';')
	args = append(args, playerID)

	if DebugSQL {
		fmt.Println(s.String())
		fmt.Println(args)
	}

	r, err := tx.Exec(s.String(), args...)
	if err != nil {
		DebugPrintln(s.String())
		DebugPrintln(args)
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
