// Package game implements the details of Wisconsin Chess host functionality related to the
// individual matches between players.
package game

import (
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// An Instance is an in-progress game between two players. The pieces on the board and other
// information needed to make move calculations are represented by the Game type in package memory.
type Instance struct {
	*memory.Game
	copy bool // if copied then memory for the game id won't be changed

	// calculated moves are cached here and reset when Moves or Promote is called
	moves []rules.MoveSet
	state rules.State
}

// Copy copies the Instance's memory.Game. The returned copy does not use the package memory
// system, so changes to it will not overwrite the original in RAM or in a file.
func (an Instance) Copy() Instance { return Instance{an.Game.Copy(), true, nil, an.state} }

// Lock and RLock return an Instance with the sync.RWMutex of the memory.Game called. This means
// the caller must call Unlock or RUnlock on the Instance.
func Lock(id memory.GameIdentifier) Instance  { return Instance{Game: memory.LockGame(id)} }
func RLock(id memory.GameIdentifier) Instance { return Instance{Game: memory.RLockGame(id)} }

// Nil is used to determine if Lock was called with an invalid identifier.
func (an Instance) Nil() bool { return an.Game == nil }

// Completed determines if the game instance is done (a checkmate, draw, or concession). The
// specific state of the game is also returned.
func (an Instance) Completed() (bool, rules.State) {
	// TODO: cache moves and state in game memory so moves called only once
	_, state := an.Moves()
	if (state == rules.Normal) || (state == rules.Promotion) || (state == rules.Check) {
		return false, state
	}
	return true, state
}

// Acknowledge acknowledges that this player is done reviewing this game. If both players have
// acknowledged then the game is marked for deletion in the background.
func (an Instance) Acknowledge(by memory.PlayerIdentifier) {
	if ((by == an.White.PlayerIdentifier) && an.Black.Acknowledge) ||
		((by == an.Black.PlayerIdentifier) && an.White.Acknowledge) || an.HasComputerPlayer() {
		an.CanDelete()
		return
	}
	if by == an.White.PlayerIdentifier {
		an.White.Acknowledge = true
	} else if by == an.Black.PlayerIdentifier {
		an.Black.Acknowledge = true
	}
}

// IsComputerGame indicates if this is a game against the computer player.
func (an Instance) IsComputerGame() bool {
	if (an.White.PlayerIdentifier == memory.ComputerPlayerIdentifier) ||
		(an.Black.PlayerIdentifier == memory.ComputerPlayerIdentifier) {
		return true
	}
	return false
}

// RewardsOf returns the left, right, and reward pieces for a player as noted in the game.
func (an Instance) RewardsOf(p memory.PlayerIdentifier) (piece.Kind, piece.Kind, piece.Kind) {
	if an.OrientationOf(p) == rules.White {
		return an.Game.White.Left, an.Game.White.Right, an.Game.White.Reward
	}
	return an.Game.Black.Left, an.Game.Black.Right, an.Game.Black.Reward
}
