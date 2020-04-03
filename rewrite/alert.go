package main

// Alert is called to notify the opponent when a game has progressed.
// Changes are expected to have already been committed to the database.
// If the opponent is the computer player then their move is made.
func Alert(in GameIdentifier, player string, ofChanges Update) {
	if player == ComputerPlayerName {
		// a wait happens if the opponent still needs to promote
		if ofChanges.State != WaitUpdate {
			Autoplay(in, ComputerPlayerName)
		}
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
