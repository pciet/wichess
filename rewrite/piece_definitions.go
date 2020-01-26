package main

import ()

const (
	piece_table = "pieces"

	piece_id       = "piece_id"
	piece_kind     = "kind"
	piece_owner    = "owner"
	piece_reserved = "reserved"
	piece_taken    = "taken"
)

type PieceIdentifier int

type Piece struct {
	ID PieceIdentifier
	rules.Piece
}

const piece_query_select = []string{
	piece_owner,
	piece_kind,
	piece_taken,
}

var piece_query = BuildSQLQuery(piece_query_select, piece_table, piece_id_key)
