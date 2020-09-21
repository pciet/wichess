package rules

import "github.com/pciet/wichess/piece"

// Piece represents a chess piece in a game, with fields that include the necessary information
// to use it in package rules calculations.
//
// The JSON encoding of Piece is used to save into files by package memory. The web interface also
// uses a JSON piece encoding, but for that the wichess.Piece type is used on the host since the
// Kind and Orientation are the only needed fields to transmit.
type Piece struct {
	piece.Kind  `json:"k"`
	Orientation `json:"o"`
	Moved       bool    `json:"m"`
	Start       Address `json:"s"`

	// These flags aren't always an exact match with a package piece Characteristic.
	flags characteristics
}

// TODO: detonate=neutralizes, guards=asserts, fortified=immaterial, locks=stops, rallies=enables,
// ghost=quick

// TODO: would packing flags into an int be a significant performance improvement?

// When a new characteristics bool is added the applyCharacteristicFlags and normalize funcs must
// also be updated.
type characteristics struct {
	neutralizes, asserts, immaterial, stops, enables,
	mustEnd, // can only move to the last square on the path
	quick, // paths can continue over other pieces, like for the regular knight
	noOverCapture, // only applies with quick, can move but can't capture by moving over pieces
	reveals, tense, fantasy, keep, protective, extricates, normalizes, orders bool
}

func applyCharacteristicFlags(to *Piece) {
	if to.Kind.Basic() == piece.Knight {
		to.flags.mustEnd = true
		if (to.Kind != piece.Line) && (to.Kind != piece.Appropriate) {
			to.flags.quick = true
		}
	}
	if to.Kind == piece.Exit {
		to.flags.quick = true
		to.flags.noOverCapture = true
	}

	charA, charB := piece.Characteristics(to.Kind)
	if charA == piece.NoCharacteristic {
		return
	}

	applyChars := func(c piece.Characteristic) bool {
		switch c {
		case piece.Neutralizes:
			to.flags.neutralizes = true
		case piece.Asserts:
			to.flags.asserts = true
		case piece.Enables:
			to.flags.enables = true
		case piece.Reveals:
			to.flags.reveals = true
		case piece.Stops:
			to.flags.stops = true
		case piece.Immaterial:
			to.flags.immaterial = true
		case piece.Tense:
			to.flags.tense = true
		case piece.Fantasy:
			to.flags.fantasy = true
		case piece.Keep:
			to.flags.keep = true
		case piece.Protective:
			to.flags.protective = true
		case piece.Extricates:
			to.flags.extricates = true
		case piece.Normalizes:
			to.flags.normalizes = true
		case piece.Orders:
			to.flags.orders = true
		}
	}

	applyChars(charA)
	applyChars(charB)
}

func normalize(s *characteristics) {
	s.flags.neutralizes = false
	s.flags.asserts = false
	s.flags.immaterial = false
	s.flags.stops = false
	s.flags.enables = false
	s.flags.reveals = false
	s.flags.tense = false
	s.flags.fantasy = false
	s.flags.keep = false
	s.flags.protective = false
	s.flags.extricates = false
	s.flags.orders = false

	// the Normalize, MustEnd, and Quick bools are left true
}

func (a *Piece) fortifiedAgainst(t *Piece) bool {
	return a.flags.immaterial && (t.Kind.Basic() == piece.Pawn)
}
