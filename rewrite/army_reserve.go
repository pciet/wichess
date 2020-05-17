package main

import (
	"database/sql"
	"fmt"

	"github.com/pciet/wichess/rules"
)

// ReserveArmies verifies the requested armies are valid then updates the player's collections.
// Requested pieces in each army are flagged in the database to be in use, and requested random
// pick piece slots are replaced with new random pieces. The white and black armies (returned in
// that order) are encoded to be ready for insertion into a games table row, and any used pick
// slots have their kinds returned.
// If either army is invalid then an error is returned and there are no database effects.
func ReserveArmies(tx *sql.Tx, wa, ba ArmyRequest,
	whiteID, blackID int) (EncodedArmy, RandomPicks, EncodedArmy, RandomPicks, error) {

	whiteReservation, whiteLeft, whiteRight, err := MakeArmyReservation(tx, whiteID, wa)
	if err != nil {
		return EncodedArmy{}, RandomPicks{}, EncodedArmy{}, RandomPicks{}, err
	}

	left, right := PickSlotsInArmyRequest(wa)
	whitePicks := RandomPicks{rules.NoKind, rules.NoKind}
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
	blackPicks := RandomPicks{rules.NoKind, rules.NoKind}
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

// MakeArmyReservation does one query to the player's database row to get information needed to
// encode the pieces for insertion into the games table. Both random pick slots are always queried
// and the kinds returned for use in ReserveArmy to replace without duplication.
func MakeArmyReservation(tx *sql.Tx,
	playerID int, r ArmyRequest) ([16]rules.PieceKind, rules.PieceKind, rules.PieceKind, error) {

	if playerID == ComputerPlayerID {
		return BasicArmy, rules.NoKind, rules.NoKind, nil
	}

	left, right := false, false
	collectionRequests := make([]CollectionSlot, 0, 4)

	for _, request := range r {
		switch request {
		case NotInCollection:
			continue
		case LeftPick:
			if left == true {
				return [16]rules.PieceKind{}, rules.NoKind, rules.NoKind,
					fmt.Errorf("multiple left pick requests for %v", playerID)
			}
			left = true
			continue
		case RightPick:
			if right == true {
				return [16]rules.PieceKind{}, rules.NoKind, rules.NoKind,
					fmt.Errorf("multiple right pick requests for %v", playerID)
			}
			right = true
			continue
		}

		if request > CollectionCount {
			return [16]rules.PieceKind{}, rules.NoKind, rules.NoKind,
				fmt.Errorf("request %v for %v out of collection bounds", request, playerID)
		}

		for _, alreadyRequested := range collectionRequests {
			if alreadyRequested == request {
				return [16]rules.PieceKind{}, rules.NoKind, rules.NoKind,
					fmt.Errorf("duplicate collection request %v from %v", request, playerID)
			}
		}

		// postgres array indices start at 1, so collection slots exactly match
		collectionRequests = append(collectionRequests, request)
	}

	collectionPieces, leftKind, rightKind := PlayerSelectedCollectionPieces(tx, playerID,
		collectionRequests)

	var out [16]rules.PieceKind
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
		if p.InUse {
			return [16]rules.PieceKind{}, rules.NoKind, rules.NoKind,
				fmt.Errorf("collection piece %v for %v in use", p, playerID)
		}
		collectionPiecesIndex++
		out[i] = p.Kind
	}

	return out, leftKind, rightKind, nil
}

// ReserveArmy updates the player's database row with the requested collection pieces flagged to
// be in use, it replaces random pick slots that are used, and all pieces, whether in the
// collection or not, are encoded for insertion into the games table.
func ReserveArmy(tx *sql.Tx, playerID int, o rules.Orientation,
	pieces [16]rules.PieceKind, left, right rules.PieceKind, r ArmyRequest) EncodedArmy {

	collectionSlotsToUpdate := make([]CollectionSlot, 0, 4)
	collectionSlotKinds := make([]rules.PieceKind, 0, 4)
	var army EncodedArmy
	replaceLeft, replaceRight := false, false

	for i, c := range r {
		army[i] = Piece{
			Piece: rules.Piece{
				Orientation: o,
				Kind:        pieces[i],
			},
		}.Encode()

		switch c {
		case NotInCollection:
			continue
		case LeftPick:
			replaceLeft = true
			continue
		case RightPick:
			replaceRight = true
			continue
		}

		collectionSlotsToUpdate = append(collectionSlotsToUpdate, c)
		collectionSlotKinds = append(collectionSlotKinds, pieces[i])
	}

	if playerID == ComputerPlayerID {
		return army
	}

	PlayerCollectionFlagInUse(tx, playerID,
		collectionSlotsToUpdate, collectionSlotKinds, left, right, replaceLeft, replaceRight)

	return army
}
