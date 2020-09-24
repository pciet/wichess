package rules

import (
	"log"

	"github.com/pciet/wichess/piece"
)

type (
	// Move represents the addressing of a piece move from a square to another.
	Move struct {
		From Address `json:"f"`
		To   Address `json:"t"`
	}

	// MoveSet represents the moves of multiple pieces that may each have more than one To.
	MoveSet struct {
		From  Address   `json:"f"`
		Moves []Address `json:"m"`
	}
)

// NoMove is the value of a Move when it doesn't indicate a move.
var NoMove = Move{NoAddress, NoAddress}

// At least these bad moves can be made with DoMove:
//   putting the king in check
//   skipping a promotion
//   moving a locked piece
//   moves that aren't in the piece's move set
//   pawn takes a fortified piece
//   en passant turns later
//   castling through threatened squares, during check, or without a rook
//   swapping with a friendly piece without having the swap ability
//   extricate a piece other than the king

// DoMove does the requested move without affecting the Board. No move legality is determined and
// many illegal moves are possible with this method. The changed squares and original values of
// squares that had captures happen to them are returned.
func (a *Board) DoMove(m Move) ([]Square, []Square) {
	// copy the Board as a workspace to do temporary changes caused by characteristics
	bcopy := a.Copy()

	// apply characteristic changes caused by other pieces

	// remove characteristics due to normalizes
	for i, s := range bcopy {
		// TODO: this loop is duplicated in Game.Moves
		if (s.Kind == piece.NoKind) || (s.flags.normalizes == false) {
			continue
		}
		for _, ss := range bcopy.surroundingSquares(AddressIndex(i).Address()) {
			if ss.Kind == piece.NoKind {
				continue
			}
			normalize(&(bcopy[ss.Address.Index()].flags))
		}
	}

	// apply orders
	for i, s := range bcopy {
		if (s.Kind == piece.NoKind) || (s.flags.orders == false) {
			continue
		}
		for _, ss := range bcopy.surroundingSquares(AddressIndex(i).Address()) {
			if ss.Kind == piece.NoKind {
				continue
			}
			bcopy[ss.Address.Index()].flags.neutralizes = true
		}
	}

	changes := make([]Square, 0, 4)
	captures := make([]Square, 0, 1)

	from := bcopy[m.From.Index()]
	if from.Empty() {
		log.Panicln("no piece for move", m, a)
	}
	to := bcopy[m.To.Index()]

	if to.NotEmpty() {
		if to.Orientation == from.Orientation {
			if to.flags.extricates {
				// captures your own piece for the opponent to get king out of check
				changes, captures = bcopy.captureMove(changes, captures, m)
			}
		} else {
			if to.flags.neutralizes {
				return bcopy.neutralizesMove(changes, captures, m)
			}
			changes, captures = bcopy.captureMove(changes, captures, m)
		}
	} else {
		if bcopy.isCastleMove(m) {
			return bcopy.castleMove(changes, m), nil
		} else if bcopy.isEnPassantMove(m) {
			changes, captures = bcopy.enPassantMove(changes, captures, m)
		} else {
			changes = bcopy.noCaptureMove(changes, m)
		}
	}

	for _, s := range bcopy.surroundingSquares(m.To) {
		if bcopy.assertsWillCapture(from, s) == false {
			continue
		}
		if from.flags.neutralizes {
			return bcopy.assertsCapturesNeutralizes(changes, captures, m, s.Address)
		}
		return bcopy.assertsChain(changes, captures, m, s.Address)
	}

	return changes, captures
}

// Forward determines if the move is toward the opponent.
func (a Move) Forward(by Orientation) bool {
	if by == White {
		if a.From.Rank < a.To.Rank {
			return true
		}
	} else if by == Black {
		if a.From.Rank > a.To.Rank {
			return true
		}
	} else {
		log.Panicln("orientation", by, "not white or black")
	}
	return false
}

func (a Move) String() string {
	return "from " + a.From.String() + " to " + a.To.String()
}

func (a *Board) noCaptureMove(changes []Square, m Move) []Square {
	s := a[m.From.Index()]
	s.Moved = true
	changes = append(changes, Square{m.From, Piece{}})
	return append(changes, Square{m.To, s})
}

func (a *Board) captureMove(changes, captures []Square, m Move) ([]Square, []Square) {
	s := a[m.From.Index()]
	s.Moved = true
	changes = append(changes, Square{m.From, Piece{}})

	t := a[m.To.Index()]
	if t.flags.fantasy && (a[t.Start.Index()].Kind == piece.NoKind) {
		changes = append(changes, Square{t.Start, Piece{
			Kind:        t.Kind,
			Orientation: t.Orientation,
			Start:       t.Start,
			Moved:       true,
		}})
	} else {
		captures = append(captures, Square{m.To, t})
	}

	return append(changes, Square{m.To, s}), captures
}
