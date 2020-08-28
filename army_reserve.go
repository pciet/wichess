package wichess

import (
	"database/sql"
	"fmt"

	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// ReserveArmies verifies the requested armies are valid and replaces used pick slots. The white
// and black armies (returned in that order) are encoded for insertion into a games table row,
// and the kinds of used pick slots are returned. If either army is invalid then an error is
// returned and there are no database effects.
func ReserveArmies(tx *sql.Tx, wa, ba ArmyRequest,
	whiteID, blackID PlayerIdentifier) (EncodedArmy, RandomPicks, EncodedArmy, RandomPicks, error) {

	whiteReservation, whiteLeft, whiteRight, err := MakeArmyReservation(tx, whiteID, wa)
	if err != nil {
		return EncodedArmy{}, RandomPicks{}, EncodedArmy{}, RandomPicks{}, err
	}

	left, right := PickSlotsInArmyRequest(wa)
	whitePicks := RandomPicks{piece.NoKind, piece.NoKind}
	if left {
		whitePicks.Left = whiteLeft
	}
	if right {
		whitePicks.Right = whiteRight
	}

	blackReservation, blackLeft, blackRight, err := MakeArmyReservation(tx, blackID, ba)
	if err != nil {
		return EncodedArmy{}, RandomPicks{}, EncodedArmy{}, RandomPicks{}, err
	}

	left, right = PickSlotsInArmyRequest(ba)
	blackPicks := RandomPicks{piece.NoKind, piece.NoKind}
	if left {
		blackPicks.Left = blackLeft
	}
	if right {
		blackPicks.Right = blackRight
	}

	return ReserveArmy(tx, whiteID, rules.White, whiteReservation, whiteLeft, whiteRight, wa),
		whitePicks,
		ReserveArmy(tx, blackID, rules.Black, blackReservation, blackLeft, blackRight, ba),
		blackPicks, nil
}

// MakeArmyReservation queries/reads the player's database row to translate ArmyRequest values into
// piece kinds. Kinds of the two picks are always queried for and returned.
func MakeArmyReservation(tx *sql.Tx, id PlayerIdentifier,
	r ArmyRequest) ([16]piece.Kind, piece.Kind, piece.Kind, error) {

	if id == ComputerPlayerID {
		return BasicArmy, piece.NoKind, piece.NoKind, nil
	}

	left, right := false, false
	collectionRequests := make([]CollectionSlot, 0, 4)

	for _, request := range r {
		switch request {
		case NotInCollection:
			continue
		case LeftPick:
			if left == true {
				return [16]piece.Kind{}, piece.NoKind, piece.NoKind,
					fmt.Errorf("multiple left pick requests for %v", id)
			}
			left = true
			continue
		case RightPick:
			if right == true {
				return [16]piece.Kind{}, piece.NoKind, piece.NoKind,
					fmt.Errorf("multiple right pick requests for %v", id)
			}
			right = true
			continue
		}

		if request > CollectionCount {
			return [16]piece.Kind{}, piece.NoKind, piece.NoKind,
				fmt.Errorf("request %v for %v out of collection bounds", request, id)
		}

		for _, alreadyRequested := range collectionRequests {
			if alreadyRequested == request {
				return [16]piece.Kind{}, piece.NoKind, piece.NoKind,
					fmt.Errorf("duplicate collection request %v from %v", request, id)
			}
		}

		// postgres array indices start at 1, so collection slots exactly match
		collectionRequests = append(collectionRequests, request)
	}

	collectionPieces, leftKind, rightKind := PlayerSelectedCollectionPieces(tx, id,
		collectionRequests)

	var out [16]piece.Kind
	collectionPiecesIndex := 0

	for i, request := range r {
		switch request {
		case NotInCollection:
			out[i] = BasicArmy[i]
			continue
		case LeftPick:
			out[i] = leftKind
			continue
		case RightPick:
			out[i] = rightKind
			continue
		}
		p := collectionPieces[collectionPiecesIndex]
		collectionPiecesIndex++
		out[i] = p.Kind
	}

	return out, leftKind, rightKind, nil
}

// ReserveArmy replaces used pick slots and encodes all pieces (whether in the collection or not)
// for insertion into the games table.
func ReserveArmy(tx *sql.Tx, id PlayerIdentifier, o rules.Orientation,
	pieces [16]piece.Kind, left, right piece.Kind, r ArmyRequest) EncodedArmy {

	var army EncodedArmy
	replaceLeft, replaceRight := false, false

	for i, c := range r {
		var start rules.AddressIndex
		if o == rules.White {
			if i < 8 {
				start = rules.AddressIndex(8 + i)
			} else {
				start = rules.AddressIndex(i - 8)
			}
		} else if o == rules.Black {
			start = rules.AddressIndex(i + 48)
		} else {
			Panic("bad orientation", o)
		}

		army[i] = Piece{
			Piece: rules.Piece{
				Orientation: o,
				Start:       start.Address(),
				Kind:        pieces[i],
			},
		}.Encode()

		switch c {
		case LeftPick:
			replaceLeft = true
		case RightPick:
			replaceRight = true
		}
	}

	if id == ComputerPlayerID {
		return army
	}

	PlayerCollectionReplacePicks(tx, id, left, right, replaceLeft, replaceRight)

	return army
}
