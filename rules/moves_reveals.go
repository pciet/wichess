package rules

import "github.com/pciet/wichess/piece"

func (a *Board) appendRevealMoves(moves []Address, at Address) []Address {
	for _, s := range a.surroundingSquares(at) {
		if (s.Kind == piece.NoKind) || (s.flags.reveals == false) ||
			(s.Orientation != a[at.Index()].Orientation) {
			continue
		}
		target := Address{(2 * s.File) - at.File, (2 * s.Rank) - at.Rank}
		if (target.File > 7) || (target.Rank > 7) || (a[target.Index()].Kind != piece.NoKind) {
			continue
		}
		moves = append(moves, target)
	}
	return moves
}
