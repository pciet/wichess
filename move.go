package wichess

import (
	"database/sql"

	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// NoMove is the initial value of the From and To address indices for a new game.
// Board addresses are 0-63, so 64 is not a normal address index.
const NoMove = 64

// Move returns the squares that changed, taken pieces, and whether a following promotion is needed.
func Move(tx *sql.Tx, id GameIdentifier, player string,
	m rules.Move, promotion piece.Kind) ([]rules.AddressedSquare, []CapturedPiece, bool) {
	g := LoadGame(tx, id, true)
	if g.Header.ID == 0 {
		Panic("game", id, "not found")
	}
	if g.Header.Active != OrientationOf(player, g.Header.White.Name, g.Header.Black.Name) {
		DebugPrintln(player, "not active", g.Header.Active)
		return nil, nil, false
	}

	if promotion != piece.NoKind {
		by, needed := g.Board.PromotionNeeded()
		if (needed == false) ||
			(by != g.Header.Active) {
			DebugPrintln("invalid promotion request by", player)
			return nil, nil, false
		}
	} else {
		if g.MoveLegal(m) == false {
			DebugPrintln("illegal move", m)
			return nil, nil, false
		}
	}

	return g.DoMove(tx, m, promotion)
}

func (g Game) MoveLegal(m rules.Move) bool {
	p := g.Board.Board[m.From.Index()]
	if p.Kind == piece.NoKind {
		DebugPrintln("no piece at move from", m, "for player", g.Header.Active)
		return false
	}

	if p.Orientation != g.Header.Active {
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
// The board updates, taken pieces, and if a promotion is needed are returned.
func (g Game) DoMove(tx *sql.Tx, m rules.Move,
	promotion piece.Kind) ([]rules.AddressedSquare, []CapturedPiece, bool) {

	var changes, takes []rules.AddressedSquare

	if promotion != piece.NoKind {
		changes = make([]rules.AddressedSquare, 0, 1)
		changes = append(changes, g.Board.DoPromotion(promotion))
		m = rules.NoMove
	} else {
		changes, takes = g.Board.DoMove(m)
	}

	uc := make([]AddressedPiece, len(changes))
	for i, s := range changes {
		uc[i] = AddressedPiece{
			Address: s.Address,
			Piece: Piece{
				Slot:  NotInCollection,
				Piece: rules.Piece(s.Square),
			},
		}
	}

	wFirst := g.Header.White.Captures.FirstAvailable()
	bFirst := g.Header.Black.Captures.FirstAvailable()
	t := make([]CapturedPiece, len(takes))
	for i, s := range takes {
		if ((s.Orientation == rules.Black) && (wFirst == -1)) ||
			((s.Orientation == rules.White) && (bFirst == -1)) {
			Panic("captured piece", s, "when capture list already full")
		}
		var slot int
		if s.Orientation == rules.Black {
			wFirst++
			slot = wFirst
		} else if s.Orientation == rules.White {
			bFirst++
			slot = bFirst
		} else {
			Panic("orientation", s.Orientation, "not white or black")
		}
		t[i] = CapturedPiece{s.Orientation, s.Kind, slot}
	}

	// normal case is the next player is the opponent, but promotion can change it
	active := Opponent(g.Header.Active, g.Header.White.Name, g.Header.Black.Name)

	(&g.Board).ApplyChanges(changes)
	promoter, promotionNeeded := g.Board.PromotionNeeded()
	promoterName := PlayerWithOrientation(promoter, g.Header.White.Name, g.Header.Black.Name)

	if promotionNeeded {
		// move into promotion makes the to-be promoting player active
		active = promoterName
	} else if promotion != piece.NoKind {
		// if the promoter was not previous active then this is a reverse promotion
		if OrientationOf(promoterName,
			g.Header.White.Name, g.Header.Black.Name) != g.Header.PreviousActive {
			active = promoterName
		}
		// otherwise a promotion does the regular active player swap
	}

	// TODO: determine draw turn count (the 0 in UpdateGame)

	UpdateGame(tx, g.Header.ID,
		OrientationOf(active, g.Header.White.Name, g.Header.Black.Name),
		g.Header.Active, 0, g.Header.Turn, m, uc, t)

	// TODO: remove this copy when ID determination is done
	// main.Piece isn't needed past UpdateGame
	squares := make([]rules.AddressedSquare, len(uc))
	for i, s := range uc {
		squares[i] = rules.AddressedSquare{
			Address: s.Address,
			Square:  rules.Square(s.Piece.Piece),
		}
	}

	return squares, t, promotionNeeded
}
