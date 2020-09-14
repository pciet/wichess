package memory

import (
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// An encodedPiece is piece information packed into 64 bits for file storage.
type encodedPiece uint64

const (
	encodedPieceOrientationBit  = 8
	encodedPieceOrientationMask = 0x1

	encodedPieceMovedBit  = 16
	encodedPieceMovedMask = 0x1

	encodedPieceStartBit  = 17
	encodedPieceStartMask = 0x3F

	encodedPieceKindBit  = 32
	encodedPieceKindMask = 0xFFFF
)

func encodePiece(p rules.Piece) encodedPiece {
	var enc uint64
	enc |= shiftedUint64(uint64(p.Orientation),
		encodedPieceOrientationMask, encodedPieceOrientationBit)

	enc |= shiftedUint64(uint64(btoi(p.Moved)), encodedPieceMovedMask, encodedPieceMovedBit)
	enc |= shiftedUint64(uint64(p.Start.Index()), encodedPieceStartMask, encodedPieceStartBit)

	enc |= shiftedUint64(uint64(p.Kind), encodedPieceKindMask, encodedPieceKindBit)

	return encodedPiece(enc)
}

func (e encodedPiece) decode() rules.Piece {
	p := uint64(e)
	return rules.Piece{
		Orientation: rules.Orientation(unshiftedUint64(p,
			encodedPieceOrientationMask, encodedPieceOrientationBit)),

		Start: rules.AddressIndex(unshiftedUint64(p,
			encodedPieceStartMask, encodedPieceStartBit)).Address(),

		Moved: itob(int(unshiftedUint64(p,
			encodedPieceMovedMask, encodedPieceMovedBit))),

		Kind: piece.Kind(unshiftedUint64(p,
			encodedPieceKindMask, encodedPieceKindBit)),
	}
}

// shiftedUint64 masks the value then shifts it.
// For example, if value is 0x89, mask is 0xF, and location is 4, then the result is 0x90.
func shiftedUint64(value, mask, location uint64) uint64 { return (value & mask) << location }

// unshiftedUint64 reverses the shiftedUint64 function.
func UnshiftedUint64(value, mask, location uint64) uint64 { return (value >> location) & mask }

func btoi(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func itob(i int) bool {
	if i == 0 {
		return false
	} else {
		return true
	}
}
