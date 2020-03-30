package main

import "github.com/pciet/wichess/rules"

// Alert is called to notify the opponent when a game has progressed.
// Changes are expected to have already been committed to the database.
// If the opponent is the computer player then their move is made.
// A zero length ofChanges slice indicates the game is completed (no moves
// could be made).
func Alert(in GameIdentifier, player string, ofChanges []rules.AddressedSquare) {
	if player == ComputerPlayerName {
		Autoplay(in, ComputerPlayerName)
		return
	}

	conn, has := Connected(in, player)
	if has == false {
		return
	}

	err := conn.WriteJSON(ofChanges)
	if err != nil {
		DebugPrintln("failed WebSocket write for", player, "in", in, ":", err)
		return
	}
}
