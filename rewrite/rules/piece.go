package rules

import "github.com/pciet/wichess/piece"

// TODO: update characteristic names, and don't repeat them here

type Piece struct {
	piece.Kind  `json:"k"`
	Orientation `json:"o"`
	Moved       bool `json:"-"`

	Swaps bool `json:"-"`

	// Neutralizes
	Detonates bool `json:"-"`

	// Asserts
	Guards bool `json:"-"`

	// Immaterial
	Fortified bool `json:"-"`

	// Stops
	Locks bool `json:"-"`

	// Enables
	Rallies bool `json:"-"`

	MustEnd bool `json:"-"`

	// Quick
	Ghost bool `json:"-"`

	Reveals bool `json:"-"`
}

var (
	WhiteKingStart      = Address{4, 0}
	BlackKingStart      = Address{4, 7}
	WhiteLeftRookStart  = Address{0, 0}
	WhiteRightRookStart = Address{7, 0}
	BlackLeftRookStart  = Address{7, 7}
	BlackRightRookStart = Address{0, 7}
)

func (a Piece) ApplyCharacteristics() Piece {
	if a.Kind.Basic() == piece.Knight {
		a.MustEnd = true
		a.Ghost = true
	}

	chars := piece.CharacteristicList[a.Kind]

	applyChars := func(c piece.Characteristic) bool {
		switch c {
		case piece.Neutralizes:
			a.Detonates = true
		case piece.Asserts:
			a.Guards = true
		case piece.Enables:
			a.Rallies = true
		case piece.Reveals:
			a.Reveals = true
		case piece.Stops:
			a.Locks = true
		case piece.Immaterial:
			a.Fortified = true
		default:
			return false
		}
		return true
	}

	if applyChars(chars.A) == false {
		return a
	}

	applyChars(chars.B)

	return a
}
