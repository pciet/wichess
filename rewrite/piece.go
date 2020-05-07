package main

import "github.com/pciet/wichess/rules"

type (
	Piece struct {
		Slot        CollectionSlot `json:"s"`
		rules.Piece `json:"p"`
	}

	AddressedPiece struct {
		rules.Address `json:"a"`
		Piece         `json:"p"`
	}
)

// ConfigurePiece verifies initial square placement of requested pieces for a new game,
// and configures a returned Piece var.
// If the collection Piece has a kind of NoKind then basic piece is returned.
// A returned Piece of NoKind indicates an invalid request.
func ConfigurePiece(collection Piece, basic rules.PieceKind, o rules.Orientation) Piece {
	if collection.Kind == rules.NoKind {
		return Piece{
			Piece: rules.Piece{
				Kind:        basic,
				Orientation: o,
			},
		}
	}

	if basic != rules.BasicKind(collection.Kind) {
		return Piece{
			Piece: rules.Piece{
				Kind: rules.NoKind,
			},
		}
	}

	return Piece{
		Piece: rules.Piece{
			Kind:        collection.Kind,
			Orientation: o,
		},
	}
}
