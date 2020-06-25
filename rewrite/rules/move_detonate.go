package rules

import "github.com/pciet/wichess/piece"

// Taking a detonating piece will cause adjacent detonators to also detonate.
func (a Board) DetonateMove(changes, takes []AddressedSquare,
	m Move) ([]AddressedSquare, []AddressedSquare) {
	takes = append(takes, AddressedSquare{m.From, a[m.From.Index()]})
	takes = append(takes, AddressedSquare{m.To, a[m.To.Index()]})
	changes = append(changes, AddressedSquare{m.From, Square{}})
	changes = append(changes, AddressedSquare{m.To, Square{}})
	(&a).ApplyChanges(changes)

	var recursiveDetonate func(Address)

	recursiveDetonate = func(detonator Address) {
		for _, s := range a.SurroundingSquares(detonator) {
			if s.Kind == piece.NoKind {
				continue
			}
			takes = append(takes, s)
			changes = append(changes, AddressedSquare{s.Address, Square{}})
			a[s.Address.Index()].Kind = piece.NoKind
			if s.Detonates {
				recursiveDetonate(s.Address)
			}
		}
	}

	recursiveDetonate(m.To)

	return changes, takes
}
