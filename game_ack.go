package wichess

import "database/sql"

// AcknowledgeGameComplete acknowledges the game for this player, then deletes it if both
// players have acknowledged it.
// AcknowledgeGameComplete works improperly if called before the game is actually complete.
func AcknowledgeGameComplete(tx *sql.Tx, id GameIdentifier, player string) {
	h := LoadGameHeader(tx, id, true)
	if ((player == h.White.Name) && h.Black.Acknowledge) ||
		((player == h.Black.Name) && h.White.Acknowledge) ||
		(h.Black.Name == ComputerPlayerName) || (h.White.Name == ComputerPlayerName) {
		DeleteGame(tx, id)
		return
	}

	var exec string
	if player == h.White.Name {
		exec = GamesAcknowledgeWhiteUpdate
	} else if player == h.Black.Name {
		exec = GamesAcknowledgeBlackUpdate
	} else {
		Panic("GameHasPlayer shows", player, "in", id, ", but player not in header")
	}

	SQLExecRow(tx, exec, true, id)
}
