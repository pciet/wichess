package main

import (
	"encoding/json"
	"io"

	"github.com/pciet/wichess/rules"
)

type (
	ArmyRequest [16]CollectionSlot
	EncodedArmy [16]EncodedPiece
)

var RegularArmyRequest = ArmyRequest{}

func DecodeArmyRequest(jsonBody io.Reader) (ArmyRequest, error) {
	var a ArmyRequest
	err := json.NewDecoder(jsonBody).Decode(&a)
	return a, err
}

// PickSlotsInArmyRequest returns if an ArmyRequest includes the left and/or right random picks.
func PickSlotsInArmyRequest(r ArmyRequest) (bool, bool) {
	left, right := false, false
	for _, s := range r {
		if s == LeftPick {
			left = true
		} else if s == RightPick {
			right = true
		}
	}
	return left, right
}

// The BasicArmy is the initial position of one side for a regular chess game, addressed by
// Wisconsin Chess index.
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
