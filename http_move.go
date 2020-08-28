package wichess

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

const MovePath = "/move/"

var MoveHandler = AuthenticRequestHandler{
	Post: GameIdentifierParsed(PlayerNamed(MovePost), MovePath),
}

// The webpage sends a move request in the MoveJSON format.
// If a promotion is requested then p is nonzero and f/t is ignored.
type MoveJSON struct {
	From      int `json:"f"`
	To        int `json:"t"`
	Promotion int `json:"p"` // piece.Kind instead of rules.AddressIndex
}

func MovePost(w http.ResponseWriter, r *http.Request, tx *sql.Tx,
	id GameIdentifier, requester Player) {
	var body bytes.Buffer
	_, err := body.ReadFrom(http.MaxBytesReader(w, r.Body, 1024))
	if err != nil {
		tx.Commit()
		DebugPrintln(MovePath, "body read failed for", requester, ":", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var mj MoveJSON
	err = json.Unmarshal(body.Bytes(), &mj)
	if err != nil {
		tx.Commit()
		DebugPrintln(MovePath, "failed to unmarshal json for", requester, ":", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if (mj.From == 0) && (mj.To == 0) && (mj.Promotion == 0) {
		tx.Commit()
		DebugPrintln(MovePath, "all zero move request from", requester)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	move, promotionKind := ParseMove(mj)
	if move == rules.NoMove {
		tx.Commit()
		DebugPrintln(MovePath, "failed to parse move by", requester)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	changes, takes, promotionNeeded := Move(tx, id, requester.Name, move, promotionKind)
	if (changes == nil) || (len(changes) == 0) {
		tx.Commit()
		DebugPrintln(MovePath,
			"bad move from", requester, "for game", id, ":", move, promotionKind)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	previousActive := GamePreviousActive(tx, id)
	opponent := GameOpponent(tx, id, requester.Name)

	tx.Commit()

	promotionWasReverse := false
	if (promotionKind != piece.NoKind) && (previousActive != requester.Name) {
		promotionWasReverse = true
	}

	alertUpdate := Update{Squares: changes, Captures: takes, FromMove: move}
	if promotionNeeded || promotionWasReverse {
		alertUpdate.State = WaitUpdate
	}
	go Alert(id, opponent, alertUpdate)

	responseUpdate := Update{Squares: changes, Captures: takes}
	if promotionNeeded {
		responseUpdate.State = PromotionNeededUpdate
	} else if promotionWasReverse {
		responseUpdate.State = ContinueUpdate
	}

	JSONResponse(w, responseUpdate)
}

func ParseMove(m MoveJSON) (rules.Move, piece.Kind) {
	var to rules.Address
	promotion := piece.Kind(m.Promotion)
	if promotion != piece.NoKind {
		if promotion.IsBasic() == false {
			DebugPrintln("promotion request", promotion, "not basic kind")
			return rules.NoMove, piece.NoKind
		}
		to = rules.NoAddress
	} else {
		to = rules.AddressIndex(m.To).Address()
	}
	return rules.Move{
		rules.AddressIndex(m.From).Address(),
		to,
	}, promotion
}
