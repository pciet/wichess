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
	err := db.QueryRow("SELECT "+piece_owner_key+", "+piece_kind_key+" FROM "+piece_table+" WHERE "+piece_id_key+"=$1;", id).Scan(&dbowner, &ki)
	if (err != nil) || (owner != dbowner) {
		// an invalid ID requested results in no piece
		return piece{}
	}
	return piece{
		Identifier: id,
		Piece: wichessing.Piece{
			Base:        wichessing.BaseForKind(wichessing.Kind(ki)),
			Kind:        wichessing.Kind(ki),
			Orientation: orientation,
		},
	}
}

func (db DB) playersFreePieces(name string) []piece {
	rows, err := db.Query("SELECT "+piece_id_key+", "+piece_kind_key+", "+piece_takes_key+" FROM "+piece_table+" WHERE "+piece_owner_key+"=$1 AND "+piece_ingame_key+"=$2;", name, false)
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
		pieces[i].Base = wichessing.BaseForKind(pieces[i].Kind)
		i++
	}
	err = rows.Err()
	if err != nil {
		panicExit(err.Error())
	}
	return pieces
}

func (db DB) playersBestPiece(name string) piece {
	rows, err := db.Query("SELECT "+piece_kind_key+", "+piece_takes_key+" FROM "+piece_table+" WHERE "+piece_owner_key+"=$1;", name)
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

func (db DB) createNewPlayersPieces(name string) {
	for i := 0; i < initial_piece_count; i++ {
		piece := randomHeroPiece()
		_, err := db.Exec("INSERT INTO "+piece_table+"("+piece_kind_key+", "+piece_owner_key+", "+piece_takes_key+", "+piece_ingame_key+") VALUES ($1, $2, $3, $4);", piece.Kind, name, 0, false)
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
