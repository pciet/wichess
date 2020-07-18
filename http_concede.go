package main

import (
	"database/sql"
	"net/http"

	"github.com/pciet/wichess/rules"
)

const ConcedePath = "/concede/"

var ConcedeHandler = AuthenticRequestHandler{
	Get: GameIdentifierParsed(PlayerNamed(ConcedeGet), ConcedePath),
}

func ConcedeGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx,
	id GameIdentifier, requester Player) {

	opp := GameOpponent(tx, id, requester.Name)
	if opp != ComputerPlayerName {
		// the conceding player doesn't reload the game, so if it's marked conceded then you know
		// it was your opponent that conceded
		MarkGameConceded(tx, id)
		go Alert(id, opp, Update{
			State:    ConcededUpdate,
			FromMove: rules.NoMove,
		})
		// the only kind of game without the computer player opponent is the people game
		UpdatePlayerActivePeopleGame(tx, requester.ID, 0)
	} else {
		PlayerResetComputerStreak(tx, requester.ID)
		DeleteGame(tx, id)
	}
	tx.Commit()

	// on a success response the web browser redirects to the index
}
