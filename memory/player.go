package memory

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/pciet/wichess/piece"
)

// NoPlayer is the value of a PlayerIdentifier var when it doesn't represent a player.
const NoPlayer = 0

// PlayerNameMaxSize is the maximum number of runes in a PlayerName.
const PlayerNameMaxSize = 21

type (
	// A PlayerIdentifier is a positive integer that uniquely identifies a player.
	PlayerIdentifier int

	// A PlayerName is the unicode name of the player. All runes must return true with
	// unicode.IsGraphic.
	PlayerName string

	// Player represents a person that is playing in games on this host. This memory includes the
	// player's collection of pieces used to customize their army before starting a new match.
	Player struct {
		sync.RWMutex `json:"-"`

		PlayerIdentifier `json:"id"`
		PlayerName       `json:"name"`

		// The session key is encoded in base64.
		EncodedSessionKey string `json:"-"`

		PeopleGame   GameIdentifier `json:"people"`
		ComputerGame GameIdentifier `json:"computer"`

		ComputerStreak     int `json:"compstreak"`
		BestComputerStreak int `json:"bestcompstreak"`

		RecentOpponents [RecentOpponentCount]PlayerIdentifier `json"recent"`

		Left             piece.Kind `json:"left"`
		Right            piece.Kind `json:"right"`
		piece.Collection `json:"collection"`
	}
)

func (a PlayerIdentifier) String() string { return strconv.Itoa(int(a)) }
func (a PlayerName) String() string       { return string(a) }

func (a Player) String() string {
	str := fmt.Sprintf("%v\n%v\n", a.PlayerIdentifier, a.PlayerName)
	str += fmt.Sprintf("session key %v\n", a.EncodedSessionKey)
	str += fmt.Sprintf("people game %v\n", a.PeopleGame)
	str += fmt.Sprintf("computer game %v\n", a.ComputerGame)
	str += fmt.Sprintf("computer streak %v, best computer streak %v\n",
		a.ComputerStreak, a.BestComputerStreak)
	str += fmt.Sprintf("recent opponents %v\n", a.RecentOpponents)
	str += fmt.Sprintf("left %v\nright %v\n", a.Left, a.Right)
	str += fmt.Sprintf("collection %v", a.Collection)
	return str
}
