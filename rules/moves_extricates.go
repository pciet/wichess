package rules

import "github.com/pciet/wichess/piece"

func (a Board) AppendExtricateMoves(moves []MoveSet, active Orientation) []MoveSet {
	king := a.KingLocation(active)
	for i, s := range a {
		if (s.Kind == piece.NoKind) || (s.Orientation != active) || (s.Extricates == false) {
			continue
		}
		moves = MoveSetSliceAdd(moves, king, AddressIndex(i).Address())
	}
	return moves
}
