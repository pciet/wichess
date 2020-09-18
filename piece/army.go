package piece

import (
	"encoding/json"
	"io"
)

// ArmySize is the number of pieces in an army. This number is the same as regular chess.
const ArmySize = 16

type (
	// An ArmyIndex is a location in an army, 0-15. Numbering starts with the player's left pawn
	// at zero and ascends to the right to seven for the far right pawn, then jumps back to the
	// left with the queenside rook at eight and ascends again to the right up to the right rook
	// at index 15.
	ArmyIndex int

	// An ArmyRequest points zero or more army indices at slots in the collection for use in a game.
	ArmyRequest [ArmySize]CollectionSlot

	// An Army is a list of the kinds of pieces in a player's army.
	Army [ArmySize]Kind
)

var (
	// A RegularArmyRequest is a request for a regular chess army without any special pieces.
	RegularArmyRequest = ArmyRequest{}

	// RegularArmy is a chess army without any special pieces.
	RegularArmy = func() Army {
		var b Army
		for i := 0; i < 8; i++ {
			b[i] = Pawn
		}

		b[8] = Rook
		b[15] = Rook

		b[9] = Knight
		b[14] = Knight

		b[10] = Bishop
		b[13] = Bishop

		b[11] = Queen
		b[12] = King

		return b
	}()
)

// DecodeArmyRequest takes a JSON encoding of an ArmyRequest and decodes it.
func DecodeArmyRequest(jsonBuf io.Reader) (ArmyRequest, error) {
	var a ArmyRequest
	err := json.NewDecoder(jsonBuf).Decode(&a)
	return a, err
}

// PicksUsed indicates if the left and/or right random picks are included.
func (an *ArmyRequest) PicksUsed() (bool, bool) {
	left, right := false, false
	for _, s := range an {
		if s == LeftPick {
			left = true
		} else if s == RightPick {
			right = true
		}
	}
	return left, right
}

func (an ArmyRequest) AvailableSlotsForKind(k Kind) []ArmyIndex {
	out := make([]ArmyIndex, 0, 8)
	switch k.Basic() {
	case Pawn:
		for i := 0; i < 8; i++ {
			if an[i] != NotInCollection {
				continue
			}
			out = append(out, ArmyIndex(i))
		}
	case Rook:
		if an[8] == NotInCollection {
			out = append(out, 8)
		}
		if an[15] == NotInCollection {
			out = append(out, 15)
		}
	case Knight:
		if an[9] == NotInCollection {
			out = append(out, 9)
		}
		if an[14] == NotInCollection {
			out = append(out, 14)
		}
	case Bishop:
		if an[10] == NotInCollection {
			out = append(out, 10)
		}
		if an[13] == NotInCollection {
			out = append(out, 13)
		}
	}
	return out
}
