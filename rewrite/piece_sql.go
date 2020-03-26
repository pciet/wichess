package main

const (
	PieceTable = "pieces"

	PieceID       = "piece_id"
	PieceKind     = "kind"
	PieceOwner    = "owner"
	PieceReserved = "reserved"
	PieceTaken    = "taken"
)

var (
	PieceQuery = SQLQuery([]string{
		PieceOwner,
		PieceKind,
		PieceTaken,
	}, PieceTable, PieceID)

	PieceInsert = SQLInsert(PieceTable, []string{
		PieceKind,
		PieceOwner,
		PieceReserved,
		PieceTaken,
	})
)
