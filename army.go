package wichess

import (
	"encoding/json"
	"io"

	"github.com/pciet/wichess/piece"
)

type (
	ArmyRequest [16]CollectionSlot
	EncodedArmy [16]EncodedPiece

	ArmySlotIndex int
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

func (an ArmyRequest) AvailableSlotsForKind(k piece.Kind) []ArmySlotIndex {
	out := make([]ArmySlotIndex, 0, 8)
	switch k.Basic() {
	case piece.Pawn:
		for i := 0; i < 8; i++ {
			if an[i] != NotInCollection {
				continue
			}
			out = append(out, ArmySlotIndex(i))
		}
	case piece.Rook:
		if an[8] == NotInCollection {
			out = append(out, 8)
		}
		if an[15] == NotInCollection {
			out = append(out, 15)
		}
	case piece.Knight:
		if an[9] == NotInCollection {
			out = append(out, 9)
		}
		if an[14] == NotInCollection {
			out = append(out, 14)
		}
	case piece.Bishop:
		if an[10] == NotInCollection {
			out = append(out, 10)
		}
		if an[13] == NotInCollection {
			out = append(out, 13)
		}
	}
	return out
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
