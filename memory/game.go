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

	Game struct {
		sync.RWMutex
		GameIdentifier
		Active, PreviousActive rules.Orientation
		White, Black           GamePlayer
		PreviousMove           rules.Move
		DrawTurns              int
		Conceded               bool
		rules.Board
	}

	GamePlayer struct {
		PlayerIdentifier
		Acknowledge bool
		Captures
		Left, Right, Reward piece.Kind
	}

	// Captures is a list of pieces a player has captured in ascending time order.
	Captures [15]piece.Kind
)

// NoGame is the value of a GameIdentifier var when it's not representing a game.
const NoGame = 0

// NoMove is the initial value of the From and To address indices for a new game.
// Board addresses are 0-63, so 64 is not a normal address index.
const NoMoveIndex = 64

func (a *Game) RulesGame() rules.Game {
	return rules.Game{
		Board:        &(a.Board),
		PreviousMove: a.PreviousMove,
	}
}

func (a *Game) CanDelete() {

}

// Changed signals that, if the game continues to exist, sometime before app shutdown its latest
// version should be written to a file.
func (a *Game) Changed() {
	id := a.GameIdentifier
	go func() {
		gameFileWritesMutex.Lock()
		defer gameFileWritesMutex.Unlock()
		for _, gid := range gameFileWrites {
			if gid == id {
				return
			}
		}
		gameFileWrites = append(gameFileWrites, id)
	}()
}

func (a GameIdentifier) Int() int       { return int(a) }
func (a GameIdentifier) String() string { return strconv.Itoa(a.Int()) }

// FirstAvailableCaptureIndex returns the first array index that's piece.NoKind.
func FirstAvailableCaptureIndex(of *Captures) int {
	for i := 0; i < len(of); i++ {
		if of[i] == piece.NoKind {
			return i
		}
	}
	panic("more than 15 captures recorded")
	return 15
}
