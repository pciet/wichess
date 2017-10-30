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

	piece_id_key     = "piece_id"
	piece_kind_key   = "kind"
	piece_owner_key  = "owner"
	piece_ingame_key = "ingame"

	free_piece_slice_hint = 8
)

type piece struct {
	wichessing.Piece
	Identifier int
	Ingame     bool
}

func (db DB) releasePieceFromGame(id int) {
	_, err := db.Exec("UPDATE "+piece_table+" SET "+piece_ingame_key+" = $1 WHERE "+piece_id_key+" = $2;", false, id)
	if err != nil {
		panicExit(err.Error())
	}
}

func (db DB) markPieceIngame(id int) {
	_, err := db.Exec("UPDATE "+piece_table+" SET "+piece_ingame_key+" = $1 WHERE "+piece_id_key+" = $2;", true, id)
	if err != nil {
		panicExit(err.Error())
	}
}

func (db DB) removePiece(id int) {
	if id == 0 {
		return
	}
	result, err := db.Exec("DELETE FROM "+piece_table+" WHERE "+piece_id_key+" = $1;", id)
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
	var ig bool
	err := db.QueryRow("SELECT "+piece_owner_key+", "+piece_kind_key+", "+piece_ingame_key+" FROM "+piece_table+" WHERE "+piece_id_key+"=$1;", id).Scan(&dbowner, &ki, &ig)
	if (err != nil) || (owner != dbowner) {
		// an invalid ID requested results in no piece
		return piece{}
	}
	return piece{
		Identifier: id,
		Ingame:     ig,
		Piece: wichessing.Piece{
			Base:        wichessing.BaseForKind(wichessing.Kind(ki)),
			Kind:        wichessing.Kind(ki),
			Orientation: orientation,
		},
	}
}

func (db DB) playersFreePieces(name string) []piece {
	rows, err := db.Query("SELECT "+piece_id_key+", "+piece_kind_key+" FROM "+piece_table+" WHERE "+piece_owner_key+"=$1 AND "+piece_ingame_key+"=$2;", name, false)
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

func (db DB) createNewPlayersPieces(name string) {
	for i := 0; i < initial_piece_count; i++ {
		piece := randomHeroPiece()
		_, err := db.Exec("INSERT INTO "+piece_table+"("+piece_kind_key+", "+piece_owner_key+", "+piece_ingame_key+") VALUES ($1, $2, $3);", piece.Kind, name, false)
		if err != nil {
			panicExit(err.Error())
		}
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
