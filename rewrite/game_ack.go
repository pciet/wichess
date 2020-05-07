package main

import "database/sql"

// AcknowledgeGameComplete acknowledges the game for this player, then
// deletes it if both players have acknowledged it.
// AcknowledgeGameComplete works improperly if called before the game
// is actually complete.
func AcknowledgeGameComplete(tx *sql.Tx, id GameIdentifier, player string) {
	h := LoadGameHeader(tx, id, true)
	if ((player == h.White.Name) && h.Black.Acknowledge) ||
		((player == h.Black.Name) && h.White.Acknowledge) ||
		(h.Black.Name == ComputerPlayerName) {
		DeleteGame(tx, id)
		return
	}

	var ackKey string
	if player == h.White.Name {
		ackKey = GamesWhiteAcknowledge
	} else if player == h.Black.Name {
		ackKey = GamesBlackAcknowledge
	} else {
		Panic("GameHasPlayer shows", player, "in", id, ", but player not in header")
	}

	r, err := tx.Exec(GamesAcknowledgeUpdate, ackKey, true, id)
	if err != nil {
		Panic(err)
	}

	c, err := r.RowsAffected()
	if err != nil {
		Panic(err)
	}
	if c != 1 {
		Panic(c, "rows affected by", GamesAcknowledgeUpdate)
	}
}
