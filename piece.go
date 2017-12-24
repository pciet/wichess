// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	prand "math/rand"

	"github.com/pciet/wichess/wichessing"
)

const (
	piece_table = "pieces"

	piece_id_key       = "piece_id"
	piece_kind_key     = "kind"
	piece_owner_key    = "owner"
	piece_reserved_key = "reserved"
	piece_taken_key    = "taken"

	free_piece_slice_hint = 8
)

type piece struct {
	wichessing.Piece
	Identifier int
}

// Reserved pieces are in a game where taking doesn't remove the player's piece. Other games that do remove a player's piece can be played concurrently. If a piece is reserved then the taken field is set if the piece is to be removed. When the reserving game is deleted then all pieces that are reserved and taken and now have a reserved count of zero are removed.
// This function expects all identifiers in the setup array to point to valid and correctly owned pieces.
func (db DB) reservePieces(setup gameSetup) {
	tx := db.Begin()
	for _, id := range setup {
		if id == 0 {
			continue
		}
		result, err := tx.Exec("UPDATE "+piece_table+" SET "+piece_reserved_key+" = "+piece_reserved_key+" + 1 WHERE "+piece_id_key+" = $1;", id)
		if err != nil {
			panic(err.Error())
		}
		count, err := result.RowsAffected()
		if err != nil {
			panic(err.Error())
		}
		if count != 1 {
			panic(fmt.Sprint(count, "rows affected by single piece reserve"))
		}
	}
	tx.Commit()
}

func (db DB) unreservePieces(setup gameSetup) {
	tx := db.Begin()
	for _, id := range setup {
		if id == 0 {
			continue
		}
		tx.unreservePiece(id)
	}
	tx.Commit()
}

func (tx TX) unreservePiece(id int) {
	if id == 0 {
		return
	}
	var taken bool
	var reserveCount int
	err := tx.QueryRow("SELECT "+piece_taken_key+", "+piece_reserved_key+" FROM "+piece_table+" WHERE "+piece_id_key+" = $1 FOR UPDATE;", id).Scan(&taken, &reserveCount)
	if err != nil {
		panic(err.Error())
	}
	if taken && (reserveCount == 1) {
		result, err := tx.Exec("DELETE FROM "+piece_table+" WHERE "+piece_id_key+" = $1;", id)
		if err != nil {
			panic(err.Error())
		}
		count, err := result.RowsAffected()
		if err != nil {
			panic(err.Error())
		}
		if count != 1 {
			panic(fmt.Sprint(count, "rows affected by piece delete"))
		}
	} else {
		result, err := tx.Exec("UPDATE "+piece_table+" SET "+piece_reserved_key+" = "+piece_reserved_key+" - 1 WHERE "+piece_id_key+" = $1;", id)
		if err != nil {
			panic(err.Error())
		}
		count, err := result.RowsAffected()
		if err != nil {
			panic(err.Error())
		}
		if count != 1 {
			panic(fmt.Sprint(count, "rows affected by single piece unreserve"))
		}
	}
}

// If the piece has a reserve count of one or more then the piece is marked as taken and will be removed when the reserve count is reduced to zero by unreservePieces.
func (db DB) removePiece(id int) {
	if id == 0 {
		return
	}
	tx := db.Begin()
	var reserveCount int
	err := tx.QueryRow("SELECT "+piece_reserved_key+" FROM "+piece_table+" WHERE "+piece_id_key+" = $1 FOR UPDATE;", id).Scan(&reserveCount)
	if err != nil {
		panic(err.Error())
	}
	if reserveCount > 1 {
		result, err := tx.Exec("UPDATE "+piece_table+" SET "+piece_taken_key+" = $1 WHERE "+piece_id_key+" = $2;", true, id)
		if err != nil {
			panic(err.Error())
		}
		count, err := result.RowsAffected()
		if err != nil {
			panic(err.Error())
		}
		if count != 1 {
			panic(fmt.Sprint(count, "rows affected by piece delete"))
		}
	} else {
		result, err := tx.Exec("DELETE FROM "+piece_table+" WHERE "+piece_id_key+" = $1;", id)
		if err != nil {
			panicExit(err.Error())
		}
		count, err := result.RowsAffected()
		if err != nil {
			panicExit(err.Error())
		}
		if count != 1 {
			panicExit(fmt.Sprintf("%v rows affected by piece delete exec", count))
		}
	}
	tx.Commit()
}

// A zero ID means use the basic piece for the specified kind.
func (db DB) pieceWithID(id int, kind wichessing.Kind, orientation wichessing.Orientation, owner string) piece {
	if id == 0 {
		return piece{
			Piece: wichessing.Piece{
				Base:        wichessing.BaseForKind(kind),
				Kind:        kind,
				Orientation: orientation,
			},
		}
	}
	var dbowner string
	var ki int
	var taken bool
	err := db.QueryRow("SELECT "+piece_owner_key+", "+piece_kind_key+", "+piece_taken_key+" FROM "+piece_table+" WHERE "+piece_id_key+"=$1;", id).Scan(&dbowner, &ki, &taken)
	if (err != nil) || (owner != dbowner) || taken {
		// an invalid ID requested results in no piece
		return piece{}
	}
	base := wichessing.BaseForKind(wichessing.Kind(ki))
	if base != kind {
		if debug {
			fmt.Println("pieceWithID: piece does not match default base (assigned to wrong square)")
		}
		// to catch cases of putting a piece in the wrong square
		return piece{}
	}
	return piece{
		Identifier: id,
		Piece: wichessing.Piece{
			Base:        base,
			Kind:        wichessing.Kind(ki),
			Orientation: orientation,
		},
	}
}

func (db DB) playersFreePieces(name string) []piece {
	rows, err := db.Query("SELECT "+piece_id_key+", "+piece_kind_key+" FROM "+piece_table+" WHERE "+piece_owner_key+"=$1 AND "+piece_taken_key+"=$2;", name, false)
	if err != nil {
		panicExit(err.Error())
	}
	defer rows.Close()
	pieces := make([]piece, 0, free_piece_slice_hint)
	i := 0
	for rows.Next() {
		pieces = append(pieces, piece{})
		err = rows.Scan(&pieces[i].Identifier, &pieces[i].Kind)
		if err != nil {
			panicExit(err.Error())
		}
		pieces[i].Base = wichessing.BaseForKind(pieces[i].Kind)
		i++
	}
	err = rows.Err()
	if err != nil {
		panicExit(err.Error())
	}
	return pieces
}

func (db DB) newPiece(kind int, player string) {
	_, err := db.Exec("INSERT INTO "+piece_table+"("+piece_kind_key+", "+piece_owner_key+", "+piece_reserved_key+", "+piece_taken_key+") VALUES ($1, $2, $3, $4);", kind, player, 0, false)
	if err != nil {
		panicExit(err.Error())
	}
}

var random *prand.Rand

func init() {
	seed, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err.Error())
	}
	random = prand.New(prand.NewSource(seed.Int64()))
}

func randomHeroPiece() piece {
	return piece{
		Piece: wichessing.Piece{
			Kind: wichessing.Kind(random.Int63n(34) + 1 + 6),
		},
	}
}
