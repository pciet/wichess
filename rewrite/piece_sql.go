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

var (
	piece_query = BuildSQLQuery([]string{
		piece_owner,
		piece_kind,
		piece_taken,
	}, piece_table, piece_id)

	piece_insert = BuildSQLInsert(piece_table, []string{
		piece_kind,
		piece_owner,
		piece_reserved,
		piece_taken,
	})
)
