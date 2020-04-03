package main

import (
	"database/sql"

	"github.com/pciet/wichess/rules"
)

// NoMove is the initial value of the From and To address indices for a new game.
// Board addresses are 0-63, so 64 is not a normal address index.
const NoMove = 64

// Move returns the squares that changed and whether a following promotion is needed.
func Move(tx *sql.Tx, id GameIdentifier, player string,
	m rules.Move, promotion rules.PieceKind) ([]rules.AddressedSquare, bool) {
	g := LoadGame(tx, id)
	if g.Header.ID == 0 {
		Panic("game", id, "not found")
	}
	if g.Header.Active != player {
		DebugPrintln(player, "not active", g.Header.Active)
		return nil, false
	}

	if promotion != rules.NoKind {
		by, needed := g.Board.PromotionNeeded()
		if (needed == false) ||
			(by != ActiveOrientation(g.Header.Active,
				g.Header.White.Name, g.Header.Black.Name)) {
			DebugPrintln("invalid promotion request by", player)
			return nil, false
		}
	} else {
		if g.MoveLegal(m) == false {
			DebugPrintln("illegal move", m)
			return nil, false
		}
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

// TODO: cleaner func signatures

// DoMove does the database interactions necessary to do a move. Illegal moves can be done.
// The board updates and if a promotion is needed is returned.
func (g Game) DoMove(tx *sql.Tx, m rules.Move,
	promotion rules.PieceKind) ([]rules.AddressedSquare, bool) {
	// TODO: changes, taken
	changes := make([]rules.AddressedSquare, 0, 1)
	if promotion != rules.NoKind {
		changes = append(changes, g.Board.DoPromotion(promotion))
		m = rules.NoMove
	} else {
		changes, _ = g.Board.DoMove(m)
	}

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

	// normal case is the next player is the opponent, but promotion can change it
	active := Opponent(g.Header.Active, g.Header.White.Name, g.Header.Black.Name)

	(&g.Board).ApplyChanges(changes)
	promoter, promotionNeeded := g.Board.PromotionNeeded()
	promoterName := PlayerWithOrientation(promoter, g.Header.White.Name, g.Header.Black.Name)

	if promotionNeeded {
		// move into promotion makes the to-be promoting player active
		active = promoterName
	} else if promotion != rules.NoKind {
		// if the promoter was not previous active then this is a reverse promotion
		if promoterName != g.Header.PreviousActive {
			active = promoterName
		}
		// otherwise a promotion does the regular active player swap
	}

	// TODO: determine draw turn count (the 0 in UpdateGame)

	UpdateGame(tx, g.Header.ID, active, g.Header.Active, 0, g.Header.Turn, m, uc)

	// TODO: update piece database for timed games
	//PieceTakesUpdate(tx, id, taken)

	// TODO: remove this copy when ID determination is done
	// main.Piece isn't needed past UpdateGame
	squares := make([]rules.AddressedSquare, len(uc))
	for i, s := range uc {
		squares[i] = rules.AddressedSquare{
			Address: s.Address,
			Square:  rules.Square(s.Piece.Piece),
		}
	}

	return squares, promotionNeeded
}
