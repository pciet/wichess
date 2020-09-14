package memory

import (
	"sync"

	"github.com/pciet/wichess/piece"
)

type (
	// A PlayerIdentifier is a positive integer that uniquely identifies a player.
	PlayerIdentifier int

	// A PlayerName is the unicode name of the player. All runes must return true with
	// unicode.IsGraphic.
	PlayerName string

	Player struct {
		sync.RWMutex
		PlayerIdentifier
		PlayerName
		PeopleGame, ComputerGame           GameIdentifier
		ComputerStreak, BestComputerStreak int
		RecentOpponents                    [RecentOpponentCount]PlayerIdentifier
		Left, Right                        piece.Kind
		piece.Collection
	}
)

const (
	// NoPlayer is the value of a PlayerIdentifier var when it doesn't represent a player.
	NoPlayer = 0

	// PlayerNameMaxSize is the maximum number of runes in a PlayerName.
	PlayerNameMaxSize = 21

	// The easy artificial computer player "punching bag" opponent has a reserved name and id.
	ComputerPlayerName       = "Computer Player"
	ComputerPlayerIdentifier = 0
)

func (a PlayerIdentifier) IsComputerPlayer() bool { return a == ComputerPlayerIdentifier }

func NewPlayer(name PlayerName, crypt []byte) PlayerIdentifier {

}

const RecentOpponentCount = 5

// AddPlayerRecentOpponent updates the list of recent opponents by inserting this one at the top.
// The list is updated to remove duplicates, and opponents past the bottom of the list are lost.
func AddPlayerRecentOpponent(player, opponent PlayerIdentifier) {
	tx := DatabaseTransaction()

	// TODO: does this need to be a FOR UPDATE query?
	rec := PlayerRecentOpponentIDs(tx, player)

	// remove possible one duplicate of this opponent then condense the list
	for i, opp := range rec {
		if opp != opponent {
			continue
		}
		for j := i; j < RecentOpponentCount; j++ {
			if j == (RecentOpponentCount - 1) {
				rec[j] = 0
				break
			}
			rec[j] = rec[j+1]
		}
		break
	}

	// insert the opponent at the start of the list
	for i := (RecentOpponentCount - 1); i > 0; i-- {
		rec[i] = rec[i-1]
	}
	rec[0] = opponent

	UpdatePlayerRecentOpponents(tx, player, rec)

	tx.Commit()
}

func (an Instance) IncrementComputerStreak() {
	an.ComputerStreak++
	if an.ComputerStreak > an.BestComputerStreak {
		an.BestComputerStreak = an.ComputerStreak
	}
	an.Changed()
}

func (an Instance) ResetComputerStreak() {
	an.ComputerStreak = 0
	an.Changed()
}
