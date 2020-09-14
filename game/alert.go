package game

import (
	"github.com/pciet/wichess/memory"
	"github.com/pciet/wichess/rules"
)

// Alert notifies a player of a game update, or nothing happens if the player isn't connected.
// The update should have already been written to host memory to avoid race conditions from browser
// responses happening before the memory is correct. If the player to alert is the computer player
// then their move is automatically made (resulting in an alert back to the originating player).
func Alert(in memory.GameIdentifier,
	po rules.Orientation, pid memory.PlayerIdentifier, ofChanges Update) {

	if pid == memory.ComputerPlayerIdentifier {
		// a wait happens if the opponent still needs to promote
		if ofChanges.State != WaitUpdate {
			err := autoplay(in, pid)
			if err != nil {
				panic(err.Error())
			}
		}
		return
	}

	conn := Connected(in, po)
	if conn == nil {
		return
	}

	err := conn.WriteJSON(ofChanges)
	if err != nil {
		RemoveConnection(in, po)
		return
	}
}
