package rules

import "github.com/pciet/wichess/piece"

// Capturing a neutralizing piece will cause adjacent neutralizers to also neutralize. This
// function changes the Board.
func (a *Board) neutralizesMove(changes, captures []Square, m Move) ([]Square, []Square) {
	captures = append(captures, Square{m.From, a[m.From.Index()]})
	captures = append(captures, Square{m.To, a[m.To.Index()]})
	changes = append(changes, Square{m.From, Piece{}})
	changes = append(changes, Square{m.To, Piece{}})
	a.applyChanges(changes)

	var recursiveNeutralize func(Address)

	recursiveNeutralize = func(neutralizes Address) {
		for _, s := range a.surroundingSquares(neutralizes) {
			if s.Kind == piece.NoKind {
				continue
			}
			captures = append(captures, s)
			changes = append(changes, Square{s.Address, Piece{}})
			a[s.Address.Index()].Kind = piece.NoKind
			if s.flags.neutralizes {
				recursiveNeutralize(s.Address)
			}
		}
	}

	recursiveNeutralize(m.To)

	return changes, captures
}
