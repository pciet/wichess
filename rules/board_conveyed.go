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

	// apply normalizes first so that pieces don't convey if normalized
	for i, s := range a {
		if (s.Kind == piece.NoKind) || (s.flags.normalizes == false) || s.is.normalized {
			continue
		}
		addr := AddressIndex(i).Address()

		normd := false
		// adjacent normalizers affect each other but negate their effects on other pieces
		for _, ss := range a.surroundingSquares(addr) {
			if (ss.Kind == piece.NoKind) || (ss.flags.normalizes == false) {
				continue
			}
			normd = true
			a[ss.Address.Index()].is.normalized = true
		}

		if normd {
			a[AddressIndex(i)].is.normalized = true
			continue
		}

		for _, ss := range a.surroundingSquares(addr) {
			if ss.Kind == piece.NoKind {
				continue
			}
			a[ss.Address.Index()].is.normalized = true
		}
	}

	for i, s := range a {
		if (s.Kind == piece.NoKind) || s.is.normalized {
			continue
		}
		addr := AddressIndex(i).Address()

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
				if (ss.Kind == piece.NoKind) || (ss.Orientation == s.Orientation) ||
					(ss.Kind == piece.King) || (ss.Kind == piece.Queen) {

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
