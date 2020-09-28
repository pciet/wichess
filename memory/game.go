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
		sync.RWMutex `json:"-"`

		// volatile caches of the current turn's result of rules.Board.Moves
		MovesCache []rules.MoveSet `json:"-"`
		StateCache rules.State     `json:"-"`

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

func (a *Game) Copy() *Game {
	c := Game{
		MovesCache:     rules.CopyMoveSetSlice(a.MovesCache),
		StateCache:     a.StateCache,
		GameIdentifier: a.GameIdentifier,
		Active:         a.Active,
		PreviousActive: a.PreviousActive,
		White:          a.White,
		Black:          a.Black,
		PreviousMove:   a.PreviousMove,
		DrawTurns:      a.DrawTurns,
		Conceded:       a.Conceded,
		Board:          a.Board,
	}
	return &c
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
	for i := 0; i < len(the); i++ {
		if the[i] == piece.NoKind {
			return i
		}
	}
	panic("more than 15 captures recorded")
	return 15
}
