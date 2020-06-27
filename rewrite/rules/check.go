package rules

import "github.com/pciet/wichess/piece"

func (a Board) InCheck(active Orientation, takes []Address) bool {
	king := a.KingLocation(active)
	for _, take := range takes {
		if take == king {
			return true
		}
	}
	return false
}

func (a Game) RemoveMovesIntoCheck(moves []MoveSet, active Orientation) []MoveSet {
	out := make([]MoveSet, 0, len(moves))

	for _, moveset := range moves {
		outset := MoveSet{moveset.From, make([]Address, 0, len(moveset.Moves))}
		for _, move := range moveset.Moves {
			ga := a.AfterMove(Move{moveset.From, move})
			threats := MovesAddressSlice(ga.NaiveTakeMoves(active.Opponent()))

			if ga.Board.NoKing(active) || ga.Board.InCheck(active, threats) ||
				ga.Board.ThreatenedDetonatorAdjacent(threats, ga.Board.KingLocation(active)) {
				continue
			}

			outset.Moves = append(outset.Moves, move)
		}
		if len(outset.Moves) == 0 {
			continue
		}
		out = append(out, outset)
	}
	return out
}

// TODO: test cases for an opponent detonator next to king, can the king be improperly removed?

func (a Board) ThreatenedDetonatorAdjacent(threats []Address, at Address) bool {
	p := a[at.Index()]
	for _, as := range a.SurroundingSquares(at) {
		s := a[as.Address.Index()]
		if (s.Kind == piece.NoKind) || (p.Orientation != s.Orientation) ||
			(s.Detonates == false) {
			continue
		}
		for _, addr := range threats {
			if addr != as.Address {
				continue
			}
			return true
		}
	}
	return false
}
