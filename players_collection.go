package wichess

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/lib/pq"
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// TODO: this file is too long

func PlayerCollectionAdd(tx *sql.Tx, id PlayerIdentifier, s CollectionSlot, p piece.Kind) {
	// TODO: should collection pieces be just the kind?
	SQLExecRow(tx, SQLUpdate(PlayersTable,
		PlayersCollection+"["+strconv.Itoa(int(s))+"]", PlayersIdentifier),
		Piece{Piece: rules.Piece{Kind: p}}.Encode(), id)
}

func PlayerPiecePicks(tx *sql.Tx, name string) (left, right piece.Kind) {
	err := tx.QueryRow(PlayersPiecePicksQuery, name).Scan(&left, &right)
	if err != nil {
		Panic(err)
	}
	return
}

func PlayerCollection(tx *sql.Tx, id PlayerIdentifier) Collection {
	var values []sql.NullInt64
	err := tx.QueryRow(PlayersCollectionQuery, id).Scan(pq.Array(&values))
	if err != nil {
		DebugPrintln(PlayersCollectionQuery, id)
		Panic(err)
	}

	if len(values) != CollectionCount {
		Panic(id, "bad collection length", len(values))
	}

	var c Collection
	for i, v := range values {
		if v.Valid == false {
			Panic(id, "sql null at", i)
		}
		c[i] = EncodedPiece(v.Int64).Decode()
	}
	return c
}

// PlayersSelectedCollectionPieces returns a decoded subset of all collection pieces for a player.
// The kinds of the left and right picks are always returned. If a requested collection slot is
// outside the array range then a panic occurs.
func PlayerSelectedCollectionPieces(tx *sql.Tx, id PlayerIdentifier,
	slots []CollectionSlot) ([]Piece, piece.Kind, piece.Kind) {

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

	collectionValues := make([]sql.NullInt64, len(slots))
	var left, right piece.Kind
	values := make([]interface{}, 0, 3)
	for i := 0; i < len(slots); i++ {
		values = append(values, &(collectionValues[i]))
	}
	values = append(values, &left)
	values = append(values, &right)

	err := tx.QueryRow(s.String(), id).Scan(values...)
	if err != nil {
		DebugPrintln(s.String())
		DebugPrintln("player ID =", id)
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

func PlayerCollectionReplacePicks(tx *sql.Tx, id PlayerIdentifier,
	left, right piece.Kind, replaceLeft, replaceRight bool) {

	if (replaceLeft == false) && (replaceRight == false) {
		return
	}

	updateLeft, updateRight := piece.NoKind, piece.NoKind
	if replaceLeft && replaceRight {
		updateLeft, updateRight = piece.TwoDifferentSpecialKinds()
	} else if replaceLeft {
		updateLeft = right.DifferentSpecialKind()
	} else if replaceRight {
		updateRight = left.DifferentSpecialKind()
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

	if updateLeft != piece.NoKind {
		args = append(args, updateLeft)
		s.WriteString(PlayersLeftKind)
		if updateRight == piece.NoKind {
			placeholder(true)
		} else {
			placeholder(false)
		}
	}

	if updateRight != piece.NoKind {
		args = append(args, updateRight)
		s.WriteString(PlayersRightKind)
		placeholder(true)
	}

	s.WriteString(" WHERE ")
	s.WriteString(PlayersIdentifier)
	s.WriteString("=$")
	s.WriteString(strconv.Itoa(i))
	s.WriteRune(';')
	args = append(args, id)

	if DebugSQL {
		fmt.Println(s.String())
		fmt.Println(args)
	}

	SQLExecRow(tx, s.String(), args...)
}
