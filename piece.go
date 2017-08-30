// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"crypto/rand"
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
	piece_takes_key  = "takes"
	piece_ingame_key = "ingame"

	free_piece_slice_hint = 8
)

type piece struct {
	wichessing.Piece
	Identifier int
	Takes      int
}

const free_pieces_query = "SELECT " + piece_id_key + ", " + piece_kind_key + ", " + piece_takes_key + " FROM " + piece_table + " WHERE " + piece_owner_key + "=$1 AND " + piece_ingame_key + "=$2"

func freePiecesForPlayerFromDatabase(name string) []piece {
	rows, err := database.Query(free_pieces_query, name, false)
	if err != nil {
		panicExit(err.Error())
	}
	defer rows.Close()
	pieces := make([]piece, 0, free_piece_slice_hint)
	i := 0
	for rows.Next() {
		pieces = append(pieces, piece{})
		err = rows.Scan(&pieces[i].Identifier, &pieces[i].Kind, &pieces[i].Takes)
		if err != nil {
			panicExit(err.Error())
		}
		i++
	}
	err = rows.Err()
	if err != nil {
		panicExit(err.Error())
	}
	return pieces
}

const players_best_pieces_query = "SELECT " + piece_kind_key + ", " + piece_takes_key + " FROM " + piece_table + " WHERE " + piece_owner_key + "=$1"

func bestPieceForPlayerFromDatabase(name string) piece {
	rows, err := database.Query(players_best_pieces_query, name)
	if err != nil {
		panicExit(err.Error())
	}
	defer rows.Close()
	p := piece{}
	newp := piece{}
	for rows.Next() {
		err = rows.Scan(&newp.Kind, &newp.Takes)
		if err != nil {
			panicExit(err.Error())
		}
		if newp.Takes > p.Takes {
			p.Takes = newp.Takes
			p.Kind = newp.Kind
		}
	}
	err = rows.Err()
	if err != nil {
		panicExit(err.Error())
	}
	return p
}

const new_piece_insert = "INSERT INTO " + piece_table + "(" + piece_kind_key + ", " + piece_owner_key + ", " + piece_takes_key + ", " + piece_ingame_key + ") VALUES ($1, $2, $3, $4)"

func newPlayerPiecesIntoDatabase(name string) {
	for i := 0; i < initial_piece_count; i++ {
		piece := randomHeroPiece()
		_, err := database.Exec(new_piece_insert, piece.Kind, name, 0, false)
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
			Kind: wichessing.Kind(random.Int63n(9) + 1 + 6),
		},
	}
}

const piece_owner_query = "SELECT " + piece_owner_key + ", " + piece_kind_key + " FROM " + piece_table + " WHERE " + piece_id_key + "=$1"

// A zero ID means use the basic piece for the specified kind.
func pieceWithID(id int, kind wichessing.Kind, orientation wichessing.Orientation, owner string) piece {
	if id == 0 {
		return piece{
			Piece: wichessing.Piece{
				Kind:        kind,
				Orientation: orientation,
			},
		}
	}
	var dbowner string
	var ki int
	err := database.QueryRow(piece_owner_query, id).Scan(&dbowner, &ki)
	if (err != nil) || (owner != dbowner) {
		// an invalid ID requested results in no piece
		return piece{}
	}
	return piece{
		Identifier: id,
		Piece: wichessing.Piece{
			Kind:        wichessing.Kind(ki),
			Orientation: orientation,
		},
	}
}
