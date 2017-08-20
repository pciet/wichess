// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pciet/wichess/wichessing"
)

const (
	database_piece_table = "pieces"

	database_piece_table_id_key     = "piece_id"
	database_piece_table_kind_key   = "kind"
	database_piece_table_owner_key  = "owner"
	database_piece_table_takes_key  = "takes"
	database_piece_table_ingame_key = "ingame"

	free_piece_slice_hint = 8
)

func freePiecesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}
	key := validSession(r)
	if key == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	name := nameFromSessionKey(key)
	if name == "" {
		clearClientSession(w)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	free := freePiecesForPlayerFromDatabase(name)
	json, err := json.Marshal(free)
	if err != nil {
		panicExit(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func freePiecesForPlayerFromDatabase(name string) []wichessing.Piece {
	rows, err := database.Query(fmt.Sprintf("SELECT %v, %v, %v FROM %v WHERE %v=$1 AND %v=$2", database_piece_table_id_key, database_piece_table_kind_key, database_piece_table_takes_key, database_piece_table, database_piece_table_owner_key, database_piece_table_ingame_key), name, false)
	if err != nil {
		panicExit(err.Error())
		return nil
	}
	defer rows.Close()
	pieces := make([]wichessing.Piece, 0, free_piece_slice_hint)
	i := 0
	for rows.Next() {
		pieces = append(pieces, wichessing.Piece{})
		err = rows.Scan(&pieces[i].Identifier, &pieces[i].Kind, &pieces[i].Takes)
		if err != nil {
			panicExit(err.Error())
			return nil
		}
		i++
	}
	err = rows.Err()
	if err != nil {
		panicExit(err.Error())
		return nil
	}
	return pieces
}

func bestPieceForPlayerFromDatabase(name string) wichessing.Piece {
	rows, err := database.Query(fmt.Sprintf("SELECT %v, %v FROM %v WHERE %v=$1", database_piece_table_kind_key, database_piece_table_takes_key, database_piece_table, database_piece_table_owner_key), name)
	if err != nil {
		panicExit(err.Error())
		return wichessing.Piece{}
	}
	defer rows.Close()
	p := wichessing.Piece{}
	newp := wichessing.Piece{}
	for rows.Next() {
		err = rows.Scan(&newp.Kind, &newp.Takes)
		if err != nil {
			panicExit(err.Error())
			return wichessing.Piece{}
		}
		if newp.Takes > p.Takes {
			wichessing.CopyFromPiece(newp, &p)
		}
	}
	err = rows.Err()
	if err != nil {
		panicExit(err.Error())
		return wichessing.Piece{}
	}
	return p
}
