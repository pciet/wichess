package main

import (
	"database/sql"

	"github.com/pciet/wichess"
)

// NoMove is the initial value of the From and To address indices for a new game.
// Board addresses are 0-63, so 64 is not a normal address index.
const NoMove = 64

func Move(tx *sql.Tx, id GameIdentifier, m rules.Move, promotion rules.Kind) []AddressedSquare {
	g := LoadGame(tx, id)
	if g.ID == 0 {
		Panic("game", id, "not found")
	}

	if g.MoveLegal(m) == false {
		return nil
	}

	return g.DoMove(tx, m, promotion)
}

func (g Game) MoveLegal(m rules.Move) bool {
	if g.Board.MovingPlayer(m.From) != g.Header.Active {
		DebugPrintln("active player", g.Header.Active, "not moving player")
		return false
	}

	// TODO: cache move calculation if the database read/write is cheaper than recalculation

	// moves are recalculated to confirm legality of the requested move
	moves, state := MovesForLoadedGame(rules.Game{
		Board: g.Board,
		Previous: rules.Move{
			From: rules.AddressIndex(g.Header.From).Address(),
			To:   rules.AddressIndex(g.Header.To).Address(),
		},
	})

	if (state != rules.Normal) && (state != rules.Check) {
		DebugPrintln(g.Header.Active, "requested move", m, "in", state, "state")
		return false
	}

	if MoveSetSliceHas(moves, m) == false {
		DebugPrintln(g.Header.Active, "requested illegal move", m)
		return false
	}

	return true
}

// DoMove does the database interactions necessary to do a move. Illegal moves can be done.
func (g Game) DoMove(tx *sql.Tx, m rules.Move, promotion rules.Kind) []AddressedSquare {
	// TODO: changes must also include where each piece moved from so the ID can be matched here
	// changes, taken
	changes, _ := g.Board.DoMove(m)

	// TODO: determine draw turn count (the 0 in UpdateGame)

	UpdateGame(tx, id, g.Header.White.Name, g.Header.Black.Name, g.Header.Active, 0, g.Header.Turn, m, changes)

	// TODO: update piece database for timed games
	//PieceTakesUpdate(tx, id, taken)

	out := make([]AddressedSquare, 0, len(changes))
	for _, change := range changes {
		out = append(out, AddressedSquare{
			Address: change.Address,
			Square: Square{
				ID:     g.Board.IdentifierAt(change.From),
				Square: change.Square,
			},
		})
	}

	return out
}
