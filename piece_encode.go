package main

import (
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// An EncodedPiece is piece information packed into 64 bits for storage in the database.
type EncodedPiece uint64

func (an EncodedPiece) Uint64() uint64 { return uint64(an) }

const (
	EncodedPieceCollectionSlotBit  = 0
	EncodedPieceCollectionSlotMask = 0xFF

	EncodedPieceOrientationBit  = 8
	EncodedPieceOrientationMask = 0x1

	EncodedPieceMovedBit  = 16
	EncodedPieceMovedMask = 0x1

	EncodedPieceStartBit  = 17
	EncodedPieceStartMask = 0x3F

	EncodedPieceKindBit  = 32
	EncodedPieceKindMask = 0xFFFF
)

func (p Piece) Encode() EncodedPiece {
	var enc uint64
	enc |= ShiftedUint64(uint64(p.Slot),
		EncodedPieceCollectionSlotMask, EncodedPieceCollectionSlotBit)

	enc |= ShiftedUint64(uint64(p.Orientation),
		EncodedPieceOrientationMask, EncodedPieceOrientationBit)

	enc |= ShiftedUint64(uint64(Btoi(p.Moved)), EncodedPieceMovedMask, EncodedPieceMovedBit)
	enc |= ShiftedUint64(uint64(p.Start.Index()), EncodedPieceStartMask, EncodedPieceStartBit)

	enc |= ShiftedUint64(uint64(p.Kind), EncodedPieceKindMask, EncodedPieceKindBit)

	return EncodedPiece(enc)
}

func (e EncodedPiece) Decode() Piece {
	return Piece{
		Slot: CollectionSlot(UnshiftedUint64(e.Uint64(),
			EncodedPieceCollectionSlotMask, EncodedPieceCollectionSlotBit)),

		Piece: rules.Piece{
			Orientation: rules.Orientation(UnshiftedUint64(e.Uint64(),
				EncodedPieceOrientationMask, EncodedPieceOrientationBit)),

			Start: rules.AddressIndex(UnshiftedUint64(e.Uint64(),
				EncodedPieceStartMask, EncodedPieceStartBit)).Address(),

			Moved: Itob(int(UnshiftedUint64(e.Uint64(),
				EncodedPieceMovedMask, EncodedPieceMovedBit))),

			Kind: piece.Kind(UnshiftedUint64(e.Uint64(),
				EncodedPieceKindMask, EncodedPieceKindBit)),
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
