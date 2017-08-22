// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
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

const free_pieces_query = "SELECT " + piece_id_key + ", " + piece_kind_key + ", " + piece_takes_key + " FROM " + piece_table + " WHERE " + piece_owner_key + "=$1 AND " + piece_ingame_key + "=$2"

func freePiecesForPlayerFromDatabase(name string) []wichessing.Piece {
	rows, err := database.Query(free_pieces_query, name, false)
	if err != nil {
		panicExit(err.Error())
	}
	defer rows.Close()
	pieces := make([]wichessing.Piece, 0, free_piece_slice_hint)
	i := 0
	for rows.Next() {
		pieces = append(pieces, wichessing.Piece{})
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

func bestPieceForPlayerFromDatabase(name string) wichessing.Piece {
	rows, err := database.Query(players_best_pieces_query, name)
	if err != nil {
		panicExit(err.Error())
	}
	defer rows.Close()
	p := wichessing.Piece{}
	newp := wichessing.Piece{}
	for rows.Next() {
		err = rows.Scan(&newp.Kind, &newp.Takes)
		if err != nil {
			panicExit(err.Error())
		}
		if newp.Takes > p.Takes {
			wichessing.CopyFromPiece(newp, &p)
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
		piece := wichessing.RandomPiece()
		_, err := database.Exec(new_piece_insert, piece.Kind, name, 0, false)
		if err != nil {
			panicExit(err.Error())
		}
	}
}
