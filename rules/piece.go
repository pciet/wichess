package rules

import "github.com/pciet/wichess/piece"

// Piece represents a chess piece in a game, with fields that include the necessary information
// to use it in package rules calculations.
//
// The JSON encoding of Piece is used to save into files by package memory. The web interface also
// uses a JSON piece encoding, but for that the game.Piece type is used on the host since the
// Kind and Orientation are the only needed fields to transmit.
type Piece struct {
	piece.Kind  `json:"k"`
	Orientation `json:"o"`
	Moved       bool    `json:"m"`
	Start       Address `json:"s"`

	// These flags aren't always an exact match with a package piece Characteristic.
	flags characteristics

	// These flags depend on other pieces on the board and are only used during move
	// and moves calculations.
	is conveyedCharacteristics
}

// NoPiece is the value of a Piece when it doesn't represent a piece.
var NoPiece = Piece{}

// NewPiece is used whenever initializing a Board with pieces. Unexported fields of the type are
// set in this function.
func NewPiece(k piece.Kind, o Orientation, moved bool, start Address) Piece {
	p := Piece{
		Kind:        k,
		Orientation: o,
		Moved:       moved,
		Start:       start,
	}
	applyCharacteristicFlags(&p)
	return p
}

func (a Piece) String() string {
	if a.Kind == piece.NoKind {
		return "none"
	}
	str := a.Orientation.Letter() + " " + a.Kind.String() + " (" + a.Start.String() + ")"
	if a.Moved {
		str += " moved"
	}
	return str
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

type conveyedCharacteristics struct {
	normalized, ordered, stopped, protected, enabled, immaterialized bool
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

	applyChars := func(c piece.Characteristic) {
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

func (a *Piece) immaterialAgainst(t *Piece) bool {
	return ((a.flags.immaterial && (a.is.normalized == false)) || a.is.immaterialized) &&
		(t.Kind.Basic() == piece.Pawn)
}
