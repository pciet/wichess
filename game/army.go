package game

import (
	"fmt"

	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
)

// reserveArmies verifies the requested armies are valid and replaces used pick slots. The white
// and black armies are returned in that order, and the kinds of used pick slots are returned. If
// either army is invalid then an error is returned and there are no memory effects. Player's
// memory must be available for reading using a memory.RLockPlayer.
func reserveArmies(wa, ba piece.ArmyRequest,
	wid, bid memory.PlayerIdentifier) (piece.Army, piece.RandomPicks,
	piece.Army, piece.RandomPicks, error) {

	// TODO: this logic is repeated

	whiteReservation, whiteLeft, whiteRight, err := makeArmyReservation(wid, wa)
	if err != nil {
		return piece.Army{}, piece.RandomPicks{}, piece.Army{}, piece.RandomPicks{}, err
	}

	left, right := wa.PicksUsed()
	whitePicks := piece.RandomPicks{piece.NoKind, piece.NoKind}
	if left {
		whitePicks.Left = whiteLeft
	}
	if right {
		whitePicks.Right = whiteRight
	}

	blackReservation, blackLeft, blackRight, err := makeArmyReservation(bid, ba)
	if err != nil {
		return piece.Army{}, piece.RandomPicks{}, piece.Army{}, piece.RandomPicks{}, err
	}

	left, right = ba.PicksUsed()
	blackPicks := piece.RandomPicks{piece.NoKind, piece.NoKind}
	if left {
		blackPicks.Left = blackLeft
	}
	if right {
		blackPicks.Right = blackRight
	}

	return whiteReservation, whitePicks, blackReservation, blackPicks, nil
}

// makeArmyReservation queries memory to translate ArmyRequest values into piece kinds. Kinds of
// the two picks are always queried for and returned.
func makeArmyReservation(id memory.PlayerIdentifier,
	r piece.ArmyRequest) (piece.Army, piece.Kind, piece.Kind, error) {

	if id == memory.ComputerPlayerIdentifier {
		return piece.RegularArmy, piece.NoKind, piece.NoKind, nil
	}

	left, right := false, false
	collectionRequests := make([]piece.CollectionSlot, 0, 4)

	for _, request := range r {
		switch request {
		case piece.NotInCollection:
			continue
		case piece.LeftPick:
			if left == true {
				return piece.Army{}, piece.NoKind, piece.NoKind,
					fmt.Errorf("multiple left pick requests for %v", id)
			}
			left = true
			continue
		case piece.RightPick:
			if right == true {
				return piece.Army{}, piece.NoKind, piece.NoKind,
					fmt.Errorf("multiple right pick requests for %v", id)
			}
			right = true
			continue
		}

		if request > piece.CollectionSize {
			return piece.Army{}, piece.NoKind, piece.NoKind,
				fmt.Errorf("request %v for %v out of collection bounds", request, id)
		}

		for _, alreadyRequested := range collectionRequests {
			if alreadyRequested == request {
				return piece.Army{}, piece.NoKind, piece.NoKind,
					fmt.Errorf("duplicate collection request %v from %v", request, id)
			}
		}

		collectionRequests = append(collectionRequests, request)
	}

	collectionPieces,
		leftKind, rightKind := selectedCollectionPieces(id, collectionRequests)

	if collectionPieces == nil {
		return piece.Army{}, piece.NoKind, piece.NoKind,
			fmt.Errorf("selectedCollectionPieces failed for %v", id)
	}

	var out piece.Army
	collectionPiecesIndex := 0

	for i, request := range r {
		switch request {
		case piece.NotInCollection:
			out[i] = piece.RegularArmy[i]
			continue
		case piece.LeftPick:
			if leftKind.Basic() != piece.RegularArmy[i] {
				return piece.Army{}, piece.NoKind, piece.NoKind,
					fmt.Errorf("requested army slot %v not correct basic kind for left pick %v",
						i, leftKind)
			}
			out[i] = leftKind
			continue
		case piece.RightPick:
			if rightKind.Basic() != piece.RegularArmy[i] {
				return piece.Army{}, piece.NoKind, piece.NoKind,
					fmt.Errorf("requested army slot %v not correct basic kind for right pick %v",
						i, rightKind)
			}
			out[i] = rightKind
			continue
		}
		if collectionPieces[collectionPiecesIndex].Basic() != piece.RegularArmy[i] {
			return piece.Army{}, piece.NoKind, piece.NoKind,
				fmt.Errorf("requested army slot %v not correct basic kind for collection piece %v",
					i, collectionPieces[collectionPiecesIndex])
		}
		out[i] = collectionPieces[collectionPiecesIndex]
		collectionPiecesIndex++
	}

	return out, leftKind, rightKind, nil
}
