package rules

import "github.com/pciet/wichess/piece"

func (a *Board) appendExtricateMoves(moves []MoveSet, active Orientation) []MoveSet {
	king := a.kingLocation(active)
	for i, s := range a {
		if (s.Kind == piece.NoKind) || (s.Orientation != active) || (s.flags.extricates == false) {
			continue
		}
		moves = moveSetSliceAdd(moves, king, AddressIndex(i).Address())
	}
	return moves
}
