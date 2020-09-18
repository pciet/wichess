package memory

import (
	"strconv"
	"sync"

	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

type (
	// Each game has a unique GameIdentifier int. These integers can be reused after a game has
	// been deleted. The first game is 1 and others are larger.
	GameIdentifier int

	// Game represents a chess board and all information needed to calculate available moves
	// and the results of moves. Methods change this memory in-place. Games are encoded as JSON
	// in the backing files.
	Game struct {
		sync.RWMutex   `json:"-"`
		GameIdentifier `json:"id"`
		Active         rules.Orientation `json:"active"`
		PreviousActive rules.Orientation `json:"prevactive"`
		White          GamePlayer        `json:"white"`
		Black          GamePlayer        `json:"black"`
		PreviousMove   rules.Move        `json:"prevmove"`
		DrawTurns      int               `json:"drawturns"`
		Conceded       bool              `json:"conceded"`
		rules.Board    `json:"board"`
	}

	// GamePlayer is a field in a Game representing one of the two players.
	GamePlayer struct {
		PlayerIdentifier `json:"id"`
		Acknowledge      bool `json:"ack"`
		Captures         `json:"captures"`
		Left             piece.Kind `json:"left"`
		Right            piece.Kind `json:"right"`
		Reward           piece.Kind `json:"reward"`
	}

	// Captures is a list of pieces a player has captured in ascending time order.
	Captures [15]piece.Kind
)

// NoGame is the value of a GameIdentifier var when it's not representing a game.
const NoGame = 0

// RulesGame returns the minimal information needed to do calulations using package rules.
func (a *Game) RulesGame() rules.Game {
	return rules.Game{
		Board:        &(a.Board),
		PreviousMove: a.PreviousMove,
	}
}

// CanDelete signals that this game is no longer needed and can be permanently deleted sometime
// by this package.
func (a *Game) CanDelete() {
	id := a.GameIdentifier
	go func() {
		activeMutex.RLock()
		gamesMutex.Lock()

		delete(gamesCache, id)

		gamesMutex.Unlock()

		deleteGameFile(id)
		activeMutex.RUnlock()
	}()
}

func (a GameIdentifier) Int() int       { return int(a) }
func (a GameIdentifier) String() string { return strconv.Itoa(a.Int()) }

// FirstAvailableIndex returns the first array index that's piece.NoKind.
func (the *Captures) FirstAvailableIndex() int {
	for i := 0; i < len(a.Captures); i++ {
		if a.Captures[i] == piece.NoKind {
			return i
		}
	}
	panic("more than 15 captures recorded")
	return 15
}
