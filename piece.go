// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"fmt"

	"github.com/pciet/wichess/wichessing"
)

const (
	database_piece_table = "pieces"

	database_piece_table_kind_key  = "kind"
	database_piece_table_owner_key = "owner"
	database_piece_table_takes_key = "takes"
)

func bestPieceForPlayerFromDatabase(name string) wichessing.Piece {
	rows, err := database.Query(fmt.Sprintf(" SELECT %v, %v FROM %v WHERE %v=$1", database_piece_table_kind_key, database_piece_table_takes_key, database_piece_table, database_piece_table_owner_key), name)
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
