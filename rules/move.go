package rules

import "github.com/pciet/wichess/piece"

type (
	Move struct {
		From Address `json:"f"`
		To   Address `json:"t"`
	}

	MoveSet struct {
		From  Address   `json:"f"`
		Moves []Address `json:"m"`
	}
)

var NoMove = Move{NoAddress, NoAddress}

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
		Panic("orientation", by, "not white or black")
	}
	return false
}

func (a Game) AfterMove(m Move) Game {
	d, _ := a.DoMove(m)
	for _, s := range d {
		a.Board[s.Address.Index()] = s.Square
	}
	a.Previous = m
	return a
}

// At least these bad moves can be made with DoMove:
//   putting the king in check
//   skipping a promotion
//   moving a locked piece
//   moves that aren't in the piece's move set
//   pawn takes a fortified piece
//   en passant turns later
//   castling through threatened squares, during check, or without a rook
//   swapping with a friendly piece without having the swap ability

// Returns the squares that changed and the squares with each piece that was taken.
// No move legality is determined, bad moves either cause a panic or happen.
func (a Board) DoMove(m Move) ([]AddressedSquare, []AddressedSquare) {
	from := a[m.From.Index()]
	if from.Empty() {
		Panic("no piece for move", m, a)
	}

	changes := make([]AddressedSquare, 0, 3)
	takes := make([]AddressedSquare, 0, 1)

	to := a[m.To.Index()]
	if to.NotEmpty() {
		if to.Orientation == from.Orientation {
			changes = a.SwapMove(changes, m)
		} else {
			if to.Detonates {
				return a.DetonateMove(changes, takes, m)
			}
			changes, takes = a.TakeMove(changes, takes, m)
		}
	} else {
		if a.IsCastleMove(m) {
			return a.CastleMove(changes, m), nil
		} else if a.IsEnPassantMove(m) {
			changes, takes = a.EnPassantMove(changes, takes, m)
		} else {
			changes = a.NoTakeMove(changes, m)
		}
	}

	for _, s := range a.SurroundingSquares(m.To) {
		if a.GuardWillTake(from, s) == false {
			continue
		}
		if from.Detonates {
			return a.GuardTakesDetonate(changes, takes, m, s.Address)
		}
		return a.GuardChain(changes, takes, m, s.Address)
	}

	return changes, takes
}

func (a Board) NoTakeMove(changes []AddressedSquare, m Move) []AddressedSquare {
	s := a[m.From.Index()]
	s.Moved = true
	changes = append(changes, AddressedSquare{m.From, Square{}})
	return append(changes, AddressedSquare{m.To, s})
}

func (a Board) TakeMove(changes, takes []AddressedSquare,
	m Move) ([]AddressedSquare, []AddressedSquare) {
	s := a[m.From.Index()]
	s.Moved = true
	changes = append(changes, AddressedSquare{m.From, Square{}})

	t := a[m.To.Index()]
	if t.Fantasy && (a[t.Start.Index()].Kind == piece.NoKind) {
		changes = append(changes, AddressedSquare{t.Start, Square{
			Kind:        t.Kind,
			Orientation: t.Orientation,
			Start:       t.Start,
			Moved:       true,
		}})
	} else {
		takes = append(takes, AddressedSquare{m.To, t})
	}

	return append(changes, AddressedSquare{m.To, s}), takes
}

func (a Move) String() string {
	return "from " + a.From.String() + " to " + a.To.String()
}
