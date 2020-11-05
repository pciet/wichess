package rules

import "github.com/pciet/wichess/piece"

func (a *Board) fantasyReturns(p Piece) bool {
	if (p.flags.fantasy && (p.is.normalized == false)) &&
		(a[p.Start.Index()].Kind == piece.NoKind) {

		return true
	}
	return false
}

func fantasyReturnChange(changes []Square, f Piece) []Square {
	return append(changes, Square{f.Start, NewPiece(f.Kind, f.Orientation, true, f.Start)})
}
