package wichess

import (
	"encoding/json"
	"io"

	"github.com/pciet/wichess/piece"
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
var BasicArmy = func() [16]piece.Kind {
	var b [16]piece.Kind
	for i := 0; i < 8; i++ {
		b[i] = piece.Pawn
	}

	b[8] = piece.Rook
	b[15] = piece.Rook

	b[9] = piece.Knight
	b[14] = piece.Knight

	b[10] = piece.Bishop
	b[13] = piece.Bishop

	b[11] = piece.Queen
	b[12] = piece.King

	return b
}()
