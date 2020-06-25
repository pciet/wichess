package main

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

func PlayerPiecePicks(tx *sql.Tx, name string) (left, right piece.Kind) {
	err := tx.QueryRow(PlayersPiecePicksQuery, name).Scan(&left, &right)
	if err != nil {
		Panic(err)
	}
	return
}

func AddGamePicksToPlayerCollection(tx *sql.Tx, pl Player, gameID GameIdentifier) {
	left, right := PicksInGameForPlayer(tx, gameID, pl.Name)
	need1 := false
	if (left == piece.NoKind) && (right == piece.NoKind) {
		return
	} else if (left == piece.NoKind) || (right == piece.NoKind) {
		need1 = true
	}

	c := PlayerCollection(tx, pl.ID)

	indexA, indexB := -1, -1
	for i, p := range c {
		if p.Kind != piece.NoKind {
			continue
		}
		if indexA == -1 {
			indexA = i
			if need1 {
				break
			}
			continue
		}
		indexB = i
		break
	}

	if indexA == -1 {
		// collection is full, don't add
		return
	}

	if (need1 == false) && (indexB == -1) {
		// want to add two pieces but only one slot available, just add one
		need1 = true
	}

	// +1 to index to correctly address SQL array index
	indexA++
	indexB++

	update := func(index int, kind piece.Kind) {
		arrStr := PlayersCollection + "[" + strconv.Itoa(index) + "]"
		u := SQLUpdate(PlayersTable, arrStr, PlayersIdentifier)
		piece := Piece{
			Piece: rules.Piece{
				Kind: kind,
			},
		}.Encode()
		r, err := tx.Exec(u, piece, pl.ID)
		if err != nil {
			DebugPrintln(u)
			DebugPrintln(kind, piece, pl.ID)
			Panic(err)
		}
		c, err := r.RowsAffected()
		if err != nil {
			Panic(err)
		}
		if c != 1 {
			Panic(c, "rows affected by", u)
		}
	}

	if need1 {
		var k piece.Kind
		if left == piece.NoKind {
			k = right
		} else {
			k = left
		}
		update(indexA, k)
	} else {
		update(indexA, left)
		update(indexB, right)
	}
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

	var collectionValues []sql.NullInt64
	var left, right piece.Kind
	values := make([]interface{}, 0, 3)
	if len(slots) != 0 {
		values = append(values, pq.Array(&collectionValues))
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

// PlayerCollectionFlagInUse flags pieces in the player's collection to be in use, and it updates
// the pick piece slots if one or both slot is indicated to be used.
func PlayerCollectionFlagInUse(tx *sql.Tx, id PlayerIdentifier,
	slots []CollectionSlot, kinds []piece.Kind,
	left, right piece.Kind, replaceLeft, replaceRight bool) {

	updateLeft, updateRight := piece.NoKind, piece.NoKind
	if replaceLeft && replaceRight {
		updateLeft, updateRight = piece.TwoDifferentSpecialKinds()
	} else if replaceLeft {
		updateLeft = right.DifferentSpecialKind()
	} else if replaceRight {
		updateRight = left.DifferentSpecialKind()
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

	if updateLeft != piece.NoKind {
		args = append(args, updateLeft)
		s.WriteString(PlayersLeftKind)
		if (updateRight == piece.NoKind) && (len(slots) == 0) {
			placeholder(true)
		} else {
			placeholder(false)
		}
	}

	if updateRight != piece.NoKind {
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
	args = append(args, id)

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
