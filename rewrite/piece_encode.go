package main

// An EncodedPiece is piece information packed into 64 bits for storage in the database.
type EncodedPiece uint64

const (
	EncodedPieceIdentifierBit  = 0
	EncodedPieceIdentifierMask = 0xFFFFFFFF

	EncodedPieceOrientationBit  = 32
	EncodedPieceOrientationMask = 0x1

	EncodedPieceMovedBit  = 33
	EncodedPieceMovedMask = 0x1

	EncodedPieceKindBit  = 47
	EncodedPieceKindMask = 0xFFFF
)

func (p Piece) Encode() EncodedPiece {
	var enc uint64
	enc |= ShiftedUint64(uint64(p.ID), EncodedPieceIdentifierMask, EncodedPieceIdentifierBit)
	enc |= ShiftedUint64(uint64(p.Orientation), EncodedPieceOrientationMask, EncodedPieceOrientationBit)
	enc |= ShiftedUint64(uint64(Btoi(p.Moved)), EncodedPieceMovedMask, EncodedPieceMovedBit)
	enc |= ShiftedUint64(uint64(p.Kind), EncodedPieceKindMask, EncodedPieceKindBit)
	return EncodedPiece(enc)
}

func (e EncodedPiece) Decode() Piece {
	return Piece{
		ID: PieceIdentifier(UnshiftedUint64(e, EncodedPieceIdentifierMask, EncodedPieceIdentifierBit)),
		Piece: rules.Piece{
			Orientation: rules.Orientation(UnshiftedUint64(e, EncodedPieceOrientationMask, EncodedPieceOrientationBit)),
			Moved:       Itob(int(UnshiftedUint64(e, EncodedPieceMovedMask, EncodedPieceMovedBit))),
			Kind:        rules.PieceKind(UnshiftedUint64(e, EncodedPieceKindMask, EncodedPieceKindBit)),
		},
	}
}

// ShiftedUint64 masks the value then shifts it.
// For example, if value is 0x89, mask is 0xF, and location is 4, then the result is 0x90.
func ShiftedUint64(value, mask, location uint64) uint64 { return (value & mask) << location }

// UnshiftedUint64 reverses the ShiftedUint64 function.
func UnshiftedUint64(value, mask, location uint64) uint64 { return (value >> location) & mask }

func Btoi(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func Itob(i int) bool {
	if i == 0 {
		return false
	} else {
		return true
	}
}
