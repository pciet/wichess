package rules

import "github.com/pciet/wichess/piece"

// applyConveyedCharacteristics writes the conveyed characteristic flags that indicate if a piece's
// characteristics have been modified by other pieces. These conveyed flags are separate from the
// piece's inherent characteristics that are described on the piece's details webpage for players
// because changing those during moves calculations is error prone.
func (a *Board) applyConveyedCharacteristics() {
	// reset any previous flagging
	for i := 0; i < 64; i++ {
		if a[i].Kind == piece.NoKind {
			continue
		}
		a[i].is = conveyedCharacteristics{}
	}

	for i, s := range a {
		if s.Kind == piece.NoKind {
			continue
		}
		addr := AddressIndex(i).Address()

		if s.flags.normalizes {
			for _, ss := range a.surroundingSquares(addr) {
				if ss.Kind == piece.NoKind {
					continue
				}
				a[ss.Address.Index()].is.normalized = true
			}
		}

		if s.flags.orders {
			for _, ss := range a.surroundingSquares(addr) {
				if ss.Kind == piece.NoKind {
					continue
				}
				a[ss.Address.Index()].is.ordered = true
			}
		}

		if s.flags.stops {
			for _, ss := range a.surroundingSquares(addr) {
				if (ss.Kind == piece.NoKind) || (ss.Orientation == s.Orientation) {
					continue
				}
				a[ss.Address.Index()].is.stopped = true
			}
		}

		if s.flags.protective && (s.is.protected == false) {
			for _, ss := range a.surroundingSquares(addr) {
				if (ss.Kind == piece.NoKind) || (ss.flags.protective == false) {
					continue
				}
				a[AddressIndex(i)].is.protected = true
				a[ss.Address.Index()].is.protected = true
			}
		}

		if s.flags.enables {
			for _, ss := range a.surroundingSquares(addr) {
				if (ss.Kind == piece.NoKind) || (ss.Orientation != s.Orientation) {
					continue
				}
				a[ss.Address.Index()].is.enabled = true
			}
		}

		if s.flags.keep {
			for _, ss := range a.surroundingSquares(addr) {
				if (ss.Kind == piece.NoKind) || (ss.Orientation != s.Orientation) {
					continue
				}
				a[ss.Address.Index()].is.immaterialized = true
			}
		}
	}
}
