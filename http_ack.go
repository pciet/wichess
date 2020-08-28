package wichess

import (
	"database/sql"
	"net/http"

	"github.com/pciet/wichess/rules"
)

const AcknowledgePath = "/acknowledge/"

var AcknowledgeHandler = AuthenticRequestHandler{
	Get: GameIdentifierParsed(PlayerNamed(AcknowledgeGet), AcknowledgePath),
}

func AcknowledgeGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx,
	id GameIdentifier, requester Player) {
	defer tx.Commit()

	// TODO: pack all reads into one query?

	done, state := GameComplete(tx, id)
	if done == false {
		DebugPrintln(AcknowledgePath, requester,
			"requested acknowledge of", id, "but not complete")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if GameHasPlayer(tx, id, ComputerPlayerName) {
		if (GameActive(tx, id) != GamePlayersOrientation(tx, id, requester.Name)) &&
			(state == rules.Checkmate) {
			PlayerComputerStreakIncrement(tx, requester.ID)
		} else {
			PlayerResetComputerStreak(tx, requester.ID)
		}
	}

	AcknowledgeGameComplete(tx, id, requester.Name)

	if PlayerActivePeopleGame(tx, requester.ID) == id {
		UpdatePlayerActivePeopleGame(tx, requester.ID, 0)
	}
}
