package game

import (
	"fmt"
	"log"

	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// reserveArmies verifies the requested armies are valid and replaces used pick slots. The white
// and black armies are returned in that order, and the kinds of used pick slots are returned. If
// either army is invalid then an error is returned and there are no memory effects.
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

	return reserveArmy(wid, rules.White, whiteReservation, whiteLeft, whiteRight, wa),
		whitePicks,
		reserveArmy(bid, rules.Black, blackReservation, blackLeft, blackRight, ba),
		blackPicks, nil
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
		leftKind, rightKind := memory.SelectedCollectionPieces(id, collectionRequests)

	var out piece.Army
	collectionPiecesIndex := 0

	for i, request := range r {
		switch request {
		case piece.NotInCollection:
			out[i] = piece.RegularArmy[i]
			continue
		case piece.LeftPick:
			out[i] = leftKind
			continue
		case piece.RightPick:
			out[i] = rightKind
			continue
		}
		p := collectionPieces[collectionPiecesIndex]
		collectionPiecesIndex++
		out[i] = p.Kind
	}

	return out, leftKind, rightKind, nil
}

// TODO: reserveArmy needs to be redone

// reserveArmy replaces used pick slots and encodes all pieces (whether in the collection or not)
// for use in game memory.
func reserveArmy(id memory.PlayerIdentifier, o rules.Orientation,
	pieces piece.Army, left, right piece.Kind, r piece.ArmyRequest) piece.Army {

	var army piece.Army
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
			log.Panicln("bad orientation", o)
		}

		army[i] = Piece{
			Piece: rules.Piece{
				Orientation: o,
				Start:       start.Address(),
				Kind:        pieces[i],
			},
		}

		switch c {
		case memory.LeftPick:
			replaceLeft = true
		case memory.RightPick:
			replaceRight = true
		}
	}

	if id == memory.ComputerPlayerIdentifier {
		return army
	}

	// TODO: interact with package player correctly
	player.ReplaceCollectionPicks(id, left, right, replaceLeft, replaceRight)

	return army
}
