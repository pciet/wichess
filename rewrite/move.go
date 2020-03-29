package main

import (
	"database/sql"

	"github.com/pciet/wichess/rules"
)

// NoMove is the initial value of the From and To address indices for a new game.
// Board addresses are 0-63, so 64 is not a normal address index.
const NoMove = 64

func Move(tx *sql.Tx, id GameIdentifier, m rules.Move, promotion rules.PieceKind) []AddressedPiece {
	g := LoadGame(tx, id)
	if g.Header.ID == 0 {
		Panic("game", id, "not found")
	}

	if g.MoveLegal(m) == false {
		DebugPrintln("g.MoveLegal(", m, ") returned false")
		return nil
	}

	return g.DoMove(tx, m, promotion)
}

func (g Game) MoveLegal(m rules.Move) bool {
	p := g.Board.Board[m.From.Index()]
	if p.Kind == rules.NoKind {
		DebugPrintln("no piece at move from", m, "for player", g.Header.Active)
		return false
	}

	if p.Orientation != ActiveOrientation(g.Header.Active, g.Header.White.Name, g.Header.Black.Name) {
		DebugPrintln("active player", g.Header.Active, "not moving player")
		return false
	}

	// TODO: cache move calculation if the database read/write is cheaper than recalculation

	// moves are recalculated to confirm legality of the requested move
	moves, state := g.Moves()

	if (state != rules.Normal) && (state != rules.Check) {
		DebugPrintln(g.Header.Active, "requested move", m, "in", state, "state")
		return false
	}

	if rules.MoveSetSliceHasMove(moves, m) == false {
		DebugPrintln(g.Header.Active, "requested illegal move", m)
		return false
	}

	return true
}

// DoMove does the database interactions necessary to do a move. Illegal moves can be done.
func (g Game) DoMove(tx *sql.Tx, m rules.Move, promotion rules.PieceKind) []AddressedPiece {
	// TODO: changes, taken
	changes, _ := g.Board.DoMove(m)

	// TODO: the piece ID must be determined here somehow so it's correctly put into the database

	uc := make([]AddressedPiece, len(changes))
	for i, s := range changes {
		uc[i] = AddressedPiece{
			Address: s.Address,
			Piece: Piece{
				ID:    0,
				Piece: rules.Piece(s.Square),
			},
		}
	}

	// TODO: determine draw turn count (the 0 in UpdateGame)

	UpdateGame(tx, g.Header.ID, g.Header.White.Name, g.Header.Black.Name, g.Header.Active,
		0, g.Header.Turn, m, uc)

	// TODO: update piece database for timed games
	//PieceTakesUpdate(tx, id, taken)

	return uc
}
