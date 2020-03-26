package main

import "github.com/pciet/wichess/rules"

type (
	PieceIdentifier int

	Piece struct {
		ID          PieceIdentifier `json:"id"`
		rules.Piece `json:"p"`
	}

	AddressedPiece struct {
		rules.Address `json:"a"`
		Piece         `json:"p"`
	}

	AddressedPieceIdentifier struct {
		ID            PieceIdentifier `json:"id"`
		rules.Address `json:"a"`
	}
)
