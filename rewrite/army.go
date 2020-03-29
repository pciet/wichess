package main

import (
	"encoding/json"
	"io"

	"github.com/pciet/wichess/rules"
)

// The BasicArmy is the initial position of one side for a regular
// chess game, addressed by Wisconsin Chess index.
var BasicArmy = func() [16]rules.PieceKind {
	var b [16]rules.PieceKind
	for i := 0; i < 8; i++ {
		b[i] = rules.Pawn
	}

	b[8] = rules.Rook
	b[15] = rules.Rook

	b[9] = rules.Knight
	b[14] = rules.Knight

	b[10] = rules.Bishop
	b[13] = rules.Bishop

	b[11] = rules.Queen
	b[12] = rules.King

	return b
}()

// When a player requests a new game they specify which
// of their pieces to include in an ArmyRequest.
// The array index is the army square address index from
// the perspective of the white player (see the BasicArmy var).
// An ID of 0 means use the regular chess piece.
type ArmyRequest [16]PieceIdentifier

func DecodeArmyRequest(jsonBody io.Reader) (ArmyRequest, error) {
	var a ArmyRequest
	err := json.NewDecoder(jsonBody).Decode(&a)
	return a, err
}
