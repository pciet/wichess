package main

import (
	"github.com/pciet/wichess/rules"
)

type PieceIdentifier int

type Piece struct {
	ID PieceIdentifier
	rules.Piece
}

type AddressedPieceIdentifier struct {
	ID PieceIdentifier
	rules.BoardAddress
}

type EncodedPiece uint64

const (
	encoded_piece_identifier_bit  = 0
	encoded_piece_identifier_mask = 0xFFFFFFFF

	encoded_piece_orientation_bit  = 32
	encoded_piece_orientation_mask = 0x1

	encoded_piece_moved_bit  = 33
	encoded_piece_moved_mask = 0x1

	encoded_piece_kind_bit  = 47
	encoded_piece_kind_mask = 0xFFFF
)
